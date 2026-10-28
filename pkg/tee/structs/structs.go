// Package structs re-exports the ABI encoding and decoding helpers from
// [github.com/flare-foundation/go-flare-common/pkg/abicoder] for backwards
// compatibility. New code should use abicoder directly.
package structs

import (
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
)

// Encode is an alias for [abicoder.Encode].
func Encode(arg abi.Argument, data any) ([]byte, error) {
	return abicoder.Encode(arg, data)
}

// Decode is an alias for [abicoder.Decode].
//
// Deprecated: Use [abicoder.Decode], or an [abicoder.Coder] from [abicoder.New]
// to decode one type repeatedly.
func Decode[T any](arg abi.Argument, data []byte) (T, error) {
	return abicoder.Decode[T](arg, data)
}

// DecodeTo is an alias for [abicoder.DecodeTo].
//
// Deprecated: Use [abicoder.DecodeTo], or an [abicoder.Coder] from
// [abicoder.New] to decode one type repeatedly.
func DecodeTo[T any](arg abi.Argument, data []byte, dest *T) error {
	return abicoder.DecodeTo(arg, data, dest)
}
