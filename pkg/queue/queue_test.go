package queue_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/flare-foundation/go-flare-common/pkg/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	size          = 10
	benchmarkSize = 100000
	numWorkers    = 4
	maxAttempts   = 3

	defaultTimeout = 10 * time.Second
)

var (
	backoffTime = 10 * time.Millisecond
)

func TestEnqueueDequeue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	for i := 0; i < size; i++ {
		err := q.Enqueue(ctx, i)
		require.NoError(t, err)
	}

	for i := 0; i < size; i++ {
		err := q.Dequeue(ctx, itemCheckCallback(i))

		require.NoError(t, err)
	}
}

func TestEnqueueDequeueAsync(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	var wg sync.WaitGroup
	list := struct {
		sync.Mutex
		nums []int
	}{}
	handler := func(ctx context.Context, i int) error {
		list.Lock()
		list.nums = append(list.nums, i)
		list.Unlock()
		wg.Done()
		return nil
	}

	for i := 0; i < size; i++ {
		err := q.Enqueue(ctx, i)
		require.NoError(t, err)
	}

	err := q.EnqueuePriority(ctx, size+1)
	require.NoError(t, err)

	for i := 0; i < size+1; i++ {
		wg.Add(1)
		err := q.DequeueAsync(ctx, handler, nil)
		require.NoError(t, err)
	}

	wg.Wait()

	require.Equal(t, size+1, len(list.nums))
}

func TestEnqueuePriority(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	err := q.Enqueue(ctx, 1)
	require.NoError(t, err)

	err = q.EnqueuePriority(ctx, 42)
	require.NoError(t, err)

	err = q.Dequeue(ctx, itemCheckCallback(42))
	require.NoError(t, err)

	err = q.Dequeue(ctx, itemCheckCallback(1))
	require.NoError(t, err)
}

func TestBlockingDequeue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := q.Dequeue(ctx, itemCheckCallback(42))
		assert.NoError(t, err)
	}()

	err := q.Enqueue(ctx, 42)
	require.NoError(t, err)
	wg.Wait()
}

func TestBlockingDequeueAsync(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := q.DequeueAsync(ctx, itemCheckCallbackWG(42, &wg), nil)
		assert.NoError(t, err)
	}()

	err := q.Enqueue(ctx, 42)
	require.NoError(t, err)
	wg.Wait()
}

func TestBlockingDequeuePriority(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := q.Dequeue(ctx, itemCheckCallback(42))
		assert.NoError(t, err)
	}()

	err := q.EnqueuePriority(ctx, 42)
	require.NoError(t, err)
	wg.Wait()
}

func TestEnqueueTimeout(t *testing.T) {
	q := queue.NewPriority[int](nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	err := q.Enqueue(ctx, 1)
	require.Error(t, err)
	require.EqualError(t, err, ctx.Err().Error())
}

func TestEnqueuePriorityTimeout(t *testing.T) {
	q := queue.NewPriority[int](nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	err := q.EnqueuePriority(ctx, 1)
	require.Error(t, err)
	require.EqualError(t, err, ctx.Err().Error())
}

func TestDequeueTimeout(t *testing.T) {
	q := queue.NewPriority[int](nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	err := q.Dequeue(ctx, nil)
	require.Error(t, err)
	require.EqualError(t, err, ctx.Err().Error())
}

func TestDequeueAsyncTimeout(t *testing.T) {
	q := queue.NewPriority[int](nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	err := q.DequeueAsync(ctx, nil, nil)
	require.Error(t, err)
	require.EqualError(t, err, ctx.Err().Error())
}

func TestDequeueRateLimitTimeout(t *testing.T) {
	ctx := context.Background()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{
		Size:                 2,
		MaxDequeuesPerSecond: 1,
	})

	for i := 0; i < 2; i++ {
		err := q.Enqueue(ctx, i)
		require.NoError(t, err)
	}

	err := q.Dequeue(ctx, itemCheckCallback(0))
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(ctx, time.Nanosecond)
	defer cancel()

	// Since we are immediately attempting to dequeue the second item, this
	// should start to block for around a second to enforce the 1-per-second
	// rate limit. However, the context will cancel long before the rate limit
	// elapses.
	start := time.Now()

	err = q.Dequeue(ctx, nil)
	require.Error(t, err)
	require.EqualError(t, err, ctx.Err().Error())

	// Check that we wait much less than a full second before exiting.
	delta := time.Since(start)
	require.Less(t, delta, 100*time.Millisecond)
}

func TestWorkersLimit(t *testing.T) {
	ctx, cancel1 := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel1()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{
		Size:       size,
		MaxWorkers: numWorkers,
	})

	require.Less(t, numWorkers, size)
	for i := 0; i < numWorkers+1; i++ {
		err := q.Enqueue(ctx, i)
		require.NoError(t, err)
	}

	// Set up some blocking workers to fill the limit.
	var readyGroup sync.WaitGroup
	var finishedGroup sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		readyGroup.Add(1)
		finishedGroup.Add(1)
		go func(ctx context.Context, i int) {
			defer finishedGroup.Done()

			err := q.Dequeue(ctx, func(ctx context.Context, item int) error {
				readyGroup.Done()

				<-ctx.Done()
				return ctx.Err()
			})

			require.Error(t, err)
			require.EqualError(t, err, context.Canceled.Error())
		}(ctx, i)
	}

	// Need to wait until all worker handlers have been called.
	readyGroup.Wait()

	// Attempting to add another worker now should block.
	ctx, cancel2 := context.WithTimeout(ctx, time.Millisecond)
	defer cancel2()

	err := q.Dequeue(ctx, func(ctx context.Context, item int) error {
		t.Log("shouldn't reach here")
		t.Fail()

		return errors.New("shouldn't reach here")
	})
	require.Error(t, err)
	require.EqualError(t, err, context.DeadlineExceeded.Error())

	cancel1()
	finishedGroup.Wait()
}

func TestMaxAttempts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size, MaxAttempts: maxAttempts})

	dlq := q.DeadLetterQueue

	err := q.Enqueue(ctx, 1)
	require.NoError(t, err)

	handlerErr := errors.New("handler error")

	for i := 0; i < maxAttempts; i++ {
		err := q.Dequeue(ctx, func(ctx context.Context, item int) error {
			return handlerErr
		})
		require.ErrorIs(t, err, handlerErr)
	}

	select {
	case item := <-dlq:
		require.Equal(t, 1, item)

	case <-ctx.Done():
		t.Fatal("timed out while reading from dead letter queue")
	}
}

