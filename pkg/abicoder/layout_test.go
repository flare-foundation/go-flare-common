package abicoder

import (
	"bytes"
	"math/big"
	"runtime"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// aliasedSlice encodes a single dynamic array of n elements whose offsets all
// point at one encoded element, elem.
func aliasedSlice(n int, elem []byte) []byte {
	data := encodedWord(32)
	data = append(data, encodedWord(n)...)
	for range n {
		data = append(data, encodedWord(n*32)...)
	}

	return append(data, elem...)
}

// aliasedPairs encodes a dynamic array of n canonically placed pairs of
// dynamic values, as (string,string)[] or string[2][], whose offsets all point
// at one encoded value, elem, after the last pair.
func aliasedPairs(n int, elem []byte) []byte {
	data := encodedWord(32)
	data = append(data, encodedWord(n)...)
	for i := range n {
		data = append(data, encodedWord(32*n+64*i)...)
	}
	for i := range n {
		at := encodedWord(64 * (n - i))
		data = append(append(data, at...), at...)
	}

	return append(data, elem...)
}

// allocated returns the bytes f allocates.
func allocated(f func()) uint64 {
	var m0, m1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m0)
	f()
	runtime.ReadMemStats(&m1)

	return m1.TotalAlloc - m0.TotalAlloc
}

// TestRejectsAmplificationBeforeAllocation confirms offsets aliasing one
// region are rejected before go-ethereum decodes, or re-encodes, the region
// once per offset.
func TestRejectsAmplificationBeforeAllocation(t *testing.T) {
	blob := append(encodedWord(30000), make([]byte, 30016)...)
	words := encodedWord(3000)
	for range 3000 {
		words = append(words, encodedWord(1)...)
	}
	pairs, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{{Name: "A", Type: "string"}, {Name: "B", Type: "string"}})
	require.NoError(t, err)

	tests := []struct {
		name string
		typ  abi.Type
		data []byte
	}{
		// unguarded, each allocates about 1 GB or more
		{"stringSlice", newType(t, "string[]"), stringSliceAmplification(3000, 30000)},
		{"bytesSlice", newType(t, "bytes[]"), aliasedSlice(3000, blob)},
		{"nestedSlice", newType(t, "uint256[][]"), aliasedSlice(3000, words)},
		{"tupleFields", pairs, aliasedPairs(1500, blob)},
		{"arrayElements", newType(t, "string[2][]"), aliasedPairs(1500, blob)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			n := allocated(func() { _, err = Decode[any](abi.Argument{Type: test.typ}, test.data) })
			require.ErrorIs(t, err, ErrNonCanonical)
			require.Less(t, n, uint64(1<<20), "Decode must reject before allocating")

			n = allocated(func() { _, err = DecodeArguments[any](abi.Arguments{{Type: test.typ}}, test.data) })
			require.ErrorIs(t, err, ErrNonCanonical)
			require.Less(t, n, uint64(1<<20), "DecodeArguments must reject before allocating")
		})
	}
}

type staticPair struct {
	A *big.Int
	B bool
}

