package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVLLengthEncodeBoundaryBytes asserts the exact byte shape at every
// branch boundary defined in rippled's Serializer::addEncoded.
//
// source: rippled src/libxrpl/protocol/Serializer.cpp:210-241 (addEncoded)
//
//	n<=192     -> [n]
//	n<=12480   -> [193+((n-193)>>8), (n-193)&0xff]
//	n<=918744  -> [241+((n-12481)>>16), ((n-12481)>>8)&0xff, (n-12481)&0xff]
func TestVLLengthEncodeBoundaryBytes(t *testing.T) {
	tests := []struct {
		n    int
		want []byte
	}{
		{0, []byte{0x00}},
		{1, []byte{0x01}},
		{192, []byte{0xC0}},       // last single-byte value
		{193, []byte{0xC1, 0x00}}, // first two-byte value
		{194, []byte{0xC1, 0x01}},
		{448, []byte{0xC1, 0xFF}},   // 193 + 255 boundary within a single C1 prefix
		{449, []byte{0xC2, 0x00}},   // next prefix byte
		{12480, []byte{0xF0, 0xFF}}, // last two-byte value
		{12481, []byte{0xF1, 0x00, 0x00}},
		{12482, []byte{0xF1, 0x00, 0x01}},
		{918744, []byte{0xFE, 0xD4, 0x17}}, // last valid three-byte value
	}

	for _, tc := range tests {
		prefix, err := lengthEncode(tc.n)
		require.NoError(t, err, tc.n)
		require.Equal(t, tc.want, prefix, tc.n)

		// Round-trip decode.
		decoded, err := lengthDecode(bytes.NewBuffer(prefix))
		require.NoError(t, err, tc.n)
		require.Equal(t, tc.n, decoded)
	}
}

// TestVLLengthOverflow verifies the upper-bound rejection above 918744.
//
// source: rippled src/libxrpl/protocol/Serializer.cpp:235-238 (Throw "lenlen")
// source: rippled src/libxrpl/protocol/Serializer.cpp:287-288 (decodeVLLength b1>254 throw)
func TestVLLengthOverflow(t *testing.T) {
	_, err := lengthEncode(918745)
	require.Error(t, err)

	_, err = lengthEncode(-1)
	require.Error(t, err)

	// b1=255 is out of range for VL length.
	_, err = lengthDecode(bytes.NewBuffer([]byte{0xFF}))
	require.Error(t, err)

	// Exactly above maximum 3-byte value.
	_, err = lengthDecode(bytes.NewBuffer([]byte{0xFE, 0xD4, 0x18}))
	require.Error(t, err)
}
