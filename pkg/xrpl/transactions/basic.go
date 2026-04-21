// Package transactions provides XRPL transaction construction and validation utilities.
package transactions

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Memo represents an XRPL transaction memo. The XRPL binary codec and JSON wire format carry each field as a hex string; in Go we keep them as bytes.
type Memo struct {
	MemoData   hexutil.Bytes
	MemoFormat hexutil.Bytes
	MemoType   hexutil.Bytes
}

// ValidateMemo reports whether every byte in b is a URL character allowed in MemoType / MemoFormat per RFC 3986.
// Mirrors rippled isMemoOkay (src/libxrpl/protocol/STTx.cpp:655-679).
func ValidateMemo(b []byte) bool {
	for _, c := range b {
		if !memoAllowedBytes[c] {
			return false
		}
	}
	return true
}

// Validate reports whether MemoType and MemoFormat contain only RFC 3986 URL characters. MemoData is unconstrained.
func (m *Memo) Validate() bool {
	return ValidateMemo(m.MemoType) && ValidateMemo(m.MemoFormat)
}

// Format returns the memo as a serializable array object with fields hex-encoded as XRPL JSON expects.
func (m *Memo) Format() types.ArrayObject {
	inner := make(types.Object)

	if len(m.MemoData) > 0 {
		inner["MemoData"] = hex.EncodeToString(m.MemoData)
	}
	if len(m.MemoFormat) > 0 {
		inner["MemoFormat"] = hex.EncodeToString(m.MemoFormat)
	}
	if len(m.MemoType) > 0 {
		inner["MemoType"] = hex.EncodeToString(m.MemoType)
	}

	return types.ArrayObject{
		"Memo": inner,
	}
}

// func commonFields()
