// Package structs provides ABI encoding and decoding utilities for TEE instruction data.
package structs

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Encode ABI-encodes data according to the given argument definition.
func Encode(arg abi.Argument, data any) ([]byte, error) {
	args := abi.Arguments{arg}
	encoded, err := args.Pack(data)
	if err != nil {
		return nil, fmt.Errorf("encoding type %T: %v", data, err)
	}

	return encoded, nil
}
