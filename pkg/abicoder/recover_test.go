package abicoder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecoverError(t *testing.T) {
	errBoom := errors.New("boom")
	panicking := func(v any) (err error) {
		defer recoverError(&err)
		panic(v)
	}

	err := panicking(errBoom)
	require.ErrorIs(t, err, errBoom)
	require.EqualError(t, err, "recovered panic: boom")

	require.EqualError(t, panicking("boom"), "recovered panic non error: boom")

	returning := func() (err error) {
		defer recoverError(&err)
		return errBoom
	}
	require.ErrorIs(t, returning(), errBoom)
}
