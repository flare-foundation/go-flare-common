package logger

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

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
