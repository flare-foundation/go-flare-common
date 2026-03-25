package convert

import (
	"errors"
	"math/big"
)

// BigToUint32Safe converts big.Int into uint32.
// It returns an error if it is out of bounds.
func BigToUint32Safe(b *big.Int) (uint32, error) {
	idNegative := b.Sign() == -1
	idOverflow := b.BitLen() > 32

	if idNegative || idOverflow {
		return 0, errors.New("not uint32")
	}

	u := uint32(b.Uint64()) //nolint:gosec // if statement above checks for under and overflow

	return u, nil
}

// BigToUint64Safe converts big.Int into uint64.
// It returns an error if it is out of bounds.
func BigToUint64Safe(b *big.Int) (uint64, error) {
	if b.IsUint64() {
		return b.Uint64(), nil
	}

	return 0, errors.New("not uint64")
}
