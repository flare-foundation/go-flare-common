// Package abicoder provides generic ABI encoding and decoding utilities.
package abicoder

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Encode ABI-encodes data according to the given argument definition.
//
// abi.Arguments.Pack panics on some malformed inputs (e.g. a nil *big.Int),
// which the deferred recover surfaces as an error.
func Encode(arg abi.Argument, data any) (encoded []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			e, ok := r.(error)
			if ok {
				err = fmt.Errorf("recovered panic: %w", e)
			} else {
				err = fmt.Errorf("recovered panic non error: %v", r)
			}
		}
	}()

	args := abi.Arguments{arg}

	encoded, err = args.Pack(data)
	if err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	return encoded, nil
}
