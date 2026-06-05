package priority

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFull(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	params := Params{}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	var wg sync.WaitGroup

	handled := struct {
		list []int
		sync.Mutex
	}{}
	handle := func(ctx context.Context, item int) error {
		handled.Lock()
		defer handled.Unlock()
		handled.list = append(handled.list, item)
		wg.Done()
		return nil
	}

	for i := range 100 {
		wg.Add(1)
		_, _ = pQueue.Add(ctx, i, wInt(i))
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	wg.Wait()
	require.Len(t, handled.list, 100)
	cancel()
}

func TestDequeue(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	perSecond := 100

	params := Params{
		MaxDequeuesPerSecond: perSecond,
		MaxAttempts:          0,
	}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	var wg sync.WaitGroup

	times := struct {
		list []time.Time
		sync.Mutex
	}{}
	handle := func(ctx context.Context, item int) error {
		times.Lock()
		defer times.Unlock()
		times.list = append(times.list, time.Now())
		wg.Done()
		return nil
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	time.Sleep(10 * time.Microsecond)

	for i := range 100 {
		wg.Add(1)
		_, _ = pQueue.Add(ctx, i, wInt(i))
	}
	for i := range 20 {
		wg.Add(1)
		_, _ = pQueue.AddFast(ctx, i, wInt(i))
	}

	wg.Wait()

	var deviationTotal time.Duration
	for j := 1; j < len(times.list)-1; j++ {
		difference := times.list[j+1].Sub(times.list[j]) - (time.Second / time.Duration(perSecond))
		deviationTotal += difference.Abs()
	}

	deviationMean := deviationTotal / time.Duration(len(times.list)-2)
	require.Less(t, deviationMean/2, time.Second/time.Duration(perSecond*5))

	cancel()
}

func TestDequeue2(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	perSecond := 100

	params := Params{
		MaxDequeuesPerSecond: perSecond,
		MaxAttempts:          0,
	}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	var wg sync.WaitGroup

	times := struct {
		list []time.Time
		sync.Mutex
	}{}
	handle := func(ctx context.Context, item int) error {
		times.Lock()
		defer times.Unlock()
		times.list = append(times.list, time.Now())
		wg.Done()
		return nil
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	time.Sleep(10 * time.Microsecond)

	// Sample 100, not 20: too few samples let one CI hiccup skew the mean deviation.
	for i := range 100 {
		wg.Add(1)
		_, _ = pQueue.AddFast(ctx, i, wInt(i))
	}

	wg.Wait()

	var deviationTotal time.Duration
	for j := 1; j < len(times.list)-1; j++ {
		difference := times.list[j+1].Sub(times.list[j]) - (time.Second / time.Duration(perSecond))
		deviationTotal += difference.Abs()
	}

	deviationMean := deviationTotal / time.Duration(len(times.list)-2)
	require.Less(t, deviationMean/2, time.Second/time.Duration(perSecond))

	cancel()
}

func TestDequeueDiscard(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	perSecond := 1

	params := Params{
		MaxDequeuesPerSecond: perSecond,
		MaxAttempts:          0,
	}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	var wg sync.WaitGroup
	handled := make(map[int]bool)

	handle := func(ctx context.Context, item int) error {
		handled[item] = true
		wg.Done()
		return nil
	}

	for i := range 100 {
		wg.Add(1)
		_, _ = pQueue.Add(ctx, i, wInt(i))
	}
	for i := range 20 {
		wg.Add(1)
		_, _ = pQueue.AddFast(ctx, i, wInt(i))
	}

	discarded := make(map[int]bool)
	discardFunc := func(ctx context.Context, i int) bool {
		if i == 99 {
			return false
		}
		discarded[i] = true
		wg.Done()
		return true
	}

	start := time.Now()
	go func() {
		for {
			pQueue.Dequeue(ctx, handle, discardFunc)
		}
	}()

	wg.Wait()
	duration := time.Since(start)

	require.Len(t, discarded, 99)
	require.Len(t, handled, 1)

	// no limits were imposed
	require.Less(t, duration, time.Millisecond)

	cancel()
}

func TestMaxAttempts(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	const maxAttempts = 2
	const numOfItems = 3

	params := Params{
		MaxAttempts: maxAttempts,
	}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	stats := struct {
		attempts map[int]int
		sync.Mutex
	}{
		attempts: make(map[int]int),
	}
	handle := func(ctx context.Context, item int) error {
		stats.Lock()
		defer stats.Unlock()
		stats.attempts[item]++
		return errors.New("test fail")
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	for i := range numOfItems {
		_, _ = pQueue.Add(ctx, i, wInt(i))
	}

	require.Eventually(t, func() bool {
		stats.Lock()
		defer stats.Unlock()
		return len(stats.attempts) == numOfItems && stats.attempts[0] == maxAttempts && stats.attempts[numOfItems-1] == maxAttempts
	}, 1*time.Second, 10*time.Millisecond)

	cancel()
}

func TestMaxWorkers(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)
	defer cancel()

	params := Params{
		MaxAttempts: 1,
		MaxWorkers:  2,
	}

	pQueue := New[int, wInt](params, "test")
	pQueue.InitiateAndRun(ctx)

	var stats struct {
		attempts map[int]int
		sync.Mutex
	}
	stats.attempts = make(map[int]int)

	started := make(chan int, 3) // signals each item's entry into the handler
	release := make(chan struct{})

	handle := func(ctx context.Context, item int) error {
		stats.Lock()
		stats.attempts[item]++
		stats.Unlock()
		started <- item
		<-release // block until the test releases the worker slot
		return nil
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	for i := range 3 {
		_, _ = pQueue.Add(ctx, i, wInt(3-i))
	}

	// Two items must enter handle (MaxWorkers=2); the third must NOT until a slot frees.
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("expected two items to enter handler within 1s")
		}
	}
	select {
	case item := <-started:
		t.Fatalf("third item %d entered handler before a worker slot was freed", item)
	case <-time.After(50 * time.Millisecond):
	}

	stats.Lock()
	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 0, stats.attempts[2])
	stats.Unlock()

	// Release the two in-flight workers; the third item must then run.
	close(release)
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("expected third item to enter handler after slot freed")
	}

	stats.Lock()
	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 1, stats.attempts[2])
	stats.Unlock()
}

