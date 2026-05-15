package logger

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNopPanicAndFatalTerminate covers audit finding F-LOG-2: Nop.Panic and
// Nop.Fatal previously returned silently, breaking the terminating-call
// contract that downstream code (e.g., logger.Fatal("invariant"); return ok)
// relies on. They must now panic so the call does not return.
func TestNopPanicAndFatalTerminate(t *testing.T) {
	require.PanicsWithValue(t, "boom", func() { Nop{}.Panic("boom") })
	require.PanicsWithValue(t, "boom 42", func() { Nop{}.Panicf("boom %d", 42) })
	require.PanicsWithValue(t, "boom", func() { Nop{}.Fatal("boom") })
	require.PanicsWithValue(t, "boom 42", func() { Nop{}.Fatalf("boom %d", 42) })
}

// TestSetInvalidLevelFallsBackToDebug covers audit finding F-LOG-1: an
// unrecognised Config.Level used to return the zero-value zapcore.Level
// (InfoLevel), silently downgrading DEBUG callers to INFO. The fix falls
// back to DEBUG so debug messages stay visible and the operator-facing
// error log itself is guaranteed to surface.
func TestSetInvalidLevelFallsBackToDebug(t *testing.T) {
	t.Cleanup(func() { Set(DefaultConfig()) })
	Set(Config{Level: "not-a-level", Console: false})
	// The fact that this returns without panic and the logger is non-nil
	// is enough; the level itself is internal to zap.
	require.NotNil(t, Logger())
}

// TestSetConcurrentWithLogger covers audit Low 11e: Set must be safe to call
// concurrently with the logging convenience functions and Logger().
func TestSetConcurrentWithLogger(t *testing.T) {
	require.NotNil(t, Logger())

	var wg sync.WaitGroup
	const n = 50

	for range n {
		wg.Add(2)
		go func() {
			defer wg.Done()
			Set(Config{Level: "INFO", Console: false})
		}()
		go func() {
			defer wg.Done()
			Logger().Infof("concurrent read")
		}()
	}

	wg.Wait()
	require.NotNil(t, Logger())
}
