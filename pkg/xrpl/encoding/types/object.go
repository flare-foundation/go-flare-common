package types

import (
	"bytes"
	"errors"
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
		return nil, fmt.Errorf("encoding object %v: %w", value, err)
	}

	bytes = append(bytes, objectEnd)

	return bytes, nil
}

// ToJSON deserializes Object Field. Equivalent to ToJSONDepth at depth 0;
// prefer ToJSONDepth in a recursive context so the codec can bound nesting.
func (s *STObject) ToJSON(b *bytes.Buffer, length int) (any, error) {
	return s.ToJSONDepth(b, length, 0)
}

// ToJSONDepth is the depth-tracked variant of ToJSON; decodeNext calls this
// to enforce the maxDecodeDepth bound on nested STObject/STArray.
func (*STObject) ToJSONDepth(b *bytes.Buffer, _ int, depth int) (any, error) {
	out := make(map[string]any)

	empty := true

	for {
		nextByte, err := b.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("reading next byte: %w", err)
		}

		if nextByte == objectEnd {
			break
		}

		err = b.UnreadByte()
		if err != nil {
			return nil, fmt.Errorf("unreading next byte: %w", err)
		}

		name, value, err := decodeNext(b, depth)
		if err != nil {
			return nil, fmt.Errorf("decoding next: %w", err)
		}

		out[name] = value

		empty = false
	}

	if empty {
		return nil, errors.New("empty object")
	}

	return out, nil
}
