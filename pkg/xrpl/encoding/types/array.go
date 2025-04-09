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
type STArray struct{}

// ToBytes serializes values of Array fields.
func (a *STArray) ToBytes(value any, signing bool) ([]byte, error) {
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
