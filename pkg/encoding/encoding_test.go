package encoding

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestUint16toBytes(t *testing.T) {
	tests := []struct {
		input    uint16
		expected [2]byte
	}{
		{
			input:    0,
			expected: [2]byte{0x00, 0x00},
		},
		{
			input:    1,
			expected: [2]byte{0x00, 0x01},
		},
		{
			input:    256,
			expected: [2]byte{0x01, 0x00},
		},
		{
			input:    65535,
			expected: [2]byte{0xFF, 0xFF},
		},
		{
			input:    0x1234,
			expected: [2]byte{0x12, 0x34},
		},
	}

	for _, tt := range tests {
		result := Uint16toBytes(tt.input)
		require.Equal(t, tt.expected, result, tt.input)
	}
}

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
