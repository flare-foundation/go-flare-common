package policy

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/big"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/pkg/errors"

	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

// Hash returns hash of the signing policy.
func (sp *SigningPolicy) Hash() []byte {
	return Hash(sp.rawBytes)
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
//
//  1. 2 bytes - size
//
//  2. 3 bytes - rewardEpochId
//
//  3. 4 bytes - startingVotingRoundId
//
//  4. 2 bytes - threshold
//
//  5. 32 bytes - randomSeed
//
//  6. array of 'size':
//
//     - 20 bytes - address
//
//     - 2 bytes - weight
//
// Total 43 + size * (20 + 2) bytes.
func FromRawBytes(b []byte) (*SigningPolicy, int, error) {
	p := 0
	size := int(decodeUint32(b[p : p+2]))
	p += 2
	expectedLength := 43 + size*(common.AddressLength+2)
	if len(b) < expectedLength {
		return nil, p, errors.Errorf("message to short for decoding signing policy: expected >=%d, got %d", expectedLength, len(b))
	}

	epoch := decodeUint32(b[p : p+3])
	p += 3
	startingRound := decodeUint32(b[p : p+4])
	p += 4
	threshold := binary.BigEndian.Uint16(b[p : p+2])
	p += 2
	seed := common.BytesToHash(b[p : p+common.HashLength])
	p += common.HashLength

	signers := make([]common.Address, size)
	weights := make([]uint16, size)
	totalWeight := 0
	for i := range size {
		address := common.BytesToAddress(b[p : p+common.AddressLength])
		p += common.AddressLength
		weight := binary.BigEndian.Uint16(b[p : p+2])
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
		rawBytes:           slices.Clone(b[:p]),
		Voters:             voters.NewSet(signers, weights, nil),
	}, p, nil
}

// decodeUint32 decodes a big-endian uint32 from a variable length byte slice of up to 4 bytes.
func decodeUint32(b []byte) uint32 {
	if len(b) > 4 {
		logger.Fatal("invalid length for decoding uint32: %d", len(b))
	}
	var tmpUint32 = make([]byte, 4)
	padding := 4 - len(b)
	copy(tmpUint32[padding:], b)

	return binary.BigEndian.Uint32(tmpUint32)
}

// Hash computes hash of a signing policy from signingPolicyBytes.
func Hash(b []byte) []byte {
	if len(b)%32 != 0 {
		b = append(b, make([]byte, 32-len(b)%32)...)
	}
	hash := crypto.Keccak256(b[:32], b[32:64])
	for i := 2; i < len(b)/32; i++ {
		hash = crypto.Keccak256(hash, b[i*32:(i+1)*32])
	}
	return hash
}
