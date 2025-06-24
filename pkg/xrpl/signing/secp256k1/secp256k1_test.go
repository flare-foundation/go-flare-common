package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"

	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	priv, err := crypto.HexToECDSA("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	require.Equal(t, "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd", PrvToAddress(priv))
	require.Equal(t, "038940A036EE380369B9FCC5929A0D431ABE789C8A44E5F48F564E2F6EB608B543", PrvToPub(priv))
}

func SignT(message []byte, privKey *ecdsa.PrivateKey) []byte {
	hash := hash.Sha512Half(message)
	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, hash)

	return sig2.Serialize()
}

func TestSignatureMarshaling(t *testing.T) {
	prv, err := crypto.HexToECDSA("8F1430C6570B7AB19F4FBB0F1F9A3071192EFB133F6216103F8B4D380B55DE75")
	require.NoError(t, err)

	for range 500 {
		for j := range 5 {
			msg := fmt.Appendf(nil, "neki%d", j)

			sig, err := sign(hash.Sha512Half(msg), prv)
			require.NoError(t, err, j)

			sigXRPL, err := SignXRPL(msg, prv)
			require.NoError(t, err, j)

			isValid := sig.Verify(hash.Sha512Half(msg), &prv.PublicKey)
			require.Truef(t, isValid, "invalid signature %v for priv key %v", j, prv.D)

			sigDER := sig.DER()

			sig2DER := SignT(msg, prv)
			require.Equal(t, sig2DER, sigDER, j)

			sigMar, err := MarshalDER(sigDER)
			require.NoError(t, err, j)

			require.Equal(t, sig.r, sigMar.r, j)
			require.Equal(t, sig.s, sigMar.s, j)

			require.Equal(t, sig2DER, sigXRPL, j)

			ok, err := Validate(msg, sigDER, hex.EncodeToString(toBytesCompressed(&prv.PublicKey)))
			require.NoError(t, err, j)

			require.True(t, ok)

			pub, err := sig.Recover(hash.Sha512Half(msg))
			require.NoError(t, err, j)

			require.Equal(t, pub, &prv.PublicKey)
		}

		prv, err = crypto.GenerateKey() // change prv for the next run
		require.NoError(t, err)
	}
}

func TestMarshaling(t *testing.T) {
	sigs := []string{"30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		"30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1"}

	for j, sig := range sigs {
		b, err := hex.DecodeString(sig)
		require.NoError(t, err, j)

		_, err = MarshalDER(b)
		require.NoError(t, err, j)
	}
}
