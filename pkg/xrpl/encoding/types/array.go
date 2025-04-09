package types

import (
	"bytes"
	"fmt"
	"reflect"
)

// todo do we allow different objects ??

type ArrayObject = map[string]map[string]any

const arrayEnd byte = 0xf1

type STArray struct{}

func (a *STArray) ToBytes(value any, signing bool) ([]byte, error) {
	array, ok := value.([]ArrayObject)
	if !ok {
		return nil, fmt.Errorf("invalid array %v is of type %v ", value, reflect.TypeOf(value))
	}

	outBuff := bytes.NewBuffer(nil)

	for i := range array {
		if len(array[i]) != 1 {
			return nil, fmt.Errorf("invalid object wrapper %v", array[i])
		}

		for key := range array[i] {
			bytes, err := encodeInner(key, array[i][key], signing)
			if err != nil {
				return nil, fmt.Errorf("encoding %v: %v", array[i], err)
			}

			_, err = outBuff.Write(bytes)
			if err != nil {
				return nil, fmt.Errorf("encoding %v: %v", array[i], err)
			}
		}
	}

	err := outBuff.WriteByte(arrayEnd)
	if err != nil {
		return nil, fmt.Errorf("writing end to %v: %v", value, err)
	}

	return outBuff.Bytes(), nil
}
