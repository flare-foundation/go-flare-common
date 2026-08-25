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

// ChainBoundHash returns the signing policy's source-bound hash for chainID.
//
// There is deliberately no no-arg variant: the two schemes are indistinguishable
// 32-byte values, so every call site must name the one it means. For the legacy
// scheme call Hash(sp.RawBytes()).
func (sp *SigningPolicy) ChainBoundHash(chainID uint64) []byte {
	return ChainBoundHash(chainID, sp.rawBytes)
}

// NewSigningPolicy creates a SigningPolicy from a SigningPolicyInitialized event.
//
// Mapping from submitAddress to signingPolicyAddress can be added if needed.
// Returns an error if the event's voters/weights are malformed (length mismatch
// or duplicate address). The smart contract guarantees both, so an error here
// indicates upstream corruption.
func NewSigningPolicy(r *relay.RelaySigningPolicyInitialized, submitToSigning map[common.Address]common.Address) (*SigningPolicy, error) {
	if r == nil || r.RewardEpochId == nil {
		return nil, errors.New("nil signing policy event")
	}

	vs, err := voters.NewSet(r.Voters, r.Weights, submitToSigning)
	if err != nil {
		return nil, fmt.Errorf("building voter set: %w", err)
	}
	return &SigningPolicy{
		RewardEpochID:      uint32(r.RewardEpochId.Uint64()), //nolint:gosec // r.RewardEpochId is uint24 in the event
		StartVotingRoundID: r.StartVotingRoundId,
		Threshold:          r.Threshold,
		Seed:               r.Seed,
		rawBytes:           slices.Clone(r.SigningPolicyBytes),
		blockTimestamp:     r.Timestamp,
		Voters:             vs,
	}, nil
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

	vs, err := voters.NewSet(signers, weights, nil)
	if err != nil {
		return nil, p, fmt.Errorf("building voter set: %w", err)
	}

	return &SigningPolicy{
		RewardEpochID:      epoch,
		StartVotingRoundID: startingRound,
		Threshold:          threshold,
		Seed:               new(big.Int).SetBytes(seed[:]),
		rawBytes:           slices.Clone(b[:p]),
		Voters:             vs,
	}, p, nil
}

// Hash computes the legacy chained-fold hash of a signing policy from
// signingPolicyBytes.
//
// Inputs shorter than 64 bytes are right-zero-padded to two 32-byte blocks.
//
// Retired on chain by RLY-23 in favour of ChainBoundHash, but still the correct
// answer for reward epochs below the new Relay's initialRewardEpochId, which
// toSigningPolicyHash delegates to the old Relay. No TEE component computes it
// any more — it is kept for consumers outside that stack.
func Hash(b []byte) []byte {
	const block = 32
	minLen := 2 * block
	switch {
	case len(b) < minLen:
		padded := make([]byte, minLen)
		copy(padded, b)
		b = padded
	case len(b)%block != 0:
		padded := make([]byte, len(b)+block-len(b)%block)
		copy(padded, b)
		b = padded
	}
	hash := crypto.Keccak256(b[:block], b[block:2*block])
	for i := 2; i < len(b)/block; i++ {
		hash = crypto.Keccak256(hash, b[i*block:(i+1)*block])
	}
	return hash
}

// ChainBoundHash returns keccak256(chainID || signingPolicyBytes) — the hash the
// Relay stores and returns from toSigningPolicyHash for reward epochs at or above
// its initialRewardEpochId, and therefore the hash data providers sign.
//
// Must stay byte-identical to Relay.sol's
// keccak256(abi.encodePacked(sourceChainId, signingPolicyBytes)): one keccak over
// the 32-byte source chain id followed by the raw 43 + 22*n encoding. Unlike Hash
// there is no padding — b is hashed exactly as given.
//
// chainID is the Relay's sourceChainId, which equals block.chainid on a home
// deployment (Relay.initialize enforces it).
func ChainBoundHash(chainID uint64, b []byte) []byte {
	var chainIDWord [32]byte
	binary.BigEndian.PutUint64(chainIDWord[24:], chainID)

	return crypto.Keccak256(chainIDWord[:], b)
}
