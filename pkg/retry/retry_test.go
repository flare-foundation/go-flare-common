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
			f: testFunction2(10 * time.Millisecond),
			params: Params{
				MaxAttempts: 0,
				Delay:       3 * time.Millisecond,
				Timeout:     13 * time.Millisecond,
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

	for _, test := range tests {
		ctx := context.Background()

		outcome := Execute(ctx, test.f, test.params)

		require.Equal(t, test.expected, outcome)
	}
}
