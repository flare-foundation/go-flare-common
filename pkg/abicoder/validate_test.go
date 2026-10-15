package abicoder

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type oracleInner struct {
	D string
	E []byte
}

type oracleOuter struct {
	A *big.Int
	B string
	C oracleInner
}

// encodedWord returns the 32-byte big-endian encoding of v.
func encodedWord(v int) []byte {
	b := make([]byte, 32)
	for i := range 8 {
		b[31-i] = byte(uint64(v) >> (8 * i))
	}

	return b
}

// stringSliceAmplification builds a string[] argument encoding whose n element
// offsets all point at one blob of blobLen bytes, so a faithful decoder copies
// n*blobLen bytes out of a much smaller input.
func stringSliceAmplification(n, blobLen int) []byte {
	data := encodedWord(32) // head: offset to the array data
	data = append(data, encodedWord(n)...)
	blobOffset := n * 32 // relative to the start after the count word
	for range n {
		data = append(data, encodedWord(blobOffset)...)
	}
	data = append(data, encodedWord(blobLen)...)
	data = append(data, make([]byte, blobLen)...)
	if pad := (32 - blobLen%32) % 32; pad != 0 {
		data = append(data, make([]byte, pad)...)
	}

	return data
}

func newType(t *testing.T, solType string) abi.Type {
	t.Helper()
	typ, err := abi.NewType(solType, "", nil)
	require.NoError(t, err)

	return typ
}

func tupleType(t *testing.T, components []abi.ArgumentMarshaling) abi.Type {
	t.Helper()
	typ, err := abi.NewType("tuple", "", components)
	require.NoError(t, err)

	return typ
}

// TestCheckArgTypeRejectsUninitialized confirms a zero-value abi.Argument is
// rejected on both the encode and decode paths instead of being silently
// treated as int256.
func TestCheckArgTypeRejectsUninitialized(t *testing.T) {
	var zero abi.Argument // zero abi.Type: IntTy, size 0

	_, errEncode := Encode(zero, big.NewInt(8))
	require.ErrorIs(t, errEncode, errUninitializedArg)

	_, errDecode := Decode[*big.Int](zero, make([]byte, 32))
	require.ErrorIs(t, errDecode, errUninitializedArg)

	var dst *big.Int
	require.ErrorIs(t, DecodeTo(zero, make([]byte, 32), &dst), errUninitializedArg)
}

// TestCheckArgTypeAllowsValidIntTypes confirms the guard does not reject any
// legitimately constructed integer type.
func TestCheckArgTypeAllowsValidIntTypes(t *testing.T) {
	for _, solType := range []string{"int8", "int256", "uint8", "uint256", "int16", "uint128"} {
		require.NoError(t, checkArgType(newType(t, solType)), solType)
	}
}

type shapeStruct3 struct {
	A string
	B uint8
	C string
}

type shapeStruct2 struct {
	A string
	B uint8
}

// TestCheckShapeRejectsArrayLengthMismatch covers silent array truncation and
// zero-fill (F2): both must now be rejected, not silently reshaped.
func TestCheckShapeRejectsArrayLengthMismatch(t *testing.T) {
	arg := abi.Argument{Type: newType(t, "bytes32[2]")}
	a := [32]byte{0xaa}
	b := [32]byte{0xbb}

	packed, err := abi.Arguments{arg}.Pack([2][32]byte{a, b})
	require.NoError(t, err)

	_, errShort := Decode[[1][32]byte](arg, packed) // would drop element 2
	require.ErrorIs(t, errShort, ErrShapeMismatch)

	_, errLong := Decode[[3][32]byte](arg, packed) // would zero-fill element 3
	require.ErrorIs(t, errLong, ErrShapeMismatch)

	var short [1][32]byte
	require.ErrorIs(t, DecodeTo(arg, packed, &short), ErrShapeMismatch)
}

// TestCheckShapeRejectsFixedBytesTruncation covers bytes32 into a shorter array.
func TestCheckShapeRejectsFixedBytesTruncation(t *testing.T) {
	arg := abi.Argument{Type: newType(t, "bytes32")}
	packed, err := abi.Arguments{arg}.Pack([32]byte{1, 2, 3})
	require.NoError(t, err)

	_, err = Decode[[20]byte](arg, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
}

// TestCheckShapeRejectsFieldCountMismatch covers extra and missing struct fields.
func TestCheckShapeRejectsFieldCountMismatch(t *testing.T) {
	arg := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
		{Name: "A", Type: "string"},
		{Name: "B", Type: "uint8"},
	})}

	packed, err := abi.Arguments{arg}.Pack(shapeStruct2{"x", 7})
	require.NoError(t, err)

	_, errExtra := Decode[shapeStruct3](arg, packed) // dest has an extra field
	require.ErrorIs(t, errExtra, ErrShapeMismatch)

	type oneField struct{ A string }
	_, errMissing := Decode[oneField](arg, packed) // dest is missing a field
	require.ErrorIs(t, errMissing, ErrShapeMismatch)
}

