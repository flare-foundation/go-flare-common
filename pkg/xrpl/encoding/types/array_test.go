package types

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayToBytes(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		expectedHex string
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid array with one memo",
			input: []any{
				map[string]any{
					"Memo": map[string]any{
						"MemoType": "01",
						"MemoData": "0102",
					},
				},
			},
			expectedHex: "ea7c01017d020102e1f1", // Memo field ID (ea) + MemoType field ID (7c) + length (01) + data (01) + MemoData field ID (7d) + length (02) + data (0102) + ObjectEnd (e1) + ArrayEnd (f1)

			expectError: false,
		},
		{
			name: "Valid array with two memos",
			input: []any{
				map[string]any{
					"Memo": map[string]any{
						"MemoData": "31",
					},
				},
				map[string]any{
					"Memo": map[string]any{
						"MemoData": "32",
					},
				},
			},
			expectedHex: "ea7d0131e1ea7d0132e1f1", // first memo (memo field id (ea) + MemoType ID (7c) + length (01) + data (31) + objectEnd (e1)) + seconde memo (memo field id (ea) + MemoType ID (7c) + length (01) + data (32) + objectEnd (e1)) + ArrayEnd (f1)
			expectError: false,
		},
		{
			name:        "Valid empty array",
			input:       []any{},
			expectedHex: "f1", // Just the array end marker
			expectError: false,
		},
		{
			name:        "Input is not an array",
			input:       "not an array",
			expectError: true,
			errorMsg:    "invalid array",
		},
		{
			name: "Array element is not an object wrapper",
			input: []any{
				"not an object",
			},
			expectError: true,
			errorMsg:    "invalid array object",
		},
		{
			name: "Object wrapper has zero keys",
			input: []any{
				map[string]any{},
			},
			expectError: true,
			errorMsg:    "invalid object wrapper",
		},
		{
			name: "Object wrapper has more than one key",
			input: []any{
				map[string]any{
					"Memo":    map[string]any{},
					"Account": "r...",
				},
			},
			expectError: true,
			errorMsg:    "invalid object wrapper",
		},
		{
			name: "Encoding inner object fails",
			input: []any{
				map[string]any{
					"UnknownField": map[string]any{},
				},
			},
			expectError: true,
			errorMsg:    "unknown field name",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := STArray.ToBytes(test.input, false)

			if test.expectError {
				require.Error(t, err)
				if test.errorMsg != "" {
					require.Contains(t, err.Error(), test.errorMsg)
				}
			} else {
				require.NoError(t, err)
				expected, err := hex.DecodeString(test.expectedHex)
				require.NoError(t, err)
				require.Equal(t, expected, result)

				b := bytes.NewBuffer(result)
				obj, err := STArray.ToJSON(b, 0)
				require.NoError(t, err)

				resultAgain, err := STArray.ToBytes(obj, false)
				require.NoError(t, err)

				require.Equal(t, result, resultAgain)
			}
		})
	}
}

func TestArrayToJSONErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name: "empty", input: "",
		},
		{
			name: "zero", input: "0000",
		},
		{
			name: "gibberish", input: "0100aa1256",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h, err := hex.DecodeString(test.input)
			require.NoError(t, err)

			b := bytes.NewBuffer(h)

			_, err = STArray.ToJSON(b, 0)
			require.Error(t, err)
		})
	}
}
