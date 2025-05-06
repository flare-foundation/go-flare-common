package types

import (
	"bytes"
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

func TestAmountDeserializeSerialize(t *testing.T) {
	tests := []struct{ input string }{{"94838D7EA4C6800000000000000000000000000055534400000000000000000000000000000000000000000000000001"}}

	for j, test := range tests {
		blob, err := hex.DecodeString(test.input)
		require.NoError(t, err, j)
		b := bytes.NewBuffer(blob)

		deserialized, err := Amount.ToJson(b, 0)
		require.NoError(t, err, j)

		fmt.Printf("deserialized: %v\n", deserialized)
	}
}

func TestAmountEncodeDecode(t *testing.T) {
	amount := "15000000000"

	value, err := Amount.ToBytes(amount, false)
	require.NoError(t, err)

	b := bytes.NewBuffer(value)

	decoded, err := Amount.ToJson(b, 0)
	require.NoError(t, err)

	require.Equal(t, amount, decoded)
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
		normalized, err := serializeTokenValue(test.input)
		require.NoError(t, err)

		require.NoError(t, err)

		normalizedHex := hex.EncodeToString(normalized)

		fmt.Printf("normalizedHex: %v\n", normalizedHex)

		expected, err := hex.DecodeString(test.output)
		require.NoError(t, err)

		require.Equal(t, expected, normalized)
	}
}

func TestDeserializeValue(t *testing.T) {
	inputsExact := []string{"0", "1", "-1", "0.1", "-0.1", "1e+95", "1e-81", "40915.87486543398"}
	for _, input := range inputsExact {
		s, err := serializeTokenValue(input)
		require.NoError(t, err)

		d, err := deserializeTokenAmount(s)
		require.NoError(t, err)

		require.Equal(t, input, d)
	}
}

func TestDecodeEncode(t *testing.T) {
	inputs := [][]byte{
		{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x20, 0xB3, 0xC8, 0x5F, 0x48, 0x25, 0x32, 0xA9, 0x57, 0x8D, 0xBB, 0x39, 0x50, 0xB8, 0x5C, 0xA0, 0x65, 0x94, 0xD1},
		{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0xd3, 0x3b, 0x18, 0xd5, 0x33, 0x85, 0xf8, 0xa3, 0x18, 0x55, 0x16, 0xc2, 0xed, 0xa5, 0xde, 0xdb, 0x8a, 0xc5, 0xc6},
		{0x40, 0x00, 0x00, 0x00, 0x00, 0x8, 0x2, 0x01},
		{0x40, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80},
	}

	for _, input := range inputs {
		b := bytes.NewBuffer(input)

		decoded, err := Amount.ToJson(b, 0)
		require.NoError(t, err)
		require.Equal(t, 0, b.Len())

		encoded, err := Amount.ToBytes(decoded, false)
		require.NoError(t, err)

		require.Equal(t, input, encoded)

	}
}

func TestDecodeFail(t *testing.T) {
	inputs := [][]byte{
		{},
		{0x4f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		{0x40},
		{0x40, 0x00},
		{0x00},
		{0x00, 0x00},
		{0x4f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		{0xff},
		{0xd4, 0x56, 0x4b},
		{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0xd3, 0x3b, 0x18, 0xd5, 0x33, 0x85, 0xf8, 0xa3, 0x18, 0x55, 0x16, 0xc2, 0xed, 0xa5, 0xde, 0xdb, 0x8a},
		{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00},
		{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	for j, input := range inputs {
		b := bytes.NewBuffer(input)

		_, err := Amount.ToJson(b, 0)
		require.Error(t, err, j)
	}
}
