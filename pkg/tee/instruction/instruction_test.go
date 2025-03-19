package instruction

import (
	"math/big"
	"testing"

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

	data.RewardEpochID = new(big.Int)

	hash, err = data.HashForSigning()
	require.NoError(t, err)
	require.NotEqual(t, hash0, hash)

	hashFixed, err := data.HashFixed()
	require.NoError(t, err)

	require.NotEqual(t, hashFixed, hash)
}
