package abicoder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ErrShapeMismatch indicates that a decode destination's array length or struct
// field count diverges from the ABI argument type.
var ErrShapeMismatch = errors.New("destination shape does not match ABI type")

// ErrNonCanonical indicates ABI input whose dynamic content overlaps so that
// decoding it would copy more bytes than the input contains, which a canonical
// encoding never does.
var ErrNonCanonical = errors.New("non-canonical ABI encoding")

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
// different lengths or both structs with different field counts, recursing
// through arrays, slices, pointers, and structs. Identical types, interface
// destinations, and any other kind divergence are left to abi.ConvertType.
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

// containsString reports whether decoding t can reach a StringTy node, the only
// ABI kind whose unpacking copies its content (bytes return a view of the input,
// scalars fit in one word). Only string-bearing types can amplify allocation.
func containsString(t abi.Type) bool {
	switch t.T {
	case abi.StringTy:
		return true
	case abi.SliceTy, abi.ArrayTy:
		return containsString(*t.Elem)
	case abi.TupleTy:
		for _, elem := range t.TupleElems {
			if containsString(*elem) {
				return true
			}
		}
	}

	return false
}

// exceedsStringBudget reports whether decoding argType from data would copy more
// than len(data) bytes of string content, which signals offset aliasing that
// amplifies allocation (see ErrNonCanonical). It allocates no string content and
// follows the same offset graph as go-ethereum's Unpack, so a canonical encoding
// — whose string regions are disjoint and therefore sum to at most len(data) —
// never trips it. A malformed offset or an exhausted visit budget defers to the
// real Unpack by reporting false.
func exceedsStringBudget(argType abi.Type, data []byte) bool {
	if !containsString(argType) {
		return false
	}

	s := stringScan{limit: len(data), budget: len(data) + 64}
	ok := s.walk(argType, data, 0)

	return ok && s.exceeded
}

// stringScan sums StringTy content lengths while walking an ABI encoding.
// limit is the input length; once sum exceeds it, exceeded is set and the walk
// short-circuits. budget caps node visits so a deeply nested type cannot make
// the scan itself expensive; exhausting it defers to the real Unpack.
type stringScan struct {
	limit    int
	budget   int
	sum      int
	visits   int
	exceeded bool
}

// walk mirrors go-ethereum's toGoType traversal for v1.17.2, accumulating string
// lengths. It returns false on any malformed offset or budget exhaustion, asking
// the caller to defer to the real Unpack rather than reject.
func (s *stringScan) walk(t abi.Type, output []byte, index int) bool {
	if s.exceeded {
		return true
	}
	s.visits++
	if s.visits > s.budget || index+32 > len(output) {
		return false
	}

	switch t.T {
	case abi.StringTy:
		_, length, ok := lengthPrefixPointsTo(index, output)
		if !ok {
			return false
		}
		s.sum += length
		if s.sum > s.limit {
			s.exceeded = true
		}
		return true
	case abi.SliceTy:
		begin, length, ok := lengthPrefixPointsTo(index, output)
		if !ok {
			return false
		}
		if !containsString(*t.Elem) {
			return true
		}
		return s.walkEach(t, output[begin:], length)
	case abi.ArrayTy:
		if !containsString(*t.Elem) {
			return true
		}
		if isDynamicType(*t.Elem) {
			offset := binary.BigEndian.Uint64(output[index+24 : index+32])
			if offset > uint64(len(output)) {
				return false
			}
			return s.walkEach(t, output[offset:], t.Size)
		}
		return s.walkEach(t, output[index:], t.Size)
	case abi.TupleTy:
		if isDynamicType(t) {
			begin, ok := tuplePointsTo(index, output)
			if !ok {
				return false
			}
			return s.walkTuple(t, output[begin:])
		}
		return s.walkTuple(t, output[index:])
	default:
		return true
	}
}

