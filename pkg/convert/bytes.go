package convert

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type SliceTooLongError struct {
	maxLen int
}

func (e SliceTooLongError) Error() string {
	return fmt.Sprintf("slice longer than %d", e.maxLen)
}

// BytesToUint64 converts a byte slice to uint64 with big-endian encoding.
// Nil or zero length slice is converted to 0.
// The returns error if the slice is longer than 8 bytes.
func BytesToUint64(b []byte) (uint64, error) {
	if len(b) > 8 {
		return 0, SliceTooLongError{8}
	}
	if len(b) < 8 {
		// Pad with zeros if less than 8 bytes
		padded := make([]byte, 8)
		copy(padded[8-len(b):], b)
		return binary.BigEndian.Uint64(padded), nil
	}
	return binary.BigEndian.Uint64(b), nil
}

// BytesToUint32 converts a byte slice to uint32 with big-endian encoding.
// Nil or zero length slice is converted to 0.
// The returns error if the slice is longer than 4 bytes.
func BytesToUint32(b []byte) (uint32, error) {
	if len(b) > 4 {
		return 0, SliceTooLongError{4}
	}
	if len(b) < 4 {
		// Pad with zeros if less than 4 bytes
		padded := make([]byte, 4)
		copy(padded[4-len(b):], b)
		return binary.BigEndian.Uint32(padded), nil
	}
	return binary.BigEndian.Uint32(b), nil
}

// BytesToUint16 converts a byte slice to uint16 with big-endian encoding.
// Nil or zero length slice is converted to 0.
// The returns error if the slice is longer than 2 bytes.
func BytesToUint16(b []byte) (uint16, error) {
	if len(b) > 2 {
		return 0, SliceTooLongError{2}
	}
	if len(b) < 2 {
		// Pad with zeros if less than 2 bytes
		padded := make([]byte, 2)
		copy(padded[2-len(b):], b)
		return binary.BigEndian.Uint16(padded), nil
	}
	return binary.BigEndian.Uint16(b), nil
}

// Uint64ToHash hex encodes uint64 and zero prefixes to 32 bytes.
func Uint64ToHash(n uint64) common.Hash {
	h := common.Hash{}
	binary.BigEndian.PutUint64(h[24:], n)
	return h
}

// Uint64ToBytes hex encodes uint64.
func Uint64ToBytes(n uint64) []byte {
	h := make([]byte, 8)
	binary.BigEndian.PutUint64(h, n)
	return h
}

// Uint32ToHash hex encodes uint32 and zero prefixes to 32 bytes.
func Uint32ToHash(n uint32) common.Hash {
	h := common.Hash{}
	binary.BigEndian.PutUint32(h[28:], n)
	return h
}

// Uin32ToBytes hex encodes uint32.
func Uint32ToBytes(n uint32) []byte {
	h := make([]byte, 4)
	binary.BigEndian.PutUint32(h, n)
	return h
}

// Uint16ToHash hex encodes uint16 and zero prefixes to 32 bytes.
func Uint16ToHash(n uint16) common.Hash {
	h := common.Hash{}
	binary.BigEndian.PutUint16(h[30:], n)
	return h
}

// Uint16ToBytes hex encodes uint16.
func Uint16ToBytes(n uint16) []byte {
	h := make([]byte, 2)
	binary.BigEndian.PutUint16(h, n)
	return h
}
