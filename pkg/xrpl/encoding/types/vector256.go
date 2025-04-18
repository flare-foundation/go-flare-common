package types

import (
	"bytes"
	"fmt"
)

type vector256 struct{}

var Vector256 = &vector256{}

func (v *vector256) ToBytes(value any, _ bool) ([]byte, error) {
	array, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid vector256: %v", value)
	}

	out := make([]byte, 0, len(array)*32)

	for j := range array {
		hash256, err := Hash256.ToBytes(array[j], false)
		if err != nil {
			return nil, fmt.Errorf("decoding entry:% v", err)
		}
		out = append(out, hash256...)
	}

	return out, nil
}

func (v *vector256) ToJson(b *bytes.Buffer, l int) (any, error) {
	if l%32 != 0 {
		return nil, fmt.Errorf("invalid length of encoded vector256 %d", l)
	}

	arrayLen := l / 32

	out := make([]any, arrayLen)

	for j := range arrayLen {
		hash, err := Hash256.ToJson(b, 0)
		if err != nil {
			return nil, fmt.Errorf("reading entry %d of %d long vector256: %v", j, l, err)
		}
		out[j] = hash
	}

	return out, nil
}
