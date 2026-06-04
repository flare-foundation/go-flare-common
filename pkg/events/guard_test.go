package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSelectorFromABIRejectsNil verifies SelectorFromABI errors instead of
// panicking on a nil ABI.
func TestSelectorFromABIRejectsNil(t *testing.T) {
	_, err := SelectorFromABI(nil, "Foo")
	require.Error(t, err)
}
