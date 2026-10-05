package relay_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	relayv1 "github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/v2/relay"
)

func TestPackMatchesV1(t *testing.T) {
	v1, err := relayv1.RelayMetaData.GetAbi()
	require.NoError(t, err)
	r := relay.NewRelay()

	policy := relay.IIRelaySigningPolicy{
		RewardEpochId:      big.NewInt(42),
		StartVotingRoundId: 1000,
		Threshold:          500,
		Seed:               big.NewInt(7),
		Voters:             []common.Address{common.HexToAddress("0x01"), common.HexToAddress("0x02")},
		Weights:            []uint16{300, 400},
	}

	tests := []struct {
		method string
		got    []byte
		args   []any
	}{
		{"getConfirmedMerkleRoot", r.PackGetConfirmedMerkleRoot(big.NewInt(100), big.NewInt(200)), []any{big.NewInt(100), big.NewInt(200)}},
		{"getRandomNumber", r.PackGetRandomNumber(), nil},
		{"setSigningPolicy", r.PackSetSigningPolicy(policy), []any{relayv1.IIRelaySigningPolicy(policy)}},
	}

	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			want, err := v1.Pack(tc.method, tc.args...)
			require.NoError(t, err)
			require.Equal(t, want, tc.got)
		})
	}
}

func TestUnpackMatchesV1(t *testing.T) {
	v1, err := relayv1.RelayMetaData.GetAbi()
	require.NoError(t, err)

	data, err := v1.Methods["getRandomNumber"].Outputs.Pack(big.NewInt(5), true, big.NewInt(9))
	require.NoError(t, err)

	got, err := relay.NewRelay().UnpackGetRandomNumber(data)
	require.NoError(t, err)
	require.Equal(t, relay.GetRandomNumberOutput{
		RandomNumber:    big.NewInt(5),
		IsSecureRandom:  true,
		RandomTimestamp: big.NewInt(9),
	}, got)
}

func TestEventMatchesV1(t *testing.T) {
	v1, err := relayv1.RelayMetaData.GetAbi()
	require.NoError(t, err)

	ev := v1.Events["ProtocolMessageRelayed"]
	root := common.HexToHash("0xabcd")
	data, err := ev.Inputs.NonIndexed().Pack(true, root)
	require.NoError(t, err)
	log := types.Log{
		Topics: []common.Hash{ev.ID, common.BigToHash(big.NewInt(3)), common.BigToHash(big.NewInt(77))},
		Data:   data,
	}

	f, err := relayv1.NewRelayFilterer(common.Address{}, nil)
	require.NoError(t, err)
	want, err := f.ParseProtocolMessageRelayed(log)
	require.NoError(t, err)

	got, err := relay.NewRelay().UnpackProtocolMessageRelayedEvent(&log)
	require.NoError(t, err)
	require.Equal(t, want.ProtocolId, got.ProtocolId)
	require.Equal(t, want.VotingRoundId, got.VotingRoundId)
	require.Equal(t, want.IsSecureRandom, got.IsSecureRandom)
	require.Equal(t, want.MerkleRoot, got.MerkleRoot)
	require.Equal(t, uint8(3), got.ProtocolId)
	require.Equal(t, [32]byte(root), got.MerkleRoot)
}
