package abicoder

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

		_, err = DecodeArguments[[2]*big.Int](abi.Arguments{slice}, packed)
		require.ErrorIs(t, err, ErrShapeMismatch, n)

		packed, err = Encode(nested, struct{ X []*big.Int }{values})
		require.NoError(t, err)

		_, err = Decode[struct{ X [2]*big.Int }](nested, packed)
		require.ErrorIs(t, err, ErrShapeMismatch, n)

		_, err = DecodeArguments[struct{ P struct{ X [2]*big.Int } }](abi.Arguments{nested}, packed)
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
	require.ErrorIs(t, checkArgsShape(abi.Arguments{uint256}, reflect.TypeFor[[1]*big.Int]()), ErrShapeMismatch)

	//nolint:asciicheck // go-ethereum copies a into A, not into the tagged field
	type hidden struct {
		Ñ *big.Int `abi:"a"`
		A [1]*big.Int
		B *big.Int
	}
	args := abi.Arguments{uint256, {Name: "b", Type: newType(t, "uint256")}}
	require.ErrorIs(t, checkArgsShape(args, reflect.TypeFor[hidden]()), ErrShapeMismatch)

	str := abi.Argument{Type: newType(t, "string")}
	packed, err := Encode(str, "BTC/USD-long")
	require.NoError(t, err)
	_, err = Decode[[4]byte](str, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = Decode[[16]byte](str, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = DecodeArguments[[3]byte](abi.Arguments{str}, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)

	strs := abi.Argument{Type: newType(t, "string[]")}
	packed, err = Encode(strs, []string{"hello world", "ab"})
	require.NoError(t, err)
	_, err = Decode[[][4]byte](strs, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
}

// convertPanics reports whether go-ethereum's ConvertType, which Decode ends
// with, panics converting in to a T.
func convertPanics[T any](in any) (panicked bool) {
	defer func() { panicked = recover() != nil }()
	abi.ConvertType(in, new(T))

	return false
}

// shapeCase is a destination type checked for arg against what go-ethereum's
// ConvertType does with the decoding of sample.
type shapeCase struct {
	name     string
	arg      abi.Argument
	sample   any
	shape    func(abi.Type) error
	panics   func(any) bool
	rejected bool
}

func newShapeCase[T any](name string, arg abi.Argument, sample any, rejected bool) shapeCase {
	return shapeCase{name, arg, sample, checkShape[T], convertPanics[T], rejected}
}

// runShapeCases checks that checkShape refuses each case exactly where
// go-ethereum panics converting into it.
func runShapeCases(t *testing.T, cases []shapeCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			args := abi.Arguments{tc.arg}
			packed, err := args.Pack(tc.sample)
			require.NoError(t, err)
			values, err := args.Unpack(packed)
			require.NoError(t, err)

			require.Equal(t, tc.rejected, tc.panics(values[0]))
			if tc.rejected {
				require.ErrorIs(t, tc.shape(tc.arg.Type), ErrShapeMismatch)
			} else {
				require.NoError(t, tc.shape(tc.arg.Type))
			}
		})
	}
}

