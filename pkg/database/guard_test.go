package database

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermanentNil(t *testing.T) {
	require.NoError(t, Permanent(nil))
}

// TestPermanentErrorNilSafe verifies a directly-constructed PermanentError with
// a nil Err does not panic on Error().
func TestPermanentErrorNilSafe(t *testing.T) {
	var e PermanentError
	require.NotPanics(t, func() { _ = e.Error() })
	require.Equal(t, "", e.Error())
}

func TestPermanentWrapsNonNil(t *testing.T) {
	sentinel := errors.New("boom")
	err := Permanent(sentinel)
	require.Error(t, err)
	require.ErrorIs(t, err, sentinel)
}
