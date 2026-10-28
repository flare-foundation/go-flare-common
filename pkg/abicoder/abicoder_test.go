package abicoder

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	relayv1 "github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	relayv2 "github.com/flare-foundation/go-flare-common/pkg/contracts/v2/relay"
	"github.com/stretchr/testify/require"
)

// packFails reports whether go-ethereum's Pack, which Encode calls, fails on v.
func packFails(arg abi.Argument, v any) (failed bool) {
	defer func() {
		if recover() != nil {
			failed = true
		}
	}()
	_, err := abi.Arguments{arg}.Pack(v)

	return err != nil
}

// nameCase is a sample of a struct type checked for arg by New against what
// go-ethereum's Pack does with it.
type nameCase struct {
	name     string
	arg      abi.Argument
	sample   any
	build    func(abi.Argument) error
	rejected bool
}

func newNameCase[T any](name string, arg abi.Argument, sample T, rejected bool) nameCase {
	build := func(arg abi.Argument) error {
		_, err := New[T](arg)
		return err
	}

	return nameCase{name, arg, sample, build, rejected}
}

// TestNewChecksNames checks that New refuses a struct exactly where
// go-ethereum's Pack can't take every tuple component from it.
func TestNewChecksNames(t *testing.T) {
	type pair struct {
		A uint64
		B string
	}
	type tagged struct {
		X uint64 `abi:"a"`
		Y string `abi:"b"`
	}
	type renamed struct {
		P uint64
		Q string
	}
	type strayTag struct {
		A uint64 `abi:"c"`
		B string
	}
	type emptyTag struct {
		A uint64 `abi:""`
		B string
	}
	type inner struct{ W uint64 }
	type nested struct {
		A uint64
		B inner
	}
	// as abigen names fields for components a_b and aB
	type suffixed struct{ AB, AB0 uint64 }
	type suffixTagged struct {
		AB uint64
		X  uint64 `abi:"aB"`
	}
	// go-ethereum's own decode type for a_b and aB
	type decodeType = struct {
		AB  uint64 `json:"a_b"`
		AB0 uint64 `json:"aB"`
	}

	components := []abi.ArgumentMarshaling{{Name: "a", Type: "uint64"}, {Name: "b", Type: "string"}}
	tuple := abi.Argument{Type: tupleType(t, components)}
	tuples, err := abi.NewType("tuple[]", "", components)
	require.NoError(t, err)
	nestedTuple := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
		{Name: "a", Type: "uint64"},
		{Name: "b", Type: "tuple", Components: []abi.ArgumentMarshaling{{Name: "v", Type: "uint64"}}},
	})}
	collision := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{{Name: "a_b", Type: "uint64"}, {Name: "aB", Type: "uint64"}})}

	tests := []nameCase{
		newNameCase("pair", tuple, pair{1, "x"}, false),
		newNameCase("tagged", tuple, tagged{1, "x"}, false),
		newNameCase("renamed", tuple, renamed{1, "x"}, true),
		newNameCase("strayTag", tuple, strayTag{1, "x"}, true),
		newNameCase("emptyTag", tuple, emptyTag{1, "x"}, true),
		newNameCase("nested", nestedTuple, nested{1, inner{2}}, true),
		newNameCase("slice", abi.Argument{Type: tuples}, []pair{{1, "x"}}, false),
		newNameCase("sliceRenamed", abi.Argument{Type: tuples}, []renamed{{1, "x"}}, true),
		newNameCase("suffixed", collision, suffixed{1, 2}, true),
		newNameCase("suffixTagged", collision, suffixTagged{1, 2}, false),
		newNameCase("decodeType", collision, decodeType{1, 2}, true),
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.rejected, packFails(tc.arg, tc.sample))
			if tc.rejected {
				require.ErrorIs(t, tc.build(tc.arg), ErrShapeMismatch)
			} else {
				require.NoError(t, tc.build(tc.arg))
			}
		})
	}

	// Decode fills fields by position, so it still takes a renamed struct
	packed, err := Encode(tuple, pair{1, "x"})
	require.NoError(t, err)
	out, err := Decode[renamed](tuple, packed)
	require.NoError(t, err)
	require.Equal(t, renamed{1, "x"}, out)
}

// TestNewRefusesUnsupportedTypes checks that New refuses the types Decode
// refuses whatever the data, before GetType panics on some of them.
func TestNewRefusesUnsupportedTypes(t *testing.T) {
	tests := []struct {
		name string
		typ  abi.Type
	}{
		{"zeroLengthArray", newType(t, "uint256[0]")},
		{"emptyTuple", newType(t, "tuple")},
		{"handBuiltTuple", abi.Type{T: abi.TupleTy}},
		{"fixedPoint", abi.Type{T: abi.FixedPointTy, Size: 128}},
		{"hash", abi.Type{T: abi.HashTy}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := New[any](abi.Argument{Type: tc.typ})
			require.ErrorIs(t, err, errUnsupportedType)
		})
	}

	// before decoding: go-ethereum decodes HashTy but can't encode it
	_, err := Decode[[32]byte](abi.Argument{Type: abi.Type{T: abi.HashTy}}, make([]byte, 32))
	require.ErrorIs(t, err, errUnsupportedType)
}

// TestNewRefusesIndexed checks that New refuses an event's indexed argument,
// which go-ethereum's Unpack skips, failing every decode.
func TestNewRefusesIndexed(t *testing.T) {
	_, err := New[uint64](abi.Argument{Type: newType(t, "uint64"), Indexed: true})
	require.ErrorIs(t, err, errIndexedArg)
}

// coderCase round-trips a sample through a Coder for arg.
type coderCase struct {
	name string
	run  func(*testing.T)
}

