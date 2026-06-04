package voters_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

// TestVoterAccessorsOutOfRange verifies the index accessors return zero values
// instead of panicking when the index is out of range (e.g. the -1 that
// VoterIndex returns for an unknown address).
func TestVoterAccessorsOutOfRange(t *testing.T) {
	vs, err := voters.NewSet(testVoters, testWeights, nil)
	require.NoError(t, err)

	require.Equal(t, common.Address{}, vs.VoterAddress(-1))
	require.Equal(t, common.Address{}, vs.VoterAddress(len(testVoters)))
	require.Equal(t, uint16(0), vs.VoterWeight(-1))
	require.Equal(t, uint16(0), vs.VoterWeight(len(testWeights)))
}
