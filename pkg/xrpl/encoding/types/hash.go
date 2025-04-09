package types

import (
	"encoding/hex"
	"fmt"
)

type hashInternal struct {
	length int
}

// ToBytes serializes byte strings.
func (h *hashInternal) ToBytes(value any, _ bool) ([]byte, error) {
	s, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("not of type string %v", value)
	}

	v, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if h.length != 0 && h.length != len(v) {
		return nil, fmt.Errorf("wrong length, expected %d bytes", h.length)
	}
	return v, nil
}

// Blob is used for serialization of Blob fields. https://xrpl.org/docs/references/protocol/binary-format#blob-fields
var Blob = &hashInternal{length: 0}

// Hash128 is used for serialization of Hash128 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash128 = &hashInternal{length: 16}

// Hash160 is used for serialization of Hash160 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash160 = &hashInternal{length: 20}

// Hash192 is used for serialization of Hash193 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash192 = &hashInternal{length: 24}

// Hash256 is used for serialization of Hash256 fields. https://xrpl.org/docs/references/protocol/binary-format#hash-fields
var Hash256 = &hashInternal{length: 32}
