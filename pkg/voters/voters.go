package voters

import (
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-fundation/go-flare-common/pkg/logger"
)

type VoterData struct {
	Index  int
	Weight uint16
}

type Set struct {
	voters      []common.Address //signingPolicyAddress
	weights     []uint16
	TotalWeight uint16
	thresholds  []uint16

	VoterDataMap           map[common.Address]VoterData //signingPolicyAddressToWeight
	SubmitToSigningAddress map[common.Address]common.Address
}

// NewSet creates Set from voters' addresses and their weights.
// Optionally, a map from submitAddresses to signingAddress can be added.
//
// There has to be the same number of voters and weights.
func NewSet(voters []common.Address, weights []uint16, SubmitToSigningAddress map[common.Address]common.Address) *Set {
	if len(voters) != len(weights) {
		logger.Errorf("New voter set: mismatched lengths: %d voters and  %d weights", len(voters), len(weights))
		return nil
	}

	vs := Set{
		voters:     voters,
		weights:    weights,
		thresholds: make([]uint16, len(weights)),
	}
	// sum does not exceed uint16 limit, guaranteed by the smart contract
	for i, w := range weights {
		vs.thresholds[i] = vs.TotalWeight
		vs.TotalWeight += w
	}

	vMap := make(map[common.Address]VoterData)
	for i, voter := range vs.voters {
		logger.Debugf("New voter: %v", voter)

		if _, ok := vMap[voter]; !ok {
			vMap[voter] = VoterData{
				Index:  i,
				Weight: vs.weights[i],
			}
		}
	}
	vs.VoterDataMap = vMap
	vs.SubmitToSigningAddress = SubmitToSigningAddress

	return &vs
}

// InitialHashSeed returns initial seed for voter selection.
//
// Seed is keccak256 hash of rewardEpochSeed, protocolID, and votingRoundID each written in its own 32-byte slot.
func InitialHashSeed(rewardEpochSeed *big.Int, protocolID byte, votingRoundID uint32) common.Hash {
	seed := make([]byte, 96)
	// 0-31 bytes are filled with the reward epoch seed
	if rewardEpochSeed != nil {
		rewardEpochSeed.FillBytes(seed[0:32])
	}
	// 32-63 bytes are filled with the protocol ID
	seed[63] = protocolID
	// 64-95 bytes are filled with the voting round ID
	binary.BigEndian.PutUint32(seed[92:96], votingRoundID)

	return crypto.Keccak256Hash(seed)
}

// SelectVoters returns a set of signingPolicyAddresses that are prioritized to finalize.
func (vs *Set) SelectVoters(rewardEpochSeed *big.Int, protocolID byte, votingRoundID uint32, thresholdBIPS uint16) (map[common.Address]bool, error) {
	seed := InitialHashSeed(rewardEpochSeed, protocolID, votingRoundID)
	return vs.RandomSelectThresholdWeightVoters(seed, thresholdBIPS)
}

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
func (vs *Set) BinarySearch(value uint16) int {
	if value > vs.TotalWeight {
		panic("Value must be between 0 and total weight")
	}
	left := 0
	right := len(vs.thresholds) - 1
	mid := 0
	if vs.thresholds[right] <= value {
		return right
	}
	for left < right {
		mid = (left + right) / 2
		if vs.thresholds[mid] < value {
			left = mid + 1
		} else if vs.thresholds[mid] > value {
			right = mid
		} else {
			return mid
		}
	}
	return left - 1
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
