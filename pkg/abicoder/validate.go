package abicoder

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ErrShapeMismatch indicates a decode destination whose shape diverges from the
// ABI argument type: an array or struct of another length or field count, or
// an array where the type has none.
var ErrShapeMismatch = errors.New("destination shape does not match ABI type")

// ErrNonCanonical indicates ABI input that is not the canonical encoding of its
// type: an offset other than the one the encoder writes, non-zero padding, a
// value out of its type's range, or data after the encoding.
var ErrNonCanonical = errors.New("non-canonical ABI encoding")

// ErrTruncated indicates ABI input that ends before the encoding its type
// needs: too short, or an offset or length past its end.
var ErrTruncated = errors.New("truncated ABI encoding")

// errUninitializedArg reports an abi.Argument left with its zero abi.Type.
var errUninitializedArg = errors.New("uninitialized abi argument: zero abi.Type (construct it with abi.NewType or ABI JSON)")

// checkArgType rejects the zero abi.Type, which would otherwise be silently
// treated as int256. A valid integer type always has a non-zero size, and the
// int/uint aliases fail to construct, so an IntTy of size zero with no string
// form is only ever the uninitialized zero value.
func checkArgType(t abi.Type) error {
	if t.T == abi.IntTy && t.Size == 0 && t.String() == "" {
		return errUninitializedArg
	}

	return nil
}

// checkShape rejects destinations whose composite shape diverges from argType.
//
// abi.ConvertType maps tuples to structs positionally by field index and copies
// only min(src,dst) array elements, so a destination with a different array
// length or struct field count decodes via silent truncation or zero-fill
// rather than an error. checkShape catches those before conversion. It does not
// compare field names (mapping is positional) and skips interface destinations.
func checkShape[T any](argType abi.Type) error {
	return shapeMatches(reflect.TypeFor[T](), argType.GetType())
}

// shapeMatches reports a mismatch when dst and abiType are both arrays of
// different lengths or both structs with different field counts, or dst is an
// array where abiType isn't, recursing through arrays, slices, pointers, and
// structs. Identical types, interface destinations, and any other kind
// divergence are left to abi.ConvertType.
func shapeMatches(dst, abiType reflect.Type) error {
	if dst == abiType || dst.Kind() == reflect.Interface {
		return nil
	}

	switch {
	case dst.Kind() == reflect.Array && abiType.Kind() == reflect.Array:
		if dst.Len() != abiType.Len() {
			return fmt.Errorf("%w: array length %d, want %d", ErrShapeMismatch, dst.Len(), abiType.Len())
		}
		return shapeMatches(dst.Elem(), abiType.Elem())
	case dst.Kind() == reflect.Array:
		// go-ethereum truncates or zero-fills a slice or string, and on a
		// *big.Int recurses until a fatal stack overflow
		return fmt.Errorf("%w: array %v, want %v", ErrShapeMismatch, dst, abiType)
	case dst.Kind() == reflect.Slice && abiType.Kind() == reflect.Slice:
		return shapeMatches(dst.Elem(), abiType.Elem())
	case dst.Kind() == reflect.Pointer && abiType.Kind() == reflect.Pointer:
		return shapeMatches(dst.Elem(), abiType.Elem())
	case dst.Kind() == reflect.Struct && abiType.Kind() == reflect.Struct:
		if dst.NumField() != abiType.NumField() {
			return fmt.Errorf("%w: struct has %d fields, want %d", ErrShapeMismatch, dst.NumField(), abiType.NumField())
		}
		for i := range dst.NumField() {
			if err := shapeMatches(dst.Field(i).Type, abiType.Field(i).Type); err != nil {
				return err
			}
		}
		return nil
	default:
		return nil
	}
}
