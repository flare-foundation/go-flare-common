package abicoder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const word = 32

var errUnsupportedType = errors.New("abi type not supported")

// checkLayout checks that args' types are supported and data is their
// canonical encoding.
func checkLayout(args abi.Arguments, data []byte) error {
	for i := range args {
		if err := checkType(&args[i].Type); err != nil {
			return err
		}
	}

	return checkData(args, data)
}

// checkData checks that data is the canonical encoding of args, whose types
// passed checkType, without decoding it: every offset is the one the encoder
// writes, every padding byte is zero, every value fits its type, and nothing
// follows. go-ethereum decodes a region once per offset pointing at it, so this
// bounds allocation by len(data).
func checkData(args abi.Arguments, data []byte) error {
	n, err := checkSequence(len(args), func(i int) *abi.Type { return &args[i].Type }, data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("%w: %d bytes after the encoding", ErrNonCanonical, len(data)-n)
	}

	return nil
}

// checkSequence checks n values laid out as a tuple from the start of b,
// typed by elem, and returns the encoding's length.
func checkSequence(n int, elem func(int) *abi.Type, b []byte) (int, error) {
	head := 0
	for i := range n {
		// every head is at least a word, so this bounds the loop by len(b)
		s := headSize(elem(i))
		if s > len(b)-head {
			return 0, fmt.Errorf("%w: %d values need more than %d bytes", ErrTruncated, n, len(b))
		}
		head += s
	}

	pos, end := 0, head
	for i := range n {
		t := elem(i)
		if !isDynamic(t) {
			s, err := checkValue(t, b[pos:])
			if err != nil {
				return 0, err
			}
			pos += s

			continue
		}

		off, err := readNumber(b[pos:], len(b))
		if err != nil {
			return 0, err
		}
		if off != end {
			return 0, fmt.Errorf("%w: offset %d, the encoder writes %d", ErrNonCanonical, off, end)
		}
		s, err := checkValue(t, b[end:])
		if err != nil {
			return 0, err
		}
		pos += word
		end += s
	}

	return end, nil
}

// checkValue checks one value of type t at the start of b and returns its
// encoding's length.
func checkValue(t *abi.Type, b []byte) (int, error) {
	switch t.T {
	case abi.StringTy, abi.BytesTy:
		l, err := readNumber(b, len(b))
		if err != nil {
			return 0, err
		}
		padded := (l + word - 1) / word * word
		if word+padded > len(b) {
			return 0, fmt.Errorf("%w: %d bytes of content need more than %d", ErrTruncated, l, len(b))
		}
		if !isZero(b[word+l : word+padded]) {
			return 0, fmt.Errorf("%w: non-zero padding", ErrNonCanonical)
		}

		return word + padded, nil
	case abi.SliceTy:
		l, err := readNumber(b, len(b))
		if err != nil {
			return 0, err
		}
		n, err := checkSequence(l, func(int) *abi.Type { return t.Elem }, b[word:])

		return word + n, err
	case abi.ArrayTy:
		return checkSequence(t.Size, func(int) *abi.Type { return t.Elem }, b)
	case abi.TupleTy:
		return checkSequence(len(t.TupleElems), func(i int) *abi.Type { return t.TupleElems[i] }, b)
	default:
		return checkWord(t, b)
	}
}

// checkWord checks a value of a static elementary type t at the start of b.
func checkWord(t *abi.Type, b []byte) (int, error) {
	if len(b) < word {
		return 0, fmt.Errorf("%w: a word needs more than %d bytes", ErrTruncated, len(b))
	}
	w := b[:word]

	var ok bool
	switch t.T {
	case abi.UintTy, abi.IntTy:
		k := word - t.Size/8
		fill := byte(0)
		if t.T == abi.IntTy && w[k]&0x80 != 0 {
			fill = 0xff
		}
		ok = !slices.ContainsFunc(w[:k], func(c byte) bool { return c != fill })
	case abi.BoolTy:
		ok = isZero(w[:word-1]) && w[word-1] <= 1
	case abi.AddressTy:
		ok = isZero(w[:word-20])
	case abi.FixedBytesTy:
		ok = isZero(w[t.Size:])
	case abi.FunctionTy:
		ok = isZero(w[24:])
	case abi.HashTy:
		ok = true
	default:
		return 0, fmt.Errorf("%w: %s", errUnsupportedType, t)
	}
	if !ok {
		return 0, fmt.Errorf("%w: %s padded with non-zero bytes or out of range", ErrNonCanonical, t)
	}

	return word, nil
}

// readNumber reads an offset or a length from the start of b, at most limit.
func readNumber(b []byte, limit int) (int, error) {
	if len(b) < word {
		return 0, fmt.Errorf("%w: a word needs more than %d bytes", ErrTruncated, len(b))
	}
	v := binary.BigEndian.Uint64(b[word-8 : word])
	if !isZero(b[:word-8]) || v > uint64(limit) {
		return 0, fmt.Errorf("%w: offset or length past the end", ErrTruncated)
	}

	return int(v), nil
}

// checkType rejects types whose encoding checkData can't bound or
// go-ethereum can't decode: zero-size arrays, empty tuples, heads overflowing
// int, and unknown kinds or sizes.
func checkType(t *abi.Type) error {
	ok := true
	switch t.T {
	case abi.UintTy, abi.IntTy:
		ok = t.Size > 0 && t.Size <= 8*word && t.Size%8 == 0
	case abi.FixedBytesTy:
		ok = t.Size > 0 && t.Size <= word
	case abi.BoolTy, abi.AddressTy, abi.FunctionTy, abi.HashTy, abi.StringTy, abi.BytesTy:
	case abi.SliceTy:
		return checkType(t.Elem)
	case abi.ArrayTy:
		if err := checkType(t.Elem); err != nil {
			return err
		}
		ok = t.Size > 0 && headSize(t.Elem) <= math.MaxInt/t.Size
	case abi.TupleTy:
		ok = len(t.TupleElems) > 0
		total := 0
		for _, e := range t.TupleElems {
			if err := checkType(e); err != nil {
				return err
			}
			s := headSize(e)
			if s > math.MaxInt-total {
				ok = false
				break
			}
			total += s
		}
	default:
		ok = false
	}
	if !ok {
		return fmt.Errorf("%w: %s", errUnsupportedType, t)
	}

	return nil
}

// headSize returns what t, which passed checkType, takes in a tuple's head:
// an offset if t is dynamic, else its value.
func headSize(t *abi.Type) int {
	if isDynamic(t) {
		return word
	}

	switch t.T {
	case abi.ArrayTy:
		return t.Size * headSize(t.Elem)
	case abi.TupleTy:
		total := 0
		for _, e := range t.TupleElems {
			total += headSize(e)
		}

		return total
	default:
		return word
	}
}

func isDynamic(t *abi.Type) bool {
	switch t.T {
	case abi.StringTy, abi.BytesTy, abi.SliceTy:
		return true
	case abi.ArrayTy:
		return isDynamic(t.Elem)
	case abi.TupleTy:
		return slices.ContainsFunc(t.TupleElems, isDynamic)
	default:
		return false
	}
}

func isZero(b []byte) bool {
	return !slices.ContainsFunc(b, func(c byte) bool { return c != 0 })
}
