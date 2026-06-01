package instruction

import (
	"bytes"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
	"github.com/stretchr/testify/require"
)

// TestValidate verifies that parsed instruction Data rejects
// oversize message bytes, broken cosigner threshold/uniqueness, and
// unrecognized op type/command pairs before the values flow into a hash
// or dispatch path.
func TestValidate(t *testing.T) {
	good := Data{
		DataFixed: DataFixed{
			OPType:                 op.Wallet.Hash(),
			OPCommand:              op.KeyGenerate.Hash(),
			Cosigners:              []common.Address{{0x01}, {0x02}, {0x03}},
			CosignersThreshold:     2,
			OriginalMessage:        hexutil.Bytes{1, 2, 3},
			InstructionID:          common.Hash{},
			TeeID:                  common.Address{},
			Timestamp:              0,
			RewardEpochID:          0,
			AdditionalFixedMessage: hexutil.Bytes{},
		},
		AdditionalVariableMessage: hexutil.Bytes{},
	}

	t.Run("good", func(t *testing.T) {
		require.NoError(t, good.Validate())
	})

	t.Run("oversize OriginalMessage", func(t *testing.T) {
		d := good
		d.OriginalMessage = bytes.Repeat([]byte{0xAA}, MaxMessageSize+1)
		require.Error(t, d.Validate())
	})

	t.Run("oversize AdditionalFixedMessage", func(t *testing.T) {
		d := good
		d.AdditionalFixedMessage = bytes.Repeat([]byte{0xAA}, MaxMessageSize+1)
		require.Error(t, d.Validate())
	})

	t.Run("oversize AdditionalVariableMessage", func(t *testing.T) {
		d := good
		d.AdditionalVariableMessage = bytes.Repeat([]byte{0xAA}, MaxMessageSize+1)
		require.Error(t, d.Validate())
	})

	t.Run("too many cosigners", func(t *testing.T) {
		d := good
		d.Cosigners = make([]common.Address, MaxCosigners+1)
		d.CosignersThreshold = 1
		require.Error(t, d.Validate())
	})

	t.Run("threshold zero", func(t *testing.T) {
		d := good
		d.CosignersThreshold = 0
		require.Error(t, d.Validate())
	})

	t.Run("threshold exceeds cosigner count", func(t *testing.T) {
		d := good
		d.CosignersThreshold = uint64(len(good.Cosigners)) + 1
		require.Error(t, d.Validate())
	})

	t.Run("duplicate cosigner", func(t *testing.T) {
		d := good
		d.Cosigners = []common.Address{{0x01}, {0x01}}
		d.CosignersThreshold = 1
		require.Error(t, d.Validate())
	})

	t.Run("invalid op pair (F-prefixed unknown command)", func(t *testing.T) {
		d := good
		d.OPCommand = op.Command("F_BOGUS").Hash()
		require.Error(t, d.Validate())
	})
}

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
		Cosigners:              []common.Address{},
		CosignersThreshold:     0,
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

	ts := time.Now().Unix()

	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	sig, err := crypto.Sign(hFixed[:], priv)
	require.NoError(t, err)

	_, err = NextVoteHash(vh, 0, sig, nil, uint64(ts))
	require.NoError(t, err)
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
		Cosigners:              []common.Address{},
		CosignersThreshold:     0,
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

func TestHashFixedIncludesChainID(t *testing.T) {
	base := DataFixed{
		InstructionID:          common.Hash{0xAA},
		TeeID:                  common.Address{0xBB},
		Timestamp:              42,
		RewardEpochID:          7,
		OPType:                 op.Wallet.Hash(),
		OPCommand:              op.KeyGenerate.Hash(),
		Cosigners:              []common.Address{{0x01}},
		CosignersThreshold:     1,
		OriginalMessage:        hexutil.Bytes{1, 2, 3},
		AdditionalFixedMessage: hexutil.Bytes{4, 5},
	}

	t.Run("different ChainID changes the hash", func(t *testing.T) {
		a := base
		a.ChainID = 1
		b := base
		b.ChainID = 2

		ha, err := a.HashFixed()
		require.NoError(t, err)
		hb, err := b.HashFixed()
		require.NoError(t, err)

		require.NotEqual(t, ha, hb, "HashFixed must depend on ChainID")
	})

	t.Run("same ChainID yields stable hash", func(t *testing.T) {
		a := base
		a.ChainID = 14
		b := base
		b.ChainID = 14

		ha, err := a.HashFixed()
		require.NoError(t, err)
		hb, err := b.HashFixed()
		require.NoError(t, err)

		require.Equal(t, ha, hb)
	})

	t.Run("ChainID zero differs from non-zero", func(t *testing.T) {
		zero := base // ChainID defaults to 0
		nonzero := base
		nonzero.ChainID = 1

		hz, err := zero.HashFixed()
		require.NoError(t, err)
		hnz, err := nonzero.HashFixed()
		require.NoError(t, err)

		require.NotEqual(t, hz, hnz,
			"ChainID=0 must hash differently than ChainID=1; "+
				"otherwise a JSON producer that omits chainId silently matches ChainID=0")
	})
}
