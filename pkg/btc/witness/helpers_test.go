package witness_test

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

// signDER signs a hash the way a TEE machine does: DER, without the trailing
// sighash-type byte, which assembly appends.
func signDER(priv *btcec.PrivateKey, hash []byte) ([]byte, error) {
	return ecdsa.Sign(priv, hash).Serialize(), nil
}
