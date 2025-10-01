package types

import (
	"bytes"
	"fmt"
)

const objectEnd byte = 0xe1

// STObject is used for serialization of Object Fields. https://xrpl.org/docs/references/protocol/binary-format#object-fields.
type STObject struct{}

// ToBytes serializes value of Object Fields.
func (*STObject) ToBytes(value any, signing bool) ([]byte, error) {
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

func (*STObject) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	out := make(map[string]any)

	for {
		nextByte, err := b.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("reading next byte: %v", err)
		}

		if nextByte == objectEnd {
			break
		}

		err = b.UnreadByte()
		if err != nil {
			return nil, fmt.Errorf("unreading next byte: %v", err)
		}

		name, value, err := decodeNext(b)
		if err != nil {
			return nil, fmt.Errorf("decoding next: %v", err)
		}

		out[name] = value
	}

	return out, nil
}