// TestCheckShapeRejectsPointers checks that a pointer to anything but a big.Int
// is refused exactly where go-ethereum, which never allocates one, panics
// decoding through it.
func TestCheckShapeRejectsPointers(t *testing.T) {
	type bigPtr *big.Int
	type uintField struct{ A *uint64 }
	type tupleField struct{ A *struct{ V uint64 } }

	u64 := abi.Argument{Name: "a", Type: newType(t, "uint64")}
	u256 := abi.Argument{Type: newType(t, "uint256")}
	tuple := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint64"}})}
	nested := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
		{Name: "a", Type: "tuple", Components: []abi.ArgumentMarshaling{{Name: "v", Type: "uint64"}}},
	})}

	runShapeCases(t, []shapeCase{
		newShapeCase[*uint64]("value", u64, uint64(1), true),
		newShapeCase[*struct{ A uint64 }]("struct", tuple, struct{ A uint64 }{1}, true),
		newShapeCase[uintField]("field", tuple, struct{ A uint64 }{1}, true),
		newShapeCase[tupleField]("tupleField", nested, struct{ A struct{ V uint64 } }{}, true),
		newShapeCase[[2]*uint64]("arrayElem", abi.Argument{Type: newType(t, "uint64[2]")}, [2]uint64{}, true),
		newShapeCase[[]*uint64]("sliceElem", abi.Argument{Type: newType(t, "uint64[]")}, []uint64{1}, true),
		newShapeCase[*common.Hash]("hash", abi.Argument{Type: newType(t, "bytes32")}, common.Hash{1}, true),
		newShapeCase[**big.Int]("bigIntPtr", u256, big.NewInt(1), true),
		newShapeCase[*big.Int]("bigInt", u256, big.NewInt(1), false),
		newShapeCase[bigPtr]("namedBigInt", u256, big.NewInt(1), false),
	})

	packed, err := Encode(u64, uint64(1))
	require.NoError(t, err)
	_, err = Decode[*uint64](u64, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = DecodeArguments[uintField](abi.Arguments{u64}, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = New[uintField](tuple)
	require.ErrorIs(t, err, ErrShapeMismatch)

	// Encode never checks the shape, so it still takes pointers
	x := uint64(1)
	_, err = Encode(tuple, uintField{&x})
	require.NoError(t, err)
}

// TestCheckShapeChecksLeaves checks that a leaf is refused exactly where
// go-ethereum can't assign the decoded value to it.
func TestCheckShapeChecksLeaves(t *testing.T) {
	type amount big.Int
	type count uint64
	type pair struct {
		A uint64
		B string
	}
	type renamed struct {
		P uint64
		Q string
	}
	type wrongKind struct {
		A string
		B string
	}
	type unexported struct {
		A uint64
		b string //nolint:unused // go-ethereum can't set it
	}

	arg := func(solType string) abi.Argument { return abi.Argument{Type: newType(t, solType)} }
	u256 := arg("uint256")
	u64 := arg("uint64")
	components := []abi.ArgumentMarshaling{{Name: "a", Type: "uint64"}, {Name: "b", Type: "string"}}
	tuple := abi.Argument{Type: tupleType(t, components)}
	tuples, err := abi.NewType("tuple[]", "", components)
	require.NoError(t, err)
	bigTuple := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint256"}})}

	runShapeCases(t, []shapeCase{
		newShapeCase[big.Int]("bigIntValue", u256, big.NewInt(1), false),
		newShapeCase[amount]("namedBigIntValue", u256, big.NewInt(1), false),
		newShapeCase[uint64]("uint64ForUint256", u256, big.NewInt(1), true),
		newShapeCase[string]("stringForUint256", u256, big.NewInt(1), true),
		newShapeCase[any]("any", u256, big.NewInt(1), false),
		newShapeCase[fmt.Stringer]("implemented", u256, big.NewInt(1), false),
		newShapeCase[error]("notImplemented", u256, big.NewInt(1), true),
		newShapeCase[count]("namedUint64", u64, uint64(1), true),
		newShapeCase[int]("int", u64, uint64(1), true),
		newShapeCase[uint32]("narrower", u64, uint64(1), true),
		newShapeCase[uint32]("bigIntWidth", arg("uint24"), big.NewInt(1), true),
		newShapeCase[common.Hash]("hash", arg("bytes32"), common.Hash{1}, false),
		newShapeCase[[]byte]("sliceForBytes32", arg("bytes32"), common.Hash{1}, true),
		newShapeCase[[20]byte]("addressArray", arg("address"), common.Address{1}, false),
		newShapeCase[hexutil.Bytes]("hexBytes", arg("bytes"), []byte{1}, false),
		newShapeCase[string]("stringForBytes", arg("bytes"), []byte{1}, true),
		newShapeCase[[]byte]("bytesForString", arg("string"), "x", true),
		newShapeCase[int]("intForBool", arg("bool"), true, true),
		newShapeCase[[]any]("anyElem", arg("uint64[]"), []uint64{1}, false),
		newShapeCase[[]int]("sliceElem", arg("uint64[]"), []uint64{1}, true),
		newShapeCase[[2]int]("arrayElem", arg("uint64[2]"), [2]uint64{1, 2}, true),
		newShapeCase[pair]("tuple", tuple, pair{1, "x"}, false),
		newShapeCase[renamed]("renamed", tuple, pair{1, "x"}, false),
		newShapeCase[wrongKind]("tupleField", tuple, pair{1, "x"}, true),
		newShapeCase[unexported]("unexported", tuple, pair{1, "x"}, true),
		newShapeCase[struct{ A big.Int }]("bigIntField", bigTuple, struct{ A *big.Int }{big.NewInt(1)}, true),
		newShapeCase[struct{ A any }]("anyField", bigTuple, struct{ A *big.Int }{big.NewInt(1)}, false),
		newShapeCase[[]wrongKind]("tupleSliceField", abi.Argument{Type: tuples}, []pair{{1, "x"}}, true),
	})

	packed, err := Encode(tuple, pair{1, "x"})
	require.NoError(t, err)
	_, err = Decode[wrongKind](tuple, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = New[wrongKind](tuple)
	require.ErrorIs(t, err, ErrShapeMismatch)
	_, err = DecodeArguments[struct{ A string }](abi.Arguments{u64}, make([]byte, 32))
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
// decode, and fields named after no component still map by position.
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

	t.Run("positionalOtherNames", func(t *testing.T) {
		arg := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
			{Name: "A", Type: "uint64"},
			{Name: "B", Type: "uint64"},
		})}
		type src struct{ A, B uint64 }
		packed, err := abi.Arguments{arg}.Pack(src{A: 1, B: 2})
		require.NoError(t, err)

		type renamed struct{ P, Q uint64 }
		out, err := Decode[renamed](arg, packed)
		require.NoError(t, err)
		require.Equal(t, renamed{P: 1, Q: 2}, out)

		type tagged struct {
			P uint64 `abi:"A"`
			Q uint64 `abi:"B"`
		}
		outTagged, err := Decode[tagged](arg, packed)
		require.NoError(t, err)
		require.Equal(t, tagged{P: 1, Q: 2}, outTagged)
	})
}

