package policy

import (
	"cmp"
	"fmt"
	"sort"
	"sync"
)

// Storage stores signingPolicies.
type Storage struct {
	// sorted list of signingPolicies, sorted by rewardEpochID (and also by startVotingRoundID)
	spList []*SigningPolicy

	// mutex
	sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		spList: make([]*SigningPolicy, 0, 10),
	}
}

// Does not lock the structure, should be called from a function that does lock.
// We assume that the list is sorted by rewardEpochID and also by startVotingRoundID.
func (s *Storage) findByVotingRoundID(votingRoundID uint32) *SigningPolicy {
	i, found := sort.Find(len(s.spList), func(i int) int {
		return cmp.Compare(votingRoundID, s.spList[i].StartVotingRoundID)
	})
	if found {
		return s.spList[i]
	}
	if i == 0 {
		return nil
	}
	return s.spList[i-1]
}

// Add adds a signingPolicy to the storage.
// The added signingPolicy must be a successor of the latest signingPolicy in the storage.
func (s *Storage) Add(sp *SigningPolicy) error {
	s.Lock()
	defer s.Unlock()

	if len(s.spList) > 0 {
		// check consistency, previous epoch should be already added
		if s.spList[len(s.spList)-1].RewardEpochID != sp.RewardEpochID-1 {
			return fmt.Errorf("missing signing policy for reward epoch ID %d", sp.RewardEpochID-1)
		}
		// should be sorted by voting round ID, should not happen
		if sp.StartVotingRoundID < s.spList[len(s.spList)-1].StartVotingRoundID {
			return fmt.Errorf("signing policy for reward epoch ID %d has larger start voting round ID than previous policy",
				sp.RewardEpochID)
		}
	}

	s.spList = append(s.spList, sp)
	return nil
}

// Return the signingPolicy for the voting round, or nil if not found.
// Also returns true if the policy is the last one or false otherwise.
func (s *Storage) ForVotingRound(votingRoundID uint32) (*SigningPolicy, bool) {
	s.Lock()
	defer s.Unlock()

	sp := s.findByVotingRoundID(votingRoundID)
	if sp == nil {
		return nil, false
	}
	return sp, sp.RewardEpochID == s.spList[len(s.spList)-1].RewardEpochID
}

// RemoveBefore removes all signingPolicies that ended strictly before votingRoundID.
// Returns the list of removed reward epoch ids.
func (s *Storage) RemoveBefore(votingRoundID uint32) []uint32 {
	s.Lock()
	defer s.Unlock()

	var removedRewardEpochIDs []uint32
	for len(s.spList) > 1 && s.spList[1].StartVotingRoundID < votingRoundID {
		removedRewardEpochIDs = append(removedRewardEpochIDs, uint32(s.spList[0].RewardEpochID))
		s.spList[0] = nil
		s.spList = s.spList[1:]
	}
	return removedRewardEpochIDs
}

// OldestStored returns the oldest signingPolicy that is in the storage or nil if the storage is empty.
func (s *Storage) OldestStored() *SigningPolicy {
	s.Lock()
	defer s.Unlock()

	if len(s.spList) == 0 {
		return nil
	}
	return s.spList[0]
}
