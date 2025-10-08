package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

		deserialized, err := Amount.ToJSON(b, 0)
		require.NoError(t, err, j)

		fmt.Printf("deserialized: %v\n", deserialized)
	}
}

func TestAmountEncodeDecode(t *testing.T) {
	amount := "15000000000"

	value, err := Amount.ToBytes(amount, false)
	require.NoError(t, err)

	b := bytes.NewBuffer(value)

	decoded, err := Amount.ToJSON(b, 0)
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
		` {
                "currency": "USD",
                "value": "10",
                "issuer": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAee"
            }`,
		` {
                "currency": 1,
                "value": "10",
                "issuer": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAee"
            }`,
		` {
                "currency": "USD",
                "value": "10"
			}`,
		` {
                "currency": "USD",
                "value": "10",
				"invalidField": 10
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

func TestNormalizeValue(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"0",
			"8000000000000000",
		},
	}

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

func TestAmountDecodeEncode(t *testing.T) {
	inputs := [][]byte{
		{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x20, 0xB3, 0xC8, 0x5F, 0x48, 0x25, 0x32, 0xA9, 0x57, 0x8D, 0xBB, 0x39, 0x50, 0xB8, 0x5C, 0xA0, 0x65, 0x94, 0xD1},
		{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0xd3, 0x3b, 0x18, 0xd5, 0x33, 0x85, 0xf8, 0xa3, 0x18, 0x55, 0x16, 0xc2, 0xed, 0xa5, 0xde, 0xdb, 0x8a, 0xc5, 0xc6},
		{0x40, 0x00, 0x00, 0x00, 0x00, 0x8, 0x2, 0x01},
		{0x40, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80},
	}

	for _, input := range inputs {
		b := bytes.NewBuffer(input)

		decoded, err := Amount.ToJSON(b, 0)
		require.NoError(t, err)
		require.Equal(t, 0, b.Len())

		encoded, err := Amount.ToBytes(decoded, false)
		require.NoError(t, err)

		require.Equal(t, input, encoded)
	}
}

func TestAmountDecodeFail(t *testing.T) {
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

		_, err := Amount.ToJSON(b, 0)
		require.Error(t, err, j)
	}
}

func TestMPT(t *testing.T) {
	tests := []map[string]any{
		{
			"value":           "1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "9223372036854775807",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
	}

	for _, test := range tests {
		b, err := Amount.ToBytes(test, false)
		require.NoError(t, err)

		buf := bytes.NewBuffer(b)

		j, err := Amount.ToJSON(buf, 0)
		require.NoError(t, err)

		_, ok := j.(map[string]any)
		require.True(t, ok)

		b2, err := Amount.ToBytes(j, false)
		require.NoError(t, err)

		require.Equal(t, b, b2)
	}
}

func TestMPTWithExpectation(t *testing.T) {
	tests := []struct {
		amount           map[string]any
		expectedEncoding string
	}{
		{
			amount: map[string]any{
				"value":           "9223372036854775807",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "607FFFFFFFFFFFFFFF00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			amount: map[string]any{
				"value":           "0",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			amount: map[string]any{
				"value":           "-0",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			amount: map[string]any{
				"value":           "100",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			amount: map[string]any{
				"value":           "0xa",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "60000000000000000A00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			amount: map[string]any{
				"value":           "0x7FFFFFFFFFFFFFFF",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedEncoding: "607FFFFFFFFFFFFFFF00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
	}

	for _, test := range tests {
		b, err := Amount.ToBytes(test.amount, false)
		require.NoError(t, err)

		bExpected, err := hex.DecodeString(test.expectedEncoding)
		require.NoError(t, err)

		require.Equal(t, bExpected, b)

		buf := bytes.NewBuffer(b)

		j, err := Amount.ToJSON(buf, 0)
		require.NoError(t, err)

		_, ok := j.(map[string]any)
		require.True(t, ok)

		b2, err := Amount.ToBytes(j, false)
		require.NoError(t, err)

		require.Equal(t, b, b2)
	}
}

func TestExoticValues(t *testing.T) {
	tests := []struct {
		amount        map[string]any
		expectedValue string
	}{
		{
			amount: map[string]any{
				"value":           "1E1",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedValue: "10",
		},
		{
			amount: map[string]any{
				"value":           "1e1",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedValue: "10",
		},
		{
			amount: map[string]any{
				"value":           "0x01",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedValue: "1",
		},
		{
			amount: map[string]any{
				"value":           "0x1",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedValue: "1",
		},
		{
			amount: map[string]any{
				"value":           "0",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedValue: "0",
		},
	}

	for _, test := range tests {
		b, err := Amount.ToBytes(test.amount, false)
		require.NoError(t, err)

		buf := bytes.NewBuffer(b)

		j, err := Amount.ToJSON(buf, 0)
		require.NoError(t, err)

		jm, ok := j.(map[string]any)
		require.True(t, ok)

		require.Equal(t, test.expectedValue, jm["value"])

		b2, err := Amount.ToBytes(j, false)
		require.NoError(t, err)

		require.Equal(t, b, b2)
	}
}

func TestMPTFail(t *testing.T) {
	tests := []map[string]any{
		{
			"value":           "9223372036854775808",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "18446744073709551615",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "-1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "10.1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "1E100",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "1E-1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "-1E1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "-1E",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "10",
			"mpt_issuance_id": "10",
		},
		{
			"value":           "1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			"currency":        "USD",
		},
		{
			"value":           "1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			"issuer":          "rrrrrrrrrrrrrrrrrrrrBZbvji",
		},
		{
			"value":           "a",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "?",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "0xFFFFFFFFFFFFFFFF",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "0x-1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
		{
			"value":           "-0x1",
			"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
		},
	}

	for _, test := range tests {
		_, err := Amount.ToBytes(test, false)
		require.Error(t, err)
	}
}

func TestMPTToJSONFail(t *testing.T) {
	inputs := [][]byte{
		{},
		{0x60},
		{0x60, 0x00},
		{0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x60, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x60, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	for j, input := range inputs {
		b := bytes.NewBuffer(input)

		_, err := Amount.ToJSON(b, 0)
		require.Error(t, err, j)
		fmt.Printf("err: %v\n", err)
	}
}