// TestCheckShapeRejectsArrayForSlice covers a Go array for a dynamic ABI array
// or bytes, which go-ethereum truncates or zero-fills to the array's length.
func TestCheckShapeRejectsArrayForSlice(t *testing.T) {
	slice := abi.Argument{Name: "x", Type: newType(t, "uint256[]")}
	nested := abi.Argument{Name: "p", Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "x", Type: "uint256[]"}})}

	for _, n := range []int{1, 2, 3} {
		values := make([]*big.Int, n)
		for i := range values {
			values[i] = big.NewInt(int64(i + 1))
		}

		packed, err := Encode(slice, values)
		require.NoError(t, err)

		_, err = Decode[[2]*big.Int](slice, packed)
		require.ErrorIs(t, err, ErrShapeMismatch, n)

		packed, err = Encode(nested, struct{ X []*big.Int }{values})
		require.NoError(t, err)

		_, err = Decode[struct{ X [2]*big.Int }](nested, packed)
		require.ErrorIs(t, err, ErrShapeMismatch, n)
	}

	bytesArg := abi.Argument{Type: newType(t, "bytes")}
	packed, err := Encode(bytesArg, []byte{1, 2})
	require.NoError(t, err)

	_, err = Decode[[2]byte](bytesArg, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
}

// TestCheckShapeRejectsArrayForNonArray covers a Go array where the ABI type has
// none. go-ethereum truncates or zero-fills a string into it, and on a *big.Int
// recurses until a fatal stack overflow, so those are checked without decoding.
func TestCheckShapeRejectsArrayForNonArray(t *testing.T) {
	uint256 := abi.Argument{Name: "a", Type: newType(t, "uint256")}
	require.ErrorIs(t, checkShape[[1]*big.Int](uint256.Type), ErrShapeMismatch)
	require.ErrorIs(t, checkShape[[16]byte](newType(t, "int128")), ErrShapeMismatch)
	require.ErrorIs(t, checkShape[[]common.Hash](newType(t, "uint256[]")), ErrShapeMismatch)
	require.ErrorIs(t, checkShape[struct{ X [1]*big.Int }](tupleType(t, []abi.ArgumentMarshaling{{Name: "x", Type: "uint256"}})), ErrShapeMismatch)

	str := abi.Argument{Type: newType(t, "string")}
	packed, err := Encode(str, "BTC/USD-long")
	require.NoError(t, err)
	_, err = Decode[[4]byte](str, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = Decode[[16]byte](str, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)

	strs := abi.Argument{Type: newType(t, "string[]")}
	packed, err = Encode(strs, []string{"hello world", "ab"})
	require.NoError(t, err)
	_, err = Decode[[][4]byte](strs, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
}

func roundTrip[T any](t *testing.T, solType string, v T) {
	t.Helper()
	arg := abi.Argument{Type: newType(t, solType)}
	packed, err := Encode(arg, v)
	require.NoError(t, err)

	got, err := Decode[T](arg, packed)
	require.NoError(t, err)
	require.Equal(t, v, got)
}

// TestCheckShapeAllowsArraysForFixedTypes checks the types go-ethereum decodes
// to Go arrays still decode into them.
func TestCheckShapeAllowsArraysForFixedTypes(t *testing.T) {
	roundTrip(t, "address", common.Address{1})
	roundTrip(t, "bytes32", common.Hash{1})
	roundTrip(t, "bytes4", [4]byte{1})
	roundTrip(t, "function", [24]byte{1})
	roundTrip(t, "bytes32[2]", [2][32]byte{{1}, {2}})
}

// TestCheckShapeAllowsMatchingAndLeafFlexibility confirms the check does not
// over-reject: matching shapes, big.Int vs *big.Int, common.Hash, and any all
// decode, and positional (name-blind) tuple mapping still works.
func TestCheckShapeAllowsMatchingAndLeafFlexibility(t *testing.T) {
	t.Run("bigIntVariants", func(t *testing.T) {
		arg := abi.Argument{Type: newType(t, "uint256")}
		packed, err := abi.Arguments{arg}.Pack(big.NewInt(42))
		require.NoError(t, err)

		ptr, err := Decode[*big.Int](arg, packed)
		require.NoError(t, err)
		require.Equal(t, int64(42), ptr.Int64())

		val, err := Decode[big.Int](arg, packed)
		require.NoError(t, err)
		require.Equal(t, int64(42), val.Int64())
	})

	t.Run("anyDest", func(t *testing.T) {
		arg := abi.Argument{Type: newType(t, "string")}
		packed, err := abi.Arguments{arg}.Pack("hi")
		require.NoError(t, err)

		out, err := Decode[any](arg, packed)
		require.NoError(t, err)
		require.Equal(t, "hi", out)
	})

	t.Run("positionalNameBlind", func(t *testing.T) {
		// Same field count and types, names swapped: still decodes positionally.
		arg := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
			{Name: "A", Type: "uint64"},
			{Name: "B", Type: "uint64"},
		})}
		type src struct{ A, B uint64 }
		packed, err := abi.Arguments{arg}.Pack(src{A: 1, B: 2})
		require.NoError(t, err)

		type swapped struct{ B, A uint64 }
		out, err := Decode[swapped](arg, packed)
		require.NoError(t, err)
		require.Equal(t, swapped{B: 1, A: 2}, out) // positional: by index, names ignored
	})
}
