package abicoder

import (
	"math"
	"math/big"
	"reflect"
	"runtime"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

// word returns the 32-byte big-endian encoding of v.
func word(v int) []byte {
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
	data := word(32) // head: offset to the array data
	data = append(data, word(n)...)
	blobOffset := n * 32 // relative to the start after the count word
	for range n {
		data = append(data, word(blobOffset)...)
	}
	data = append(data, word(blobLen)...)
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

// TestExceedsStringBudgetRejectsAmplification confirms the pre-Unpack guard
// rejects offset-aliasing string encodings before any large allocation.
func TestExceedsStringBudgetRejectsAmplification(t *testing.T) {
	arg := abi.Argument{Name: "x", Type: newType(t, "string[]")}
	data := stringSliceAmplification(3000, 30000)

	require.True(t, exceedsStringBudget(arg.Type, data), "aliasing string[] must be flagged")

	_, err := Decode[[]string](arg, data)
	require.ErrorIs(t, err, ErrNonCanonical)

	var dst []string
	require.ErrorIs(t, DecodeTo(arg, data, &dst), ErrNonCanonical)
}

// TestRejectsAmplificationBeforeAllocation confirms the guard fixes the DoS,
// not merely the result: rejecting a 124 KB aliasing input that would otherwise
// materialize ~90 MB of strings must itself stay far below that.
func TestRejectsAmplificationBeforeAllocation(t *testing.T) {
	arg := abi.Argument{Name: "x", Type: newType(t, "string[]")}
	data := stringSliceAmplification(3000, 30000) // would decode to 90 MB

	var m0, m1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m0)
	_, err := Decode[[]string](arg, data)
	runtime.ReadMemStats(&m1)

	require.ErrorIs(t, err, ErrNonCanonical)
	require.Less(t, m1.TotalAlloc-m0.TotalAlloc, uint64(10<<20), "rejection must not allocate the amplified payload")
}

// TestExceedsStringBudgetAmplificationInTuple covers a string nested in a tuple
// field and in a dynamic array, the other reachable amplification shapes.
func TestExceedsStringBudgetAmplificationInTuple(t *testing.T) {
	// (string a, string b) where both offsets alias one blob.
	tup := tupleType(t, []abi.ArgumentMarshaling{
		{Name: "A", Type: "string"},
		{Name: "B", Type: "string"},
	})

	blobLen := 40000
	body := word(64)                 // field A offset (within the tuple body)
	body = append(body, word(64)...) // field B offset -> same blob as A
	body = append(body, word(blobLen)...)
	body = append(body, make([]byte, blobLen)...)
	if pad := (32 - blobLen%32) % 32; pad != 0 {
		body = append(body, make([]byte, pad)...)
	}
	data := append(word(32), body...) // head pointer to the dynamic tuple

	require.True(t, exceedsStringBudget(tup, data), "aliasing tuple strings must be flagged")
}

// TestExceedsStringBudgetAllowsCanonical is the no-false-rejection guard: a
// genuinely large but canonical string[] must decode, not be rejected.
func TestExceedsStringBudgetAllowsCanonical(t *testing.T) {
	arg := abi.Argument{Name: "x", Type: newType(t, "string[]")}

	in := make([]string, 500)
	for i := range in {
		in[i] = "a string element that is reasonably long to exercise the path"
	}

	packed, err := abi.Arguments{arg}.Pack(in)
	require.NoError(t, err)

	require.False(t, exceedsStringBudget(arg.Type, packed), "canonical string[] must not be flagged")

	out, err := Decode[[]string](arg, packed)
	require.NoError(t, err)
	require.Equal(t, in, out)
}

// TestExceedsStringBudgetDefersOnMalformed confirms the guard never rejects on
// malformed or truncated input: it returns false so the real Unpack surfaces
// the error, and ErrNonCanonical is reserved for confirmed amplification.
func TestExceedsStringBudgetDefersOnMalformed(t *testing.T) {
	sliceType := newType(t, "string[]")

	tests := []struct {
		name string
		data []byte
	}{
		{"empty", nil},
		{"shortHead", word(32)[:16]},
		{"offsetPastEnd", word(1 << 40)},
		{"truncatedCount", append(word(32), word(5)[:8]...)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.False(t, exceedsStringBudget(sliceType, test.data))
		})
	}
}

