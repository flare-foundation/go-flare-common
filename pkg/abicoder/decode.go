package abicoder

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/panics"
)

var errNonPointer = errors.New("dest is not a non nil pointer")

// DecodeTo decodes ABI-encoded data and writes it to dest.
//
// dest must be a non-nil pointer. DecodeTo behaves like [Decode] and writes the
// result through dest, leaving dest unchanged on error.
func DecodeTo[T any](arg abi.Argument, data []byte, dest *T) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errNonPointer
	}

	t, err := Decode[T](arg, data)
	if err != nil {
		return err
	}

	*dest = t

	return nil
}

// Decode decodes ABI-encoded data into a value of type T.
//
// T's array lengths and struct field counts must match the ABI argument, T
// may have a Go array only where the argument has a fixed-size array, bytesN,
// address or function, no pointer but a *big.Int and no unexported field, and
// each decoded value must be assignable to its place in T, else
// ErrShapeMismatch is returned. Tuples map to structs by position, and a
// struct field whose name or abi tag matches a component at another position,
// where [Encode] would take it from, is ErrShapeMismatch. Data that isn't the
// canonical encoding of a value is rejected with ErrNonCanonical, and data too
// short for one with ErrTruncated, both before decoding. When T is an
// interface the decoded value must implement it, and tuples decode to
// go-ethereum's json-tagged anonymous struct.
func Decode[T any](arg abi.Argument, data []byte) (t T, err error) {
	defer panics.Recover(&err)

	values, err := unpack(abi.Arguments{arg}, data)
	if err != nil {
		return t, err
	}

	if err := checkShape[T](arg.Type); err != nil {
		return t, err
	}

	x := abi.ConvertType(values[0], new(T))

	tp, ok := x.(*T)
	if !ok {
		return t, errors.New("invalid type assertion")
	}
	t = *tp

	return t, nil
}

// unpack checks that data is the canonical encoding of args and decodes it.
func unpack(args abi.Arguments, data []byte) ([]any, error) {
	for _, arg := range args {
		if err := checkArgType(arg.Type); err != nil {
			return nil, err
		}
	}

	if err := checkLayout(args, data); err != nil {
		return nil, err
	}

	values, err := args.Unpack(data)
	if err != nil {
		return nil, err
	}

	if err := checkEncodeDecode(&args, data, values); err != nil {
		return nil, err
	}

	return values, nil
}

// checkEncodeDecode encodes decodedSlice using args and compares it to data.
// Its errors never quote data, which may be secret.
func checkEncodeDecode(args *abi.Arguments, data []byte, decodedSlice []any) error {
	encoded, err := args.Pack(decodedSlice...)
	if err != nil {
		return fmt.Errorf("re-encoding the decoded values: %w", err)
	}

	if !bytes.Equal(data, encoded) {
		return fmt.Errorf("%w: re-encoding the decoded values gives other bytes", ErrNonCanonical)
	}

	return nil
}

// DecodeArguments decodes data, the encoding of the argument list args, as
// Solidity's abi.encode(a, b, ...) writes it: unlike the encoding of one tuple,
// it has no leading offset.
//
// Several arguments are copied, as abi.Arguments.Copy copies them, into a
// struct with a field per argument, matched by name, or an array with an
// element per argument, in order. A single argument is copied into T, or the
// first field of a struct T. Data is checked as Decode checks it, and the field
// or element each argument is copied into as Decode checks T. Fields of T no
// argument is copied into are left zero.
func DecodeArguments[T any](args abi.Arguments, data []byte) (t T, err error) {
	defer panics.Recover(&err)

	values, err := unpack(args, data)
	if err != nil {
		return t, err
	}

	if err := checkArgsShape(args, reflect.TypeFor[T]()); err != nil {
		return t, err
	}

	// a copy that fails part way leaves fields written, so t stays zero until it succeeds
	var out T
	if err := args.Copy(&out, values); err != nil {
		return t, fmt.Errorf("copying into %v: %w", reflect.TypeFor[T](), err)
	}

	return out, nil
}

// DecodeArgumentsTo decodes data, the encoding of the argument list args, and
// writes it to dest, a non-nil pointer. DecodeArgumentsTo behaves like
// [DecodeArguments] and leaves dest unchanged on error.
func DecodeArgumentsTo[T any](args abi.Arguments, data []byte, dest *T) error {
	if dest == nil {
		return errNonPointer
	}

	t, err := DecodeArguments[T](args, data)
	if err != nil {
		return err
	}

	*dest = t

	return nil
}
