package structs

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/tee/constants"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/verification"
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

func TestDecodeFail(t *testing.T) {
	stringType, err := abi.NewType("string", "string", nil)
	require.NoError(t, err)

	argStr := abi.Argument{
		Name:    "string",
		Type:    stringType,
		Indexed: false,
	}

	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	argUint8 := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	tests := []struct {
		inputType  abi.Argument
		input      any
		outputType abi.Argument
	}{
		{
			inputType:  argUint8,
			input:      uint8(8),
			outputType: argStr,
		},
		{
			inputType:  argStr,
			input:      "osem",
			outputType: argUint8,
		},
	}

	for j, test := range tests {
		packed, err := abi.Arguments{test.inputType}.Pack(test.input)
		require.NoError(t, err)

		var unpacked0 any
		var unpacked1 any

		_, err = Decode[any](test.outputType, packed)
		require.Error(t, err, j)

		err = DecodeTo(test.outputType, packed, &unpacked0)
		require.Error(t, err, j)

		err = DecodeTo2(test.outputType, packed, &unpacked1)
		require.Error(t, err, j)
	}
}

func TestUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	argUint8 := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	eight := uint8(8)

	packed, err := abi.Arguments{argUint8}.Pack(eight)
	require.NoError(t, err)

	packed = append(packed, 2)

	var unpacked0 uint8
	var unpacked1 uint8

	_, err = Decode[uint8](argUint8, packed)
	require.Error(t, err)

	err = DecodeTo(argUint8, packed, &unpacked0)
	require.Error(t, err)

	err = DecodeTo2(argUint8, packed, &unpacked1)
	require.Error(t, err)
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
	arg := verification.MessageArguments[constants.TEEAttestation]

	id := common.HexToAddress("6e656b69")
	h := common.HexToHash("6e656b69")

	x := common.HexToHash("0x1234")

	pre := verification.ITeeVerificationTeeAttestation{
		TeeMachine: verification.ITeeRegistryTeeMachineWithAttestationData{
			TeeId:        id,
			InitialTeeId: id,
			Url:          "moj.com",
			CodeHash:     h,
			Platform:     h,
		},
		Challenge: x,
	}

	encoded, err := abi.Arguments{arg}.Pack(pre)
	require.NoError(t, err)

	var unpacked1 verification.ITeeVerificationTeeAttestation
	var unpacked2 verification.ITeeVerificationTeeAttestation

	err = DecodeTo(arg, encoded, &unpacked1)
	require.NoError(t, err)
	require.Equal(t, pre, unpacked1)

	err = DecodeTo2(arg, encoded, &unpacked2)
	require.NoError(t, err)
	require.Equal(t, pre, unpacked2)

	unpacked3, err := Decode[verification.ITeeVerificationTeeAttestation](arg, encoded)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked3)
}

func TestMissMatchedStructs(t *testing.T) {
	abiJSON0 := `{
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

	arg0 := abi.Argument{}
	err := arg0.UnmarshalJSON([]byte(abiJSON0))
	require.NoError(t, err)

	type Neki struct {
		A string
		B uint8
		C string
	}

	pre := Neki{"a", 2, "c"}

	type Drugo struct {
		A []byte
	}

	packed, err := abi.Arguments{arg0}.PackValues([]any{pre})
	require.NoError(t, err)

	var unpacked Drugo
	var unpacked3 Drugo

	err = DecodeTo(arg0, packed, &unpacked)
	require.Error(t, err)

	_, err = Decode[Drugo](arg0, packed)
	require.Error(t, err)

	err = DecodeTo2(arg0, packed, &unpacked3)
	require.Error(t, err)
}
