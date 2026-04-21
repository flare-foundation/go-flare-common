package transactions

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func hexOf(s string) string {
	return hex.EncodeToString([]byte(s))
}

func TestMemoValidate(t *testing.T) {
	tests := []struct {
		name  string
		memo  Memo
		valid bool
	}{
		{name: "empty fields", memo: Memo{}, valid: true},
		{name: "alphanumeric MemoType", memo: Memo{MemoType: hexOf("application/json")}, valid: true},
		{name: "alphanumeric MemoFormat", memo: Memo{MemoFormat: hexOf("text/plain")}, valid: true},
		{name: "both fields valid", memo: Memo{MemoType: hexOf("application/json"), MemoFormat: hexOf("utf-8")}, valid: true},
		{name: "MemoData arbitrary bytes", memo: Memo{MemoData: "00ff"}, valid: true},
		{name: "space in MemoType", memo: Memo{MemoType: hexOf("hello world")}, valid: false},
		{name: "space in MemoFormat", memo: Memo{MemoFormat: hexOf("hello world")}, valid: false},
		{name: "newline in MemoType", memo: Memo{MemoType: hexOf("foo\nbar")}, valid: false},
		{name: "null byte in MemoFormat", memo: Memo{MemoFormat: hexOf("foo\x00bar")}, valid: false},
		{name: "non-ASCII in MemoType", memo: Memo{MemoType: hexOf("héllo")}, valid: false},
		{name: "MemoType valid, MemoFormat invalid", memo: Memo{MemoType: hexOf("valid"), MemoFormat: hexOf("invalid char \x01")}, valid: false},
		{name: "MemoType invalid, MemoFormat valid", memo: Memo{MemoType: hexOf("bad char \x01"), MemoFormat: hexOf("valid")}, valid: false},
		{name: "MemoType not valid hex", memo: Memo{MemoType: "zz"}, valid: false},
		{name: "MemoData not valid hex", memo: Memo{MemoData: "xx"}, valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, tt.memo.Validate())
		})
	}
}

// source: rippled src/libxrpl/protocol/STTx.cpp:655-679 isMemoOkay allowed-symbols
func TestValidateMemoAllowedSet(t *testing.T) {
	allowed := "0123456789" +
		"-._~:/?#[]@!$&'()*+,;=%" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz"

	assert.True(t, ValidateMemo([]byte(allowed)))
	assert.True(t, ValidateMemo(nil))
	assert.True(t, ValidateMemo([]byte{}))

	for _, bad := range []byte{0x00, 0x01, ' ', '\n', '\t', 0x7f, 0x80, 0xff, '"', '<', '>', '\\', '^', '`', '{', '}'} {
		assert.False(t, ValidateMemo([]byte{bad}), "byte %#x should be rejected", bad)
	}
}

func TestMemoFormat(t *testing.T) {
	cases := []struct {
		name string
		memo Memo
		keys []string
	}{
		{
			name: "all fields",
			memo: Memo{MemoData: "AA", MemoFormat: "BB", MemoType: "CC"},
			keys: []string{"MemoData", "MemoFormat", "MemoType"},
		},
		{
			name: "only MemoData",
			memo: Memo{MemoData: "AA"},
			keys: []string{"MemoData"},
		},
		{
			name: "empty",
			memo: Memo{},
			keys: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := c.memo.Format()
			inner, ok := out["Memo"].(map[string]any)
			require.True(t, ok)
			require.Len(t, inner, len(c.keys))
			for _, k := range c.keys {
				_, ok := inner[k]
				assert.True(t, ok, "missing key %s", k)
			}
		})
	}
}
