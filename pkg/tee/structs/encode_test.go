package structs

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/wallet"
	"github.com/stretchr/testify/require"
)

func TestEncodeUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	eight := uint8(8)

	packed, err := Encode(arg, eight)
	require.NoError(t, err)

	expected := make([]byte, 32)
	expected[31] = eight

	require.Equal(t, expected, packed)
}

func TestEncodeFailUint8(t *testing.T) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint8",
		Type:    uint8Type,
		Indexed: false,
	}

	notUint8 := uint16(8)

	_, err = Encode(arg, notUint8)
	require.Error(t, err)

	stillNotUint8 := 257
	_, err = Encode(arg, stillNotUint8)
	require.Error(t, err)
}

func TestEncodeStruct(t *testing.T) {
	abiJson := `{
          "components": [
            {
              "internalType": "address",
              "name": "A",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "B",
              "type": "uint8"
            },
            {
              "internalType": "address",
              "name": "C",
              "type": "uint8"
            }
          ],
          "internalType": "struct IEntityManager.VoterAddresses",
          "name": "Neki",
          "type": "tuple"
        }`

	arg := abi.Argument{}
	err := arg.UnmarshalJSON([]byte(abiJson))
	require.NoError(t, err)

	type neki struct {
		A uint8
		B uint8
		C uint8
	}

	n := neki{8, 8, 8}

	encoded, err := Encode(arg, n)
	require.NoError(t, err)

	expected := make([]byte, 96)
	nonzero := []int{31, 63, 95}

	for _, j := range nonzero {
		expected[j] = uint8(8)
	}

	require.Equal(t, expected, encoded)

	// Fail

	nd := struct {
		a string
	}{"a"}

	_, err = Encode(arg, nd)
	require.Error(t, err)
}

func TestEncodeString(t *testing.T) {
	stringType, err := abi.NewType("string", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "string",
		Type:    stringType,
		Indexed: false,
	}

	s := "abc"

	packed, err := Encode(arg, s)
	require.NoError(t, err)

	expected := make([]byte, 96)
	expected[31] = uint8(32)
	expected[63] = uint8(len(s))
	copy(expected[64:68], []byte("abc"))

	require.Equal(t, expected, packed)

	decoded, err := Decode[string](arg, packed)
	require.NoError(t, err)
	require.Equal(t, s, decoded)
}

func TestKeyExistenceStruct(t *testing.T) {
	s := wallet.IWalletKeyManagerKeyExistence{
		TeeId:       common.HexToAddress("dead"),
		WalletId:    crypto.Keccak256Hash([]byte("dead")),
		KeyId:       0,
		KeyType:     [32]byte{},
		SigningAlgo: [32]byte{},
		PublicKey:   []byte{},
		Nonce:       big.NewInt(10),
		Restored:    false,
		ConfigConstants: wallet.IWalletKeyManagerKeyConfigConstants{
			AdminsPublicKeys:   []wallet.PublicKey{},
			AdminsThreshold:    1,
			Cosigners:          []common.Address{},
			CosignersThreshold: 0,
		},
		SettingsVersion: [32]byte{},
		Settings:        []byte{},
	}

	encoded, err := Encode(wallet.KeyExistenceStructArg, s)
	require.NoError(t, err)

	decoded, err := Decode[wallet.IWalletKeyManagerKeyExistence](wallet.KeyExistenceStructArg, encoded)
	require.NoError(t, err)

	require.Equal(t, s, decoded)
}

func TestEncodeBigInt(t *testing.T) {
	uint256Type, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	arg := abi.Argument{
		Name:    "uint256",
		Type:    uint256Type,
		Indexed: false,
	}

	in := big.NewInt(8)

	packed, err := Encode(arg, in)
	require.NoError(t, err)

	expected := make([]byte, 32)
	expected[31] = 8

	require.Equal(t, expected, packed)

	x := big.NewInt(1231223123123333332)
	packed, err = Encode(arg, x)
	require.NoError(t, err)

	decoded, err := Decode[*big.Int](arg, packed)
	require.NoError(t, err)

	require.Equal(t, 0, x.Cmp(decoded))
}
