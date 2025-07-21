package instruction

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestHashForSigning(t *testing.T) {
	var data Data

	hash0, err := data.HashForSigning()
	require.NoError(t, err)

	require.Nil(t, data.AdditionalFixedMessage)
	require.Nil(t, data.AdditionalVariableMessage)

	data.AdditionalFixedMessage = []byte{}
	data.AdditionalVariableMessage = hexutil.Bytes{}

	hash, err := data.HashForSigning()
	require.NoError(t, err)
	require.Equal(t, hash0, hash)

	data.RewardEpochId = 1

	hash, err = data.HashForSigning()
	require.NoError(t, err)
	require.NotEqual(t, hash0, hash)

	hashFixed, err := data.HashFixed()
	require.NoError(t, err)

	require.NotEqual(t, hashFixed, hash)
}

func TestHash(t *testing.T) {
	var data Data

	dataFixed := DataFixed{
		InstructionId:          common.Hash{},
		TeeId:                  common.Address{},
		Timestamp:              0,
		RewardEpochId:          0,
		OpType:                 common.Hash{},
		OpCommand:              common.Hash{},
		OriginalMessage:        hexutil.Bytes{1},
		AdditionalFixedMessage: hexutil.Bytes{1},
	}

	data.DataFixed = dataFixed
	data.AdditionalVariableMessage = []byte{1}

	hFull, err := data.HashFixed()
	require.NoError(t, err)

	hFixed, err := dataFixed.HashFixed()
	require.NoError(t, err)

	require.Equal(t, hFixed, hFull)

	vh, err := data.InitialVoteHash()
	require.NoError(t, err)

	vhFixed, err := data.InitialVoteHash()
	require.NoError(t, err)

	require.Equal(t, vhFixed, vh)
}
