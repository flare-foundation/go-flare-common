package types

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashError(t *testing.T) {
	tests := []struct {
		t      Encoder
		inputs []any
	}{
		{
			t:      Blob,
			inputs: []any{"a", "0g", "-19", 12, 0x11},
		},
		{
			t:      Hash128,
			inputs: []any{"a", "0g", "-19", "00"},
		},
		{
			t:      Hash160,
			inputs: []any{"a", "0g", "-19", "00"},
		},
		{
			t:      Hash192,
			inputs: []any{"a", "0g", "-19", "00"},
		},
		{
			t:      Hash256,
			inputs: []any{"a", "0g", "-19", "00"},
		},
	}

	for _, test := range tests {
		for _, input := range test.inputs {
			_, err := test.t.ToBytes(input, false)
			require.Error(t, err, input)
		}

	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		t      Encoder
		input  any
		output string
	}{
		{
			t:      Blob,
			input:  "",
			output: "",
		},
		{
			t:      Blob,
			input:  "AA",
			output: "aa",
		},
		{
			t:      Blob,
			input:  "00",
			output: "00",
		},
		{
			t:      Hash128,
			input:  "a0000000000000000000000000000000",
			output: "A0000000000000000000000000000000",
		},
	}

	for _, test := range tests {
		output, err := test.t.ToBytes(test.input, false)
		require.NoError(t, err)

		outputBytes, err := hex.DecodeString(test.output)
		require.NoError(t, err)
		require.Equal(t, outputBytes, output)
	}
}
