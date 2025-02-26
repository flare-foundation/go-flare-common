package structs

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

	err = Decode(arg, packed, &unpacked)
	require.NoError(t, err)

	require.Equal(t, eight, unpacked)
}

type Neki struct {
	A string
	B uint8
	C string
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

	pre := Neki{"a", 2, "c"}

	packed, err := abi.Arguments{arg}.PackValues([]any{pre})
	require.NoError(t, err)

	var unpacked Neki

	err = Decode(arg, packed, &unpacked)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked)
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

	unpacked := [2][32]byte{}

	err = Decode(arg, packed, &unpacked)
	require.NoError(t, err)
	require.Equal(t, pre, unpacked)
}
