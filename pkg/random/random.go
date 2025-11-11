package random

import (
	"crypto/rand"
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/common"
)

// Hash returns cryptographically random common.Hash.
func Hash() (common.Hash, error) {
	x := common.Hash{}

	_, err := io.ReadFull(rand.Reader, x[:])

	return x, err
}

// Bytes returns cryptographically random byte slice of specified length.
func Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("negative length")
	}

	b := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, err
	}

	return b, nil
}
