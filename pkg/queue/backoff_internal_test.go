package queue

import (
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/require"
)

// TestNewBackoffNilCallbackFallsBack verifies that a caller Backoff callback
// returning nil falls back to a constant backoff, so handleError never calls
// NextBackOff on a nil value.
func TestNewBackoffNilCallbackFallsBack(t *testing.T) {
	q := PriorityQueue[int]{
		backoff: func() backoff.BackOff { return nil },
	}

	b := q.newBackoff()
	require.NotNil(t, b)
	require.NotPanics(t, func() { b.NextBackOff() })
}
