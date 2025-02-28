package structs

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/registry"
	"github.com/stretchr/testify/require"
)

func TestDecodeUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	eight := uint8(8)

	packed, err := abi.Arguments{arg}.Pack(eight)
	require.NoError(t, err)

	var unpacked uint8
	var unpacked2 uint8
	var unpacked3 uint8

	unpacked2, err = Decode[uint8](arg, packed)
	require.NoError(t, err)

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	err = DecodeTo2(arg, packed, &unpacked3)
	require.NoError(t, err)

	require.Equal(t, eight, unpacked)
	require.Equal(t, eight, unpacked2)
	require.Equal(t, eight, unpacked3)

}

func TestDecodeStruct(t *testing.T) {
	abiJson := `{
          "components": [
            {
              "internalType": "address",
              "name": "A",
              "type": "string"
            },
            {
              "internalType": "address",
              "name": "B",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "C",
              "type": "string"
            }
          ],
          "internalType": "struct IEntityManager.VoterAddresses",
          "name": "Neki",
          "type": "tuple"
        }`

	arg := abi.Argument{}
	err := arg.UnmarshalJSON([]byte(abiJson))
	require.NoError(t, err)

	type Neki struct {
		A string
		B uint8
		C string
	}

	pre := Neki{"a", 2, "c"}

	packed, err := abi.Arguments{arg}.PackValues([]any{pre})
	require.NoError(t, err)

	var unpacked Neki
	var unpacked2 Neki
	var unpacked3 Neki

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[Neki](arg, packed)
	require.NoError(t, err)

	err = DecodeTo2(arg, packed, &unpacked3)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
	require.Equal(t, pre, unpacked3)
}

func TestDecodeStructNested(t *testing.T) {
	abiJson := `{
		"components": [
		  {
			"internalType": "string",
			"name": "A",
			"type": "string"
		  },
		  {
			"internalType": "uint8",
			"name": "B",
			"type": "uint8"
		  },
		  {
			"components": [
			  {
				"internalType": "string",
				"name": "A",
				"type": "string"
			  },
			  {
				"internalType": "uint8",
				"name": "B",
				"type": "uint8"
			  }
			],
			"internalType": "struct",
			"name": "C",
			"type": "tuple"
		  }
		],
		"internalType": "struct IEntityManager.VoterAddresses",
		"name": "Neki2",
		"type": "tuple"
	  }`

	arg := abi.Argument{}
	err := arg.UnmarshalJSON([]byte(abiJson))
	require.NoError(t, err)

	type Neki struct {
		A string
		B uint8
	}

	type Neki2 struct {
		A string
		B uint8
		C Neki
	}

	pre := Neki2{"a", 2, Neki{"a", 2}}

	packed, err := abi.Arguments{arg}.PackValues([]any{pre})
	require.NoError(t, err)

	var unpacked Neki2
	var unpacked2 Neki2
	var unpacked3 Neki2

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[Neki2](arg, packed)
	require.NoError(t, err)

	err = DecodeTo2(arg, packed, &unpacked3)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
	require.Equal(t, pre, unpacked3)
}

func TestDecodeArray(t *testing.T) {
	abiJson := `{
		"internalType": "bytes32[]",
		"name": "_types",
		"type": "bytes32[2]"
	  }`

	arg := abi.Argument{}
	err := arg.UnmarshalJSON([]byte(abiJson))
	require.NoError(t, err)

	a := common.HexToHash("a")
	pre := [2][32]byte{a, a}

	packed, err := abi.Arguments{arg}.Pack(pre)
	require.NoError(t, err)

	var unpacked [2][32]byte
	var unpacked2 [2][32]byte
	var unpacked3 [2][32]byte

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[[2][32]byte](arg, packed)
	require.NoError(t, err)

	err = DecodeTo2(arg, packed, &unpacked3)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
	require.Equal(t, pre, unpacked3)
}

func TestDecodeInstructionMessage(t *testing.T) {
	arg := registry.MessageArguments[registry.ToPauseForUpgrade]

	id := common.HexToAddress("6e656b69")

	pre := registry.ITeeRegistryPauseForUpgrade{TeeId: id}

	encoded, err := abi.Arguments{arg}.Pack(pre)
	require.NoError(t, err)

	var unpacked registry.ITeeRegistryPauseForUpgrade

	err = DecodeTo(arg, encoded, &unpacked)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
}
