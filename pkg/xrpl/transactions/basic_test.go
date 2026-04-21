package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoValidate(t *testing.T) {
	tests := []struct {
		name  string
		memo  Memo
		valid bool
	}{
		{name: "empty fields", memo: Memo{}, valid: true},
		{name: "alphanumeric MemoType", memo: Memo{MemoType: []byte("application/json")}, valid: true},
		{name: "alphanumeric MemoFormat", memo: Memo{MemoFormat: []byte("text/plain")}, valid: true},
		{name: "both fields valid", memo: Memo{MemoType: []byte("application/json"), MemoFormat: []byte("utf-8")}, valid: true},
		{name: "MemoData arbitrary bytes", memo: Memo{MemoData: []byte{0x00, 0xff}}, valid: true},
		{name: "space in MemoType", memo: Memo{MemoType: []byte("hello world")}, valid: false},
		{name: "space in MemoFormat", memo: Memo{MemoFormat: []byte("hello world")}, valid: false},
		{name: "newline in MemoType", memo: Memo{MemoType: []byte("foo\nbar")}, valid: false},
		{name: "null byte in MemoFormat", memo: Memo{MemoFormat: []byte("foo\x00bar")}, valid: false},
		{name: "non-ASCII in MemoType", memo: Memo{MemoType: []byte("héllo")}, valid: false},
		{name: "MemoType valid, MemoFormat invalid", memo: Memo{MemoType: []byte("valid"), MemoFormat: []byte("invalid char \x01")}, valid: false},
		{name: "MemoType invalid, MemoFormat valid", memo: Memo{MemoType: []byte("bad char \x01"), MemoFormat: []byte("valid")}, valid: false},
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
		name    string
		memo    Memo
		wantHex map[string]string
	}{
		{
			name:    "all fields",
			memo:    Memo{MemoData: []byte{0xAA}, MemoFormat: []byte{0xBB}, MemoType: []byte{0xCC}},
			wantHex: map[string]string{"MemoData": "aa", "MemoFormat": "bb", "MemoType": "cc"},
		},
		{
			name:    "only MemoData",
			memo:    Memo{MemoData: []byte{0xAA}},
			wantHex: map[string]string{"MemoData": "aa"},
		},
		{
			name:    "empty",
			memo:    Memo{},
			wantHex: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := c.memo.Format()
			inner, ok := out["Memo"].(map[string]any)
			require.True(t, ok)
			require.Len(t, inner, len(c.wantHex))
			for k, v := range c.wantHex {
				assert.Equal(t, v, inner[k], "key %s", k)
			}
		})
	}
}