// TestCheckLayoutAllowsCanonical is the no-false-rejection guard: go-ethereum's
// encoding of every shape passes and decodes back to its value.
func TestCheckLayoutAllowsCanonical(t *testing.T) {
	long := make([]string, 500)
	for i := range long {
		long[i] = "a string element that is reasonably long to exercise the path"
	}
	nested := []abi.ArgumentMarshaling{
		{Name: "A", Type: "uint256"},
		{Name: "B", Type: "string"},
		{Name: "C", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "D", Type: "string"},
			{Name: "E", Type: "bytes"},
		}},
	}
	outer := oracleOuter{A: big.NewInt(7), B: "outer", C: oracleInner{D: "inner", E: []byte{9, 9}}}
	pair := []abi.ArgumentMarshaling{{Name: "A", Type: "uint256"}, {Name: "B", Type: "bool"}}
	pairArray, err := abi.NewType("tuple[2]", "", pair)
	require.NoError(t, err)
	outerSlice, err := abi.NewType("tuple[]", "", nested)
	require.NoError(t, err)

	tests := []struct {
		name  string
		typ   abi.Type
		value any
	}{
		{"uint8", newType(t, "uint8"), uint8(7)},
		{"int8", newType(t, "int8"), int8(-3)},
		{"int128", newType(t, "int128"), big.NewInt(-5)},
		{"uint128Max", newType(t, "uint128"), new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))},
		{"bool", newType(t, "bool"), true},
		{"address", newType(t, "address"), common.HexToAddress("0xc0")},
		{"bytes4", newType(t, "bytes4"), [4]byte{1, 2, 3, 4}},
		{"bytes", newType(t, "bytes"), bytes.Repeat([]byte{0xff}, 33)},
		{"emptyString", newType(t, "string"), ""},
		{"stringSlice", newType(t, "string[]"), long},
		{"emptyStringSlice", newType(t, "string[]"), []string{}},
		{"stringArray", newType(t, "string[2]"), [2]string{"xx", "yyyy"}},
		{"stringNestedSlice", newType(t, "string[][]"), [][]string{{"a", "bb"}, {"ccc"}}},
		{"bytesNestedSlice", newType(t, "bytes[][]"), [][][]byte{{{1}, {}}, {}, {{2, 3}}}},
		{"uintArray", newType(t, "uint256[3]"), [3]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
		{"arraySlice", newType(t, "uint8[2][]"), [][2]uint8{{1, 2}, {3, 4}}},
		{"sliceArray", newType(t, "string[][2]"), [2][]string{{"a"}, {"b", "c"}}},
		{"tuple", tupleType(t, nested), outer},
		{"tupleSlice", outerSlice, []oracleOuter{outer, outer}},
		// a dynamic value whose offset exceeds what follows its head slot
		{"stringAfterThreeWords", tupleType(t, []abi.ArgumentMarshaling{{Name: "A", Type: "uint256"}, {Name: "B", Type: "uint256"}, {Name: "C", Type: "uint256"}, {Name: "D", Type: "string"}}), struct {
			A, B, C *big.Int
			D       string
		}{big.NewInt(1), big.NewInt(2), big.NewInt(3), "x"}},
		{"staticTupleArray", pairArray, [2]staticPair{{big.NewInt(1), true}, {big.NewInt(2), false}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arg := abi.Argument{Type: test.typ}
			packed, err := abi.Arguments{arg}.Pack(test.value)
			require.NoError(t, err)

			require.NoError(t, checkLayout(abi.Arguments{arg}, packed))

			decoded, err := Decode[any](arg, packed)
			require.NoError(t, err)
			again, err := abi.Arguments{arg}.Pack(decoded)
			require.NoError(t, err)
			require.Equal(t, packed, again)
		})
	}
}

// TestCheckLayoutRejectsValuesOutOfRange covers values go-ethereum decodes
// without checking: integers wider than their type, and dirty padding.
func TestCheckLayoutRejectsValuesOutOfRange(t *testing.T) {
	high := encodedWord(0)
	high[15] = 1 // 2^128
	positive := bytes.Repeat([]byte{0xff}, 32)
	positive[16] = 0x7f // int128's leading bytes say negative, its value positive
	dirty := encodedWord(0)
	dirty[31] = 1
	dirtyAddress := encodedWord(0)
	dirtyAddress[0] = 1

	tests := []struct {
		name    string
		solType string
		data    []byte
	}{
		{"uint128", "uint128", high},
		{"int128", "int128", positive},
		{"uint8", "uint8", encodedWord(256)},
		{"bool", "bool", encodedWord(2)},
		{"bytes4", "bytes4", dirty},
		{"address", "address", dirtyAddress},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Decode[any](abi.Argument{Type: newType(t, test.solType)}, test.data)
			require.ErrorIs(t, err, ErrNonCanonical)
		})
	}
}

