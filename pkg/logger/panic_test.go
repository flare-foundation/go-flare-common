package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPanicAlwaysPanics pins the contract that the package-level Panic and
// Panicf panic regardless of configuration — even with no cores or a level
// filter above PANIC. Invariant checks elsewhere (e.g. pkg/queue worker
// accounting) terminate through this guarantee.
func TestPanicAlwaysPanics(t *testing.T) {
	defer Set(DefaultConfig())

	configs := map[string]Config{
		"no cores":    {Console: false},
		"level FATAL": {Level: "FATAL", Console: false},
	}

	for name, cfg := range configs {
		t.Run(name, func(t *testing.T) {
			Set(cfg)
			require.Panics(t, func() { Panic("invariant") })
			require.Panics(t, func() { Panicf("invariant %d", 1) })
		})
	}
}
