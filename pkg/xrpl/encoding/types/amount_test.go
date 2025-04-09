package types

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXRPAmount(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "0",
			output: "4000000000000000",
		},
		{
			input:  "1",
			output: "4000000000000001",
		},
		{
			input:  "10",
			output: "400000000000000a",
		},
		{
			input:  "100",
			output: "4000000000000064",
		},
		{
			input:  "1000",
			output: "40000000000003e8",
		},
		{
			input:  "10000",
			output: "4000000000002710",
		},
		{
			input:  "100000",
			output: "40000000000186a0",
		},
		{
			input:  "1000000",
			output: "40000000000F4240",
		},
		{
			input:  "10000000",
			output: "4000000000989680",
		},
		{
			input:  "100000000",
			output: "4000000005F5E100",
		},
		{
			input:  "100000000000000000",
			output: "416345785D8A0000",
		},
	}

	for j, test := range tests {
		output, err := xrpAmountToBytes(test.input)
		require.NoError(t, err)

		to, err := hex.DecodeString(test.output)
		require.NoError(t, err)

		require.Equal(t, to, output, fmt.Sprintf("test %d failed. input: %s ", j, test.input))
	}

	testsFail := []struct {
		input string
	}{
		{
			input: "",
		},
		{
			input: "-1",
		},
		{
			input: "100000000000000001",
		},
		{
			input: "a",
		},
	}

	for j, test := range testsFail {
		_, err := xrpAmountToBytes(test.input)
		require.Error(t, err, fmt.Sprintf("fail test %d failed. input: %s ", j, test.input))
	}
}

func TestExtractString(t *testing.T) {
	testMap := map[string]any{
		"ok":    "ok",
		"":      "",
		"int":   10,
		"bool":  false,
		"bytes": []byte("bytes"),
		"nil":   nil,
	}

	extracted, err := extractString(testMap, "ok")
	require.NoError(t, err)
	require.Equal(t, "ok", extracted)

	extracted, err = extractString(testMap, "")
	require.NoError(t, err)
	require.Equal(t, "", extracted)

	for _, key := range []string{"int", "bool", "bytes", "unknown", "nil"} {
		_, err := extractString(testMap, key)
		require.Error(t, err, fmt.Sprintf("fail test failed: %s", key))
	}
}

func TestSerializeCurrency(t *testing.T) {
	testsStandard := []string{"ABC", "1()"}

	zeros12 := make([]byte, 12)
	zeros5 := make([]byte, 5)

	for _, code := range testsStandard {
		output, err := serializeCurrency(code)
		require.NoError(t, err)

		require.Equal(t, zeros12, output[:12])
		require.Equal(t, zeros5, output[15:])

		require.Equal(t, []byte(code), output[12:15])
	}

	testsNonStandard := []string{"436152526F747300000000000000000000000000", "0x584A4F5900000000000000000000000000000000"}

	for _, code := range testsNonStandard {
		output, err := serializeCurrency(code)
		require.NoError(t, err)

		expected, err := hex.DecodeString(strings.TrimPrefix(code, "0x"))
		require.NoError(t, err)

		require.Equal(t, expected, output)
	}

	testsFail := []string{"XRP", "AB", "", "0000000000000000000000005852500000000000", "0x0000000000000000000000000000000000000000", "asdf", "⌘⌘⌘", "'''"}

	for _, code := range testsFail {
		_, err := serializeCurrency(code)
		require.Error(t, err)
	}
}

// todo
func TestNormalizeValue(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{}

	for _, test := range tests {
		normalized, err := normalizeValue(test.input)
		require.NoError(t, err)

		require.NoError(t, err)

		normalizedHex := hex.EncodeToString(normalized)

		fmt.Printf("normalizedHex: %v\n", normalizedHex)

		expected, err := hex.DecodeString(test.output)
		require.NoError(t, err)

		require.Equal(t, expected, normalized)
	}
}
