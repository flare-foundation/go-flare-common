package types

import (
	"fmt"
)

const objectEnd byte = 0xe1

// STObject is used for serialization of Object Fields. https://xrpl.org/docs/references/protocol/binary-format#object-fields
type STObject struct{}

// ToBytes serializes value of Object Fields
func (o *STObject) ToBytes(value any, signing bool) ([]byte, error) {
	valuerObj, ok := value.(Object)
	if !ok {
		return nil, fmt.Errorf("value %v is not an object", value)
	}

	bytes, err := Encode(valuerObj, signing)
	if err != nil {
		return nil, fmt.Errorf("encoding object %v: %v", value, err)
	}

	bytes = append(bytes, objectEnd)

	return bytes, nil
}
