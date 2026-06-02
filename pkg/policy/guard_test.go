package policy_test

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/stretchr/testify/require"
)

func TestNewSigningPolicyRejectsNil(t *testing.T) {
	_, err := policy.NewSigningPolicy(nil, nil)
	require.Error(t, err)

	// Non-nil event with a nil RewardEpochId must also be rejected.
	_, err = policy.NewSigningPolicy(&relay.RelaySigningPolicyInitialized{}, nil)
	require.Error(t, err)
}

func TestStorageAddRejectsNil(t *testing.T) {
	s := policy.NewStorage()
	require.Error(t, s.Add(nil))
}
