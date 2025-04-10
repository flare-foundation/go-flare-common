package types

import (
	"encoding/hex"
	"encoding/json"
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
		output, err := xrpToBytes(test.input)
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
		_, err := xrpToBytes(test.input)
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

func TestAmount(t *testing.T) {
	tests := []struct {
		json   string
		output string
	}{
		{
			json:   `"0"`,
			output: "4000000000000000",
		},
		{
			json:   `"1"`,
			output: "4000000000000001",
		},
		{
			json: ` {
                "currency": "USD",
                "value": "-1",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
			output: "94838D7EA4C6800000000000000000000000000055534400000000000000000000000000000000000000000000000001",
		},
		{
			json: `{
                "currency": "USD",
                "value": "1",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
			output: "D4838D7EA4C6800000000000000000000000000055534400000000000000000000000000000000000000000000000001",
		},
		{
			json: `{
                "currency": "USD",
                "value": "0",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
			output: "800000000000000000000000000000000000000055534400000000000000000000000000000000000000000000000001",
		},
		{
			json: `{
                "currency": "USD",
                "value": "0.000000000000000000000000000000000000000000000000000000000000000000000000001",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
			output: "C1C38D7EA4C6800000000000000000000000000055534400000000000000000000000000000000000000000000000001",
		},
		{
			json: `{
                "currency": "USD",
                "value": "0.000000000000000000000000000000000000000000000000000000000000000000000000001",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
			output: "C1C38D7EA4C6800000000000000000000000000055534400000000000000000000000000000000000000000000000001",
		},
	}

	for j := range tests {
		var value any

		err := json.Unmarshal([]byte(tests[j].json), &value)
		require.NoError(t, err)

		serialized, err := Amount.ToBytes(value, false)
		require.NoError(t, err, value)

		bytesOutput, err := hex.DecodeString(tests[j].output)
		require.NoError(t, err)

		require.Equal(t, bytesOutput, serialized)

	}
}

func TestAmountFail(t *testing.T) {
	tests := []string{`{
                "currency": "USD",
                "value": "9999999999999999000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
		`{
		        "currency": "USD",
		        "value": "1111111111111111.1",
		        "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
		    }`,
		`"-10000000000000000000000000"`,
		`"1000000000000000000"`,
		`[1]`,
		`{
		        "currency": "USDC",
		        "value": "1111111111111111.1",
		        "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
		    }`,
		`{
		        "currency": "USDC",
		        "value": "1111111111111111.1",
		        "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji",
				"extra": 10
		    }`,
		` {
                "currency": "USD",
                "value": 10,
                "issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji"
            }`,
	}

	for j := range tests {
		var value any

		err := json.Unmarshal([]byte(tests[j]), &value)
		require.NoError(t, err)

		_, err = Amount.ToBytes(value, false)
		require.Error(t, err, value)
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
