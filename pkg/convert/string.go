package convert

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// StringToCommonHash returns Solidity's bytes32(s) ([]byte(s) appended with zeros to length 32)
// String s can be at most 32 characters long, otherwise an error is returned.
func StringToCommonHash(s string) (common.Hash, error) {
	if len(s) > 32 {
		return common.Hash{}, errors.New("string too long. At most 32 characters allowed")
	}
	x := [32]byte{}
	copy(x[:], s)

	return x, nil
}

// CommonHashToString is the inverse of StringToCommonHash. It trims trailing zero bytes and converts the rest to string using utf-8.
func CommonHashToString(h common.Hash) string {
	return strings.TrimRight(string(h[:]), "\x00")
}

func Hex32StringToCommonHash(s string) (common.Hash, error) {
	st := removeHexPrefix(s)

	b, err := hex.DecodeString(st)
	if err != nil {
		return common.Hash{}, fmt.Errorf("invalid hex string %s: %w", s, err)
	}

	if len(b) != 32 {
		return common.Hash{}, fmt.Errorf("invalid length for hex string %s: expected 32 bytes, got %d", s, len(b))
	}

	var h common.Hash
	copy(h[:], b)

	return h, nil
}

func removeHexPrefix(s string) string {
	return strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")
}
