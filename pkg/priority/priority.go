// Package priority provides a heap-based priority queue with rate limiting and retry support.
package priority

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"golang.org/x/time/rate"
)

type backOff func(int) time.Duration // custom function defining timeOut after attempt

const bucketSize = 2

// Logger is the subset of logger.Logger used by PriorityQueue.
// logger.Nop satisfies it directly (Panic and Panicf propagate as panics).
type Logger interface {
	Infof(string, ...any)
	Errorf(string, ...any)
	Panic(...any)
	Panicf(string, ...any)
}

// Params for the priority queue.
//
// WARNING: there is no size cap on the underlying heap. If producers
// call Add/AddFast faster than Dequeue drains the queue, the heap slice
// grows without bound and the process can OOM. Callers that ingest from
// untrusted or burst sources are responsible for rate-limiting at the
// producer side, or for using [pkg/queue] which exposes a Size knob.
type Params struct {
	MaxDequeuesPerSecond int           `toml:"max_dequeues_per_second"` // Set to 0 to disable rate-limiting.
	MaxWorkers           int           `toml:"max_workers"`             // Set to 0 for unlimited workers.
	MaxAttempts          int           `toml:"max_attempts"`            // Maximal number of retries attempted after the first try. If not positive, no retries are made.
	TimeOff              time.Duration `toml:"time_off"`                // TimeOff between attempts
	ErrorChan            bool          `toml:"error_chan"`              // If true, errors on final attempts are pushed to the channel

}

// Wrapped is a wrapper around the item that is stored in the queue.
// It contains the item itself, the number of attempts left to process it, and a flag indicating if it's a fast lane item.
type Wrapped[T any] struct {
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
	regular QueueMutex[Wrapped[T], W]
	fast    QueueMutex[Wrapped[T], W]
	in      chan *Item[Wrapped[T], W]
	inFast  chan *Item[Wrapped[T], W]
	workers chan bool
	Errors  chan error

	maxAttempts int
	timeOff     time.Duration
	// bo is read by the retry goroutine in handleRetry and may be written
	// by SetBackOff at any time, including after InitiateAndRun. Hold via
	// atomic.Pointer so the read/write pair is race-free.
	bo atomic.Pointer[backOff]

	limiter *rate.Limiter

	logger Logger
}

// New returns new Priority from params.
func New[T any, W weight[W]](params Params, name string) *PriorityQueue[T, W] {
	return NewWithLogger[T, W](params, name, logger.Nop{})
}

// NewWithLogger returns new Priority from params with logger.
func NewWithLogger[T any, W weight[W]](params Params, name string, l Logger) *PriorityQueue[T, W] {
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

	// The empty signal channels are capacity 1 (not unbuffered) so a
	// producer's non-blocking signal after a push always lands in the
	// channel buffer — surviving until the consumer arms its wait. With
	// an unbuffered channel, a signal raced against an in-transition
	// consumer (between draining the channel and entering the final
	// select) fell through `default` and was lost, stranding the item
	// in the heap.
	return &PriorityQueue[T, W]{
		name:    name,
		regular: QueueMutex[Wrapped[T], W]{empty: make(chan bool, 1)},
		fast:    QueueMutex[Wrapped[T], W]{empty: make(chan bool, 1)},
		in:      make(chan *Item[Wrapped[T], W]),
		inFast:  make(chan *Item[Wrapped[T], W]),
		workers: workers,
		Errors:  errors,

		maxAttempts: params.MaxAttempts,
		timeOff:     params.TimeOff,

		limiter: limiter,
		logger:  l,
	}
}

// SetBackOff sets the backOff function for the queue.
func (p *PriorityQueue[T, W]) SetBackOff(bo backOff) {
	p.bo.Store(&bo)
}

// InitiateAndRun starts the goroutines that move incoming items into the
// regular and fast lanes. The underlying channels are created by New/NewWithLogger,
// so Add/AddFast may be called before InitiateAndRun — the send will block until
// the processing goroutines start. Call InitiateAndRun at most once per queue.
func (p *PriorityQueue[T, W]) InitiateAndRun(ctx context.Context) {
	go p.processIn(ctx)
	go p.processInFast(ctx)
}

