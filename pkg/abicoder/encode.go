// Package abicoder provides generic ABI encoding and decoding utilities.
//
// It handles ABI argument values such as structs, tuples, and primitives. It
// does not handle events: indexed (topic) arguments are not supported.
package abicoder

import (
	"fmt"
	"math/big"
	"reflect"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/panics"
)

// Encode ABI-encodes data according to the given argument definition.
//
// It refuses types [Decode] doesn't support. Values that don't fit their type,
// such as an int256 at or above 2^255, or an address from more than 32 bytes or
// from 32 with the high ones set, fail with ErrNonCanonical.
//
// abi.Arguments.Pack panics on some malformed inputs (e.g. a nil *big.Int),
// which the deferred recover surfaces as an error.
func Encode(arg abi.Argument, data any) (encoded []byte, err error) {
	defer panics.Recover(&err)

	if err := checkArgType(arg.Type); err != nil {
		return nil, err
	}

	// go-ethereum wraps big integers modulo 2^256 and writes long addresses and functions whole
	if err := checkRange(&arg.Type, reflect.ValueOf(data)); err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	args := abi.Arguments{arg}

	encoded, err = args.Pack(data)
	if err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	// go-ethereum writes some values out of range, such as an address from 32 bytes
	if err := checkLayout(args, encoded); err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	return encoded, nil
}

var bigIntPtr = reflect.TypeFor[*big.Int]()

// checkRange checks that every *big.Int in v fits its integer type in t, and
// every address or function array fits a word. It walks v as
// abi.Arguments.Pack does and leaves what Pack refuses to Pack.
func checkRange(t *abi.Type, v reflect.Value) error {
	var seen []reflect.Type
	for v.Kind() == reflect.Pointer && v.Type() != bigIntPtr {
		// a recursive type such as type P *P would loop, and overflow Pack's stack
		if slices.Contains(seen, v.Type()) {
			return fmt.Errorf("recursive pointer type %v", v.Type())
		}
		seen = append(seen, v.Type())
		v = v.Elem()
	}
	if !v.IsValid() {
		return nil
	}

	switch t.T {
	case abi.AddressTy, abi.FunctionTy:
		// go-ethereum writes a longer one whole, shifting what follows
		if v.Kind() == reflect.Array && v.Len() > word {
			return fmt.Errorf("%w: %d bytes for %s", ErrNonCanonical, v.Len(), t)
		}
	case abi.IntTy, abi.UintTy:
		// native Go integers fit by type
		if v.Type() != bigIntPtr || v.IsNil() || !v.CanInterface() {
			return nil
		}
		n, ok := reflect.TypeAssert[*big.Int](v)
		if ok && !fitsInt(t, n) {
			return fmt.Errorf("%w: %s out of range", ErrNonCanonical, t)
		}
	case abi.SliceTy, abi.ArrayTy:
		if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
			return nil
		}
		for i := range v.Len() {
			if err := checkRange(t.Elem, v.Index(i)); err != nil {
				return err
			}
		}
	case abi.TupleTy:
		if v.Kind() != reflect.Struct {
			return nil
		}
		for i, e := range t.TupleElems {
			// Pack takes each component from the field its name matches
			f, ok := argField(v.Type(), t.TupleRawNames[i])
			if !ok {
				continue
			}
			fv, err := v.FieldByIndexErr(f.Index)
			if err != nil {
				continue
			}
			if err := checkRange(e, fv); err != nil {
				return err
			}
		}
	}

	return nil
}

// fitsInt reports whether n is in the range of the integer type t.
func fitsInt(t *abi.Type, n *big.Int) bool {
	if t.T == abi.UintTy {
		return n.Sign() >= 0 && n.BitLen() <= t.Size
	}

	m := n
	if n.Sign() < 0 {
		m = new(big.Int).Not(n) // -n-1
	}

	return m.BitLen() < t.Size
}
