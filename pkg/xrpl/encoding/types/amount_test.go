package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAmountEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{
			name:  "XRP amount as string",
			value: "15000000000",
		},
		{
			name: "Token amount",
			value: map[string]any{
				"currency": "USD",
				"value":    "-1",
				"issuer":   "rrrrrrrrrrrrrrrrrrrrBZbvji",
			},
		},
		{
			name: "MPT amount 0",
			value: map[string]any{
				"value":           "9223372036854775807",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
		},
		{
			name: "MPT amount 1",
			value: map[string]any{
				"value":           "1",
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Encode
			encoded, err := Amount.ToBytes(test.value, false)
			require.NoError(t, err)

			// Decode
			buffer := bytes.NewBuffer(encoded)
			decoded, err := Amount.ToJSON(buffer, 0)
			require.NoError(t, err)
			require.Zero(t, buffer.Len(), "buffer should be fully read")

			// Verify
			require.Equal(t, test.value, decoded)

			// Re-encode and check for stability
			reEncoded, err := Amount.ToBytes(decoded, false)
			require.NoError(t, err)
			require.Equal(t, encoded, reEncoded)
		})
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

func TestAmountEncoding(t *testing.T) {
	tests := []struct {
		name        string
		jsonValue   string
		expectedHex string
		expectError bool
	}{
		// XRP Amounts
		{name: "XRP zero", jsonValue: `"0"`, expectedHex: "4000000000000000", expectError: false},
		{name: "XRP one", jsonValue: `"1"`, expectedHex: "4000000000000001", expectError: false},
		{name: "XRP 100", jsonValue: `"100"`, expectedHex: "4000000000000064", expectError: false},
		{name: "XRP 1000", jsonValue: `"1000"`, expectedHex: "40000000000003e8", expectError: false},
		{name: "XRP 10000", jsonValue: `"10000"`, expectedHex: "4000000000002710", expectError: false},
		{name: "XRP large", jsonValue: `"100000000000000000"`, expectedHex: "416345785D8A0000", expectError: false},

		// Token Amounts
		{name: "Token positive", jsonValue: `{"currency":"USD","value":"1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectedHex: "d4838d7ea4c6800000000000000000000000000055534400000000000000000000000000000000000000000000000001", expectError: false},
		{name: "Token negative", jsonValue: `{"currency":"USD","value":"-1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectedHex: "94838d7ea4c6800000000000000000000000000055534400000000000000000000000000000000000000000000000001", expectError: false},
		{name: "Token zero", jsonValue: `{"currency":"USD","value":"0","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectedHex: "800000000000000000000000000000000000000055534400000000000000000000000000000000000000000000000001", expectError: false},
		{name: "Token tiny value", jsonValue: `{"currency":"USD","value":"0.000000000000000000000000000000000000000000000000000000000000000000000000001","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectedHex: "c1c38d7ea4c6800000000000000000000000000055534400000000000000000000000000000000000000000000000001", expectError: false},

		// MPT Amounts
		{name: "MPT max value", jsonValue: `{"value":"9223372036854775807","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "607fffffffffffffff00002403c84a0a28e0190e208e982c352bbd5006600555cf", expectError: false},
		{name: "MPT max hex", jsonValue: `{"value":"0x7FFFFFFFFFFFFFFF","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "607fffffffffffffff00002403c84a0a28e0190e208e982c352bbd5006600555cf", expectError: false},
		{name: "MPT zero value", jsonValue: `{"value":"0","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF", expectError: false},
		{name: "MPT negative zero value", jsonValue: `{"value":"-0","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF", expectError: false},
		{name: "MPT hex value", jsonValue: `{"value":"0xa","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "60000000000000000a00002403c84a0a28e0190e208e982c352bbd5006600555cf", expectError: false},
		{name: "MPT scientific notation", jsonValue: `{"value":"1E1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectedHex: "60000000000000000a00002403c84a0a28e0190e208e982c352bbd5006600555cf", expectError: false},

		// Invalid Cases
		{name: "empty string", jsonValue: `""`, expectError: true},
		{name: "Invalid amount type", jsonValue: `[1]`, expectError: true},
		{name: "string", jsonValue: `"a"`, expectError: true},

		{name: "XRP too large", jsonValue: `"100000000000000001"`, expectError: true},
		{name: "XRP negative", jsonValue: `"-1"`, expectError: true},

		{name: "Token precision overflow", jsonValue: `{"currency":"USD","value":"1111111111111111.1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "Token invalid currency", jsonValue: `{"currency":"USDC","value":"1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "Token XRP currency", jsonValue: `{"currency":"XRP","value":"1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "Token wrong value type", jsonValue: `{"currency":"USD","value":10,"issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "Token invalid issuer", jsonValue: `{"currency":"USD","value":"10","issuer":"rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAee"}`, expectError: true},
		{name: "Token missing issuer", jsonValue: `{"currency":"USD","value":"10"}`, expectError: true},
		{name: "Token missing currency", jsonValue: `{"value":"10","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "Token missing currency and unknown field", jsonValue: `{"value":"10","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji", "invalidField": 10}`, expectError: true},
		{name: "Token additional field", jsonValue: `{"currency":"USD","value":"1","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji","invalidField": 10}`, expectError: true},
		{name: "Token wrong field", jsonValue: `{"currency":"USD","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji","invalidField": 10}`, expectError: true},

		{name: "MPT too large 0", jsonValue: `{"value":"9223372036854775808","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT too large 1", jsonValue: `{"value":"18446744073709551615","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true}, {name: "MPT negative", jsonValue: `{"value":"-1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid issuance id 0", jsonValue: `{"value":"10","mpt_issuance_id":"10"}`, expectError: true},
		{name: "MPT invalid issuance id 1", jsonValue: `{"value":"10","mpt_issuance_id":"xx"}`, expectError: true},
		{name: "MPT missing issuance id", jsonValue: `{"value":"10"}`, expectError: true},
		{name: "MPT missing value", jsonValue: `{"mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT missing value and invalid field", jsonValue: `{"mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF","invalidField": 10}`, expectError: true},
		{name: "MPT invalid value scientific to large", jsonValue: `{"value":"1E100","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value scientific to small", jsonValue: `{"value":"1e-100","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value scientific negative", jsonValue: `{"value":"-1e1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value scientific invalid", jsonValue: `{"value":"1e","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value empty", jsonValue: `{"value":"","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value string 0", jsonValue: `{"value":"a","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value string 1", jsonValue: `{"value":"?","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value hex to large", jsonValue: `{"value":"0xFFFFFFFFFFFFFFFF","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value hex negative 0", jsonValue: `{"value":"0x-1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value hex negative 1", jsonValue: `{"value":"-0x1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value decimal 0", jsonValue: `{"value":"10.0","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value decimal 1", jsonValue: `{"value":"10.1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value decimal negative", jsonValue: `{"value":"-1.1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value decimal zero 0", jsonValue: `{"value":"0.0","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},
		{name: "MPT invalid value decimal zero 1", jsonValue: `{"value":".0","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF"}`, expectError: true},

		{name: "MPT additional field 0 ", jsonValue: `{"value":"1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF","issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"}`, expectError: true},
		{name: "MPT additional field 1 ", jsonValue: `{"value":"1","mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF","currency":"USD"}`, expectError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var value any
			err := json.Unmarshal([]byte(test.jsonValue), &value)
			require.NoError(t, err)

			encoded, err := Amount.ToBytes(value, false)

			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				expectedBytes, err := hex.DecodeString(test.expectedHex)
				require.NoError(t, err)
				require.Equal(t, expectedBytes, encoded)
			}
		})
	}
}

func TestAmountDecoding(t *testing.T) {
	tests := []struct {
		name        string
		hexInput    string
		expectJSON  string
		expectError bool
	}{
		// Valid decoding
		{
			name:       "Token amount",
			hexInput:   "d55920ac9391400000000000000000000000000055534400000000000a20b3c85f482532a9578dbb3950b85ca06594d1",
			expectJSON: `{"currency":"USD","issuer":"rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B","value":"7072.8"}`,
		},
		{
			name:       "XRP amount",
			hexInput:   "4000000000989680",
			expectJSON: `"10000000"`,
		},
		{
			name:       "MPT amount",
			hexInput:   "607fffffffffffffff00002403c84a0a28e0190e208e982c352bbd5006600555cf",
			expectJSON: `{"mpt_issuance_id":"00002403C84A0A28E0190E208E982C352BBD5006600555CF","value":"9223372036854775807"}`,
		},

		// Invalid decoding
		{name: "Empty input", hexInput: "", expectError: true},
		{name: "Zero input", hexInput: "00", expectError: true},
		{name: "FF input", hexInput: "ff", expectError: true},
		{name: "More zero input", hexInput: "0000", expectError: true},
		{name: "XRP too large", hexInput: "4fffffffffffffff", expectError: true},
		{name: "XRP incomplete 0", hexInput: "40", expectError: true},
		{name: "XRP incomplete 1", hexInput: "4000", expectError: true},
		{name: "Token incomplete amount", hexInput: "d4564b", expectError: true},
		{name: "Token incomplete currency", hexInput: "d4564b964a845ac00000000000000000000000005553440000000000", expectError: true},
		{name: "Token incomplete issuer", hexInput: "d4564b964a845ac0000000000000000000000000555344000000000069d33b18d53385f8a3185516c2eda5dedb8ac5", expectError: true},
		{name: "MPT incomplete 0", hexInput: "60", expectError: true},
		{name: "MPT incomplete 1", hexInput: "6000", expectError: true},
		{name: "MPT incomplete 2", hexInput: "6000000000000000", expectError: true},
		{name: "MPT incomplete 3", hexInput: "600000000000000000", expectError: true},
		{name: "MPT incomplete 4", hexInput: "6000000000000000000000000000000000000000000000", expectError: true},
		{name: "MPT incomplete 5", hexInput: "608000000000000000000000000000000000000000000000000000000000000000", expectError: true},
		{name: "MPT too large", hexInput: "60800000000000000000002403c84a0a28e0190e208e982c352bbd5006600555cf", expectError: true},
		{name: "MPT too large again", hexInput: "60ffffffffffffffff000000000000000000000000000000000000000000000000", expectError: true},

		{name: "Invalid amount type bits", hexInput: "0100000000000000", expectError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputBytes, err := hex.DecodeString(test.hexInput)
			require.NoError(t, err)

			buffer := bytes.NewBuffer(inputBytes)
			decoded, err := Amount.ToJSON(buffer, 0)

			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Zero(t, buffer.Len())

				// Marshal decoded JSON to string for comparison
				decodedBytes, err := json.Marshal(decoded)
				require.NoError(t, err)

				// Unmarshal expected JSON to re-marshal, ensuring consistent key order and spacing
				var expectedObj any
				err = json.Unmarshal([]byte(test.expectJSON), &expectedObj)
				require.NoError(t, err)
				expectedBytes, err := json.Marshal(expectedObj)
				require.NoError(t, err)

				require.JSONEq(t, string(expectedBytes), string(decodedBytes))
			}
		})
	}
}

func TestTokenAmountRoundTrip(t *testing.T) {
	tests := []string{
		"0",
		"1",
		"-1",
		"0.1",
		"-0.1",
		"1e+95", // upper exponent bound
		"1e-80", // lower exponent bound
		"40915.87486543398",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			// Serialize
			serialized, err := serializeTokenValue(test)
			require.NoError(t, err)

			// Deserialize
			deserialized, err := deserializeTokenAmount(serialized)
			require.NoError(t, err)

			// Verify
			require.Equal(t, test, deserialized)
		})
	}
}

func TestTokenAmountSerializationFailure(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "Too precise", value: "1111111111111111.1"},
		{name: "Exponent too high", value: "1e97"},
		{name: "Exponent too low", value: "1e-97"},
		{name: "Not a number", value: "abc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := serializeTokenValue(test.value)
			require.Error(t, err)
		})
	}
}

func TestExtractString(t *testing.T) {
	testMap := map[string]any{
		"ok":    "ok",
		"empty": "",
		"int":   10,
		"bool":  false,
		"bytes": []byte("bytes"),
		"nil":   nil,
	}

	t.Run("Valid strings", func(t *testing.T) {
		extracted, err := extractString(testMap, "ok")
		require.NoError(t, err)
		require.Equal(t, "ok", extracted)

		extracted, err = extractString(testMap, "empty")
		require.NoError(t, err)
		require.Equal(t, "", extracted)
	})

	t.Run("Invalid extractions", func(t *testing.T) {
		failKeys := []string{"int", "bool", "bytes", "nil", "nonexistent"}
		for _, key := range failKeys {
			t.Run(key, func(t *testing.T) {
				_, err := extractString(testMap, key)
				require.Error(t, err)
			})
		}
	})
}

func TestMPTValueExotic(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    string
		expectedValue string
	}{
		{name: "Scientific notation E", inputValue: "1E1", expectedValue: "10"},
		{name: "Scientific notation e", inputValue: "1e1", expectedValue: "10"},
		{name: "Hex with 0x prefix", inputValue: "0x1", expectedValue: "1"},
		{name: "Hex with 0x0 prefix", inputValue: "0x01", expectedValue: "1"},
		{name: "Zero", inputValue: "0", expectedValue: "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			amount := map[string]any{
				"value":           test.inputValue,
				"mpt_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			}

			encoded, err := Amount.ToBytes(amount, false)
			require.NoError(t, err)

			buffer := bytes.NewBuffer(encoded)
			decoded, err := Amount.ToJSON(buffer, 0)
			require.NoError(t, err)

			decodedMap, ok := decoded.(map[string]any)
			require.True(t, ok)
			require.Equal(t, test.expectedValue, decodedMap["value"])
		})
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

		expected, err := hex.DecodeString(test.output)
		require.NoError(t, err)

		require.Equal(t, expected, normalized)
	}
}
