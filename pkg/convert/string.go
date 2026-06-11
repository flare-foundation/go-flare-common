package convert

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// StringToCommonHash returns Solidity's bytes32(s) ([]byte(s) appended with zeros to length 32)
// String s can be at most 32 bytes long, otherwise an error is returned.
func StringToCommonHash(s string) (common.Hash, error) {
	if len(s) > 32 {
		return common.Hash{}, errors.New("string too long: at most 32 bytes allowed")
	}
	x := [32]byte{}
	copy(x[:], s)

	return x, nil
}

// CommonHashToString trims trailing zero bytes from h and returns the rest as a
// string of raw bytes, without UTF-8 validation. It inverts [StringToCommonHash]
// for strings that do not end with a zero byte.
func CommonHashToString(h common.Hash) string {
	return strings.TrimRight(string(h[:]), "\x00")
}

// Hex32StringToCommonHash decodes a 32-byte hex string (with at most one optional
// 0x or 0X prefix) into a common.Hash.
func Hex32StringToCommonHash(s string) (common.Hash, error) {
	st := removeHexPrefix(s)

	b, err := hex.DecodeString(st)
	if err != nil {
		return common.Hash{}, fmt.Errorf("invalid hex string %s: %w", truncateForError(s), err)
	}

	if len(b) != 32 {
		return common.Hash{}, fmt.Errorf("invalid length for hex string %s: expected 32 bytes, got %d", truncateForError(s), len(b))
	}

	var h common.Hash
	copy(h[:], b)

	return h, nil
}

// removeHexPrefix strips at most one 0x or 0X prefix.
func removeHexPrefix(s string) string {
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}

	return s
}

// truncateForError bounds s for inclusion in an error message.
func truncateForError(s string) string {
	const limit = 80
	if len(s) <= limit {
		return s
	}

	return fmt.Sprintf("%s... (%d bytes total)", s[:limit], len(s))
}
