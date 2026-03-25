package types

import (
	"bytes"
	"fmt"
	"reflect"
)

// ArrayObject is a wrapped element of form map[string]map[string]any{"nameOfTheObject: Object"}.
type ArrayObject = map[string]any

const arrayEnd byte = 0xf1

// STArray is used for serialization of Array Fields https://xrpl.org/docs/references/protocol/binary-format#array-fields
type stArray struct{}

var STArray = &stArray{}

// ToBytes serializes values of Array fields.
func (*stArray) ToBytes(value any, signing bool) ([]byte, error) {
	array, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid array %v is of type %v ", value, reflect.TypeOf(value))
	}

	outBuff := bytes.NewBuffer(nil)

	for i := range array {
		arrayObj, ok := array[i].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid array object %v is of type %v ", array[i], reflect.TypeOf(array[i]))
		}

		if len(arrayObj) != 1 {
			return nil, fmt.Errorf("invalid object wrapper %v", array[i])
		}

		for key := range arrayObj {
			bytes, err := encodeInner(key, arrayObj[key], signing)
			if err != nil {
				return nil, fmt.Errorf("encoding %v: %w", array[i], err)
			}

			_, err = outBuff.Write(bytes)
			if err != nil {
				return nil, fmt.Errorf("writing to buffer %v: %w", array[i], err)
			}
		}
	}

	err := outBuff.WriteByte(arrayEnd)
	if err != nil {
		return nil, fmt.Errorf("writing array end to %v: %w", value, err)
	}

	return outBuff.Bytes(), nil
}

// ToJSON decodes an encoded array.
func (*stArray) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	out := make([]any, 0, 2)

	for {
		nextByte, err := b.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("reading next byte: %w", err)
		}
		if nextByte == arrayEnd {
			break
		}

		err = b.UnreadByte()
		if err != nil {
			return nil, fmt.Errorf("unreading next byte: %w", err)
		}

		name, value, err := decodeNext(b)
		if err != nil {
			return nil, fmt.Errorf("decoding next: %w", err)
		}

		wrapped := make(map[string]any)

		wrapped[name] = value
		out = append(out, wrapped)
	}

	return out, nil
}