func TestCheckLayoutRejectsTruncated(t *testing.T) {
	tests := []struct {
		name    string
		solType string
		data    []byte
	}{
		{"empty", "string[]", nil},
		{"shortHead", "string[]", encodedWord(32)[:16]},
		{"offsetPastEnd", "string[]", encodedWord(1 << 40)},
		{"truncatedCount", "string[]", append(encodedWord(32), encodedWord(5)[:8]...)},
		{"lengthPastEnd", "bytes", append(encodedWord(32), encodedWord(1<<40)...)},
		{"hugeLength", "uint256[]", append(encodedWord(32), bytes.Repeat([]byte{0xff}, 32)...)},
		{"moreElementsThanWords", "uint256[]", append(append(encodedWord(32), encodedWord(3)...), make([]byte, 64)...)},
		{"missingPadding", "bytes", append(append(encodedWord(32), encodedWord(1)...), 7)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Decode[any](abi.Argument{Type: newType(t, test.solType)}, test.data)
			require.ErrorIs(t, err, ErrTruncated)
		})
	}
}

func TestCheckLayoutRejectsTrailingData(t *testing.T) {
	arg := abi.Argument{Type: newType(t, "string")}
	packed, err := abi.Arguments{arg}.Pack("hello")
	require.NoError(t, err)

	for _, extra := range [][]byte{{0}, encodedWord(0)} {
		_, err := Decode[string](arg, append(bytes.Clone(packed), extra...))
		require.ErrorIs(t, err, ErrNonCanonical)
	}
}

// TestCheckLayoutRejectsUnsupportedType confirms unsupported types are
// rejected whatever the data, even data a count of zero leaves unread.
func TestCheckLayoutRejectsUnsupportedType(t *testing.T) {
	noElements := append(encodedWord(32), encodedWord(0)...)
	huge := "576460752303423488" // 2^59: 32 bytes each overflow int
	half := "144115188075855872" // 2^57: two of 32 bytes each overflow int
	hugePairs, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{{Name: "A", Type: "uint256[" + half + "]"}, {Name: "B", Type: "uint256[" + half + "]"}})
	require.NoError(t, err)
	emptyTuples, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{})
	require.NoError(t, err)

	tests := []struct {
		name string
		typ  abi.Type
		data []byte
	}{
		{"fixedPoint", abi.Type{T: abi.FixedPointTy, Size: 128}, encodedWord(1)},
		{"zeroArray", newType(t, "uint256[0]"), nil},
		{"dynamicZeroArray", newType(t, "string[0]"), encodedWord(32)},
		{"zeroArraySlice", newType(t, "uint256[0][]"), noElements},
		{"emptyTupleSlice", emptyTuples, noElements},
		{"hugeArraySlice", newType(t, "uint256["+huge+"][]"), noElements},
		{"hugeDynamicArray", newType(t, "string["+huge+"]"), encodedWord(32)},
		{"hugeTupleSlice", hugePairs, noElements},
		// its head overflows int, so this must not loop
		{"hugeNestedArray", newType(t, "uint256["+huge+"]["+huge+"]"), make([]byte, 64)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.ErrorIs(t, checkLayout(abi.Arguments{{Type: test.typ}}, test.data), errUnsupportedType)
			_, err := Decode[any](abi.Argument{Type: test.typ}, test.data)
			require.ErrorIs(t, err, errUnsupportedType)
			_, err = DecodeArguments[any](abi.Arguments{{Type: test.typ}}, test.data)
			require.ErrorIs(t, err, errUnsupportedType)
		})
	}
}

type sendArgs struct {
	Sender  common.Address
	Message []byte
	Names   []string
	Amount  *big.Int
}

var sendComponents = []abi.ArgumentMarshaling{
	{Name: "sender", Type: "address"},
	{Name: "message", Type: "bytes"},
	{Name: "names", Type: "string[]"},
	{Name: "amount", Type: "uint256"},
}

var sendArguments = func() abi.Arguments {
	var args abi.Arguments
	for _, c := range sendComponents {
		typ, err := abi.NewType(c.Type, "", nil)
		if err != nil {
			panic(err)
		}
		args = append(args, abi.Argument{Name: c.Name, Type: typ})
	}

	return args
}()