type wTup [2]int

func (x wTup) Self() wTup {
	return x
}

func (x wTup) Less(y wTup) bool {
	return x[0] < y[0] || (x[0] == y[0] && x[1] < y[1])
}

// TestDequeueReturnsOnCtxCancel verifies that next() honors ctx when both
// lanes are empty, instead of blocking forever.
func TestDequeueReturnsOnCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	pQueue := New[int, wInt](Params{}, "test")
	pQueue.InitiateAndRun(ctx)

	done := make(chan struct{})
	go func() {
		pQueue.Dequeue(ctx, func(context.Context, int) error { return nil }, nil)
		close(done)
	}()

	cancel()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Dequeue did not return after ctx cancel")
	}
}

// TestAddBeforeInitiateAndRun verifies that Add does not race against (nor
// hang forever on) a not-yet-initialized channel. The channels exist from
// construction, so Add blocks only until the processing goroutines start,
// then proceeds normally.
func TestAddBeforeInitiateAndRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	pQueue := New[int, wInt](Params{}, "test")

	added := make(chan struct{})
	go func() {
		_, _ = pQueue.Add(ctx, 42, wInt(42))
		close(added)
	}()

	pQueue.InitiateAndRun(ctx)

	handled := make(chan int, 1)
	handle := func(_ context.Context, item int) error {
		handled <- item
		return nil
	}
	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	select {
	case <-added:
	case <-ctx.Done():
		t.Fatal("Add did not unblock after InitiateAndRun")
	}

	select {
	case got := <-handled:
		require.Equal(t, 42, got)
	case <-ctx.Done():
		t.Fatal("item never processed")
	}
}

// TestSetBackOffConcurrent verifies that SetBackOff writing p.bo concurrently
// with the retry goroutine in handleRetry reading it does not flag a data
// race under go test -race.
func TestSetBackOffConcurrent(t *testing.T) {
	pQueue := New[int, wInt](Params{MaxAttempts: 5, TimeOff: time.Millisecond}, "test")

	done := make(chan struct{})
	go func() {
		defer close(done)
		for range 1000 {
			pQueue.SetBackOff(func(_ int) time.Duration { return time.Microsecond })
		}
	}()

	// Read concurrently by exercising the unexported backOff() method.
	for range 1000 {
		_ = pQueue.backOff(1)
	}
	<-done
}

// TestAddReturnsAfterCtxCancel verifies that the caller's ctx aborts the
// send. Add and AddFast must not send unconditionally on the internal
// unbuffered channels: once InitiateAndRun's ctx is cancelled, the
// processIn / processInFast goroutines have returned and a follow-up Add
// would otherwise block forever, leaking a goroutine.
func TestAddReturnsAfterCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())

	pQueue := New[int, wInt](Params{}, "test")
	pQueue.InitiateAndRun(ctx)

	cancel()

	// Once the producer goroutines exit, no one receives on p.in/p.inFast, so
	// Add/AddFast must surface their ctx error. Poll instead of sleeping a fixed
	// interval, which is unreliable under CI scheduling load.
	require.Eventually(t, func() bool {
		probeCtx, probeCancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer probeCancel()
		_, err := pQueue.Add(probeCtx, 1, wInt(1))
		return err != nil
	}, 5*time.Second, 20*time.Millisecond, "Add must return after the queue's ctx was cancelled")

	require.Eventually(t, func() bool {
		probeCtx, probeCancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer probeCancel()
		_, err := pQueue.AddFast(probeCtx, 1, wInt(1))
		return err != nil
	}, 5*time.Second, 20*time.Millisecond, "AddFast must return after the queue's ctx was cancelled")
}

func TestDoubleWeights(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)

	params := Params{}

	pQueue := New[int, wTup](params, "test")
	pQueue.InitiateAndRun(ctx)

	var wg sync.WaitGroup

	handled := struct {
		list []int
		sync.Mutex
	}{}
	handle := func(ctx context.Context, item int) error {
		handled.Lock()
		defer handled.Unlock()
		handled.list = append(handled.list, item)
		wg.Done()
		return nil
	}

	for i := range 100 {
		wg.Add(1)
		_, _ = pQueue.Add(ctx, i, wTup{-i, i})
	}

	for range 100 {
		pQueue.Dequeue(ctx, handle, nil)
	}

	wg.Wait()
	require.Len(t, handled.list, 100)

	cancel()
}
