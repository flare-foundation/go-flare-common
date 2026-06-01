package retry

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const errorMsg = "still error"

var errRetry = errors.New(errorMsg)

// testFunction returns a function that returns err until it is called k+1 times.
func testFunction(k int) func() (int, error) {
	return func() (int, error) {
		if k > 0 {
			k--
			return 0, errRetry
		}

		return 100, nil
	}
}

// testFunction2 returns a function that returns err until the duration d has passed from its creation.
func testFunction2(d time.Duration) func() (int, error) {
	start := time.Now()

	return func() (int, error) {
		if time.Since(start) > d {
			return 100, nil
		}

		return 0, errRetry
	}
}

func TestExecuteWithRetry(t *testing.T) {
	tests := []struct {
		name     string
		f        func() (int, error)
		params   Params
		expected ExecuteStatus[int]
	}{
		{
			name: "success within max attempts",
			f:    testFunction(3),
			params: Params{
				MaxAttempts: 4,
				Delay:       0,
				Timeout:     0,
			},
			expected: ExecuteStatus[int]{
				Success: true,
				Err:     nil,
				Value:   100,
			},
		},
		{
			name: "success with unlimited attempts",
			f:    testFunction(3),
			params: Params{
				MaxAttempts: 0,
				Delay:       0,
				Timeout:     0,
			},
			expected: ExecuteStatus[int]{
				Success: true,
				Err:     nil,
				Value:   100,
			},
		},
		{
			name: "failure when attempts exhausted",
			f:    testFunction(3),
			params: Params{
				MaxAttempts: 2,
				Delay:       0,
				Timeout:     0,
			},
			expected: ExecuteStatus[int]{
				Success: false,
				Err:     fmt.Errorf("all attempts failed. Last error: %w", errRetry),
				Value:   0,
			},
		},
		{
			name: "success before timeout",
			f:    testFunction2(5 * time.Millisecond),
			params: Params{
				MaxAttempts: 0,
				Delay:       3 * time.Millisecond,
				Timeout:     15 * time.Millisecond,
			},
			expected: ExecuteStatus[int]{
				Success: true,
				Err:     nil,
				Value:   100,
			},
		},
		{
			name: "context deadline exceeded",
			f:    testFunction2(20 * time.Millisecond),
			params: Params{
				MaxAttempts: 0,
				Delay:       2 * time.Millisecond,
				Timeout:     5 * time.Millisecond,
			},
			expected: ExecuteStatus[int]{
				Success: false,
				Err:     fmt.Errorf("context error mid retry: %w. Last error: %w", context.DeadlineExceeded, errRetry),
				Value:   0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			outcome := Execute(ctx, test.f, test.params)
			require.Equal(t, test.expected, outcome)
		})
	}
}

// TestExecutePreservesLastValueOnExhaustion verifies that Value holds the last
// attempt's value when all attempts fail, rather than the zero value.
func TestExecutePreservesLastValueOnExhaustion(t *testing.T) {
	f := func() (int, error) {
		return 42, errRetry
	}
	res := Execute(t.Context(), f, Params{MaxAttempts: 3})
	require.False(t, res.Success)
	require.Error(t, res.Err)
	require.Equal(t, 42, res.Value)
}

func TestIngrainAttempt(t *testing.T) {
	identity := func(i int) (int, error) {
		return i, nil
	}

	ingrained := ingrainAttempt(identity)

	for j := range 10 {
		x, err := ingrained()
		require.NoError(t, err)
		require.Equal(t, j, x)
	}
}

func TestIngrainAttemptConcurrent(t *testing.T) {
	called := struct {
		c map[int]int

		sync.Mutex
	}{}

	called.c = make(map[int]int)

	identity := func(i int) (int, error) {
		called.Lock()
		defer called.Unlock()
		called.c[i]++
		return i, nil
	}

	ingrained := ingrainAttempt(identity)

	var wg sync.WaitGroup

	const n = 100
	for range n {
		wg.Go(func() {
			_, err := ingrained()
			require.NoError(t, err)
		})
	}

	wg.Wait()

	require.Len(t, called.c, n)

	for j := range n {
		require.Equal(t, 1, called.c[j])
	}
}

