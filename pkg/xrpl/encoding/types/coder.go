// Package types implements XRPL binary serialization and deserialization for all field types.
package types

import (
	"bytes"
	"errors"
	"fmt"
	"slices"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/defs"
)

// TODOError is returned for types that are not yet implemented.
type TODOError struct {
	xt defs.XType
}

func (e *TODOError) Error() string {
	return fmt.Sprintf("To be implemented %v", e.xt)
}

// UnsupportedError is returned for types that are not supported by this implementation.
type UnsupportedError struct {
	xt defs.XType
}

func (e *UnsupportedError) Error() string {
	return fmt.Sprintf("Unsupported type %v", e.xt)
}

// IllegalError is returned for types that must not appear in serialized data.
type IllegalError struct {
	xt defs.XType
}

func (e *IllegalError) Error() string {
	return fmt.Sprintf("Illegal type %v", e.xt)
}

// Encode serializes an object.
func Encode(value Object, signing bool) ([]byte, error) {
	if len(value) == 0 {
		return nil, errors.New("empty object")
	}

	outBuff := bytes.NewBuffer(nil)

	names := keys(value)
	sortedNames, err := sortFields(names)
	if err != nil {
		return nil, fmt.Errorf("cannot sort: %w", err)
	}

	for _, name := range sortedNames {
		bytes, err := encodeInner(name, value[name], signing)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		_, err = outBuff.Write(bytes)
		if err != nil {
			return nil, fmt.Errorf("cannot add %s to buffer: %w", name, err)
		}
	}

	return outBuff.Bytes(), nil
}

// Coder defines the interface for encoding and decoding XRPL field values.
type Coder interface {
	ToBytes(value any, signing bool) ([]byte, error)
	ToJSON(b *bytes.Buffer, length int) (any, error)
}

func typeCoder(xt defs.XType) (Coder, error) {
	switch xt {
	case defs.Hash160:
		return Hash160, nil
	case defs.Hash192:
		return Hash192, nil
	case defs.Transaction:
		return nil, &UnsupportedError{xt}
	case defs.UInt512:
		return nil, &UnsupportedError{xt}
	case defs.UInt8:
		return &UInt8{}, nil
	case defs.XChainBridge:
		return &XChainBridge{}, nil
	case defs.Vector256:
		return Vector256, nil
	case defs.AccountID:
		return AccountID, nil
	case defs.Currency:
		return Currency, nil
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
		return Issue, nil
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

// lengthEncode computes length prefix.
func lengthEncode(n int) ([]byte, error) {
	var prefix []byte
	switch {
	case n < 0 || n > 918744:
		return nil, fmt.Errorf("unsupported Variable Length encoding: %d", n)
	case n <= 192:
		prefix = []byte{uint8(n)}
	case n <= 12480:
		n -= 193
		prefix = []byte{193 + uint8(n>>8), uint8(n)} //nolint:gosec // checked for overflow + expected cutoff
	case n <= 918744:
		n -= 12481
		prefix = []byte{241 + uint8(n>>16), uint8(n >> 8), uint8(n)} //nolint:gosec // checked for overflow + expected cutoff
	}

	return prefix, nil
}

// lengthDecode decodes length prefix.
func lengthDecode(b *bytes.Buffer) (int, error) {
	byte1, err := b.ReadByte()
	if err != nil {
		return -1, fmt.Errorf("cannot read first byte %w", err)
	}

	switch {
	case byte1 < 193:
		return int(byte1), nil
	case byte1 < 241:
		byte2, err := b.ReadByte()
		if err != nil {
			return -1, fmt.Errorf("cannot read second byte %w", err)
		}

		return 193 + ((int(byte1) - 193) * 256) + int(byte2), nil
	case byte1 < 255:
		bytes := make([]byte, 2)
		n, err := b.Read(bytes)
		if n != 2 || err != nil {
			return -1, fmt.Errorf("cannot read second and third byte %w", err)
		}

		out := 12481 + ((int(byte1) - 241) * 65536) + (int(bytes[0]) * 256) + int(bytes[1])

		if out > 918744 {
			return -1, fmt.Errorf("length prefix overflow %d can be at most %d", out, 918744)
		}
		return out, nil
	}

	return -1, fmt.Errorf("invalid first length byte %d", byte1)
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

	encoder, err := typeCoder(field.Type)
	if err != nil {
		return nil, fmt.Errorf("no encoder: %w", err)
	}

	valueBytes, err := encoder.ToBytes(value, signing)
	if err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(valueBytes)+5)

	id, err := field.ID()
	if err != nil {
		return nil, fmt.Errorf("invalid field id %s: %w", name, err)
	}

	out = append(out, id...)

	if field.IsVLEncoded {
		prefix, err := lengthEncode(len(valueBytes))
		if err != nil {
			return nil, fmt.Errorf("invalid len of field %v: %w", name, err)
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

// keys extract keys from map and returns them into an array.
// The order of keys is not deterministic.
func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func decodeNext(b *bytes.Buffer) (string, any, error) {
	idPair, err := defs.ReadID(b)
	if err != nil {
		return "", nil, fmt.Errorf("reading id: %w", err)
	}

	fName, ok := defs.IDToName[idPair]
	if !ok {
		return "", nil, fmt.Errorf("unknown pair. Field: %v, type: %v", idPair.F, idPair.T)
	}

	coder, err := typeCoder(idPair.T)
	if err != nil {
		return "", nil, fmt.Errorf("missing coder for %v: %w", idPair.T, err)
	}

	field, ok := defs.NameToField[fName]
	if !ok {
		return "", nil, fmt.Errorf("missing field for %v", fName)
	}

	length := 0
	if field.IsVLEncoded {
		length, err = lengthDecode(b)
		if err != nil {
			return "", nil, fmt.Errorf("reading length %v: %w", idPair.T, err)
		}
	}

	value, err := coder.ToJSON(b, length)
	if err != nil {
		return "", nil, fmt.Errorf("decoding %v: %w", fName, err)
	}

	if fName == "TransactionType" {
		value, err = Uint16ToTxType(value)
		if err != nil {
			return "", nil, fmt.Errorf("invalid tx type: %w", err)
		}
	}

	return fName, value, nil
}

// Decode deserializes an XRPL binary blob into a map of field names to values.
func Decode(blob []byte) (map[string]any, error) {
	out := make(map[string]any)

	b := bytes.NewBuffer(blob)

	for b.Len() > 0 {
		name, value, err := decodeNext(b)
		if err != nil {
			return nil, fmt.Errorf("decoding next: %w", err)
		}

		out[name] = value
	}

	return out, nil
}

// outOfBytes formats error when buffer does not hold the expected number of bytes.
func outOfBytes(value string, expected, got int) error {
	return fmt.Errorf("cannot read %s from buffer: expected %d bytes got %d", value, expected, got)
}
