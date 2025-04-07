package policy

import (
	"bytes"
	"encoding/binary"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/pkg/errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

type SigningPolicy struct {
	RewardEpochID      int64
	StartVotingRoundID uint32
	Threshold          uint16
	Seed               *big.Int
	rawBytes           []byte
	blockTimestamp     uint64

	// The set of all Voters and their weights
	Voters *voters.Set
}

func (sp *SigningPolicy) RawBytes() []byte {
	return sp.rawBytes
}

func NewSigningPolicy(r *relay.RelaySigningPolicyInitialized, submitToSigning map[common.Address]common.Address) *SigningPolicy {
	return &SigningPolicy{
		RewardEpochID:      r.RewardEpochId.Int64(),
		StartVotingRoundID: r.StartVotingRoundId,
		Threshold:          r.Threshold,
		Seed:               r.Seed,
		rawBytes:           r.SigningPolicyBytes,
		blockTimestamp:     r.Timestamp,
		Voters:             voters.NewSet(r.Voters, r.Weights, submitToSigning),
	}
}

// Equals compares two SigningPolicy objects based on their rawBytes.
func (sp *SigningPolicy) Equals(other *SigningPolicy) bool {
	if other == nil {
		return false
	}
	return bytes.Equal(sp.rawBytes, other.rawBytes)
}

// FromRawBytes decodes a SigningPolicy from a byte slice, returning the decoded SigningPolicy and the number of bytes read.
//
// Signing policy byte encoding structure:
// 2 bytes - size
// 3 bytes - rewardEpochId
// 4 bytes - startingVotingRoundId
// 2 bytes - threshold
// 32 bytes - randomSeed
// array of 'size':
// - 20 bytes address
// - 2 bytes weight
// Total 43 + size * (20 + 2) bytes
func FromRawBytes(bytes []byte) (*SigningPolicy, int, error) {
	p := 0
	size := int(decodeUint32(bytes[p : p+2]))
	p += 2
	expectedLength := 43 + size*(common.AddressLength+2)
	if len(bytes) < expectedLength {
		return nil, p, errors.Errorf("message to short for decoding signing policy: expected >=%d, got %d", expectedLength, len(bytes))
	}

	epoch := decodeUint32(bytes[p : p+3])
	p += 3
	startingRound := decodeUint32(bytes[p : p+4])
	p += 4
	threshold := binary.BigEndian.Uint16(bytes[p : p+2])
	p += 2
	seed := common.BytesToHash(bytes[p : p+common.HashLength])
	p += common.HashLength

	signers := make([]common.Address, size)
	weights := make([]uint16, size)
	totalWeight := 0
	for i := 0; i < size; i++ {
		address := common.BytesToAddress(bytes[p : p+common.AddressLength])
		p += common.AddressLength
		weight := binary.BigEndian.Uint16(bytes[p : p+2])
		p += 2

		signers[i] = address
		weights[i] = weight
		totalWeight += int(weight)
	}

	if totalWeight > math.MaxUint16 {
		return nil, p, errors.New("total weight exceeds maximum uint16 value")
	}

	return &SigningPolicy{
		RewardEpochID:      int64(epoch),
		StartVotingRoundID: startingRound,
		Threshold:          threshold,
		Seed:               new(big.Int).SetBytes(seed[:]),
		rawBytes:           append([]byte(nil), bytes[:p]...),
		Voters:             voters.NewSet(signers, weights, nil),
	}, p, nil
}

// decodeUint32 decodes a big-endian uint32 from a variable length byte slice of up to 4 bytes.
func decodeUint32(bytes []byte) uint32 {
	if len(bytes) > 4 {
		logger.Fatal("invalid length for decoding uint32: %d", len(bytes))
	}
	var tmpUint32 = make([]byte, 4)
	padding := 4 - len(bytes)
	copy(tmpUint32[padding:], bytes)

	return binary.BigEndian.Uint32(tmpUint32[:])
}
