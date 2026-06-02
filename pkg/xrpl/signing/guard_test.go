package signing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestJoinMultisigRejectsNilTx verifies the Join helpers return an error
// instead of panicking on a nil tx map.
func TestJoinMultisigRejectsNilTx(t *testing.T) {
	_, err := JoinMultisig(nil, nil)
	require.Error(t, err)

	_, err = JoinMultisigJSON(nil, nil)
	require.Error(t, err)
}
