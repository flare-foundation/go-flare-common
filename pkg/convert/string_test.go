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
	})

	t.Run("valid 32 character string", func(t *testing.T) {
		input := "12345678901234567890123456789012"
		expected := common.Hash{}
		copy(expected[:], input)

		result, err := StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""
		expected := common.Hash{}

		result, err := StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("string too long", func(t *testing.T) {
		input := "123456789012345678901234567890123" // 33 characters

		_, err := StringToCommonHash(input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "string too long")
	})
}

func TestHex32StringToCommonHash(t *testing.T) {
	t.Run("valid hex string with 0x prefix", func(t *testing.T) {
		input := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

		result, err := Hex32StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", result.Hex()[2:])
	})

	t.Run("valid hex string with 0X prefix", func(t *testing.T) {
		input := "0X1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"

		result, err := Hex32StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", result.Hex()[2:])
	})

	t.Run("valid hex string without prefix", func(t *testing.T) {
		input := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

		result, err := Hex32StringToCommonHash(input)
		require.NoError(t, err)
		require.Equal(t, input, result.Hex()[2:])
	})

	t.Run("invalid hex characters", func(t *testing.T) {
		input := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdeg"

		_, err := Hex32StringToCommonHash(input)
		require.Error(t, err)
	})

	t.Run("wrong length - too short", func(t *testing.T) {
		input := "0x1234567890abcdef"

		_, err := Hex32StringToCommonHash(input)
		require.Error(t, err)
	})

	t.Run("wrong length - too long", func(t *testing.T) {
		input := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef12"

		_, err := Hex32StringToCommonHash(input)
		require.Error(t, err)
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""

		_, err := Hex32StringToCommonHash(input)
		require.Error(t, err)
	})
}
