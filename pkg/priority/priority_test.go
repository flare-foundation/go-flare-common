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

	params := Params{
		MaxAttempts: 1,
		MaxWorkers:  2,
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
		stats.attempts[item]++
		stats.Unlock()
		time.Sleep(30 * time.Millisecond)

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

	time.Sleep(20 * time.Millisecond)

	stats.Lock()
	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 0, stats.attempts[2])
	stats.Unlock()

	time.Sleep(30 * time.Millisecond)

	stats.Lock()
	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 1, stats.attempts[2])
	stats.Unlock()

	cancel()
}

type wTup [2]int

func (x wTup) Self() wTup {
	return x
}

func (x wTup) Less(y wTup) bool {
	return x[0] < y[0] || (x[0] == y[0] && x[1] < y[1])
}

// TestDequeueReturnsOnCtxCancel covers audit Low 11c: next() must honor ctx
// when both lanes are empty, instead of blocking forever.
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

// TestAddBeforeInitiateAndRun covers audit finding M28: Add must not race
// against (nor hang forever on) a not-yet-initialized channel. After the fix
// the channels exist from construction, so Add blocks only until the
// processing goroutines start, then proceeds normally.
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

// TestSetBackOffConcurrent covers audit finding F-PRI-3: SetBackOff used to
// write p.bo without synchronization while the retry goroutine in
// handleRetry read it. With go test -race this must not flag a data race.
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

// TestAddReturnsAfterCtxCancel covers audit finding F-PRI-2: Add and
// AddFast used to send unconditionally on the internal unbuffered
// channels. Once InitiateAndRun's ctx was cancelled, the processIn /
// processInFast goroutines had returned and a follow-up Add blocked
// forever — a permanent goroutine leak. With the new signature, the
// caller's ctx aborts the send.
func TestAddReturnsAfterCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())

	pQueue := New[int, wInt](Params{}, "test")
	pQueue.InitiateAndRun(ctx)

	cancel()
	// Settle so processIn/processInFast observe ctx.Done() and exit;
	// select is non-deterministic when both p.in and ctx.Done() are ready,
	// so without a brief pause the producer goroutines may still consume
	// one Add even after cancel().
	time.Sleep(50 * time.Millisecond)

	// Fresh ctx with a short deadline so the test fails fast if Add
	// blocks: the producer goroutines are gone, so the only way Add can
	// return is via the new ctx.Done() select branch.
	addCtx, addCancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer addCancel()

	_, err := pQueue.Add(addCtx, 1, wInt(1))
	require.Error(t, err, "Add must return after the queue's ctx was cancelled")

	_, err = pQueue.AddFast(addCtx, 1, wInt(1))
	require.Error(t, err, "AddFast must return after the queue's ctx was cancelled")
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

	time.Sleep(100 * time.Millisecond)

	for range 100 {
		time.Sleep(time.Millisecond)
		pQueue.Dequeue(ctx, handle, nil)
	}

	wg.Wait()
	require.Len(t, handled.list, 100)

	cancel()
}