// processIn enqueues items from in channel to the regular lane.
func (p *PriorityQueue[T, W]) processIn(ctx context.Context) {
	for {
		select {
		case item := <-p.in:
			p.regular.Lock()
			heapt.Push(&p.regular.Queue, item)
			select {
			case p.regular.empty <- false:
			default:
			}
			p.regular.Unlock()
		case <-ctx.Done():
			p.logger.Infof("processIn exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

// processInFast enqueues items from inFast channel to the fast lane.
func (p *PriorityQueue[T, W]) processInFast(ctx context.Context) {
	for {
		select {
		case item := <-p.inFast:
			p.fast.Lock()
			heapt.Push(&p.fast.Queue, item)
			select {
			case p.fast.empty <- false:
			default:
			}
			p.fast.Unlock()
		case <-ctx.Done():
			p.logger.Infof("processInFast exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

// AddFast adds value with weight to inFast channel.
//
// Returns ctx.Err() if ctx is cancelled before the send completes. After
// the queue's InitiateAndRun ctx is cancelled the producer goroutines have
// already returned, so a send on the unbuffered inFast channel would block
// forever; callers must pass a ctx whose cancellation matches the queue's
// lifetime to avoid a goroutine leak.
func (p *PriorityQueue[T, W]) AddFast(ctx context.Context, value T, weight W) (*Item[Wrapped[T], W], error) {
	item := &Item[Wrapped[T], W]{
		value: Wrapped[T]{
			item:         value,
			attemptsLeft: 0,
			fast:         true,
		},
		weight: weight,
		index:  0,
	}

	select {
	case p.inFast <- item:
		return item, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Add adds value with weight to in channel.
//
// Returns ctx.Err() if ctx is cancelled before the send completes; see
// AddFast for the goroutine-leak hazard this prevents.
func (p *PriorityQueue[T, W]) Add(ctx context.Context, value T, weight W) (*Item[Wrapped[T], W], error) {
	item := &Item[Wrapped[T], W]{
		value: Wrapped[T]{
			item:         value,
			attemptsLeft: p.maxAttempts,
			fast:         false,
		},
		weight: weight,
	}

	select {
	case p.in <- item:
		return item, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// next returns the item that is next in line, or (nil, false) when ctx is
// cancelled while both lanes are empty.
//
// The two empty channels are capacity 1 (see New); after draining stale
// signals under the lane lock and releasing it, any concurrent push lands
// its signal in the buffer where the subsequent select picks it up — so
// no signal can be lost. Wakeups loop back to re-check under lock because
// the signal may be stale (e.g., the producer pushed an item that we
// already drained, or two pushes raced into a single buffered slot).
func (p *PriorityQueue[T, W]) next(ctx context.Context) (*Item[Wrapped[T], W], bool) {
	for {
		p.fast.Lock()
		if p.fast.Len() > 0 {
			item, _ := heapt.Pop(&p.fast)
			p.fast.Unlock()
			return item, true
		}
		// Drain any stale signal so the upcoming wait only fires on a
		// genuine push that happens AFTER this drain.
		select {
		case <-p.fast.empty:
		default:
		}
		p.fast.Unlock()

		p.regular.Lock()
		if p.regular.Len() > 0 {
			item, _ := heapt.Pop(&p.regular)
			p.regular.Unlock()
			return item, true
		}
		select {
		case <-p.regular.empty:
		default:
		}
		p.regular.Unlock()

		select {
		case <-p.fast.empty:
		case <-p.regular.empty:
		case <-ctx.Done():
			return nil, false
		}
		// Wakeup: loop back to re-check under lock. The signal may be
		// real (loop pops the item) or stale (loop arms a fresh wait).
	}
}

// Dequeue gets next item and processes it with discard and handler function.
// Items that are discarded do not affect rate limit.
// If handler returns an error, item (from regular lane) is retried until success or maxAttempts is reached.
//
// Delivery is at-most-once. Once an item is popped from the heap, it is no
// longer in the queue; if ctx is cancelled before the handler runs (between
// pop and limiter.Wait / incrementWorkers / dispatch), the item is logged
// and dropped. Callers that need at-least-once semantics across shutdown
// must drain the queue before cancelling the ctx — for example by signalling
// producers to stop, waiting for Length to reach zero, and only then
// cancelling.
func (p *PriorityQueue[T, W]) Dequeue(ctx context.Context, handler func(context.Context, T) error, discard func(context.Context, T) bool) {
	wItem, ok := p.next(ctx)
	if !ok {
		return
	}

	if handler == nil {
		return
	}

	if discard != nil && discard(ctx, wItem.value.item) {
		return
	}

	err := p.limiter.Wait(ctx)
	if err != nil {
		p.logger.Errorf("queue %s wait error %v", p.Name(), err)
		return
	}

	if p.workers != nil {
		if err := p.incrementWorkers(ctx); err != nil {
			return
		}
	}

	go func() {
		err := handler(ctx, wItem.value.item)

		if p.workers != nil {
			p.decrementWorkers()
		}

		if err != nil {
			p.handleRetry(ctx, wItem, err)
		}
	}()
}

// Name of the PriorityQueue. Returns "unnamed" if not set.
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
		p.logger.Panicf("queue %s: decrementWorkers called with no worker in flight", p.Name())
	}
}

// handleRetry re-enqueues an item from the regular lane after timeOff if maxAttempts has not been reached.
func (p *PriorityQueue[T, W]) handleRetry(ctx context.Context, item *Item[Wrapped[T], W], err error) {
	item.value.attemptsLeft--
	if item.value.attemptsLeft <= 0 || item.value.fast {
		p.addError(err)
		return
	}

	go func() {
		to := p.backOff(p.maxAttempts - item.value.attemptsLeft)
		select {
		case <-time.After(to):
		case <-ctx.Done():
			return
		}
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

// backOff returns the backoff duration for the j-th retry attempt.
func (p *PriorityQueue[T, W]) backOff(j int) time.Duration {
	if v := p.bo.Load(); v != nil {
		return (*v)(j)
	}
	return p.timeOff
}
