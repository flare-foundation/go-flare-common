package priority

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFull(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

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

	for i := 0; i < 100; i++ {
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

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

	for i := 0; i < 100; i++ {
		wg.Add(1)
		pQueue.Add(i, wInt(i))
	}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		pQueue.AddFast(i, wInt(i))
	}

	go func() {
		for {
			pQueue.Dequeue(ctx, handle, nil)
		}
	}()

	wg.Wait()

	// require
	var deviationTotal time.Duration = 0
	for j := 1; j < len(times.list)-1; j++ {
		difference := times.list[j+1].Sub(times.list[j]) - (time.Second / time.Duration(perSecond))
		deviationTotal = +difference.Abs()
	}

	deviationMean := deviationTotal / time.Duration(len(times.list)-2)
	require.Less(t, deviationMean, time.Second/time.Duration(perSecond*10))

	cancel()
}

func TestMaxAttempts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	params := Params{
		MaxAttempts: 2,
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

	for i := 0; i < 3; i++ {
		pQueue.Add(i, wInt(i))
	}

	time.Sleep(50 * time.Millisecond)
	require.Equal(t, 2, stats.attempts[0])

	cancel()
}

func TestMaxWorkers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

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

	for i := 0; i < 3; i++ {
		pQueue.Add(i, wInt(3-i))
	}

	time.Sleep(20 * time.Millisecond)

	stats.Lock()
	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 0, stats.attempts[2])
	stats.Unlock()

	time.Sleep(30 * time.Millisecond)

	require.Equal(t, 1, stats.attempts[0])
	require.Equal(t, 1, stats.attempts[1])
	require.Equal(t, 1, stats.attempts[2])

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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

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

	for i := 0; i < 100; i++ {
		wg.Add(1)
		pQueue.Add(i, wTup{-i, i})
	}

	time.Sleep(100 * time.Millisecond)

	for j := 0; j < 100; j++ {
		time.Sleep(time.Millisecond)
		pQueue.Dequeue(ctx, handle, nil)
	}

	wg.Wait()
	require.Len(t, handled.list, 100)

	fmt.Printf("handled.list: %v\n", handled.list)
	cancel()
}
