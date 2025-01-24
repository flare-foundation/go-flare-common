package queue

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"golang.org/x/time/rate"
)

const defaultMaxAttempts uint64 = 10

const bucketSize = 10

const NotRatedDequeue string = "!!NotRatedDequeue!!"

// PriorityQueue is made up of two sub-queues - one regular and one with
// higher priority. Items can be enqueued in either queue and when dequeueing
// items from the priority queue are returned first.
type PriorityQueue[T any] struct {
	name            string
	regular         chan priorityQueueItem[T]
	priority        chan priorityQueueItem[T]
	limiter         *rate.Limiter
	workersSem      chan struct{}
	maxAttempts     uint64
	DeadLetterQueue chan T
	backoff         func() backoff.BackOff
	timeOff         time.Duration
}

type priorityQueueItem[T any] struct {
	value    T
	backoff  backoff.BackOff
	priority bool
}

// PriorityQueueInput values are used to construct a new PriorityQueue.
type PriorityQueueParams struct {
	Size                 int           `toml:"size"`                    // Capacity of the queue and priority queue
	MaxDequeuesPerSecond int           `toml:"max_dequeues_per_second"` // Set to 0 to disable rate-limiting.
	MaxWorkers           int           `toml:"max_workers"`             // Set to 0 for unlimited workers.
	MaxAttempts          int32         `toml:"max_attempts"`            // Set to negative for unlimited attempts. If unset or set to 0, the default value (10) is applied.
	TimeOff              time.Duration `toml:"time_off"`                // Only relevant if Backoff is not set.

	// Pass a callback to specify the backoff policy which affects when items
	// are returned to the queue after an error. By default, items are
	// re-queued after TimeOff.
	Backoff func() backoff.BackOff
}

// NewPriority constructs a new PriorityQueue.
func NewPriority[T any](input *PriorityQueueParams) PriorityQueue[T] {
	if input == nil {
		input = new(PriorityQueueParams)
	}

	var limiter *rate.Limiter

	if input.MaxDequeuesPerSecond > 0 {
		limiter = rate.NewLimiter(rate.Limit(input.MaxDequeuesPerSecond), bucketSize)
	} else {
		limiter = rate.NewLimiter(rate.Inf, 0)
	}

	q := PriorityQueue[T]{
		regular:  make(chan priorityQueueItem[T], input.Size),
		priority: make(chan priorityQueueItem[T], input.Size),
		limiter:  limiter,
		backoff:  input.Backoff,
		timeOff:  input.TimeOff,
	}

	if input.MaxWorkers > 0 {
		q.workersSem = make(chan struct{}, input.MaxWorkers)
	}

	if input.MaxAttempts > -1 {
		q.maxAttempts = uint64(input.MaxAttempts)
		if input.MaxAttempts == 0 {
			q.maxAttempts = defaultMaxAttempts
		}
		q.DeadLetterQueue = make(chan T, input.Size)
	}

	return q
}

// Enqueue adds an item to the queue with regular priority.
func (q *PriorityQueue[T]) Enqueue(ctx context.Context, item T) error {
	return q.enqueue(ctx, priorityQueueItem[T]{value: item, backoff: q.newBackoff(), priority: false})
}

