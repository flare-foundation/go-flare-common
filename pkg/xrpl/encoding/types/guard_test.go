package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVector256ToJSONRejectsNegativeLength verifies a negative length errors
// instead of panicking on make([]any, negative). A negative multiple of 32
// passes the %32 check, so the explicit l < 0 guard is required.
func TestVector256ToJSONRejectsNegativeLength(t *testing.T) {
	_, err := Vector256.ToJSON(bytes.NewBuffer(nil), -32)
	require.Error(t, err)
}
