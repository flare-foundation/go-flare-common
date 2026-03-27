// Package transactions provides XRPL transaction construction and validation utilities.
package transactions

import (
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

const memoAllowedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~:/?#[]@!$&'()*+,;=%"

// CommonFields holds fields shared across XRPL transaction types.
type CommonFields struct {
	Account         string
	TransactionType string
	Fee             string
	Sequence        uint
	Memos           []Memo
}

// Memo represents an XRPL transaction memo with optional data, format, and type fields.
type Memo struct {
	MemoData   string
	MemoFormat string
	MemoType   string
}

// Validate checks that the memo fields contain valid characters.
// The MemoType and MemoFormat fields should only consist of the following characters:
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~:/?#[]@!$&'()*+,;=%
func (m *Memo) Validate() bool {
	return validMemoField(m.MemoType) && validMemoField(m.MemoFormat)
}

// validMemoField reports whether all characters in s are in memoAllowedChars.
func validMemoField(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !strings.ContainsRune(memoAllowedChars, r)
	}) == -1
}

// Format returns the memo as a serializable array object.
func (m *Memo) Format() types.ArrayObject {
	inner := make(types.Object)

	if len(m.MemoData) > 0 {
		inner["MemoData"] = m.MemoData
	}
	if len(m.MemoFormat) > 0 {
		inner["MemoFormat"] = m.MemoFormat
	}
	if len(m.MemoType) > 0 {
		inner["MemoType"] = m.MemoType
	}

	return types.ArrayObject{
		"Memo": inner,
	}
}

// func commonFields()
