// Package voters provides voter set management and weighted random selection for Flare signing policies.
package voters

import (
	"encoding/binary"
	"errors"
	"fmt"
	"maps"
	"math"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// VoterData holds the index and weight of a voter in the signing policy.
type VoterData struct {
	Index  int
	Weight uint16
}

// Set represents a weighted set of voters for a signing policy.
type Set struct {
	voters      []common.Address // signingPolicyAddress
	weights     []uint16
	TotalWeight uint16
	thresholds  []uint16

	VoterDataMap           map[common.Address]VoterData // signingPolicyAddressToWeight
	SubmitToSigningAddress map[common.Address]common.Address
}

// NewSet creates Set from voters' addresses and their weights.
// Optionally, a map from submitAddresses to signingAddress can be added.
//
// Returns an error if voters and weights differ in length or if voters contains
// a duplicate address. The signing policy emitted by the Flare smart contracts
// is guaranteed to be free of duplicates, so a duplicate here indicates a
// caller-side bug or corrupt input that would otherwise silently double-count
// weight in the slice-indexed selection path while VoterDataMap dedupes.
func NewSet(voters []common.Address, weights []uint16, submitToSigningAddress map[common.Address]common.Address) (*Set, error) {
	if len(voters) != len(weights) {
		return nil, fmt.Errorf("voters/weights length mismatch: %d vs %d", len(voters), len(weights))
	}

	// Defensive copies: callers' slice/map mutations after NewSet must not corrupt the Set.
	vs := Set{
		voters:     slices.Clone(voters),
		weights:    slices.Clone(weights),
		thresholds: make([]uint16, len(weights)),
	}
	// Accumulate into a wider type so we can detect overflow before it
	// wraps TotalWeight and silently corrupts thresholds (non-monotonic
	// after wrap). The smart-contract path enforces this at decode time
	// (see policy.FromRawBytes); NewSet must enforce it for every caller
	// — overflowing here would silently disagree with chain consensus on
	// the selected voter.
	var sum uint32
	for i, w := range weights {
		vs.thresholds[i] = uint16(sum)
		sum += uint32(w)
		if sum > math.MaxUint16 {
			return nil, fmt.Errorf("total weight exceeds uint16 max: %d", sum)
		}
	}
	if sum == 0 {
		return nil, errors.New("total weight is zero")
	}
	vs.TotalWeight = uint16(sum)

	vMap := make(map[common.Address]VoterData, len(voters))
	for i, voter := range vs.voters {
		if _, ok := vMap[voter]; ok {
			return nil, fmt.Errorf("duplicate voter address at index %d: %s", i, voter.Hex())
		}
		vMap[voter] = VoterData{
			Index:  i,
			Weight: vs.weights[i],
		}
	}
	vs.VoterDataMap = vMap
	vs.SubmitToSigningAddress = maps.Clone(submitToSigningAddress)

	return &vs, nil
}

// InitialHashSeed returns initial seed for voter selection.
//
// Seed is keccak256 hash of rewardEpochSeed, protocolID, and votingRoundID each written in its own 32-byte slot.
// Returns an error if rewardEpochSeed is negative or exceeds 256 bits.
func InitialHashSeed(rewardEpochSeed *big.Int, protocolID byte, votingRoundID uint32) (common.Hash, error) {
	seed := make([]byte, 96)
	// 0-31 bytes are filled with the reward epoch seed
	if rewardEpochSeed != nil {
		if rewardEpochSeed.Sign() < 0 || rewardEpochSeed.BitLen() > 256 {
			return common.Hash{}, fmt.Errorf("rewardEpochSeed out of 256-bit range (bitlen=%d)", rewardEpochSeed.BitLen())
		}
		rewardEpochSeed.FillBytes(seed[0:32])
	}
	// 32-63 bytes are filled with the protocol ID
	seed[63] = protocolID
	// 64-95 bytes are filled with the voting round ID
	binary.BigEndian.PutUint32(seed[92:96], votingRoundID)

	return crypto.Keccak256Hash(seed), nil
}

// SelectVoters returns a set of signingPolicyAddresses that are prioritized to finalize.
func (vs *Set) SelectVoters(rewardEpochSeed *big.Int, protocolID byte, votingRoundID uint32, thresholdBIPS uint16) (map[common.Address]bool, error) {
	seed, err := InitialHashSeed(rewardEpochSeed, protocolID, votingRoundID)
	if err != nil {
		return nil, fmt.Errorf("computing seed: %w", err)
	}
	return vs.RandomSelectThresholdWeightVoters(seed, thresholdBIPS)
}

// RandomSelectThresholdWeightVoters randomly selects voters until the accumulated weight reaches the threshold.
func (vs *Set) RandomSelectThresholdWeightVoters(randomSeed common.Hash, thresholdBIPS uint16) (map[common.Address]bool, error) {
	// We limit the threshold to 5000 BIPS to avoid long running loops
	// In practice it will be used with around 1000 BIPS or lower.
	if thresholdBIPS > 5000 {
		return nil, errors.New("threshold must be between 0 and 5000 BIPS")
	}

	selectedWeight := uint16(0)
	thresholdWeight := uint16(uint64(vs.TotalWeight) * uint64(thresholdBIPS) / 10000)
	currentSeed := randomSeed

	selected := make(map[common.Address]bool)

	// If threshold weight is not too big, the loop should end quickly
	for selectedWeight < thresholdWeight {
		index := vs.selectVoterIndex(currentSeed)
		selectedAddress := vs.voters[index]
		if !selected[selectedAddress] {
			selected[selectedAddress] = true
			selectedWeight += vs.weights[index]
		}
		currentSeed = crypto.Keccak256Hash(currentSeed.Bytes())
	}
	return selected, nil
}

// selectVoterIndex selects a random voter based on the provided randomNumber.
func (vs *Set) selectVoterIndex(randomNumber common.Hash) int {
	randomWeight := randomNumber.Big()
	randomWeight = randomWeight.Mod(randomWeight, big.NewInt(int64(vs.TotalWeight)))
	return vs.BinarySearch(uint16(randomWeight.Uint64()))
}

// BinarySearch finds the highest index of the threshold that is less than or equal to the value.
//
// value has to be lower then vs.totalWeight otherwise, the function will panic.
// Returns -1 if the voter set is empty.
func (vs *Set) BinarySearch(value uint16) int {
	if len(vs.thresholds) == 0 {
		return -1
	}
	if value > vs.TotalWeight {
		panic("Value must be between 0 and total weight")
	}
	left := 0
	right := len(vs.thresholds) - 1
	var mid int
	if vs.thresholds[right] <= value {
		return right
	}
	for left < right {
		mid = (left + right) / 2

		switch {
		case vs.thresholds[mid] < value:
			left = mid + 1
		case vs.thresholds[mid] > value:
			right = mid
		default:
			return mid
		}
	}
	return left - 1
}

// VoterAddress returns the signing policy address of the voter at the specified index.
func (vs *Set) VoterAddress(index int) common.Address {
	return vs.voters[index]
}

// VoterWeight returns the weight of the voter with index.
func (vs *Set) VoterWeight(index int) uint16 {
	return vs.weights[index]
}

// Count returns the number of voters.
func (vs *Set) Count() int {
	return len(vs.voters)
}

// VoterIndex returns the signing policy index of the signingPolicy address.
//
// If address is not among the signingPolicy addresses, it returns -1.
func (vs *Set) VoterIndex(address common.Address) int {
	voterData, ok := vs.VoterDataMap[address]
	if !ok {
		return -1
	}
	return voterData.Index
}

// VoterWeightForAddress returns the weight for the signing policy address.
func (vs *Set) VoterWeightForAddress(a common.Address) uint16 {
	d, exists := vs.VoterDataMap[a]
	if !exists {
		return 0
	}
	return d.Weight
}

// Voters returns a fresh copy of voter addresses; mutating the result does not affect the Set.
func (vs *Set) Voters() []common.Address {
	out := make([]common.Address, len(vs.voters))
	copy(out, vs.voters)
	return out
}
