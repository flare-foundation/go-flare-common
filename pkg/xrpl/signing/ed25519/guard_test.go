package ed25519

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSignTxMultisigRejectsNilTx verifies SignTxMultisig returns an error
// instead of panicking on a nil tx map.
func TestSignTxMultisigRejectsNilTx(t *testing.T) {
	_, err := SignTxMultisig(nil, nil)
	require.Error(t, err)
}
