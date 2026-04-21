// Package transactions provides XRPL transaction construction and validation utilities.
package transactions

import (
	"encoding/hex"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

var memoAllowedBytes = func() [256]bool {
	var a [256]bool
	const symbols = "0123456789" +
		"-._~:/?#[]@!$&'()*+,;=%" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz"
	for i := range len(symbols) {
		a[symbols[i]] = true
	}
	return a
}()

// CommonFields holds fields shared across XRPL transaction types.
type CommonFields struct {
	Account         string
	TransactionType string
	Fee             string
	Sequence        uint
	Memos           []Memo
}

// Memo represents an XRPL transaction memo. MemoData, MemoFormat and MemoType hold hex-encoded bytes.
type Memo struct {
	MemoData   string
	MemoFormat string
	MemoType   string
}

// ValidateMemo reports whether every byte in b is a URL character allowed in MemoType / MemoFormat per RFC 3986.
// Mirrors rippled isMemoOkay (src/libxrpl/protocol/STTx.cpp:655-679): the check runs on the hex-decoded bytes.
func ValidateMemo(b []byte) bool {
	for _, c := range b {
		if !memoAllowedBytes[c] {
			return false
		}
	}
	return true
}

// Validate reports whether the memo fields are well-formed per rippled isMemoOkay:
// every field must be valid hex, and the hex-decoded MemoType / MemoFormat must contain only RFC 3986 URL characters.
// MemoData is only required to be valid hex.
func (m *Memo) Validate() bool {
	if _, err := hex.DecodeString(m.MemoData); err != nil {
		return false
	}
	typeBytes, err := hex.DecodeString(m.MemoType)
	if err != nil {
		return false
	}
	formatBytes, err := hex.DecodeString(m.MemoFormat)
	if err != nil {
		return false
	}
	return ValidateMemo(typeBytes) && ValidateMemo(formatBytes)
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
