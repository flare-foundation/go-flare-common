// Package policy provides signing policy management, parsing, and storage for Flare reward epochs.
package policy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/flare-foundation/go-flare-common/pkg/convert"

	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

// SigningPolicy holds the signing policy data for a reward epoch.
type SigningPolicy struct {
	RewardEpochID      uint32
	StartVotingRoundID uint32
	Threshold          uint16
	Seed               *big.Int
	rawBytes           []byte
	blockTimestamp     uint64

	// The set of all Voters and their weights
	Voters *voters.Set
}

// RawBytes returns the raw signing policy bytes.
func (sp *SigningPolicy) RawBytes() []byte {
	return sp.rawBytes
}

// Hash returns hash of the signing policy.
func (sp *SigningPolicy) Hash() []byte {
	return Hash(sp.rawBytes)
}

// NewSigningPolicy creates a SigningPolicy from a SigningPolicyInitialized event.
//
// Mapping from submitAddress to signingPolicyAddress can be added if needed.
func NewSigningPolicy(r *relay.RelaySigningPolicyInitialized, submitToSigning map[common.Address]common.Address) *SigningPolicy {
	return &SigningPolicy{
		RewardEpochID:      uint32(r.RewardEpochId.Uint64()), //nolint:gosec // r.RewardEpochId is uint24 in the event
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
	if len(b) < 2 {
		return nil, 0, errors.New("message too short for decoding signing policy")
	}

	p := 0
	size32, err := convert.BytesToUint32(b[p : p+2])
	if err != nil {
		return nil, p, fmt.Errorf("reading size: %w", err)
	}
	size := int(size32)
	p += 2

	expectedLength := 43 + size*(common.AddressLength+2)
	if len(b) < expectedLength {
		return nil, p, fmt.Errorf("message too short for decoding signing policy: expected >=%d, got %d", expectedLength, len(b))
	}

	epoch, err := convert.BytesToUint32(b[p : p+3])
	if err != nil {
		return nil, p, fmt.Errorf("reading epoch: %w", err)
	}
	p += 3

	startingRound, err := convert.BytesToUint32(b[p : p+4])
	if err != nil {
		return nil, p, fmt.Errorf("reading starting round: %w", err)
	}
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
		RewardEpochID:      epoch,
		StartVotingRoundID: startingRound,
		Threshold:          threshold,
		Seed:               new(big.Int).SetBytes(seed[:]),
		rawBytes:           slices.Clone(b[:p]),
		Voters:             voters.NewSet(signers, weights, nil),
	}, p, nil
}

// Hash computes hash of a signing policy from signingPolicyBytes.
func Hash(b []byte) []byte {
	if len(b)%32 != 0 {
		padded := make([]byte, len(b)+32-len(b)%32)
		copy(padded, b)
		b = padded
	}
	hash := crypto.Keccak256(b[:32], b[32:64])
	for i := 2; i < len(b)/32; i++ {
		hash = crypto.Keccak256(hash, b[i*32:(i+1)*32])
	}
	return hash
}
