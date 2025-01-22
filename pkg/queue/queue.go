package queue

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const defaultMaxAttempts uint64 = 10

const NotRatedDequeue string = "!!NotRatedDequeue!!"

type lastDequeueMutex struct {
	time time.Time

	// mutex
	sync.RWMutex
}

// PriorityQueue is made up of two sub-queues - one regular and one with
// higher priority. Items can be enqueued in either queue and when dequeueing
// items from the priority queue are returned first.
type PriorityQueue[T any] struct {
	Name            string
	regular         chan priorityQueueItem[T]
	priority        chan priorityQueueItem[T]
	minDequeueDelta time.Duration
	lastDequeue     *lastDequeueMutex
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
	Size                 int           `toml:"size"`
	MaxDequeuesPerSecond int           `toml:"max_dequeues_per_second"` // Set to 0 to disable rate-limiting.
	MaxWorkers           int           `toml:"max_workers"`             // Set to 0 for unlimited workers.
	MaxAttempts          int32         `toml:"max_attempts"`            // Set to negative for unlimited retry attempts. If unset or set to 0, the default value (10) is applied.
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

	lastDequeue := new(lastDequeueMutex)
	lastDequeue.time = time.Now()

	q := PriorityQueue[T]{
		regular:     make(chan priorityQueueItem[T], input.Size),
		priority:    make(chan priorityQueueItem[T], input.Size),
		lastDequeue: lastDequeue,
		backoff:     input.Backoff,
		timeOff:     input.TimeOff,
	}

	if input.MaxDequeuesPerSecond > 0 {
		q.minDequeueDelta = time.Second / time.Duration(input.MaxDequeuesPerSecond)
		logger.Info("minDequeueDelta:", q.minDequeueDelta)
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
	var lastDequeue time.Time
	if q.minDequeueDelta > 0 {
		q.lastDequeue.RLock()
		lastDequeue = q.lastDequeue.time
		q.lastDequeue.RUnlock()
	}
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

	// do not retry and do not affect rate limit
	if notRated(err) {
		if q.minDequeueDelta > 0 {
			q.lastDequeue.Lock()
			q.lastDequeue.time = lastDequeue
			q.lastDequeue.Unlock()

		}
		return nil
	}

	// If there was any error, we re-queue the item for processing again.
	if err != nil {
		q.handleError(ctx, item)
		return err
	}

	return nil
}

// Dequeue removes an item from the queue and processes it using the provided handler
// function in a goroutine. If configured, rate limits and concurrent worker limits will be enforced.
// This function will block if an item is not immediately available for
// processing or if necessary to enforce limits.
func (q *PriorityQueue[T]) DequeueAsync(ctx context.Context, handler func(context.Context, T) error) error {
	var lastDequeue time.Time
	if q.minDequeueDelta > 0 {
		q.lastDequeue.RLock()
		lastDequeue = q.lastDequeue.time
		q.lastDequeue.RUnlock()
	}
	item, err := q.dequeueWithRateLimit(ctx)
	if err != nil {
		return err
	}

	// Avoid panic if the handler is nil - could be used to pop an item without processing.
	if handler == nil {
		return nil
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

		// do not retry and do not affect rate limit
		if notRated(err) {
			if q.minDequeueDelta > 0 {
				q.lastDequeue.Lock()
				q.lastDequeue.time = lastDequeue
				q.lastDequeue.Unlock()
			}
			return
		}

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
			logger.Debugf("max retry attempts reached in queue %s, sent item to dead letter queue: %v", q.Name, item.value)

		default:
			logger.Debugf("DeadLetterQueue in queue %s full, discarding item: %v", item.value)

		}
		return
	}

	go func() {
		if waitDuration > 0 {
			logger.Debugf("queue %s sleeping for %v before retrying", q.Name, waitDuration)

			select {
			case <-time.After(waitDuration):

			case <-ctx.Done():
				logger.Errorf("context cancelled while waiting to retry item %v in queue %s", item.value, q.Name)
				return
			}
		}

		if item.priority {
			err := q.enqueuePriority(ctx, item)
			if err != nil {
				logger.Errorf("error enqueing to %s priority item %v for retry: %v", q.Name, item.value, err)

			}
		} else if err := q.enqueue(ctx, item); err != nil {
			logger.Errorf("error enqueing to %s item %v for retry: %v", q.Name, item.value, err)
		}
	}()

}

func (q *PriorityQueue[T]) dequeueWithRateLimit(ctx context.Context) (result priorityQueueItem[T], err error) {
	if q.minDequeueDelta > 0 {
		if err = q.enforceRateLimit(ctx); err != nil {
			return result, err
		}

		defer func() {
			if err == nil {
				q.lastDequeue.Lock()
				q.lastDequeue.time = time.Now()
				q.lastDequeue.Unlock()

			}
		}()
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

func (q *PriorityQueue[T]) enforceRateLimit(ctx context.Context) error {
	now := time.Now()
	q.lastDequeue.RLock()
	delta := now.Sub(q.lastDequeue.time)
	q.lastDequeue.RUnlock()

	if delta >= q.minDequeueDelta {
		return nil
	}

	sleepDuration := q.minDequeueDelta - delta
	logger.Debugf("enforcing rate limit - sleeping for %s", sleepDuration)

	select {
	case <-time.After(sleepDuration):
		return nil

	case <-ctx.Done():
		return ctx.Err()
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

// notRated checks whether error returned by handle should be included in the rate limit
func notRated(err error) bool {
	if err != nil {
		return strings.Contains(err.Error(), NotRatedDequeue)
	}

	return false
}

// Length return the number of item in the queue
func (q *PriorityQueue[T]) Length() int {
	return len(q.priority) + len(q.regular)
}
