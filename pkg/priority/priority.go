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
	MaxAttempts          int           `toml:"max_attempts"`            // If not positive or unset, it defaults to attempt.
	TimeOff              time.Duration `toml:"time_off"`                // TimeOff between attempts
	ErrorChan            bool          `toml:"error_chan"`              // If true, errors, error on final attempt are pushed to the channel
}

type wrapped[T any] struct {
	item         T
	attemptsLeft int
	fast         bool
}

// PriorityQueue is a priority queue with a regular and a fast lane.
// When dequeuing, items from the fast lane are returned first.
// The ordering of the queue is based on weight. Items with higher weight are returned first.
// A rate limit and/or maximal items handled at any time can be set.
type PriorityQueue[T any, W weight[W]] struct {
	name    string
	regular QueueMutex[wrapped[T], W]
	fast    QueueMutex[wrapped[T], W]
	in      chan *Item[T, W]
	inFast  chan *Item[T, W]
	workers chan bool
	Errors  chan error

	maxAttempts int
	timeOff     time.Duration

	limiter *rate.Limiter
}

// New returns new Priority from params.
func New[T any, W weight[W]](params Params, name string) PriorityQueue[T, W] {
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

	var errors chan error
	if params.ErrorChan {
		errors = make(chan error, 10)
	}

	return PriorityQueue[T, W]{
		name:    name,
		regular: QueueMutex[wrapped[T], W]{},
		fast:    QueueMutex[wrapped[T], W]{},
		workers: workers,
		Errors:  errors,

		maxAttempts: params.MaxAttempts,
		timeOff:     params.TimeOff,

		limiter: limiter,
	}
}

// InitiateAndRun starts accepting new items to priority queue
func (p *PriorityQueue[T, W]) InitiateAndRun(ctx context.Context) {
	in := make(chan *Item[T, W])
	inFast := make(chan *Item[T, W])
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
func (p *PriorityQueue[T, W]) processIn(ctx context.Context) {
	for {
		select {
		case item := <-p.in:
			wItem := Item[wrapped[T], W]{
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

// processInFast enqueues items from inFast channel to the fast lane.
func (p *PriorityQueue[T, W]) processInFast(ctx context.Context) {
	for {
		select {
		case item := <-p.inFast:
			wItem := Item[wrapped[T], W]{
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
func (p *PriorityQueue[T, W]) AddFast(value T, weight W) {
	p.inFast <- &Item[T, W]{
		value:  value,
		weight: weight,
	}
}

// Add adds value with weight to in channel.
func (p *PriorityQueue[T, W]) Add(value T, weight W) {
	p.in <- &Item[T, W]{
		value:  value,
		weight: weight,
	}
}

// next returns the item that is next in line.
func (p *PriorityQueue[T, W]) next() *Item[wrapped[T], W] {
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

// Dequeue gets next item and processes it with discard and handler function.
// Items that are discarded do not affect rate limit.
// If handler returns an error, item (from regular late) is retried until success or maxAttempts is reached.
func (p *PriorityQueue[T, W]) Dequeue(ctx context.Context, handler func(context.Context, T) error, discard func(context.Context, T) bool) {
	wItem := p.next()

	if handler == nil {
		return
	}

	if discard != nil && discard(ctx, wItem.value.item) {
		return
	}

	err := p.limiter.Wait(ctx)
	if err != nil {
		logger.Errorf("queue %s wait error %v", p.Name(), err)
		return
	}

	if p.workers != nil {
		if err := p.incrementWorkers(ctx); err != nil {
			return
		}
	}

	go func() {
		err = handler(ctx, wItem.value.item)

		if p.workers != nil {
			p.decrementWorkers()
		}

		if err != nil {
			p.handleRetry(wItem, err)
		}
	}()
}

// Name of the PriorityQueue. "Unnamed" if not set.
func (p *PriorityQueue[T, W]) Name() string {
	if len(p.name) > 0 {
		return p.name
	}
	return "unnamed"
}

// incrementWorkers indicates a worker is taken. Blocks if no worker is available.
func (p *PriorityQueue[T, W]) incrementWorkers(ctx context.Context) error {
	select {
	case p.workers <- true:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

// decrementWorkers indicates a worker is not taken anymore.
func (p *PriorityQueue[T, W]) decrementWorkers() {
	select {
	case <-p.workers:
		return

	default:
		logger.Panic("should never block")
	}
}

// handleRetry re-enqueues an item from the regular lane after timeOff if maxAttempts has not been reached.
func (p *PriorityQueue[T, W]) handleRetry(item *Item[wrapped[T], W], err error) {
	item.value.attemptsLeft--
	if item.value.attemptsLeft <= 0 || item.value.fast {
		p.addError(err)
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

// addError adds error to the Errors channel if it is enabled.
// If the channel is full, the error is lost.
func (p *PriorityQueue[T, W]) addError(err error) {
	if p.Errors == nil {
		return
	}

	select {
	case p.Errors <- err:
	default:
	}
}
