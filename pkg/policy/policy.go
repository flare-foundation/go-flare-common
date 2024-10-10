package policy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/contracts/relay"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/voters"
)

type SigningPolicy struct {
	RewardEpochID      int64
	StartVotingRoundID uint32
	Threshold          uint16
	seed               *big.Int
	rawBytes           []byte
	blockTimestamp     uint64

	// The set of all Voters and their weights
	Voters *voters.Set
}

func (sp SigningPolicy) RawBytes() []byte {
	return sp.rawBytes
}

func NewSigningPolicy(r *relay.RelaySigningPolicyInitialized, submitToSigning map[common.Address]common.Address) *SigningPolicy {
	return &SigningPolicy{
		RewardEpochID:      r.RewardEpochId.Int64(),
		StartVotingRoundID: r.StartVotingRoundId,
		Threshold:          r.Threshold,
		seed:               r.Seed,
		rawBytes:           r.SigningPolicyBytes,
		blockTimestamp:     r.Timestamp,
		Voters:             voters.NewSet(r.Voters, r.Weights, submitToSigning),
	}
}
