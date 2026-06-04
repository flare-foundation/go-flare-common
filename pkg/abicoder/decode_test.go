package abicoder

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

	unpacked2, err = Decode[uint8](arg, packed)
	require.NoError(t, err)

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	require.Equal(t, eight, unpacked)
	require.Equal(t, eight, unpacked2)
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
		name       string
		inputType  abi.Argument
		input      any
		outputType abi.Argument
	}{
		{
			name:       "uint8 decoded as string",
			inputType:  argUint8,
			input:      uint8(8),
			outputType: argStr,
		},
		{
			name:       "string decoded as uint8",
			inputType:  argStr,
			input:      "osem",
			outputType: argUint8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			packed, err := abi.Arguments{test.inputType}.Pack(test.input)
			require.NoError(t, err)

			var unpacked0 any

			_, err = Decode[any](test.outputType, packed)
			require.Error(t, err)

			err = DecodeTo(test.outputType, packed, &unpacked0)
			require.Error(t, err)
		})
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

	_, err = Decode[uint8](argUint8, packed)
	require.Error(t, err)

	err = DecodeTo(argUint8, packed, &unpacked0)
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

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[Neki](arg, packed)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
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

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[Neki2](arg, packed)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
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

	err = DecodeTo(arg, packed, &unpacked)
	require.NoError(t, err)

	unpacked2, err = Decode[[2][32]byte](arg, packed)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
	require.Equal(t, pre, unpacked2)
}

func TestMismatchedStructs(t *testing.T) {
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

	err = DecodeTo(arg0, packed, &unpacked)
	require.Error(t, err)

	_, err = Decode[Drugo](arg0, packed)
	require.Error(t, err)
}

// TestDecodeToRejectsSubtleShapeMismatch verifies that DecodeTo errors
// when a destination has the right field count but a wrong field type,
// which strict ABI shape verification must reject.
func TestDecodeToRejectsSubtleShapeMismatch(t *testing.T) {
	abiJSON := `{
        "components": [
          {"internalType": "string", "name": "A", "type": "string"},
          {"internalType": "uint8",  "name": "B", "type": "uint8"}
        ],
        "internalType": "struct Mismatch.Subtle",
        "name": "subtle",
        "type": "tuple"
    }`

	arg := abi.Argument{}
	require.NoError(t, arg.UnmarshalJSON([]byte(abiJSON)))

	type Source struct {
		A string
		B uint8
	}
	// Dest has the right field count but B's type is wrong (uint8 -> []byte).
	type Dest struct {
		A string
		B []byte
	}

	packed, err := abi.Arguments{arg}.PackValues([]any{Source{"hello", 7}})
	require.NoError(t, err)

	var dst Dest
	err = DecodeTo(arg, packed, &dst)
	require.Error(t, err, "subtle B-field type mismatch must be rejected")
}
