package types

import (
	"encoding/hex"
	"fmt"
)

type hash struct {
	length int
}

func (h *hash) ToBytes(value any, _ bool) ([]byte, error) {
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

var Blob = &hash{length: 0}
var Hash128 = &hash{length: 16}
var Hash160 = &hash{length: 20}
var Hash192 = &hash{length: 24}
var Hash256 = &hash{length: 32}
