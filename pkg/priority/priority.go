package priority

import (
	"context"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"golang.org/x/time/rate"
)

const bucketSize = 2

type Params struct {
	MaxDequeuesPerSecond int           `toml:"max_dequeues_per_second"` // Set to 0 to disable rate-limiting.
	MaxWorkers           int           `toml:"max_workers"`             // Set to 0 for unlimited workers.
	MaxAttempts          int           `toml:"max_attempts"`            // Set to negative for unlimited attempts. If unset or set to 0, the default value (10) is applied.
	TimeOff              time.Duration `toml:"time_off"`                // TimeOff between attempts
}

type wrapped[T any] struct {
	item         T
	attemptsLeft int
	fast         bool
}

// PriorityQueue is a priority queue with a regular and a fast lane.
// When dequeuing, items from the fast lane are returned first.
// The order of the queue is based on weight.
// A rate limit and/or maximal items handled at any time can be set.
type PriorityQueue[T any] struct {
	name    string
	regular QueueMutex[wrapped[T]]
	fast    QueueMutex[wrapped[T]]
	in      chan *Item[T]
	inFast  chan *Item[T]
	workers chan bool

	maxAttempts int
	timeOff     time.Duration

	limiter *rate.Limiter
}

// New returns new Priority from params.
func New[T any](params Params, name string) PriorityQueue[T] {
	var limiter *rate.Limiter

	if params.MaxDequeuesPerSecond > 0 {
		limiter = rate.NewLimiter(rate.Limit(params.MaxDequeuesPerSecond), bucketSize)
	} else {
		limiter = rate.NewLimiter(rate.Inf, 0)
	}

	var workers chan bool
	if params.MaxWorkers > 0 {
		workers = make(chan bool, params.MaxWorkers)
	}

	return PriorityQueue[T]{
		name:    name,
		regular: QueueMutex[wrapped[T]]{},
		fast:    QueueMutex[wrapped[T]]{},
		workers: workers,

		maxAttempts: params.MaxAttempts,
		timeOff:     params.TimeOff,

		limiter: limiter,
	}
}

// InitiateAndRun starts accepting new items to priority queue
func (p *PriorityQueue[T]) InitiateAndRun(ctx context.Context) {
	in := make(chan *Item[T])
	inFast := make(chan *Item[T])
	emptyR := make(chan bool)
	emptyF := make(chan bool)

	p.in = in
	p.inFast = inFast
	p.regular.empty = emptyR
	p.fast.empty = emptyF

	go p.processIn(ctx)
	go p.processInFast(ctx)
}

// processIn enqueues items from in channel to the regular lane.
func (p *PriorityQueue[T]) processIn(ctx context.Context) {
	for {
		select {
		case item := <-p.in:
			wItem := Item[wrapped[T]]{
				value: wrapped[T]{
					item:         item.value,
					attemptsLeft: p.maxAttempts,
					fast:         false,
				},
				weight: item.weight,
			}
			p.regular.Lock()
			heapt.Push(&p.regular.Queue, &wItem)
			select {
			case p.regular.empty <- false:
			default:
			}
			p.regular.Unlock()
		case <-ctx.Done():
			logger.Infof("processIn exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

// processIn enqueues items from inFast channel to the fast lane.
func (p *PriorityQueue[T]) processInFast(ctx context.Context) {
	for {
		select {
		case item := <-p.inFast:
			wItem := Item[wrapped[T]]{
				value: wrapped[T]{
					item:         item.value,
					attemptsLeft: p.maxAttempts,
					fast:         true,
				},
				weight: item.weight,
			}

			p.fast.Lock()
			heapt.Push(&p.fast.Queue, &wItem)
			select {
			case p.fast.empty <- false:
			default:
			}
			p.fast.Unlock()
		case <-ctx.Done():
			logger.Infof("processInFast exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

// AddFast adds value with weight to inFast channel.
func (p *PriorityQueue[T]) AddFast(value T, weight int) {
	p.inFast <- &Item[T]{
		value:  value,
		weight: weight,
	}
}

// Add adds value with weight to in channel.
func (p *PriorityQueue[T]) Add(value T, weight int) {
	p.in <- &Item[T]{
		value:  value,
		weight: weight,
	}
}

// next returns the item that is next in line.
func (p *PriorityQueue[T]) next() *Item[wrapped[T]] {
	p.fast.Lock()
	if p.fast.Len() > 0 {
		item, _ := heapt.Pop(&p.fast)
		p.fast.Unlock()
		return item
	} else {
		// vacate the channel wait for the signal
		select {
		case <-p.fast.empty:
		default:
		}
		p.fast.Unlock()
	}

	p.regular.Lock()
	if p.regular.Len() > 0 {
		item, _ := heapt.Pop(&p.regular)
		p.regular.Unlock()
		return item
	} else {
		// vacate the channel wait for the signal
		select {
		case <-p.regular.empty:
		default:
		}
		p.regular.Unlock()
	}

	select {
	case <-p.fast.empty:
		p.fast.Lock()
		item, _ := heapt.Pop(&p.fast)
		p.fast.Unlock()
		return item
	case <-p.regular.empty:
		p.regular.Lock()
		item, _ := heapt.Pop(&p.regular)
		p.regular.Unlock()
		return item
	}
}

// Dequeue gets next item and process it with discard and handler function.
// Items that are discarded are do not affect rate limit.
// If handler returns an error, item (from regular late) is retried until success of maxAttempts is reached.
func (p *PriorityQueue[T]) Dequeue(ctx context.Context, handler func(context.Context, T) error, discard func(context.Context, T) bool) {
	wItem := p.next()

	if handler == nil {
		return
	}

	if discard != nil && discard(ctx, wItem.value.item) {
		return
	}

	if p.workers != nil {
		if err := p.incrementWorkers(ctx); err != nil {
			return
		}
	}

	err := p.limiter.Wait(ctx)
	if err != nil {
		logger.Errorf("queue %s wait error %v", p.Name(), err)
		return
	}

	go func() {
		err = handler(ctx, wItem.value.item)

		if p.workers != nil {
			p.decrementWorkers()
		}

		if err != nil {
			p.handleRetry(wItem)
		}
	}()
}

// Name of the PriorityQueue. "Unnamed" if not set.
func (p *PriorityQueue[T]) Name() string {
	if len(p.name) > 0 {
		return p.name
	}
	return "unnamed"
}

// incrementWorkers indicates a worker is taken. Blocks if no worker is available.
func (p *PriorityQueue[T]) incrementWorkers(ctx context.Context) error {
	select {
	case p.workers <- true:
		return nil

	case <-ctx.Done():
		return ctx.Err()

	default:
		select {
		case p.workers <- true:
			return nil

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// decrementWorkers indicates a worker is not taken anymore.
func (p *PriorityQueue[T]) decrementWorkers() {
	select {
	case <-p.workers:
		return

	default:
		logger.Panic("should never block")
	}
}

// handleRetry re-enqueues an item from the regular lane after timeOff if maxAttempts has not been reached.
func (p *PriorityQueue[T]) handleRetry(item *Item[wrapped[T]]) {
	item.value.attemptsLeft--
	if item.value.attemptsLeft <= 0 || item.value.fast {
		return
	}

	go func() {
		time.Sleep(p.timeOff)
		p.regular.Lock()
		heapt.Push(&p.regular.Queue, item)
		select {
		case p.regular.empty <- false:
		default:
		}
		p.regular.Unlock()
	}()
}
