// Package defs provides XRPL protocol field definitions and binary serialization metadata.
package defs

import (
	"bytes"
	"fmt"
)

// XType represents an XRPL binary serialization type code.
type XType int16

// Field describes the serialization properties of an XRPL protocol field.
type Field struct {
	IsSerialized   bool
	IsSigningField bool
	IsVLEncoded    bool
	Nth            int16 // Field code
	Type           XType // Type code
}

// IDPair uniquely identifies a field by its field code and type code.
type IDPair struct {
	F int16
	T XType
}

// IDToName maps pair of (field code, type code) to field name.
var IDToName map[IDPair]string

// ValueToTxType maps a transaction type numeric value to its name.
var ValueToTxType map[int32]string

func init() {
	IDToName = make(map[IDPair]string)

	for name := range NameToField {
		fCode := NameToField[name].Nth
		tCode := NameToField[name].Type

		key := IDPair{F: fCode, T: tCode}

		IDToName[key] = name
	}

	ValueToTxType = make(map[int32]string)

	for key, value := range TxTypeToValue {
		ValueToTxType[value] = key
	}
}

// ID returns the ID of a field according to https://xrpl.org/docs/references/protocol/binary-format#field-ids.
func (f Field) ID() ([]byte, error) {
	if f.Type < 0 || f.Type > 255 || f.Nth < 0 || f.Nth > 255 {
		return nil, fmt.Errorf("invalid type or name %d, %d", f.Type, f.Nth)
	}

	t := uint8(f.Type)
	n := uint8(f.Nth)

	var b []byte
	var err error
	switch {
	case t < 16 && n < 16:
		b = []byte{t<<4 | n}
	case t < 16: // n >= 16
		b = []byte{t << 4, n}
	case n < 16: // t >= 16
		b = []byte{n, t}
	default: // t,n >= 16
		b = []byte{0, t, n}
	}
	return b, err
}

// ReadID reads a field/type ID pair from the buffer according to the XRPL binary format.
func ReadID(b *bytes.Buffer) (IDPair, error) {
	byte1, err := b.ReadByte()
	if err != nil {
		return IDPair{-1, NotPresent}, fmt.Errorf("cannot read first byte %v", err)
	}

	tCode := byte1 & 0xf0
	tCode >>= 4

	fCode := byte1 & 0x0f

	if tCode == 0 {
		tCode, err = b.ReadByte()
		if err != nil {
			return IDPair{-1, NotPresent}, fmt.Errorf("cannot read type code byte %v", err)
		}
		if tCode < 16 {
			return IDPair{-1, NotPresent}, fmt.Errorf("invalid encoding type code %v", []byte{byte1, tCode})
		}
	}

	if fCode == 0 {
		fCode, err = b.ReadByte()
		if err != nil {
			return IDPair{-1, NotPresent}, fmt.Errorf("cannot read field code byte %v", err)
		}
		if fCode < 16 {
			return IDPair{-1, NotPresent}, fmt.Errorf("invalid encoding field code %v", []byte{byte1, fCode})
		}
	}

	return IDPair{F: int16(fCode), T: XType(tCode)}, nil
}

// Less returns
//
//   - -1 if f should be sorted before g
//   - 0 if f is equal to g
//   - 1 if f should be sorted after g
func (f *Field) Less(g *Field) int {
	switch {
	case f.Type < g.Type:
		return -1
	case f.Type == g.Type && f.Nth < g.Nth:
		return -1
	case f.Type == g.Type && f.Nth == g.Nth:
		return 0
	default:
		return 1
	}
}
