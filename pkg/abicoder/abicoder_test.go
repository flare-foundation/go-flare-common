package abicoder

import (
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
