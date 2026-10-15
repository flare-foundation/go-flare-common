package abicoder

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
// T's array lengths and struct field counts must match the ABI argument, and T
// may have a Go array only where the argument has a fixed-size array, bytesN,
// address or function, else ErrShapeMismatch is returned; tuple-to-struct
// mapping is positional by index and ignores field names. Data that isn't the
// canonical encoding of a value is rejected with ErrNonCanonical, and data too
// short for one with ErrTruncated, both before decoding. When T is an interface
// no shape check is performed and tuples decode to go-ethereum's json-tagged
// anonymous struct.
func Decode[T any](arg abi.Argument, data []byte) (t T, err error) {
	defer recoverError(&err)

	if err := checkArgType(arg.Type); err != nil {
		return t, err
	}

	args := abi.Arguments{arg}

	if err := checkLayout(args, data); err != nil {
		return t, err
	}

	if err := checkShape[T](arg.Type); err != nil {
		return t, err
	}

	decodedSlice, err := args.Unpack(data)
	if err != nil {
		return t, err
	}

	err = checkEncodeDecode(&args, data, decodedSlice)
	if err != nil {
		return t, err
	}

	x := abi.ConvertType(decodedSlice[0], new(T))

	tp, ok := x.(*T)
	if !ok {
		return t, errors.New("invalid type assertion")
	}
	t = *tp

	return t, nil
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
