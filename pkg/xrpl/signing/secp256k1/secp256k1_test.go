package secp256k1

import (
	"crypto/ecdsa"

	"fmt"
	"testing"

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

func ECDSASign(message []byte, privKey *ecdsa.PrivateKey) []byte {
	msg := hash.Sha512Half(message)

	sign, _ := crypto.Sign(msg, privKey)

	return sign
}

func Verify(pub *ecdsa.PublicKey, message, sig []byte) bool {
	return ecdsa.VerifyASN1(pub, message, sig)
}

func TestSECtoDER(t *testing.T) {
	prv, _ := crypto.GenerateKey()

	for j := range 10003 {
		msg := fmt.Appendf(nil, "neki%d", j)

		sig1 := Sign(msg, prv)
		sig2 := ECDSASign(msg, prv)

		sig2E := SECToDER(sig2)

		require.Equal(t, sig1, sig2E, j)
	}
}
