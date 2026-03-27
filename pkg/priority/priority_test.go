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
		pQueue.Add(i, wInt(i))
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
		pQueue.Add(i, wInt(i))
	}
	for i := range 20 {
		wg.Add(1)
		pQueue.AddFast(i, wInt(i))
	}

	wg.Wait()

	var deviationTotal time.Duration
	for j := 1; j < len(times.list)-1; j++ {
		difference := times.list[j+1].Sub(times.list[j]) - (time.Second / time.Duration(perSecond))
		deviationTotal += difference.Abs()
	}

	deviationMean := deviationTotal / time.Duration(len(times.list)-2)
	require.Less(t, deviationMean, time.Second/time.Duration(perSecond*10))

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
		pQueue.AddFast(i, wInt(i))
	}

	wg.Wait()

	var deviationTotal time.Duration
	for j := 1; j < len(times.list)-1; j++ {
		difference := times.list[j+1].Sub(times.list[j]) - (time.Second / time.Duration(perSecond))
		deviationTotal += difference.Abs()
	}

	deviationMean := deviationTotal / time.Duration(len(times.list)-2)
	require.Less(t, deviationMean, time.Second/time.Duration(perSecond*10))

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
		pQueue.Add(i, wInt(i))
	}
	for i := range 20 {
		wg.Add(1)
		pQueue.AddFast(i, wInt(i))
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
		pQueue.Add(i, wInt(i))
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
		pQueue.Add(i, wInt(3-i))
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
		pQueue.Add(i, wTup{-i, i})
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
