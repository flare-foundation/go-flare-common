package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
