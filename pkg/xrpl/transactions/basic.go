// Package transactions provides XRPL transaction construction and validation utilities.
package transactions

import "github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"

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
func (m Memo) Validate() bool {
	return true // TODO check characters in MemoFormat and MemoType
}

// Format returns the memo as a serializable array object.
func (m Memo) Format() types.ArrayObject {
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
