package types

import (
	"bytes"
	"fmt"
)

type vector256 struct{}

var Vector256 = &vector256{}

func (*vector256) ToBytes(value any, _ bool) ([]byte, error) {
	array, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid vector256: %v", value)
	}

	out := make([]byte, 0, len(array)*32)

	for j := range array {
		hash256, err := Hash256.ToBytes(array[j], false)
		if err != nil {
			return nil, fmt.Errorf("decoding entry: %w", err)
		}
		out = append(out, hash256...)
	}

	return out, nil
}

func (*vector256) ToJSON(b *bytes.Buffer, l int) (any, error) {
	if l%32 != 0 {
		return nil, fmt.Errorf("invalid length of encoded vector256 %d", l)
	}

	arrayLen := l / 32

	out := make([]any, arrayLen)

	for j := range arrayLen {
		hash, err := Hash256.ToJSON(b, 0)
		if err != nil {
			return nil, fmt.Errorf("reading entry %d of %d long vector256: %w", j, l, err)
		}
		out[j] = hash
	}

	return out, nil
}
