package ed25519

import (
	"crypto/ed25519"
	"encoding/hex"
	"strings"
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

// Vectors ported from rippled src/test/protocol/SecretKey_test.cpp ed25519TestVectors.
// Each entry: (16-byte seed, 33-byte 0xED-prefixed pub key, 32-byte raw priv, address).
// Rippled's generateSecretKey(KeyType::ed25519, seed) = sha512Half(seed) (SecretKey.cpp),
// derivePublicKey prepends 0xED to the 32-byte ed25519 public key, and calcAccountID
// is RIPEMD160(SHA256(pubkey)) (AccountID.cpp).
func TestKeyDerivationVectors(t *testing.T) {
	tests := []struct {
		seed   string
		pubKey string
		secKey string
		addr   string
	}{
		{
			seed:   "AF41FF66F75EBD3A6B18FB7A1DF61C97",
			pubKey: "ED48CBBBE0EE7B8686A7DE9F0A0159734E65F9C369947F2E2696232B461E553213",
			secKey: "1A1097FCD9CE4E1DA24666B698879766E1757547D1D4E364B64355F7C84BA0F3",
			addr:   "rVAEQBhWT6nZ4woEifdN3TMMdUZaxeXnR",
		},
		{
			seed:   "140C1D081319339C799DC6A165951BE1",
			pubKey: "ED3BC82EF45F8909CC00F8B7AAF05931681411758C117124875066C28398FE156D",
			secKey: "FE3E5A82B80DD82E915F7638942A332CE3068879740C7E90E220A4FB0B37CEC8",
			addr:   "rK57dJ9533WtoY8NNwVWGY7ffuAc8WCcPE",
		},
		{
			seed:   "865323B6E45AD6EEFA279FA584852ED4",
			pubKey: "EDE83422EBE075707312661E4B033A29BA86386230500CF28CF165F3C21E906D00",
			secKey: "A3C4E243A4644E738A247A59AABB5C89E4090D1B7302F245826487C838DA6989",
			addr:   "rfZiEDieHSHsQJ1UNfv2jYDuQawdRSBFwz",
		},
	}

	for _, tc := range tests {
		t.Run(tc.addr, func(t *testing.T) {
			seed, err := hex.DecodeString(tc.seed)
			require.NoError(t, err)

			rawPriv := hash.Sha512Half(seed)
			require.Equal(t, tc.secKey, strings.ToUpper(hex.EncodeToString(rawPriv)))

			priv := ed25519.NewKeyFromSeed(rawPriv)

			pub, err := PrvToPub(priv)
			require.NoError(t, err)
			require.Equal(t, tc.pubKey, pub)

			addr, err := PrvToAddress(priv)
			require.NoError(t, err)
			require.Equal(t, tc.addr, addr)
		})
	}
}

// Vector from rippled src/test/protocol/Seed_test.cpp testKeypairGenerationAndSigning:
// passphrase "masterpassphrase", KeyType::ed25519 -> address rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf.
// generateSeed(passphrase) = sha512Half(passphrase)[:16]; generateSecretKey(ed25519, seed) = sha512Half(seed).
func TestKeyDerivationPassphrase(t *testing.T) {
	seed := hash.Sha512Half([]byte("masterpassphrase"))[:16]
	rawPriv := hash.Sha512Half(seed)
	priv := ed25519.NewKeyFromSeed(rawPriv)

	addr, err := PrvToAddress(priv)
	require.NoError(t, err)
	require.Equal(t, "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", addr)
}

// Cross-reference: rippled derivePublicKey(KeyType::ed25519) (SecretKey.cpp)
// returns a 33-byte key with buf[0] = 0xED. Verify Validate rejects any other prefix.
func TestValidateRejectsWrongPrefix(t *testing.T) {
	_, priv, err := ed25519.GenerateKey(nil)
	require.NoError(t, err)

	msg := []byte("payload")
	sig := ed25519.Sign(priv, msg)

	pub, err := PrvToPub(priv)
	require.NoError(t, err)

	bad := "02" + pub[2:]
	_, err = Validate(msg, sig, bad)
	require.Error(t, err)

	shortPub := pub[:len(pub)-2]
	_, err = Validate(msg, sig, shortPub)
	require.Error(t, err)

	_, err = Validate(msg, sig[:len(sig)-1], pub)
	require.Error(t, err)
}
