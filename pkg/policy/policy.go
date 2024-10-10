package policy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/contracts/relay"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/voters"
)

type SigningPolicy struct {
	RewardEpochID      int64
	startVotingRoundID uint32
	threshold          uint16
	seed               *big.Int
	rawBytes           []byte
	blockTimestamp     uint64

	// The set of all Voters and their weights
	Voters *voters.Set
}

func NewSigningPolicy(r *relay.RelaySigningPolicyInitialized, submitToSigning map[common.Address]common.Address) *SigningPolicy {
	return &SigningPolicy{
		RewardEpochID:      r.RewardEpochId.Int64(),
		startVotingRoundID: r.StartVotingRoundId,
		threshold:          r.Threshold,
		seed:               r.Seed,
		rawBytes:           r.SigningPolicyBytes,
		blockTimestamp:     r.Timestamp,
		Voters:             voters.NewSet(r.Voters, r.Weights, submitToSigning),
	}
}