// walkEach mirrors go-ethereum's forEachUnpack: it visits size elements of the
// slice or array type t laid out from the start of output.
func (s *stringScan) walkEach(t abi.Type, output []byte, size int) bool {
	if size < 0 || 32*size > len(output) {
		return false
	}

	elemSize := getTypeSize(*t.Elem)
	for j := range size {
		if s.exceeded {
			return true
		}
		if !s.walk(*t.Elem, output, j*elemSize) {
			return false
		}
	}

	return true
}

// walkTuple mirrors go-ethereum's forTupleUnpack, tracking the virtual-argument
// offset that static array and tuple fields introduce.
func (s *stringScan) walkTuple(t abi.Type, output []byte) bool {
	virtualArgs := 0
	for index, elem := range t.TupleElems {
		if s.exceeded {
			return true
		}
		if containsString(*elem) {
			if !s.walk(*elem, output, (index+virtualArgs)*32) {
				return false
			}
		}
		if (elem.T == abi.ArrayTy || elem.T == abi.TupleTy) && !isDynamicType(*elem) {
			virtualArgs += getTypeSize(*elem)/32 - 1
		}
	}

	return true
}

// isDynamicType mirrors go-ethereum's abi.isDynamicType for v1.17.2.
func isDynamicType(t abi.Type) bool {
	if t.T == abi.TupleTy {
		for _, elem := range t.TupleElems {
			if isDynamicType(*elem) {
				return true
			}
		}
		return false
	}

	return t.T == abi.StringTy || t.T == abi.BytesTy || t.T == abi.SliceTy ||
		(t.T == abi.ArrayTy && isDynamicType(*t.Elem))
}

// getTypeSize mirrors go-ethereum's abi.getTypeSize for v1.17.2.
func getTypeSize(t abi.Type) int {
	if t.T == abi.ArrayTy && !isDynamicType(*t.Elem) {
		if t.Elem.T == abi.ArrayTy || t.Elem.T == abi.TupleTy {
			return t.Size * getTypeSize(*t.Elem)
		}
		return t.Size * 32
	} else if t.T == abi.TupleTy && !isDynamicType(t) {
		total := 0
		for _, elem := range t.TupleElems {
			total += getTypeSize(*elem)
		}
		return total
	}

	return 32
}

// lengthPrefixPointsTo mirrors go-ethereum's abi.lengthPrefixPointsTo for
// v1.17.2, resolving a length-prefixed offset to its content start and length.
// ok is false when the offset or length is out of range.
func lengthPrefixPointsTo(index int, output []byte) (start, length int, ok bool) {
	if index+32 > len(output) {
		return 0, 0, false
	}

	bigOffsetEnd := new(big.Int).SetBytes(output[index : index+32])
	bigOffsetEnd.Add(bigOffsetEnd, big.NewInt(32))
	outputLength := big.NewInt(int64(len(output)))
	if bigOffsetEnd.Cmp(outputLength) > 0 || bigOffsetEnd.BitLen() > 63 {
		return 0, 0, false
	}

	offsetEnd := int(bigOffsetEnd.Uint64())
	lengthBig := new(big.Int).SetBytes(output[offsetEnd-32 : offsetEnd])

	totalSize := new(big.Int).Add(bigOffsetEnd, lengthBig)
	if totalSize.BitLen() > 63 || totalSize.Cmp(outputLength) > 0 {
		return 0, 0, false
	}

	return offsetEnd, int(lengthBig.Uint64()), true
}

// tuplePointsTo mirrors go-ethereum's abi.tuplePointsTo for v1.17.2.
func tuplePointsTo(index int, output []byte) (start int, ok bool) {
	if index+32 > len(output) {
		return 0, false
	}

	offset := new(big.Int).SetBytes(output[index : index+32])
	if offset.Cmp(big.NewInt(int64(len(output)))) > 0 || offset.BitLen() > 63 {
		return 0, false
	}

	return int(offset.Uint64()), true
}
