// Package retry provides generic retry execution with configurable attempts, delay, and timeout.
package retry

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ExecuteStatus holds the result of an Execute call, including success state, value, and any error.
type ExecuteStatus[T any] struct {
	Success bool
	Err     error
	Value   T
}

// Params configures retry behavior for Execute.
type Params struct {
	MaxAttempts int           // if non positive, number of attempts is not limited.
	Delay       time.Duration // minimal delay between each attempts
	Timeout     time.Duration // total maximal duration of the execution. If zero, there is no Timeout. For a single execution, the function should handle timeout.
}

// Execute executes function f and retries on error according to params.
func Execute[T any](ctx context.Context, f func() (T, error), params Params) ExecuteStatus[T] {
	var cancel context.CancelFunc

	if params.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, params.Timeout)
		defer cancel()
	}

	var ticker *time.Ticker

	if params.Delay > 0 {
		ticker = time.NewTicker(params.Delay)
	}
	var result ExecuteStatus[T]

	var err error
	var r T

	cond := func(j, attempts int) bool {
		return j < attempts
	}
	if params.MaxAttempts <= 0 { // unlimited attempts
		cond = func(_, _ int) bool {
			return true
		}
	}

	for j := 0; cond(j, params.MaxAttempts); j++ {
		if cerr := ctx.Err(); cerr != nil {
			result.Err = fmt.Errorf("context error mid retry: %v. Last error: %w", cerr, err)
			return result
		}

		r, err = f()
		if err == nil {
			result.Success = true
			result.Value = r
			return result
		}

		if params.Delay > 0 {
			<-ticker.C
		}
	}

	result.Err = fmt.Errorf("max retries reached. Last error: %v", err)

	return result
}

// ExecuteAttempt executes function f that takes the index of the attempt as a parameter and retries on error according to params.
func ExecuteAttempt[T any](ctx context.Context, f func(j int) (T, error), params Params) ExecuteStatus[T] {
	h := ingrainAttempt(f)
	return Execute(ctx, h, params)
}

// ingrainAttempt takes a function f that takes an integer for the param and returns a function that on the j-th call returns f(j-1).
//
// The returned function should be used with caution in a concurrent scenario.
func ingrainAttempt[T any](f func(int) (T, error)) func() (T, error) {
	x := &sync.Mutex{}
	j := -1
	return func() (T, error) {
		x.Lock()
		j++
		i := j
		x.Unlock()
		return f(i)
	}
}
