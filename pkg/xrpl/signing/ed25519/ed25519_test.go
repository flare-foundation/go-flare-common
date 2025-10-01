package ed25519

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/stretchr/testify/require"
)

func TestAddressED(t *testing.T) {
	seed, err := hex.DecodeString("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	fullPriv := ed25519.NewKeyFromSeed(seed)

	require.Equal(t, seed, fullPriv.Seed())

	addr, err := PrvToAddress(fullPriv)
	require.NoError(t, err)

	require.Equal(t, "rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8", addr)

	pub, err := PrvToPub(fullPriv)
	require.NoError(t, err)

	require.Equal(t, "ED9AA5EBF4FB942A7D81545C42ADDB86B389E56FBAF3F1295A5B6E1CDBCE2424BB", pub)
}

func TestSignVerify(t *testing.T) {
	_, priv, err := ed25519.GenerateKey(nil)
	require.NoError(t, err)

	msg := []byte("my message")

	signature := ed25519.Sign(priv, msg)

	pkStr, err := PrvToPub(priv)
	require.NoError(t, err)

	ok, err := Validate(msg, signature, pkStr)
	require.NoError(t, err)
	require.True(t, ok)
}

func TestPrivateKeyFromSecret(t *testing.T) {
	secret := "sEd7TybX8deEyDjeH5WX2rJ527jvfRs"

	priv, err := PrivKeyFromSecret(secret)
	require.NoError(t, err)

	// Verify the private key is valid
	require.Equal(t, 64, len(priv))

	// Test that we can derive address from the private key
	addr, err := PrvToAddress(priv)
	require.NoError(t, err)
	require.NotEmpty(t, addr)

	expectedAddr := "rLjVxe9ub7SDyouaWVjdjCxNmBE9TBW5qi"
	require.Equal(t, expectedAddr, addr)

	// Test with invalid secret (wrong checksum)
	invalidSecret := "sEdTM1uX8pu2do5XvTnutH6HsouMaM3"
	_, err = PrivKeyFromSecret(invalidSecret)
	require.Error(t, err)

	// Test with invalid secret format
	invalidFormat := "invalid_secret"
	_, err = PrivKeyFromSecret(invalidFormat)
	require.Error(t, err)

	// Zero string
	_, err = PrivKeyFromSecret("")
	require.Error(t, err)

	// Invalid starts
	z := []byte("random")
	cs := hash.Checksum(z)
	z = append(z, cs...)
	invalidStart := base58.XRPLCoder.Encode(z)
	_, err = PrivKeyFromSecret(invalidStart)
	require.Error(t, err)
}
