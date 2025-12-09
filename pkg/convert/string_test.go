package convert

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestStringToCommonHash(t *testing.T) {
	t.Run("valid short string", func(t *testing.T) {
		input := "hello"
		expected := common.Hash{}
		copy(expected[:], "hello")

		result, err := StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, expected, result)

		back := CommonHashToString(result)
		require.Equal(t, input, back)
	})

	t.Run("valid 32 character string", func(t *testing.T) {
		input := "12345678901234567890123456789012"
		expected := common.Hash{}
		copy(expected[:], input)

		result, err := StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, expected, result)

		back := CommonHashToString(result)
		require.Equal(t, input, back)
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""
		expected := common.Hash{}

		result, err := StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, expected, result)

		back := CommonHashToString(result)
		require.Equal(t, input, back)
	})

	t.Run("string too long", func(t *testing.T) {
		input := "123456789012345678901234567890123" // 33 characters

		_, err := StringToCommonHash(input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "string too long")
	})
}

func TestHex32StringToCommonHash(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
		err         bool
	}{
		{
			description: "valid hex string with 0x prefix",
			input:       "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			expected:    "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			err:         false,
		},
		{
			description: "valid hex string with 0X prefix",
			input:       "0X1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
			expected:    "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			err:         false,
		},
		{
			description: "valid hex string without prefix",
			input:       "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			expected:    "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			err:         false,
		},
		{
			description: "invalid hex characters",
			input:       "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdeg",
			expected:    "",
			err:         true,
		},
		{
			description: "wrong length - too short",
			input:       "0x1234567890abcdef",
			expected:    "",
			err:         true,
		},
		{
			description: "wrong length - too long",
			input:       "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef12",
			expected:    "",
			err:         true,
		},
		{
			description: "empty string",
			input:       "",
			expected:    "",
			err:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := Hex32StringToCommonHash(test.input)
			if !test.err {
				require.NoError(t, err)
				require.Equal(t, test.expected, result.Hex()[2:])
			} else {
				require.Error(t, err)
			}
		})
	}
}
