package queue_test

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/queue"
	"github.com/stretchr/testify/require"
)

// TestNewPriorityNegativeSize verifies a negative configured size is clamped
// rather than panicking the channel make at construction.
func TestNewPriorityNegativeSize(t *testing.T) {
	require.NotPanics(t, func() {
		_ = queue.NewPriority[int](&queue.PriorityQueueParams{Size: -5})
	})
}