// TestExceedsStringBudgetSkipsStringFreeTypes confirms the guard is a no-op for
// types that cannot reach a string (every in-repo usage), so their decode path
// is unchanged.
func TestExceedsStringBudgetSkipsStringFreeTypes(t *testing.T) {
	for _, solType := range []string{"uint256", "bytes", "bytes32", "uint256[]", "bytes[]"} {
		require.False(t, containsString(newType(t, solType)), solType)
		require.False(t, exceedsStringBudget(newType(t, solType), make([]byte, 64)), solType)
	}
}

// reflectStringBytes sums the byte length of every string reachable in v.
func reflectStringBytes(v reflect.Value) int {
	switch v.Kind() {
	case reflect.String:
		return v.Len()
	case reflect.Slice, reflect.Array:
		total := 0
		for i := range v.Len() {
			total += reflectStringBytes(v.Index(i))
		}
		return total
	case reflect.Struct:
		total := 0
		for i := range v.NumField() {
			total += reflectStringBytes(v.Field(i))
		}
		return total
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return 0
		}
		return reflectStringBytes(v.Elem())
	default:
		return 0
	}
}

// TestStringScanMatchesGoEthereum is the faithfulness oracle: for canonical
// encodings of every string-bearing shape, the walker's string-byte sum must
// equal the actual bytes go-ethereum allocates, proving the guard cannot
// over-count and therefore cannot false-reject canonical input.
func TestStringScanMatchesGoEthereum(t *testing.T) {
	nestedTuple := []abi.ArgumentMarshaling{
		{Name: "A", Type: "uint256"},
		{Name: "B", Type: "string"},
		{Name: "C", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "D", Type: "string"},
			{Name: "E", Type: "bytes"},
		}},
	}

	tests := []struct {
		name  string
		typ   abi.Type
		value any
	}{
		{"string", newType(t, "string"), "hello world"},
		{"emptyString", newType(t, "string"), ""},
		{"stringSlice", newType(t, "string[]"), []string{"a", "bb", "ccc"}},
		{"emptyStringSlice", newType(t, "string[]"), []string{}},
		{"stringArray", newType(t, "string[2]"), [2]string{"xx", "yyyy"}},
		{"stringNestedSlice", newType(t, "string[][]"), [][]string{{"a", "bb"}, {"ccc"}}},
		{"bytesNoStrings", newType(t, "bytes"), []byte{1, 2, 3, 4}},
		{"uintSliceNoStrings", newType(t, "uint256[]"), []*big.Int{big.NewInt(1), big.NewInt(2)}},
		{"tupleWithString", tupleType(t, nestedTuple), oracleOuter{
			A: big.NewInt(7),
			B: "outer",
			C: oracleInner{D: "inner", E: []byte{9, 9}},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			packed, err := abi.Arguments{{Type: test.typ}}.Pack(test.value)
			require.NoError(t, err)

			decoded, err := abi.Arguments{{Type: test.typ}}.Unpack(packed)
			require.NoError(t, err)
			wantSum := reflectStringBytes(reflect.ValueOf(decoded[0]))

			s := stringScan{limit: math.MaxInt, budget: math.MaxInt}
			ok := s.walk(test.typ, packed, 0)
			require.True(t, ok, "walk should not defer on canonical data")
			require.Equal(t, wantSum, s.sum, "walker string-byte sum must equal go-ethereum's")

			require.False(t, exceedsStringBudget(test.typ, packed))
		})
	}
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
