package types

import (
	"bytes"
	"fmt"
	"reflect"
)

// ArrayObject is a wrapped element of form map[string]map[string]any{"nameOfTheObject: Object"}
type ArrayObject = map[string]any

const arrayEnd byte = 0xf1

// STArray is used for serialization of Array Fields https://xrpl.org/docs/references/protocol/binary-format#array-fields
type stArray struct{}

var STArray = &stArray{}

// ToBytes serializes values of Array fields.
func (a *stArray) ToBytes(value any, signing bool) ([]byte, error) {
	array, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid array %v is of type %v ", value, reflect.TypeOf(value))
	}

	outBuff := bytes.NewBuffer(nil)

	var uniqueKeyCheck string

	for i := range array {
		arrayObj, ok := array[i].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid array object %v is of type %v ", array[i], reflect.TypeOf(array[i]))
		}

		if len(arrayObj) != 1 {
			return nil, fmt.Errorf("invalid object wrapper %v", array[i])
		}

		for key := range arrayObj {
			// only one wrapper key is allowed
			if uniqueKeyCheck != "" && key != uniqueKeyCheck {
				return nil, fmt.Errorf("invalid array multiple wrapper keys")
			} else if uniqueKeyCheck == "" {
				uniqueKeyCheck = key
			}

			bytes, err := encodeInner(key, arrayObj[key], signing)
			if err != nil {
				return nil, fmt.Errorf("encoding %v: %v", array[i], err)
			}

			_, err = outBuff.Write(bytes)
			if err != nil {
				return nil, fmt.Errorf("writing to buffer %v: %v", array[i], err)
			}
		}
	}

	err := outBuff.WriteByte(arrayEnd)
	if err != nil {
		return nil, fmt.Errorf("writing array end to %v: %v", value, err)
	}

	return outBuff.Bytes(), nil
}

// ToJson decodes an encoded array.
func (a *stArray) ToJson(b *bytes.Buffer, _ int) (any, error) {
	out := make([]any, 0)

	nameCheck := "" // to check that all elements have the same wrapper

	for {
		nextByte, err := b.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("reading next byte: %v", err)
		}
		if nextByte == arrayEnd {
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

		if nameCheck == "" {
			nameCheck = name
		} else if nameCheck != name {
			return nil, fmt.Errorf("different wrappers in same array, %s, %s", nameCheck, name)
		}

		wrapped := make(map[string]any)

		wrapped[name] = value
		out = append(out, wrapped)
	}

	return out, nil
}