func (q *PriorityQueue[T]) enqueue(ctx context.Context, item priorityQueueItem[T]) error {
	select {
	case q.regular <- item:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnqueuePriority adds an item to the queue with high priority.
func (q *PriorityQueue[T]) EnqueuePriority(ctx context.Context, item T) error {
	return q.enqueuePriority(ctx, priorityQueueItem[T]{value: item, backoff: q.newBackoff(), priority: true})
}

func (q *PriorityQueue[T]) enqueuePriority(ctx context.Context, item priorityQueueItem[T]) error {
	select {
	case q.priority <- item:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (q *PriorityQueue[T]) newBackoff() (bOff backoff.BackOff) {
	if q.backoff == nil {
		bOff = backoff.NewConstantBackOff(q.timeOff)
	} else {
		bOff = q.backoff()
	}

	if q.maxAttempts > 0 {
		bOff = backoff.WithMaxRetries(bOff, q.maxAttempts-1)
	}

	return bOff
}

// Dequeue removes an item from the queue and processes it using the provided handler
// function. If configured, rate limits and concurrent worker limits will be enforced.
// This function will block if an item is not immediately available for
// processing or if necessary to enforce limits.
func (q *PriorityQueue[T]) Dequeue(ctx context.Context, handler func(context.Context, T) error) error {
	item, err := q.dequeueWithRateLimit(ctx)
	if err != nil {
		return err
	}

	if q.workersSem != nil {
		if err := q.incrementWorkers(ctx); err != nil {
			return err
		}
		defer q.decrementWorkers()
	}

	// Avoid panic if the handler is nil - could be used to pop an item without processing.
	if handler == nil {
		return nil
	}

	err = handler(ctx, item.value)

	// If there was any error, we re-queue the item for processing again.
	if err != nil {
		q.handleError(ctx, item)
		return err
	}

	return nil
}

// DequeueAsync removes an item from the queue and processes it using the provided handler
// function in a goroutine. If configured, rate limits and concurrent worker limits will be enforced.
// Function discard checks if the item can be discarded.
// If it is discarded dequeue does not affect the rate limit.
// This function will block if an item is not immediately available for
// processing or if necessary to enforce limits.
func (q *PriorityQueue[T]) DequeueAsync(ctx context.Context, handler func(context.Context, T) error, discard func(context.Context, T) bool) error {
	item, err := q.dequeue(ctx)
	if err != nil {
		return err
	}

	// Avoid panic if the handler is nil - could be used to pop an item without processing.
	if handler == nil {
		return nil
	}

	if discard != nil && discard(ctx, item.value) {
		return nil
	}

	// apply rate limit
	err = q.wait(ctx)
	if err != nil {
		return err
	}

	if q.workersSem != nil {
		if err := q.incrementWorkers(ctx); err != nil {
			return err
		}
	}

	go func() {
		if q.workersSem != nil {
			defer q.decrementWorkers()
		}
		err = handler(ctx, item.value)

		// If there was any error, we re-queue the item for processing again.
		if err != nil {
			q.handleError(ctx, item)
		}
	}()

	return nil
}

func (q *PriorityQueue[T]) handleError(ctx context.Context, item priorityQueueItem[T]) {
	waitDuration := item.backoff.NextBackOff()

	if waitDuration == backoff.Stop {
		// Attempt to send the item to the dead letter queue, but do not block if it is full -
		// in that case the item will be discarded.
		select {
		case q.DeadLetterQueue <- item.value:
			logger.Debugf("max retry attempts reached in queue %s, sent item to dead letter queue: %v", q.Name(), item.value)
		default:
		}
		return
	}

	go func() {
		if waitDuration > 0 {
			logger.Debugf("queue %s sleeping for %v before retrying", q.Name(), waitDuration)

			select {
			case <-time.After(waitDuration):

			case <-ctx.Done():
				logger.Errorf("context cancelled while waiting to retry item %v in queue %s", item.value, q.Name())
				return
			}
		}

		if item.priority {
			err := q.enqueuePriority(ctx, item)
			if err != nil {
				logger.Errorf("error enqueing to %s priority item %v for retry: %v", q.Name(), item.value, err)

			}
		} else if err := q.enqueue(ctx, item); err != nil {
			logger.Errorf("error enqueing to %s item %v for retry: %v", q.Name(), item.value, err)
		}
	}()

}

func (q *PriorityQueue[T]) dequeueWithRateLimit(ctx context.Context) (result priorityQueueItem[T], err error) {
	err = q.wait(ctx)
	if err != nil {
		return result, err
	}

	// Set the err variable so that the deferred function can read it.
	result, err = q.dequeue(ctx)
	return result, err
}

func (q *PriorityQueue[T]) dequeue(ctx context.Context) (priorityQueueItem[T], error) {
	var result priorityQueueItem[T]

	select {
	case result = <-q.priority:
		return result, nil

	default:
		select {
		case result = <-q.priority:
			return result, nil

		case result = <-q.regular:
			return result, nil

		case <-ctx.Done():
			return result, ctx.Err()
		}
	}
}

// This operation may block until a worker slot is available.
func (q *PriorityQueue[T]) incrementWorkers(ctx context.Context) error {
	logger.Debugf("incrementing workers")

	select {
	case q.workersSem <- struct{}{}:
		return nil

	case <-ctx.Done():
		return ctx.Err()

	default:
		logger.Debug("enforcing workers limit")

		select {
		case q.workersSem <- struct{}{}:
			return nil

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// This operation should never block - if it does that indicates that decrement
// has been called too many times.
func (q *PriorityQueue[T]) decrementWorkers() {
	logger.Debugf("decrementing workers")

	select {
	case <-q.workersSem:
		return

	default:
		logger.Panic("should never block")
	}
}

func (q *PriorityQueue[T]) wait(ctx context.Context) error {
	return q.limiter.Wait(ctx)
}

// Length return the number of item in the queue
func (q *PriorityQueue[T]) Length() int {
	return len(q.priority) + len(q.regular)
}

// Name return the nam of the queue or "unnamed"
func (q *PriorityQueue[T]) Name() string {
	if len(q.name) > 0 {
		return q.name
	}
	return "unnamed"
}
