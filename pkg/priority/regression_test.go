package priority

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// quietLogger records Panicf calls without panicking.
type quietLogger struct{ panics []string }

func (l *quietLogger) Infof(string, ...any)  {}
func (l *quietLogger) Errorf(string, ...any) {}
func (l *quietLogger) Panic(...any)          {}
func (l *quietLogger) Panicf(format string, args ...any) {
	l.panics = append(l.panics, fmt.Sprintf(format, args...))
}

// TestDecrementWorkersPanicsWithNonPanickingLogger verifies that the worker
// invariant panics even when the injected logger's Panicf does not.
func TestDecrementWorkersPanicsWithNonPanickingLogger(t *testing.T) {
	l := &quietLogger{}
	p := NewWithLogger[int, wInt](Params{MaxWorkers: 1}, "test", l)

	require.Panics(t, func() { p.decrementWorkers() })
	require.Len(t, l.panics, 1)
	require.Contains(t, l.panics[0], "no worker in flight")
}

// TestSetBackOffNilResets verifies that SetBackOff(nil) resets to the default
// TimeOff instead of planting a nil-function call in the retry path.
func TestSetBackOffNilResets(t *testing.T) {
	p := New[int, wInt](Params{MaxAttempts: 3, TimeOff: 5 * time.Millisecond}, "test")

	p.SetBackOff(nil)
	require.NotPanics(t, func() {
		require.Equal(t, 5*time.Millisecond, p.backOff(1))
	})

	p.SetBackOff(func(int) time.Duration { return time.Second })
	require.Equal(t, time.Second, p.backOff(1))

	p.SetBackOff(nil)
	require.Equal(t, 5*time.Millisecond, p.backOff(1))
}

// TestNextResignalsWaiters verifies that a consumer popping an item passes the
// wakeup on when more items remain, so a coalesced signal (two pushes, one
// buffered signal) cannot strand a second blocked consumer.
func TestNextResignalsWaiters(t *testing.T) {
	p := New[int, wInt](Params{}, "test")
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	var handled atomic.Int32
	h := func(context.Context, int) error { handled.Add(1); return nil }

	for range 2 {
		go p.Dequeue(ctx, h, nil)
	}
	// Let both consumers block in next before items arrive.
	require.Eventually(t, func() bool { return p.Length() == 0 }, time.Second, 5*time.Millisecond)
	time.Sleep(50 * time.Millisecond)

	// Two items but a single buffered signal — the coalesced state that racing
	// pushes can produce.
	p.fast.Lock()
	p.fast.AddValue(Wrapped[int]{item: 1, fast: true}, wInt(1))
	p.fast.AddValue(Wrapped[int]{item: 2, fast: true}, wInt(2))
	select {
	case p.fast.empty <- false:
	default:
	}
	p.fast.Unlock()

	require.Eventually(t, func() bool { return handled.Load() == 2 }, 2*time.Second, 10*time.Millisecond)
	require.Equal(t, 0, p.Length())
}

// TestLength verifies Length counts items across both lanes.
func TestLength(t *testing.T) {
	p := New[int, wInt](Params{}, "test")
	require.Equal(t, 0, p.Length())

	p.fast.Lock()
	p.fast.AddValue(Wrapped[int]{item: 1, fast: true}, wInt(1))
	p.fast.Unlock()

	p.regular.Lock()
	p.regular.AddValue(Wrapped[int]{item: 2}, wInt(2))
	p.regular.AddValue(Wrapped[int]{item: 3}, wInt(3))
	p.regular.Unlock()

	require.Equal(t, 3, p.Length())
}

// TestDequeueNilHandlerKeepsItem verifies that a nil handler returns without
// consuming an item.
func TestDequeueNilHandlerKeepsItem(t *testing.T) {
	p := New[int, wInt](Params{}, "test")

	p.fast.Lock()
	p.fast.AddValue(Wrapped[int]{item: 1, fast: true}, wInt(1))
	p.fast.Unlock()

	p.Dequeue(t.Context(), nil, nil)
	require.Equal(t, 1, p.Length())
}