func TestDecodeArguments(t *testing.T) {
	want := sendArgs{common.HexToAddress("0xc0"), []byte("hello"), []string{"a", "b"}, big.NewInt(42)}
	packed, err := sendArguments.Pack(want.Sender, want.Message, want.Names, want.Amount)
	require.NoError(t, err)

	got, err := DecodeArguments[sendArgs](sendArguments, packed)
	require.NoError(t, err)
	require.Equal(t, want, got)

	var dest sendArgs
	require.NoError(t, DecodeArgumentsTo(sendArguments, packed, &dest))
	require.Equal(t, want, dest)

	// one tuple of the same values has a leading offset, so it isn't this encoding
	_, err = Decode[any](abi.Argument{Type: tupleType(t, sendComponents)}, packed)
	require.ErrorIs(t, err, ErrNonCanonical)
}

// TestDecodeArgumentsShape confirms a destination go-ethereum's Copy would
// truncate or zero-fill is rejected, wherever the argument lands in it.
func TestDecodeArgumentsShape(t *testing.T) {
	args := abi.Arguments{{Name: "names", Type: newType(t, "string[2]")}, {Name: "amount", Type: newType(t, "uint256")}}
	packed, err := args.Pack([2]string{"a", "b"}, big.NewInt(5))
	require.NoError(t, err)
	pairArgs := abi.Arguments{
		{Name: "pair", Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a", Type: "uint256"}, {Name: "b", Type: "bool"}})},
		{Name: "n", Type: newType(t, "uint256")},
	}
	pairPacked, err := pairArgs.Pack(staticPair{big.NewInt(1), true}, big.NewInt(2))
	require.NoError(t, err)
	single := abi.Arguments{{Name: "x", Type: newType(t, "uint256[3]")}}
	singlePacked, err := single.Pack([3]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)})
	require.NoError(t, err)

	type reordered struct {
		Amount *big.Int
		Names  [2]string
	}
	got, err := DecodeArguments[reordered](args, packed)
	require.NoError(t, err)
	require.Equal(t, reordered{big.NewInt(5), [2]string{"a", "b"}}, got)

	type tagged struct {
		X      [2]string `abi:"names"`
		Amount *big.Int
	}
	gotTagged, err := DecodeArguments[tagged](args, packed)
	require.NoError(t, err)
	require.Equal(t, tagged{[2]string{"a", "b"}, big.NewInt(5)}, gotTagged)

	inOrder, err := DecodeArguments[[2]any](args, packed)
	require.NoError(t, err)
	require.Equal(t, [2]any{[2]string{"a", "b"}, big.NewInt(5)}, inOrder)

	type shorter struct {
		Names  [1]string
		Amount *big.Int
	}
	type longer struct {
		Names  [3]string
		Amount *big.Int
	}
	type shorterTagged struct {
		X      [1]string `abi:"names"`
		Amount *big.Int
	}
	type fewerFields struct {
		Pair struct{ A *big.Int }
		N    *big.Int
	}

	tests := []struct {
		name   string
		decode func() error
	}{
		{"shorterArray", decodeInto[shorter](args, packed)},
		{"longerArray", decodeInto[longer](args, packed)},
		{"taggedField", decodeInto[shorterTagged](args, packed)},
		{"fewerTupleFields", decodeInto[fewerFields](pairArgs, pairPacked)},
		{"extraArrayElement", decodeInto[[3]any](args, packed)},
		{"missingArrayElement", decodeInto[[1]any](args, packed)},
		{"single", decodeInto[[1]*big.Int](single, singlePacked)},
		{"singleFirstField", decodeInto[struct{ X [1]*big.Int }](single, singlePacked)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.ErrorIs(t, test.decode(), ErrShapeMismatch)
		})
	}
}

// decodeInto returns a call of DecodeArguments into a T.
func decodeInto[T any](args abi.Arguments, data []byte) func() error {
	return func() error {
		_, err := DecodeArguments[T](args, data)
		return err
	}
}

