package secp256k1

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestSignTxMultisigRejectsNilTx(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)

	_, err = SignTxMultisig(nil, prv)
	require.Error(t, err)
}

func TestSignTxMultisigRejectsNilKey(t *testing.T) {
	_, err := SignTxMultisig(map[string]any{}, nil)
	require.Error(t, err)
}

// TestSignTxMultisigRejectsDLessKey covers a non-nil key with no secret
// scalar, which previously panicked in PrvToID -> secp256k1PrvToPub.
func TestSignTxMultisigRejectsDLessKey(t *testing.T) {
	_, err := SignTxMultisig(map[string]any{}, &ecdsa.PrivateKey{})
	require.Error(t, err)
}

func TestToBytesCompressedNilSafe(t *testing.T) {
	require.Nil(t, toBytesCompressed(nil))
	require.Nil(t, toBytesCompressed(&ecdsa.PublicKey{}))
}
