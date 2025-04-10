package types

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/defs"
)

type TODOError struct {
	xt defs.XType
}

func (e *TODOError) Error() string {
	return fmt.Sprintf("To be implemented %v", e.xt)
}

type UnsupportedError struct {
	xt defs.XType
}

func (e *UnsupportedError) Error() string {
	return fmt.Sprintf("Unsupported type %v", e.xt)
}

type IllegalError struct {
	xt defs.XType
}

func (e *IllegalError) Error() string {
	return fmt.Sprintf("Illegal type %v", e.xt)
}

type Encoder interface {
	ToBytes(value any, signing bool) ([]byte, error)
}

func typeEncoder(xt defs.XType) (Encoder, error) {
	switch xt {
	case defs.Hash160:
		return Hash160, nil
	case defs.Hash192:
		return nil, &UnsupportedError{xt}
	case defs.Transaction:
		return nil, &UnsupportedError{xt}
	case defs.UInt512:
		return nil, &UnsupportedError{xt}
	case defs.UInt8:
		return &UInt8{}, nil
	case defs.XChainBridge:
		return &XChainBridge{}, nil
	case defs.Vector256:
		return nil, &UnsupportedError{xt}
	case defs.AccountID:
		return AccountID, nil
	case defs.Currency:
		return nil, &TODOError{xt}
	case defs.Hash128:
		return Hash128, nil
	case defs.Metadata:
		return nil, &UnsupportedError{xt}
	case defs.PathSet:
		return &PathSet{}, nil
	case defs.UInt384:
		return nil, &UnsupportedError{xt}
	case defs.Blob:
		return Blob, nil
	case defs.NotPresent:
		return nil, &IllegalError{xt}
	case defs.UInt32:
		return &UInt32{}, nil
	case defs.Unknown:
		return nil, &IllegalError{xt}
	case defs.Amount:
		return Amount, nil
	case defs.Done:
		return nil, &IllegalError{xt}
	case defs.Hash256:
		return Hash256, nil
	case defs.Issue:
		return &Issue{}, nil
	case defs.LedgerEntry:
		return nil, &UnsupportedError{xt}
	case defs.Number:
		return nil, &UnsupportedError{xt}
	case defs.STArray:
		return STArray, nil
	case defs.STObject:
		return &STObject{}, nil
	case defs.UInt16:
		return &UInt16{}, nil
	case defs.UInt64:
		return &UInt64{}, nil
	case defs.UInt96:
		return nil, &UnsupportedError{xt}
	case defs.Validation:
		return nil, &UnsupportedError{xt}
	default:
		return nil, fmt.Errorf("unknown type %v", xt)
	}
}

// lengthPrefix computes prefix for a bytes array of length n.
func lengthPrefix(n int) ([]byte, error) {
	var prefix []byte
	switch {
	case n < 0 || n > 918744:
		return nil, fmt.Errorf("unsupported Variable Length encoding: %d", n)
	case n <= 192:
		prefix = []byte{uint8(n)}
	case n <= 12480:
		n -= 193
		prefix = []byte{193 + uint8(n>>8), uint8(n)}
	case n <= 918744:
		n -= 12481
		prefix = []byte{241 + uint8(n>>16), uint8(n >> 8), uint8(n)}
	}

	return prefix, nil
}

// encodeInner encodes value according to definition that corresponds to its name.
// If the field is not serialized or the signing is false and the field is not a signing field, empty byte slice is returned.
func encodeInner(name string, value any, signing bool) ([]byte, error) {
	field, ok := defs.NameToField[name]
	if !ok {
		return nil, fmt.Errorf("unknown field name %s", name)
	}

	// do not encode
	if !field.IsSerialized {
		return []byte{}, nil
	}

	// do not encode for signing
	if signing && !field.IsSigningField {
		return []byte{}, nil
	}

	encoder, err := typeEncoder(field.Type)
	if err != nil {
		return nil, fmt.Errorf("no encoder: %v", err)
	}

	valueBytes, err := encoder.ToBytes(value, signing)
	if err != nil {
		return nil, fmt.Errorf("invalid field %s: %v", name, err)
	}

	out := make([]byte, 0, len(valueBytes)+5)

	id, err := field.ID()
	if err != nil {
		return nil, fmt.Errorf("invalid field id %s: %v", name, err)
	}

	out = append(out, id...)

	if field.IsVLEncoded {
		prefix, err := lengthPrefix(len(valueBytes))
		if err != nil {
			return nil, fmt.Errorf("invalid len of field %v: %v", name, err)
		}

		out = append(out, prefix...)
	}

	out = append(out, valueBytes...)

	return out, nil
}

// sortFunc returns a function that is used to sort fields.
// If an unknown field is entered, the value of the provided err is set.
func sortFunc(err *error) func(string, string) int {
	return func(a, b string) int {
		aField, ok := defs.NameToField[a]
		if !ok {
			*err = fmt.Errorf("unknown field %s", a)
			return 0
		}
		bField, ok := defs.NameToField[b]
		if !ok {
			*err = fmt.Errorf("unknown field %s", b)
			return 0
		}

		return aField.Less(&bField)
	}
}

// sortFields returns sorted names of fields in order in which they should be serialized.
// If an unknown field is among names, an error is returned.
func sortFields(names []string) ([]string, error) {
	var err error

	slices.SortFunc(names, sortFunc(&err))
	if err != nil {
		return nil, err
	}

	return names, nil
}

type Object = map[string]any

// Encode serializes an object.
func Encode(value any, signing bool) ([]byte, error) {
	valueObj, ok := value.(Object)
	if !ok {
		return nil, fmt.Errorf("invalid object %v", value)
	}

	outBuff := bytes.NewBuffer(nil)

	names := keys(valueObj)
	sortedNames, err := sortFields(names)
	if err != nil {
		return nil, fmt.Errorf("cannot sort: %v", err)
	}

	for _, name := range sortedNames {
		bytes, err := encodeInner(name, valueObj[name], signing)
		if err != nil {
			return nil, fmt.Errorf("cannot encode %s: %v", name, err)
		}
		_, err = outBuff.Write(bytes)
		if err != nil {
			return nil, fmt.Errorf("cannot add %s to buffer: %v", name, err)
		}
	}

	return outBuff.Bytes(), nil
}

// keys extract keys from map and returns them into an array.
// The order of keys is not deterministic.
func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
