package structs

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/wallet"
	"github.com/stretchr/testify/require"
)

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
