package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyBlob(t *testing.T) {
	x, err := Blob.ToBytes("", false)
	require.NoError(t, err)

	fmt.Printf("x: %v\n", x)

	prefixed, err := encodeInner("SigningPubKey", "", true)
	require.NoError(t, err)

	fmt.Printf("prefixed: %v\n", prefixed)

}
