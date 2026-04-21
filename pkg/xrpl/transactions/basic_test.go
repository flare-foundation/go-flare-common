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
		{
			name:  "empty fields",
			memo:  Memo{},
			valid: true,
		},
		{
			name:  "all allowed characters",
			memo:  Memo{MemoType: memoAllowedChars},
			valid: true,
		},
		{
			name:  "alphanumeric MemoType",
			memo:  Memo{MemoType: "application/json"},
			valid: true,
		},
		{
			name:  "alphanumeric MemoFormat",
			memo:  Memo{MemoFormat: "text/plain"},
			valid: true,
		},
		{
			name:  "both fields valid",
			memo:  Memo{MemoType: "application/json", MemoFormat: "utf-8"},
			valid: true,
		},
		{
			name:  "MemoData not checked",
			memo:  Memo{MemoData: "\x00\xff"},
			valid: true,
		},
		{
			name:  "space in MemoType",
			memo:  Memo{MemoType: "hello world"},
			valid: false,
		},
		{
			name:  "space in MemoFormat",
			memo:  Memo{MemoFormat: "hello world"},
			valid: false,
		},
		{
			name:  "newline in MemoType",
			memo:  Memo{MemoType: "foo\nbar"},
			valid: false,
		},
		{
			name:  "null byte in MemoFormat",
			memo:  Memo{MemoFormat: "foo\x00bar"},
			valid: false,
		},
		{
			name:  "non-ASCII in MemoType",
			memo:  Memo{MemoType: "héllo"},
			valid: false,
		},
		{
			name:  "MemoType valid but MemoFormat invalid",
			memo:  Memo{MemoType: "valid", MemoFormat: "invalid char \x01"},
			valid: false,
		},
		{
			name:  "MemoType invalid but MemoFormat valid",
			memo:  Memo{MemoType: "bad char \x01", MemoFormat: "valid"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, tt.memo.Validate())
		})
	}
}

// Source: rippled src/libxrpl/protocol/STTx.cpp isMemoOkay allowed symbols (lines 655-667).
// Verifies that the Go memoAllowedChars constant matches the rippled RFC 3986 URL-safe subset exactly.
func TestMemoAllowedCharsMatchesRippled(t *testing.T) {
	rippledSymbols := "0123456789" +
		"-._~:/?#[]@!$&'()*+,;=%" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz"

	for _, c := range rippledSymbols {
		assert.Contains(t, memoAllowedChars, string(c), "rippled allowed char %q missing from Go set", c)
	}
	require.Equal(t, len(rippledSymbols), len(memoAllowedChars))
}

// Deviation note (not a failure): rippled's isMemoOkay validates the *hex-decoded* bytes
// of MemoType / MemoFormat against allowedSymbols (STTx.cpp lines 669-679), while Go's
// validMemoField validates the raw Go string before any hex decoding. The character *set*
// is identical but the validation domain differs.
func TestMemoValidateOperatesOnRawString(t *testing.T) {
	// "deadbeef" is valid hex but contains only letters a-f which are in memoAllowedChars.
	m := Memo{MemoType: "deadbeef"}
	assert.True(t, m.Validate())

	// "hi there" decoded would fail in rippled too, but here we show Go rejects
	// due to the space regardless of hex-decode semantics.
	m2 := Memo{MemoType: "hi there"}
	assert.False(t, m2.Validate())
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
