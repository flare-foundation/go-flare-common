package signing

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	prefix := common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	dataHash := common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222")
	chainID := big.NewInt(14)

	p := Payload{Prefix: prefix, ChainID: chainID, DataHash: dataHash}

	got, err := p.Hash()
	require.NoError(t, err)

	// The tuple(bytes32, uint256, bytes32) encoding is the three 32-byte
	// words concatenated, so build it independently of PayloadArgument.
	encoded := make([]byte, 0, 96)
	encoded = append(encoded, prefix.Bytes()...)
	encoded = append(encoded, common.LeftPadBytes(chainID.Bytes(), 32)...)
	encoded = append(encoded, dataHash.Bytes()...)
	want := crypto.Keccak256Hash(encoded)

	require.Equal(t, [32]byte(want), got)
}

func TestHashNilChainID(t *testing.T) {
	p := Payload{
		Prefix:   common.HexToHash("0x01"),
		ChainID:  nil,
		DataHash: common.HexToHash("0x02"),
	}

	_, err := p.Hash()
	require.Error(t, err)
}

func TestHashDeterministic(t *testing.T) {
	p := Payload{
		Prefix:   common.HexToHash("0xaa"),
		ChainID:  big.NewInt(1),
		DataHash: common.HexToHash("0xbb"),
	}

	first, err := p.Hash()
	require.NoError(t, err)

	second, err := p.Hash()
	require.NoError(t, err)

	require.Equal(t, first, second)
}