func TestDecodeArgumentsSingle(t *testing.T) {
	args := abi.Arguments{{Name: "pair", Type: newType(t, "string")}}
	packed, err := args.Pack("BTC/USD")
	require.NoError(t, err)

	got, err := DecodeArguments[string](args, packed)
	require.NoError(t, err)
	require.Equal(t, "BTC/USD", got)
}

func TestDecodeArgumentsRejects(t *testing.T) {
	packed, err := sendArguments.Pack(common.HexToAddress("0xc0"), []byte("hello"), []string{"a"}, big.NewInt(1))
	require.NoError(t, err)
	dirty := bytes.Clone(packed)
	dirty[0] = 1 // the address's padding

	_, err = DecodeArguments[sendArgs](sendArguments, dirty)
	require.ErrorIs(t, err, ErrNonCanonical)

	_, err = DecodeArguments[sendArgs](sendArguments, packed[:len(packed)-32])
	require.ErrorIs(t, err, ErrTruncated)

	_, err = DecodeArguments[struct{ Other []byte }](sendArguments, packed)
	require.Error(t, err)

	require.ErrorIs(t, DecodeArgumentsTo[sendArgs](sendArguments, packed, nil), errNonPointer)

	// Copy copies a single tuple into the struct's first field, here an array
	pair := []abi.ArgumentMarshaling{{Name: "A", Type: "uint256"}, {Name: "B", Type: "bool"}}
	wrapped := abi.Arguments{{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "X", Type: "tuple[2]", Components: pair}})}}
	_, err = DecodeArguments[struct{ X [2]staticPair }](wrapped, make([]byte, 128))
	require.ErrorIs(t, err, ErrShapeMismatch)

	// go-ethereum's Copy panics on this one: a struct with no first field
	_, err = DecodeArguments[struct{}](abi.Arguments{{Type: newType(t, "uint256")}}, make([]byte, 32))
	require.ErrorContains(t, err, "recovered panic")

	kept := sendArgs{Amount: big.NewInt(9)}
	require.Error(t, DecodeArgumentsTo(sendArguments, dirty, &kept))
	require.Equal(t, sendArgs{Amount: big.NewInt(9)}, kept)

	// Copy writes the first three fields before failing on the last
	type badAmount struct {
		Sender  common.Address
		Message []byte
		Names   []string
		Amount  string
	}
	partial, err := DecodeArguments[badAmount](sendArguments, packed)
	require.Error(t, err)
	require.Equal(t, badAmount{}, partial)
	keptBad := badAmount{Amount: "kept"}
	require.Error(t, DecodeArgumentsTo(sendArguments, packed, &keptBad))
	require.Equal(t, badAmount{Amount: "kept"}, keptBad)

	_, err = DecodeArguments[sendArgs](abi.Arguments{{}}, packed)
	require.ErrorIs(t, err, errUninitializedArg)
}

// TestCheckEncodeDecodeOmitsData confirms a round-trip mismatch names no
// input bytes, which may be secret, and is matchable.
func TestCheckEncodeDecodeOmitsData(t *testing.T) {
	args := abi.Arguments{{Type: newType(t, "uint256")}}
	data := bytes.Repeat([]byte{0xab}, 32)

	err := checkEncodeDecode(&args, data, []any{big.NewInt(1)})
	require.ErrorIs(t, err, ErrNonCanonical)
	require.NotContains(t, err.Error(), "171") // 0xab
}

