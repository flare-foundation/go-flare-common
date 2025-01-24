package priority

import (
	"context"

	"github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"golang.org/x/time/rate"
)

type Priority[T any] struct {
	name    string
	regular QueueMutex[T]
	fast    QueueMutex[T]
	in      chan *Item[T]
	inFast  chan *Item[T]
	limiter *rate.Limiter
}

func (p *Priority[T]) InitiateAndRun(ctx context.Context) {
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

func (p *Priority[T]) processIn(ctx context.Context) {
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
			logger.Infof("processIn exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

func (p *Priority[T]) processInFast(ctx context.Context) {
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
			logger.Infof("processInFast exiting in queue %s, %v", p.Name(), ctx.Err())
			return
		}
	}
}

func (p *Priority[T]) getNext() T {
	p.fast.Lock()
	if p.fast.Len() > 0 {
		item, _ := heapt.Pop(&p.fast)
		p.fast.Unlock()
		return item.value
	} else {
		// vacant the channel wait for the signal
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
		return item.value
	} else {
		// vacant the channel wait for the signal
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
		return item.value
	case <-p.regular.empty:
		p.regular.Lock()
		item, _ := heapt.Pop(&p.regular)
		p.regular.Unlock()
		return item.value
	}
}

func (p *Priority[T]) Name() string {
	if len(p.name) > 0 {
		return p.name
	}
	return "unnamed"
}
