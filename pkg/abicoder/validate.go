package abicoder

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ErrShapeMismatch indicates a decode destination whose shape diverges from the
// ABI argument type: an array or struct of another length or field count, an
// array where the type has none, a field named for a tuple component at
// another position, an unexported field, a pointer to anything but a big.Int,
// or a type the decoded value can't be assigned to.
var ErrShapeMismatch = errors.New("destination shape does not match ABI type")

// ErrNonCanonical indicates ABI input that is not the canonical encoding of its
// type: an offset other than the one the encoder writes, non-zero padding, a
// value out of its type's range, or data after the encoding. Encode returns it
// for a value it can't encode canonically, such as one out of its type's range.
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
// rather than an error. checkShape catches those before conversion. Mapping is
// positional, but a field whose name or abi tag matches a component at another
// position, as Encode matches them, is rejected. So is a leaf ConvertType
// can't assign the decoded value to, which would fail every decode.
func checkShape[T any](argType abi.Type) error {
	dst, src := reflect.TypeFor[T](), argType.GetType()
	// ConvertType converts a *big.Int straight to a *T for a T based on big.Int
	if src.Kind() == reflect.Pointer && src.ConvertibleTo(reflect.PointerTo(dst)) {
		return nil
	}

	return shapeMatches(dst, src)
}

// checkArgsShape is checkShape for each destination abi.Arguments.Copy writes
// args to within dst.
func checkArgsShape(args abi.Arguments, dst reflect.Type) error {
	if len(args) == 1 {
		if dst.Kind() == reflect.Struct && dst.NumField() > 0 {
			dst = dst.Field(0).Type
		}
		return shapeMatches(dst, args[0].Type.GetType())
	}

	if dst.Kind() == reflect.Array && dst.Len() != len(args) {
		return fmt.Errorf("%w: array length %d, want %d arguments", ErrShapeMismatch, dst.Len(), len(args))
	}

	for _, arg := range args {
		var field reflect.Type
		switch dst.Kind() {
		case reflect.Struct:
			f, ok := argField(dst, arg.Name)
			if !ok {
				continue // Copy fails on it
			}
			field = f.Type
		case reflect.Slice, reflect.Array:
			field = dst.Elem()
		default:
			return nil
		}
		if err := shapeMatches(field, arg.Type.GetType()); err != nil {
			return fmt.Errorf("matching argument %q: %w", arg.Name, err)
		}
	}

	return nil
}

// argField returns the field of struct type t that abi.Arguments.Copy writes
// the argument name to: the one tagged abi:"name", else the camel-cased name.
func argField(t reflect.Type, name string) (reflect.StructField, bool) {
	for f := range t.Fields() {
		// go-ethereum's test, which also skips names with a non-ASCII initial
		exported := f.Name[:1] == strings.ToUpper(f.Name[:1])
		if tag, ok := f.Tag.Lookup("abi"); ok && tag == name && exported {
			return f, true
		}
	}

	return t.FieldByName(abi.ToCamelCase(name))
}

// shapeMatches checks dst against abiType as go-ethereum's set, which
// ConvertType and Copy end in, would fill a zero dst. It reports a mismatch
// when dst and abiType are both arrays of different lengths, or both structs
// with different field counts, an unexported field, or a field named after a
// component at another position, or dst is an array where abiType isn't, a
// pointer to anything but a big.Int, or any other type abiType isn't
// assignable to, recursing through arrays, slices, and structs.
func shapeMatches(dst, abiType reflect.Type) error {
	if dst == abiType {
		return nil
	}

	switch {
	case dst.Kind() == reflect.Interface:
		if !abiType.AssignableTo(dst) {
			return fmt.Errorf("%w: %v does not implement %v", ErrShapeMismatch, abiType, dst)
		}
		return nil
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
	case dst.Kind() == reflect.Pointer && dst.Elem() != reflect.TypeFor[big.Int]():
		// go-ethereum decodes through a pointer without allocating it, so it panics
		return fmt.Errorf("%w: pointer %v, want %v", ErrShapeMismatch, dst, abiType)
	case abiType.AssignableTo(dst):
		return nil
	case dst.Kind() == reflect.Struct && abiType.Kind() == reflect.Struct:
		if dst.NumField() != abiType.NumField() {
			return fmt.Errorf("%w: struct has %d fields, want %d", ErrShapeMismatch, dst.NumField(), abiType.NumField())
		}
		for i := range dst.NumField() {
			if !dst.Field(i).IsExported() {
				return fmt.Errorf("%w: unexported field %s", ErrShapeMismatch, dst.Field(i).Name)
			}
			// go-ethereum tags each tuple field json:"<component name>"
			name := abiType.Field(i).Tag.Get("json")
			if f, ok := argField(dst, name); ok && (len(f.Index) != 1 || f.Index[0] != i) {
				return fmt.Errorf("%w: component %q decodes into field %s but names field %s", ErrShapeMismatch, name, dst.Field(i).Name, f.Name)
			}
			if err := shapeMatches(dst.Field(i).Type, abiType.Field(i).Type); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("%w: %v, want %v", ErrShapeMismatch, dst, abiType)
	}
}
