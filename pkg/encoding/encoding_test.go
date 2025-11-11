package encoding

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
	"github.com/stretchr/testify/require"
)

func TestSignatureConversion(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)

	msg := crypto.Keccak256([]byte("test"))

	sig, err := crypto.Sign(msg, prv)
	require.NoError(t, err)

	vrs := TransformSignatureRSVtoVRS(sig)
	rsv := TransformSignatureVRStoRSV(vrs)

	require.Equal(t, sig, rsv)
}

func TestEncodeSignature(t *testing.T) {
	x := []byte("payload")

	h := accounts.TextHash(x)

	tests := []struct {
		indexes []int
		fail    bool
	}{
		{
			indexes: nil,
			fail:    false,
		},
		{
			indexes: []int{0, 1},
			fail:    false,
		},
		{
			indexes: []int{1, 2},
			fail:    false,
		},
		{
			indexes: []int{1, 10, 100},
			fail:    false,
		},
		{
			indexes: []int{0, 0},
			fail:    true,
		},
		{
			indexes: []int{-1, 0},
			fail:    true,
		},
		{
			indexes: []int{10, 11, 9},
			fail:    true,
		},
	}

	m := 0
	for j := range tests {
		if n := len(tests[j].indexes); n > m {
			m = n
		}
	}

	sigs := make([][]byte, m)

	for j := range sigs {
		priv, err := crypto.GenerateKey()
		require.NoError(t, err)

		sigs[j], err = crypto.Sign(h, priv)
		require.NoError(t, err)
	}

	for j, test := range tests {
		inSigs := make([]IndexedSignature, len(test.indexes))
		for j := range inSigs {
			inSigs[j] = IndexedSignature{
				Index:     test.indexes[j],
				Signature: sigs[j],
			}
		}

		eSigs, err := EncodeSignatures(inSigs)

		if test.fail {
			require.Error(t, err, j)
		} else {
			require.NoError(t, err, j)

			l, err := convert.BytesToUint16(eSigs[0:2])
			require.NoError(t, err, j)

			require.Equal(t, int(l), len(test.indexes), j)
		}
	}
}
