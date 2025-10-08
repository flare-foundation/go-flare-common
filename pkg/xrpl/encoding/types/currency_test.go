package types

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerializeCurrency(t *testing.T) {
	testsStandard := []string{"ABC", "1()"}

	zeros12 := make([]byte, 12)
	zeros5 := make([]byte, 5)

	for _, code := range testsStandard {
		output, err := serializeCurrency(code, disallowedCodes)
		require.NoError(t, err)

		require.Equal(t, zeros12, output[:12])
		require.Equal(t, zeros5, output[15:])

		require.Equal(t, []byte(code), output[12:15])
	}

	testsNonStandard := []string{"436152526F747300000000000000000000000000", "0x584A4F5900000000000000000000000000000000"}

	for _, code := range testsNonStandard {
		output, err := serializeCurrency(code, disallowedCodes)
		require.NoError(t, err)

		expected, err := hex.DecodeString(strings.TrimPrefix(code, "0x"))
		require.NoError(t, err)

		require.Equal(t, expected, output)
	}

	testsFail := []string{"XRP", "AB", "", "0000000000000000000000005852500000000000", "0x0000000000000000000000000000000000000000", "asdf", "⌘⌘⌘", "'''"}

	for _, code := range testsFail {
		_, err := serializeCurrency(code, disallowedCodes)
		require.Error(t, err)
	}
}

func TestCurrency(t *testing.T) {
	// Test ToBytes with standard code
	t.Run("ToBytes standard code", func(t *testing.T) {
		b, err := Currency.ToBytes("USD", false)
		require.NoError(t, err)
		require.Len(t, b, 20)
		require.Equal(t, []byte("USD"), b[12:15])
	})

	// Test ToBytes with nonstandard code
	t.Run("ToBytes nonstandard code", func(t *testing.T) {
		code := "436152526F747300000000000000000000000000"
		b, err := Currency.ToBytes(code, false)
		require.NoError(t, err)
		require.Len(t, b, 20)
		expected, _ := hex.DecodeString(code)
		require.Equal(t, expected, b)
	})

	// Test ToBytes with invalid code
	t.Run("ToBytes invalid code", func(t *testing.T) {
		_, err := Currency.ToBytes("XaRP", false)
		require.Error(t, err)
	})

	// Test ToBytes with non-string value
	t.Run("ToBytes non-string value", func(t *testing.T) {
		_, err := Currency.ToBytes(123, false)
		require.Error(t, err)
	})

	// Test ToJSON with standard code bytes
	t.Run("ToJSON standard code", func(t *testing.T) {
		buf := make([]byte, 20)
		copy(buf[12:], []byte("EUR"))
		b := bytes.NewBuffer(buf)
		val, err := Currency.ToJSON(b, 0)
		require.NoError(t, err)
		require.Equal(t, "EUR", val)
	})

	// Test ToJSON with nonstandard code bytes
	t.Run("ToJSON nonstandard code", func(t *testing.T) {
		code := "584A4F5900000000000000000000000000000000"
		bb, _ := hex.DecodeString(code)
		b := bytes.NewBuffer(bb)
		val, err := Currency.ToJSON(b, 0)
		require.NoError(t, err)
		require.Equal(t, strings.ToUpper(code), val)
	})

	// Test ToJSON with insufficient bytes
	t.Run("ToJSON insufficient bytes", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{1, 2, 3})
		_, err := Currency.ToJSON(b, 0)
		require.Error(t, err)
	})

	// Test ToJSON with XRP code
	t.Run("ToJSON XRP code", func(t *testing.T) {
		buf := make([]byte, 20)
		copy(buf[12:], []byte("XRP"))
		b := bytes.NewBuffer(buf)
		_, err := Currency.ToJSON(b, 0)
		require.Error(t, err)
	})
}
