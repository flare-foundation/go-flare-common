// Package retry provides generic retry execution with configurable attempts, delay, and timeout.
package retry

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
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
	MaxAttempts int           // if non-positive, unlimited; caller must cancel ctx or set Timeout to stop.
	Delay       time.Duration // initial delay between attempts
	Timeout     time.Duration // total maximal duration. If zero, no timeout. Per-call timeout is the function's responsibility.

	Multiplier float64       // if >1, the delay is multiplied by this factor each attempt. Default ≤1 keeps delay constant.
	MaxDelay   time.Duration // optional cap on the computed delay. Zero means no cap.
	Jitter     float64       // optional randomization in [0,1]; applies ±Jitter*delay wobble. Zero is deterministic.
}

// Execute executes function f and retries on error according to params.
func Execute[T any](ctx context.Context, f func() (T, error), params Params) ExecuteStatus[T] {
	var cancel context.CancelFunc

	if params.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, params.Timeout)
		defer cancel()
	}

	var timer *time.Timer
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()

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
			result.Err = fmt.Errorf("context error mid retry: %w. Last error: %w", cerr, err)
			return result
		}

		r, err = f()
		if err == nil {
			result.Success = true
			result.Value = r
			return result
		}

		if params.Delay > 0 && cond(j+1, params.MaxAttempts) {
			d := nextDelay(params, j)
			if timer == nil {
				timer = time.NewTimer(d)
			} else {
				timer.Reset(d)
			}
			select {
			case <-timer.C:
			case <-ctx.Done():
			}
		}
	}

	result.Err = fmt.Errorf("all attempts failed. Last error: %w", err)

	return result
}

// nextDelay computes the delay before the next attempt (after attempt index j, 0-based).
// Applies Multiplier^j, then MaxDelay cap, then ±Jitter wobble. Saturates on overflow.
func nextDelay(p Params, j int) time.Duration {
	d := p.Delay
	if p.Multiplier > 1 && j > 0 {
		factor := math.Pow(p.Multiplier, float64(j))
		if math.IsInf(factor, 0) || float64(d)*factor > float64(math.MaxInt64) {
			d = math.MaxInt64
		} else {
			d = time.Duration(float64(d) * factor)
		}
	}
	if p.MaxDelay > 0 && d > p.MaxDelay {
		d = p.MaxDelay
	}
	if p.Jitter > 0 {
		jitter := p.Jitter
		if jitter > 1 {
			jitter = 1
		}
		wobble := 1.0 + (rand.Float64()*2-1)*jitter
		d = max(time.Duration(float64(d)*wobble), 0)
	}
	return d
}

// ExecuteAttempt executes function f that takes the index of the attempt as a parameter and retries on error according to params.
func ExecuteAttempt[T any](ctx context.Context, f func(j int) (T, error), params Params) ExecuteStatus[T] {
	h := ingrainAttempt(f)
	return Execute(ctx, h, params)
}

// ingrainAttempt takes a function f that takes an integer and returns a function that on call number j (starting from 0) returns f(j).
//
// The returned function should be used with caution in a concurrent scenario.
func ingrainAttempt[T any](f func(int) (T, error)) func() (T, error) {
	mu := &sync.Mutex{}
	j := -1
	return func() (T, error) {
		mu.Lock()
		j++
		i := j
		mu.Unlock()
		return f(i)
	}
}