// TestArgFieldSkipsNonASCIITags covers a field tagged with a component name
// whose own name starts with a non-ASCII letter: go-ethereum ignores the tag
// and uses the field with the camel-cased name, so the checks must look there.
func TestArgFieldSkipsNonASCIITags(t *testing.T) {
	//nolint:asciicheck // the non-ASCII initial is the case under test
	type tagged struct {
		Ñ *big.Int `abi:"a"`
		A *big.Int
	}
	single := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint256"}})}
	_, err := Encode(single, tagged{big.NewInt(1), new(big.Int).Add(pow2(256), big.NewInt(5))})
	require.ErrorIs(t, err, ErrNonCanonical)

	//nolint:asciicheck // the non-ASCII initial is the case under test
	type swapped struct {
		Ñ uint64 `abi:"a"`
		A uint64
	}
	pair := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint64"}, {Name: "b", Type: "uint64"}})}
	packed, err := Encode(pair, struct{ A, B uint64 }{1, 2})
	require.NoError(t, err)
	_, err = Decode[swapped](pair, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)

	//nolint:asciicheck // the non-ASCII initial is the case under test
	type truncated struct {
		Ñ []*big.Int `abi:"a"`
		A [2]*big.Int
		B *big.Int
	}
	args := abi.Arguments{{Name: "a", Type: newType(t, "uint256[]")}, {Name: "b", Type: newType(t, "uint256")}}
	packed, err = args.Pack([]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}, big.NewInt(9))
	require.NoError(t, err)
	_, err = DecodeArguments[truncated](args, packed)
	require.ErrorIs(t, err, ErrShapeMismatch)
}

func decodeOne[T any](arg abi.Argument, data []byte) func() error {
	return func() error {
		_, err := Decode[T](arg, data)
		return err
	}
}

// TestCheckShapeRejectsReorderedFields covers structs whose field names or abi
// tags put a component at another position than Decode fills: Encode reads
// them by name, so Decode would swap the values.
func TestCheckShapeRejectsReorderedFields(t *testing.T) {
	pair := []abi.ArgumentMarshaling{{Name: "A", Type: "uint64"}, {Name: "B", Type: "uint64"}}
	arg := abi.Argument{Name: "p", Type: tupleType(t, pair)}
	type src struct{ A, B uint64 }
	packed, err := Encode(arg, src{A: 1, B: 2})
	require.NoError(t, err)

	type swapped struct{ B, A uint64 }
	type tagged struct {
		X uint64 `abi:"B"`
		Y uint64 `abi:"A"`
	}
	type partial struct{ B, Q uint64 }

	tests := []struct {
		name   string
		decode func() error
	}{
		{"swapped", decodeOne[swapped](arg, packed)},
		{"tagged", decodeOne[tagged](arg, packed)},
		{"partial", decodeOne[partial](arg, packed)},
		{"decodeArguments", decodeInto[struct{ P swapped }](abi.Arguments{arg}, packed)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.ErrorIs(t, tc.decode(), ErrShapeMismatch)
		})
	}

	t.Run("nested", func(t *testing.T) {
		list := abi.Argument{Name: "ps", Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "ps", Type: "tuple[]", Components: pair}})}
		packed, err := Encode(list, struct{ Ps []src }{[]src{{1, 2}}})
		require.NoError(t, err)

		_, err = Decode[struct{ Ps []swapped }](list, packed)
		require.ErrorIs(t, err, ErrShapeMismatch)
	})
}
