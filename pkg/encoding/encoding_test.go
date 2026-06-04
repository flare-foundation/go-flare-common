package encoding

import (
	"math"
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

	vrs, err := TransformSignatureRSVtoVRS(sig)
	require.NoError(t, err)
	rsv, err := TransformSignatureVRStoRSV(vrs)
	require.NoError(t, err)

	require.Equal(t, sig, rsv)
}

func TestTransformSignatureRejectsBadInput(t *testing.T) {
	_, err := TransformSignatureVRStoRSV(make([]byte, 64))
	require.Error(t, err)
	_, err = TransformSignatureVRStoRSV(make([]byte, 66))
	require.Error(t, err)

	// V == 0 (already normalised) must not underflow.
	v0 := make([]byte, 65)
	_, err = TransformSignatureVRStoRSV(v0)
	require.Error(t, err)

	_, err = TransformSignatureRSVtoVRS(make([]byte, 64))
	require.Error(t, err)
}

// TestTransformSignatureRangeChecks verifies that VRStoRSV rejects V > 28
// (out-of-range Ethereum recid), and RSVtoVRS rejects recid > 1.
func TestTransformSignatureRangeChecks(t *testing.T) {
	// V byte values that previously slipped through: EIP-155 chainID*2+35+{0,1},
	// and the all-ones byte.
	for _, v := range []byte{29, 37, 38, 0xFE, 0xFF} {
		vrs := make([]byte, 65)
		vrs[0] = v
		_, err := TransformSignatureVRStoRSV(vrs)
		require.Error(t, err, "V=%d must be rejected", v)
	}

	// Recid byte values that previously slipped through.
	for _, rec := range []byte{2, 3, 27, 100, 0xFE, 0xFF} {
		rsv := make([]byte, 65)
		rsv[64] = rec
		_, err := TransformSignatureRSVtoVRS(rsv)
		require.Error(t, err, "recid=%d must be rejected", rec)
	}
}

// TestEncodeSignatureRejectsLargeIndex verifies that an out-of-range Index is
// rejected. Index is encoded as uint16, but the prior monotonicity check compared
// full int values. A sorted, distinct int pair like [65535, 65536] passed
// the check and then truncated to [65535, 0] on the wire — silently
// non-monotonic and possibly duplicate.
func TestEncodeSignatureRejectsLargeIndex(t *testing.T) {
	priv, err := crypto.GenerateKey()
	require.NoError(t, err)
	sig, err := crypto.Sign(accounts.TextHash([]byte("x")), priv)
	require.NoError(t, err)

	for _, idx := range []int{math.MaxUint16 + 1, math.MaxUint16 + 10, math.MaxInt32} {
		_, err := EncodeSignatures([]IndexedSignature{{Index: idx, Signature: sig}})
		require.Error(t, err, "Index %d must be rejected", idx)
	}

	// Boundary: math.MaxUint16 must be accepted.
	_, err = EncodeSignatures([]IndexedSignature{{Index: math.MaxUint16, Signature: sig}})
	require.NoError(t, err)
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
