package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const errorMsg = "still error"

var errRetry error = errors.New(errorMsg)

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

// testFunction returns a function that returns err until the duration d has passed from its creation.
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
		f        func() (int, error)
		params   Params
		expected ExecuteStatus[int]
	}{
		{
			f: testFunction(3),
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
			f: testFunction(3),
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
			f: testFunction(3),
			params: Params{
				MaxAttempts: 2,
				Delay:       0,
				Timeout:     0,
			},
			expected: ExecuteStatus[int]{
				Success: false,
				Err:     fmt.Errorf("max retries reached: %v", errRetry),
				Value:   0,
			},
		},
		{
			f: testFunction2(5 * time.Millisecond),
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
			f: testFunction2(20 * time.Millisecond),
			params: Params{
				MaxAttempts: 0,
				Delay:       2 * time.Millisecond,
				Timeout:     5 * time.Millisecond,
			},
			expected: ExecuteStatus[int]{
				Success: false,
				Err:     fmt.Errorf("context error mid retry: %v", context.DeadlineExceeded),
				Value:   0,
			},
		},
	}

	for j, test := range tests {
		ctx := context.Background()

		outcome := Execute(ctx, test.f, test.params)
		require.Equal(t, test.expected, outcome, j)
	}
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

func TestExecuteAttempt(t *testing.T) {
	attempts := make(map[int]int)

	limit := 3

	testFunc := func(i int) (bool, error) {
		attempts[i]++
		if i < limit {
			return false, errRetry
		}

		return true, nil
	}

	ctx := context.Background()

	tests := []struct {
		p Params
		e ExecuteStatus[bool]
	}{
		{
			p: Params{
				MaxAttempts: 2,
				Delay:       0,
				Timeout:     0,
			},
			e: ExecuteStatus[bool]{
				Success: false,
				Err:     fmt.Errorf("max retries reached: %v", errRetry),
				Value:   false,
			},
		},
		{
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

	for j, test := range tests {
		attempts = make(map[int]int)

		res := ExecuteAttempt(ctx, testFunc, test.p)
		require.Equal(t, test.e, res, j)

		maxAttempts := test.p.MaxAttempts
		if maxAttempts == 0 {
			maxAttempts = limit + 1
		}

		require.Len(t, attempts, min(maxAttempts, 4)) // test.p.MaxAttempts should not be 0 in test.

		for _, value := range attempts {
			require.Equal(t, value, 1)
		}
	}
}