func TestNextDelay(t *testing.T) {
	base := 10 * time.Millisecond

	tests := []struct {
		name   string
		params Params
		j      int
		want   time.Duration
	}{
		{
			name:   "constant when multiplier <= 1",
			params: Params{Delay: base},
			j:      5,
			want:   base,
		},
		{
			name:   "first retry uses base delay (j=0 unaffected by multiplier)",
			params: Params{Delay: base, Multiplier: 2},
			j:      0,
			want:   base,
		},
		{
			name:   "exponential doubling at j=3",
			params: Params{Delay: base, Multiplier: 2},
			j:      3,
			want:   80 * time.Millisecond, // base * 2^3
		},
		{
			name:   "MaxDelay caps growth",
			params: Params{Delay: base, Multiplier: 2, MaxDelay: 50 * time.Millisecond},
			j:      10,
			want:   50 * time.Millisecond,
		},
		{
			name:   "saturates on overflow without panic",
			params: Params{Delay: time.Hour, Multiplier: 1e9, MaxDelay: 2 * time.Second},
			j:      100,
			want:   2 * time.Second,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, nextDelay(tc.params, tc.j))
		})
	}
}

func TestNextDelayJitter(t *testing.T) {
	base := 100 * time.Millisecond
	p := Params{Delay: base, Jitter: 0.5}

	for range 200 {
		d := nextDelay(p, 0)
		require.GreaterOrEqual(t, d, 50*time.Millisecond)
		require.LessOrEqual(t, d, 150*time.Millisecond)
	}
}

func TestExecuteReturnsEarlyOnCtxCancelDuringDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	delay := 100 * time.Millisecond
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	res := Execute(ctx, func() (int, error) { return 0, errRetry }, Params{
		MaxAttempts: 10,
		Delay:       delay,
	})
	elapsed := time.Since(start)

	require.False(t, res.Success)
	require.ErrorIs(t, res.Err, context.Canceled)
	require.ErrorIs(t, res.Err, errRetry)
	require.Less(t, elapsed, delay, "Execute must return before the next delay completes (elapsed=%s)", elapsed)
}

func TestExecuteBackoffDelays(t *testing.T) {
	// 3 attempts with base 5ms, multiplier 2: delays before attempts 2 and 3 are 5ms and 10ms => >= 15ms.
	start := time.Now()
	res := Execute(context.Background(), testFunction(2), Params{
		MaxAttempts: 3,
		Delay:       5 * time.Millisecond,
		Multiplier:  2,
		Timeout:     500 * time.Millisecond,
	})
	elapsed := time.Since(start)
	require.True(t, res.Success)
	require.GreaterOrEqual(t, elapsed, 15*time.Millisecond, "backoff did not grow as expected (elapsed=%s)", elapsed)
}

func TestExecuteAttempt(t *testing.T) {
	limit := 3

	tests := []struct {
		name string
		p    Params
		e    ExecuteStatus[bool]
	}{
		{
			name: "failure when attempts exhausted",
			p: Params{
				MaxAttempts: 2,
				Delay:       0,
				Timeout:     0,
			},
			e: ExecuteStatus[bool]{
				Success: false,
				Err:     fmt.Errorf("all attempts failed. Last error: %w", errRetry),
				Value:   false,
			},
		},
		{
			name: "success within max attempts",
			p: Params{
				MaxAttempts: 5,
				Delay:       0,
				Timeout:     0,
			},
			e: ExecuteStatus[bool]{
				Success: true,
				Err:     nil,
				Value:   true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			attempts := make(map[int]int)

			testFunc := func(i int) (bool, error) {
				attempts[i]++
				if i < limit {
					return false, errRetry
				}
				return true, nil
			}

			ctx := context.Background()
			res := ExecuteAttempt(ctx, testFunc, test.p)
			require.Equal(t, test.e, res)

			maxAttempts := test.p.MaxAttempts
			if maxAttempts == 0 {
				maxAttempts = limit + 1
			}

			require.Len(t, attempts, min(maxAttempts, 4))

			for _, value := range attempts {
				require.Equal(t, 1, value)
			}
		})
	}
}
