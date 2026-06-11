package convert

import (
	"errors"
	"fmt"
	"math/big"
)

// ErrNilBigInt indicates a nil *big.Int input.
var ErrNilBigInt = errors.New("nil big.Int")

// ErrOutOfRange indicates a value that does not fit the target type.
var ErrOutOfRange = errors.New("value out of range")

// BigToUint32Safe converts big.Int into uint32.
// It returns ErrNilBigInt if b is nil and ErrOutOfRange if b is out of bounds.
func BigToUint32Safe(b *big.Int) (uint32, error) {
	if b == nil {
		return 0, ErrNilBigInt
	}
	negative := b.Sign() == -1
	overflow := b.BitLen() > 32

	if negative || overflow {
		return 0, fmt.Errorf("uint32: %w", ErrOutOfRange)
	}

	u := uint32(b.Uint64()) //nolint:gosec // if statement above checks for under and overflow

	return u, nil
}

// BigToUint64Safe converts big.Int into uint64.
// It returns ErrNilBigInt if b is nil and ErrOutOfRange if b is out of bounds.
func BigToUint64Safe(b *big.Int) (uint64, error) {
	if b == nil {
		return 0, ErrNilBigInt
	}
	if b.IsUint64() {
		return b.Uint64(), nil
	}

	return 0, fmt.Errorf("uint64: %w", ErrOutOfRange)
}