// FuzzCheckLayout checks checkLayout against go-ethereum on scalars, dynamic
// values, and slices and arrays of them.
func FuzzCheckLayout(f *testing.F) {
	pair := []abi.ArgumentMarshaling{{Name: "S", Type: "string"}, {Name: "B", Type: "bytes[2]"}}
	flag := []abi.ArgumentMarshaling{{Name: "N", Type: "uint256"}, {Name: "F", Type: "bool"}}
	type pairValue struct {
		S string
		B [2][]byte
	}
	type flagValue struct {
		N *big.Int
		F bool
	}

	fuzzLayout(f, []abi.ArgumentMarshaling{
		{Name: "a", Type: "uint256"},
		{Name: "b", Type: "string[]"},
		{Name: "c", Type: "bytes"},
		{Name: "d", Type: "bytes32[2]"},
		{Name: "e", Type: "uint256[][]"},
		{Name: "f", Type: "tuple[]", Components: pair},
		{Name: "g", Type: "string[2]"},
		{Name: "h", Type: "tuple", Components: flag},
		{Name: "i", Type: "int16"},
		{Name: "j", Type: "uint8"},
		{Name: "k", Type: "address"},
	}, big.NewInt(1), []string{"x", ""}, []byte{1, 2},
		[2][32]byte{{1}, {2}}, [][]*big.Int{{big.NewInt(3)}, {}},
		[]pairValue{{"p", [2][]byte{{1}, {}}}}, [2]string{"", "g"}, flagValue{big.NewInt(4), true},
		int16(-2), uint8(5), common.HexToAddress("0xc0"))
}

// FuzzCheckLayoutNested checks checkLayout against go-ethereum on static
// values nested in slices, arrays and tuples, fixed arrays of dynamic tuples,
// and the remaining elementary types.
func FuzzCheckLayoutNested(f *testing.F) {
	entry := []abi.ArgumentMarshaling{{Name: "N", Type: "uint64"}, {Name: "A", Type: "address"}}
	record := []abi.ArgumentMarshaling{{Name: "F", Type: "uint32[2]"}, {Name: "S", Type: "bytes"}, {Name: "T", Type: "tuple", Components: entry}}
	type entryValue struct {
		N uint64
		A common.Address
	}
	type recordValue struct {
		F [2]uint32
		S []byte
		T entryValue
	}

	fuzzLayout(f, []abi.ArgumentMarshaling{
		{Name: "a", Type: "tuple[]", Components: entry},
		{Name: "b", Type: "uint16[2][3]"},
		{Name: "c", Type: "tuple[2]", Components: record},
		{Name: "d", Type: "bytes32[]"},
		{Name: "e", Type: "function"},
		{Name: "f", Type: "string[][2][]"},
		{Name: "g", Type: "int8"},
		{Name: "h", Type: "bytes1"},
	}, []entryValue{{1, common.HexToAddress("0x01")}, {2, common.HexToAddress("0x02")}},
		[3][2]uint16{{1, 2}, {3, 4}, {5, 6}},
		[2]recordValue{{[2]uint32{1, 2}, []byte{7}, entryValue{3, common.HexToAddress("0x03")}}, {}},
		[][32]byte{{1}}, [24]byte{1, 2}, [][2][]string{{{"a"}, {}}, {{}, {"b", "cc"}}},
		int8(-1), [1]byte{9})
}

// fuzzLayout checks that checkLayout accepts exactly what go-ethereum decodes
// and re-encodes unchanged, for the arguments components describe, seeded
// with the encoding of values. go-ethereum must range-check every value of
// their types.
func fuzzLayout(f *testing.F, components []abi.ArgumentMarshaling, values ...any) {
	f.Helper()
	args := make(abi.Arguments, 0, len(components))
	for _, c := range components {
		typ, err := abi.NewType(c.Type, "", c.Components)
		require.NoError(f, err)
		args = append(args, abi.Argument{Name: c.Name, Type: typ})
	}

	packed, err := args.Pack(values...)
	require.NoError(f, err)
	f.Add(packed)
	f.Add(packed[:len(packed)-32])
	f.Add(append(bytes.Clone(packed), 0))

	f.Fuzz(func(t *testing.T, data []byte) {
		values, err := args.Unpack(data)
		canonical := err == nil
		if canonical {
			again, err := args.Pack(values...)
			canonical = err == nil && bytes.Equal(again, data)
		}
		require.Equal(t, canonical, checkLayout(args, data) == nil)
	})
}
