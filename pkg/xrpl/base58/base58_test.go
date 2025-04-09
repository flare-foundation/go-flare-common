package base58

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct{ input []byte }{
		{
			[]byte{},
		},
		{
			[]byte("random"),
		},
		{
			[]byte{0, 0, 0},
		},
		{
			[]byte{0, 1, 0},
		},
	}

	for _, test := range tests {
		encoded := XRPLCoder.Encode(test.input)

		decoded, err := XRPLCoder.Decode(encoded)
		require.NoError(t, err)

		require.Equal(t, test.input, decoded)
	}
}

func TestDecodeEncode(t *testing.T) {
	tests := []struct {
		input string
		ok    bool
	}{
		{
			"r",
			true,
		},
		{
			"rMdpzP6SSRcNDSgHQGGx4hV3qeRiuRoqQM",
			true,
		},
		{
			"rrrr",
			true,
		},
		{
			"rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz",
			true,
		},
		{
			"l",
			false,
		},
		{
			"日本語",
			false,
		},
	}

	for _, test := range tests {
		decoded, err := XRPLCoder.Decode(test.input)
		if test.ok {
			require.NoError(t, err)
			encoded := XRPLCoder.Encode(decoded)

			require.Equal(t, test.input, encoded)
		} else {
			require.Error(t, err)
		}
	}
}
