package defs

import (
	"fmt"
)

type XType int16

type TransactionType int16

type Field struct {
	IsSerialized   bool
	IsSigningField bool
	IsVLEncoded    bool
	Nth            int16 // Field code
	Type           XType // Type code
}

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
