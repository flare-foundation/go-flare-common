package instruction

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	data.RewardEpochID = 1

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
		InstructionID:          common.Hash{},
		TeeID:                  common.Address{},
		Timestamp:              0,
		RewardEpochID:          0,
		OPType:                 common.Hash{},
		OPCommand:              common.Hash{},
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

func TestSignAndRecover(t *testing.T) {
	pk, err := crypto.GenerateKey()

	require.NoError(t, err)

	var data Data

	dataFixed := DataFixed{
		InstructionID:          common.Hash{},
		TeeID:                  common.Address{},
		Timestamp:              0,
		RewardEpochID:          0,
		OPType:                 common.Hash{},
		OPCommand:              common.Hash{},
		OriginalMessage:        hexutil.Bytes{1},
		AdditionalFixedMessage: hexutil.Bytes{1},
	}

	data.DataFixed = dataFixed
	data.AdditionalVariableMessage = []byte{1}

	hash, err := data.HashForSigning()
	require.NoError(t, err)

	sig, err := SignInstructionHash(hash, pk)
	require.NoError(t, err)

	i := Instruction{
		Data:      data,
		Signature: sig,
	}

	pub, err := i.RecoverSignersPubKey()
	require.NoError(t, err)

	require.Equal(t, *pub, pk.PublicKey)
}