func TestMaxAttemptsAsync(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size, MaxAttempts: maxAttempts})

	dlq := q.DeadLetterQueue

	err := q.Enqueue(ctx, 1)
	require.NoError(t, err)

	handlerErr := errors.New("handler error")

	for i := 0; i < maxAttempts; i++ {
		err := q.DequeueAsync(ctx, func(ctx context.Context, item int) error {
			return handlerErr
		}, nil)
		require.NoError(t, err)
	}

	select {
	case item := <-dlq:
		require.Equal(t, 1, item)

	case <-ctx.Done():
		t.Fatal("timed out while reading from dead letter queue")
	}
}

func TestConstantBackOff(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size, Backoff: func() backoff.BackOff {
		return backoff.NewConstantBackOff(backoffTime)
	}})

	err := q.Enqueue(ctx, 1)
	require.NoError(t, err)

	handlerErr := errors.New("handler error")
	start := time.Now()

	for i := 0; i < maxAttempts; i++ {
		err := q.Dequeue(ctx, func(ctx context.Context, item int) error {
			return handlerErr
		})
		require.ErrorIs(t, err, handlerErr)
	}

	elapsed := time.Since(start)

	require.GreaterOrEqual(t, elapsed, (maxAttempts-1)*backoffTime)
}

func TestQueue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{
		Size:                 size,
		MaxDequeuesPerSecond: 0,
		MaxWorkers:           0,
		MaxAttempts:          maxAttempts,
		TimeOff:              20 * time.Millisecond,
	})

	nuOfRequest := 4

	for j := range nuOfRequest {
		err := q.Enqueue(ctx, j)
		require.NoError(t, err)
	}

	errs := map[int]bool{3: true}
	counter := make(map[int]int)
	activity := make(chan bool, 10)

	go func() {
		timer := time.NewTimer(100 * time.Millisecond)
		for {
			select {
			case <-activity:
				timer.Reset(100 * time.Millisecond)
			case <-timer.C:
				cancel()
				return
			}
		}
	}()

	run(ctx, q, testHandlerFactory(errs, counter, activity))
	require.Len(t, counter, nuOfRequest)

	for j := range counter {
		if errs[j] {
			require.Equal(t, counter[j], maxAttempts)
		} else {
			require.Equal(t, counter[j], 1)
		}
	}
}

func run[T any](ctx context.Context, q queue.PriorityQueue[T], handler func(context.Context, T) error) {
	for {
		if ctx.Err() != nil {
			return
		}

		err := q.Dequeue(ctx, handler)

		if err != nil {
			continue
		}
	}
}

func testHandlerFactory(errs map[int]bool, counter map[int]int, activity chan bool) func(context.Context, int) error {
	return func(ctx context.Context, j int) error {
		fmt.Printf("j: %v\n", j)
		activity <- true
		counter[j]++
		if errs[j] {
			return fmt.Errorf("%v returns error", j)
		}
		return nil
	}
}

func itemCheckCallback(expected int) func(context.Context, int) error {
	return func(ctx context.Context, item int) error {
		if item != expected {
			return fmt.Errorf("%d != %d", item, expected)
		}

		return nil
	}
}

func itemCheckCallbackWG(expected int, wg *sync.WaitGroup) func(context.Context, int) error {
	return func(ctx context.Context, item int) error {
		defer wg.Done()
		if item != expected {
			return fmt.Errorf("%d != %d", item, expected)
		}

		return nil
	}
}

func BenchmarkPriorityQueue(b *testing.B) {
	ctx := context.Background()

	q := queue.NewPriority[int](&queue.PriorityQueueParams{Size: size})

	for n := 0; n < b.N; n++ {
		_ = q.Enqueue(ctx, 1)
		_ = q.EnqueuePriority(ctx, 2)
		_ = q.Dequeue(ctx, nil)
		_ = q.Dequeue(ctx, nil)
	}
}
