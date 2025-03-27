package retry

import (
	"context"
	"fmt"
	"time"
)

type ExecuteStatus[T any] struct {
	Success bool
	Err     error
	Value   T
}

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

	increment := 1
	attempts := params.MaxAttempts
	if params.MaxAttempts <= 0 {
		increment = 0
		attempts = 1
	}

	for j := 0; j < attempts; j += increment {
		if err = ctx.Err(); err != nil {
			result.Err = fmt.Errorf("context error mid retry: %v", err)
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

	result.Err = fmt.Errorf("max retries reached: %v", err)

	return result
}
