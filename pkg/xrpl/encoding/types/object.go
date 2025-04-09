package types

import (
	"fmt"
)

const objectEnd byte = 0xe1

type STObject struct{}

func (o *STObject) ToBytes(value any, signing bool) ([]byte, error) {
	valuerObj, ok := value.(Object)
	if !ok {
		return nil, fmt.Errorf("value %v is ont an object", value)
	}

	bytes, err := Encode(valuerObj, signing)

	bytes = append(bytes, objectEnd)
	if err != nil {
		return nil, fmt.Errorf("encoding object %v: %v", value, err)
	}

	return bytes, nil
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
