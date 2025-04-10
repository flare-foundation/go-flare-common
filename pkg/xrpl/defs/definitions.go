package defs

import (
	"bytes"
	"fmt"
)

type XType int16

type Field struct {
	IsSerialized   bool
	IsSigningField bool
	IsVLEncoded    bool
	Nth            int16 // Field code
	Type           XType // Type code
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

func IDDecode(b *bytes.Buffer) (int16, XType, error) {
	byte1, err := b.ReadByte()
	if err != nil {
		return -1, NotPresent, fmt.Errorf("cannot read first byte %v", err)
	}

	tCode := byte1 & 0xf0
	tCode >>= 4

	fCode := byte1 & 0x0f

	if tCode == 0 {
		tCode, err = b.ReadByte()
		if err != nil {
			return -1, NotPresent, fmt.Errorf("cannot read type code byte %v", err)
		}
		if tCode < 16 {
			return -1, NotPresent, fmt.Errorf("invalid encoding type code %v", []byte{byte1, tCode})
		}
	}

	if fCode == 0 {
		fCode, err = b.ReadByte()
		if err != nil {
			return -1, NotPresent, fmt.Errorf("cannot read field code byte %v", err)
		}
		if fCode < 16 {
			return -1, NotPresent, fmt.Errorf("invalid encoding field code%v", []byte{byte1, fCode})
		}
	}

	return int16(fCode), XType(tCode), nil
}

// Less returns
//
//   - -1 if f should be sorted before g
//   - 0 if f is equal to g
//   - 1 if f should be sorted after g
func (f *Field) Less(g *Field) int {
	if f.Type < g.Type {
		return -1
	} else if f.Type == g.Type && f.Nth < g.Nth {
		return -1
	} else if f.Type == g.Type && f.Nth == g.Nth {
		return -1
	}
	return 1
}