func newCoderCase[T any](name string, arg abi.Argument, v T) coderCase {
	run := func(t *testing.T) {
		t.Helper()
		c, err := New[T](arg)
		require.NoError(t, err)
		packed, err := c.Encode(v)
		require.NoError(t, err)
		want, err := Encode(arg, v)
		require.NoError(t, err)
		require.Equal(t, want, packed)
		got, err := c.Decode(packed)
		require.NoError(t, err)
		require.Equal(t, v, got)
	}

	return coderCase{name, run}
}

// TestNewAcceptsTuplesAndBasicTypes checks that a Coder takes tuples and basic
// types, and encodes them as Encode does.
func TestNewAcceptsTuplesAndBasicTypes(t *testing.T) {
	type pair struct {
		A uint64
		B string
	}
	type nested struct {
		A []pair
		B [2][]byte
	}

	arg := func(solType string) abi.Argument { return abi.Argument{Type: newType(t, solType)} }
	components := []abi.ArgumentMarshaling{{Name: "a", Type: "uint64"}, {Name: "b", Type: "string"}}
	tuple := abi.Argument{Type: tupleType(t, components)}
	tuples, err := abi.NewType("tuple[]", "", components)
	require.NoError(t, err)
	nestedTuple := abi.Argument{Type: tupleType(t, []abi.ArgumentMarshaling{
		{Name: "a", Type: "tuple[]", Components: components},
		{Name: "b", Type: "bytes[2]"},
	})}

	tests := []coderCase{
		newCoderCase("string", arg("string"), "x"),
		newCoderCase("bytes", arg("bytes"), []byte{1, 2}),
		newCoderCase("bytes32", arg("bytes32"), common.Hash{1}),
		newCoderCase("bytes32AsArray", arg("bytes32"), [32]byte{1}),
		newCoderCase("bytes4", arg("bytes4"), [4]byte{1}),
		newCoderCase("bool", arg("bool"), true),
		newCoderCase("uint8", arg("uint8"), uint8(1)),
		newCoderCase("uint64", arg("uint64"), uint64(1)),
		newCoderCase("int8", arg("int8"), int8(-1)),
		newCoderCase("uint24", arg("uint24"), big.NewInt(1)),
		newCoderCase("uint256", arg("uint256"), big.NewInt(1)),
		newCoderCase("int256", arg("int256"), big.NewInt(-1)),
		newCoderCase("address", arg("address"), common.Address{1}),
		newCoderCase("function", arg("function"), [24]byte{1}),
		newCoderCase("uint8Slice", arg("uint8[]"), []uint8{1, 2}),
		newCoderCase("stringSlice", arg("string[]"), []string{"x", "y"}),
		newCoderCase("bytesSlice", arg("bytes[]"), [][]byte{{1}, {2}}),
		newCoderCase("addressSlice", arg("address[]"), []common.Address{{1}}),
		newCoderCase("uint8Array", arg("uint8[3]"), [3]uint8{1, 2, 3}),
		newCoderCase("bytes32Pair", arg("bytes32[2]"), [2][32]byte{{1}, {2}}),
		newCoderCase("tuple", tuple, pair{1, "x"}),
		newCoderCase("tupleSlice", abi.Argument{Type: tuples}, []pair{{1, "x"}, {2, "y"}}),
		newCoderCase("nested", nestedTuple, nested{[]pair{{1, "x"}}, [2][]byte{{1}, {2}}}),
	}
	for _, tc := range tests {
		t.Run(tc.name, tc.run)
	}
}

// TestCoderChecksData checks that the Coder's calls still check the data.
func TestCoderChecksData(t *testing.T) {
	c, err := New[uint64](abi.Argument{Type: newType(t, "uint64")})
	require.NoError(t, err)
	packed, err := c.Encode(1)
	require.NoError(t, err)

	dirty := bytes.Clone(packed)
	dirty[0] = 1 // padding
	_, err = c.Decode(dirty)
	require.ErrorIs(t, err, ErrNonCanonical)
	_, err = c.Decode(packed[:31])
	require.ErrorIs(t, err, ErrTruncated)
	_, err = c.Decode(append(bytes.Clone(packed), 0))
	require.ErrorIs(t, err, ErrNonCanonical)

	// go-ethereum writes a 32-byte address whole, high bytes and all
	addr, err := New[any](abi.Argument{Type: newType(t, "address")})
	require.NoError(t, err)
	_, err = addr.Encode(common.MaxHash)
	require.ErrorIs(t, err, ErrNonCanonical)
}

// TestNewAcceptsAbigenStructs checks abigen's structs, v1 and v2, against the
// tuple they were generated for.
func TestNewAcceptsAbigenStructs(t *testing.T) {
	v1, err := relayv1.RelayMetaData.GetAbi()
	require.NoError(t, err)
	v2, err := relayv2.RelayMetaData.ParseABI()
	require.NoError(t, err)

	policy := relayv1.IIRelaySigningPolicy{
		RewardEpochId:      big.NewInt(1),
		StartVotingRoundId: 2,
		Threshold:          3,
		Seed:               big.NewInt(4),
		Voters:             []common.Address{{5}},
		Weights:            []uint16{6},
	}

	c1, err := New[relayv1.IIRelaySigningPolicy](v1.Methods["setSigningPolicy"].Inputs[0])
	require.NoError(t, err)
	packed, err := c1.Encode(policy)
	require.NoError(t, err)
	got, err := c1.Decode(packed)
	require.NoError(t, err)
	require.Equal(t, policy, got)

	c2, err := New[relayv2.IIRelaySigningPolicy](v2.Methods["setSigningPolicy"].Inputs[0])
	require.NoError(t, err)
	packed2, err := c2.Encode(relayv2.IIRelaySigningPolicy(policy))
	require.NoError(t, err)
	require.Equal(t, packed, packed2)
}
