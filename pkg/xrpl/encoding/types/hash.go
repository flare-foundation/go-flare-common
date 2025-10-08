package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

type hash struct {
	length int
}

// ToBytes serializes hexadecimal strings.
func (h *hash) ToBytes(value any, _ bool) ([]byte, error) {
	s, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("not of type string %v", value)
	}

	v, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("not hex string %v, %v", value, err)
	}
	if h.length != 0 && h.length != len(v) {
		return nil, fmt.Errorf("wrong length, expected %d bytes", h.length)
	}
	return v, nil
}

// ToJson decodes hash to hexadecimal string.
//
// If length is fixed for type, the parameter is shadowed.
func (h *hash) ToJSON(b *bytes.Buffer, length int) (any, error) {
	l := h.length
	if l == 0 {
		l = length
		if l < 0 {
			return nil, fmt.Errorf("invalid length %d", l)
		}
	}

	value := make([]byte, l)

	n, err := b.Read(value)
	if err != nil {
		return nil, fmt.Errorf("cannot read hash from buffer: %v", err)
	}
	if n != l {
		return nil, outOfBytes("hash", l, n)
	}

	return strings.ToUpper(hex.EncodeToString(value)), nil
}

// Blob is used for serialization of Blob fields. https://xrpl.org/docs/references/protocol/binary-format#blob-fields
var Blob = &hash{length: 0}

// Hash128 is used for serialization of Hash128 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash128 = &hash{length: 16}

// Hash160 is used for serialization of Hash160 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash160 = &hash{length: 20}

// Hash192 is used for serialization of Hash192 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash192 = &hash{length: 24}

// Hash256 is used for serialization of Hash256 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash256 = &hash{length: 32}
