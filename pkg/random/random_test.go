package random

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	hash, err := Hash()
	require.NoError(t, err)

	zeroHash := common.Hash{}
	require.NotEqual(t, zeroHash, hash)

	hash2, err := Hash()
	require.NoError(t, err)

	require.NotEqual(t, hash, hash2)
}

func TestBytes(t *testing.T) {
	lengths := []int{0, 3, 7}

	for _, l := range lengths {
		b0, err := Bytes(l)
		require.NoError(t, err)
		b1, err := Bytes(l)
		require.NoError(t, err)

		z := make([]byte, l)

		require.Len(t, b0, l)
		require.Len(t, b1, l)

		if l > 0 {
			require.NotEqual(t, z, b0)
			require.NotEqual(t, z, b1)
			require.NotEqual(t, b0, b1)
		}
	}

	_, err := Bytes(-1)
	require.Error(t, err, "negative length")
}
