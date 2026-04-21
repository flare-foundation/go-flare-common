package secp256k1

import (
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/seed"
)

// maxAttempts bounds the validity-check loop; rippled caps at 128 (SecretKey.cpp:85, 144).
const maxAttempts = 128

// PrivKeyFromSeed derives an XRPL secp256k1 private key from a 16-byte rippled family seed, ordinal 0.
// Mirrors rippled detail::Generator in SecretKey.cpp:120-200.
func PrivKeyFromSeed(familySeed []byte) (*ecdsa.PrivateKey, error) {
	if len(familySeed) != seed.Length {
		return nil, fmt.Errorf("invalid seed length %d, want %d", len(familySeed), seed.Length)
	}

	root, err := deriveRootScalar(familySeed)
	if err != nil {
		return nil, fmt.Errorf("deriving root: %w", err)
	}

	generator := scalarToCompressedPub(root)

	tweak, err := deriveTweakScalar(generator, 0)
	if err != nil {
		return nil, fmt.Errorf("deriving tweak: %w", err)
	}

	account, err := addScalarsModN(root, tweak)
	if err != nil {
		return nil, fmt.Errorf("applying tweak: %w", err)
	}

	return scalarToPrivKey(account), nil
}

// PrivKeyFromFamilySeed derives an XRPL secp256k1 private key from a rippled-native base58 family seed string.
func PrivKeyFromFamilySeed(s string) (*ecdsa.PrivateKey, error) {
	payload, err := seed.DecodeFamilySeed(s)
	if err != nil {
		return nil, fmt.Errorf("decoding family seed: %w", err)
	}
	return PrivKeyFromSeed(payload)
}

// deriveRootScalar implements rippled detail::deriveDeterministicRootKey (SecretKey.cpp:67-99).
//
//	buf[0:16]  = seed
//	buf[16:20] = big-endian uint32 counter
//	loop seq: candidate = sha512Half(buf); valid iff 0 < candidate < N.
func deriveRootScalar(familySeed []byte) (*big.Int, error) {
	var buf [20]byte
	copy(buf[:16], familySeed)

	for seq := range uint32(maxAttempts) {
		binary.BigEndian.PutUint32(buf[16:], seq)
		candidate := hash.Sha512Half(buf[:])
		if k, ok := validScalar(candidate); ok {
			return k, nil
		}
	}
	return nil, errors.New("unable to derive root from seed")
}

// deriveTweakScalar implements rippled detail::Generator::calculateTweak (SecretKey.cpp:126-158).
//
//	buf[0:33]  = generator (compressed pubkey of root)
//	buf[33:37] = big-endian uint32 ordinal (outer seq)
//	buf[37:41] = big-endian uint32 retry (inner subseq)
//	loop subseq: candidate = sha512Half(buf); valid iff 0 < candidate < N.
func deriveTweakScalar(generator []byte, ordinal uint32) (*big.Int, error) {
	if len(generator) != 33 {
		return nil, fmt.Errorf("generator must be 33 compressed bytes, got %d", len(generator))
	}
	var buf [41]byte
	copy(buf[:33], generator)
	binary.BigEndian.PutUint32(buf[33:37], ordinal)

	for subseq := range uint32(maxAttempts) {
		binary.BigEndian.PutUint32(buf[37:], subseq)
		candidate := hash.Sha512Half(buf[:])
		if k, ok := validScalar(candidate); ok {
			return k, nil
		}
	}
	return nil, errors.New("unable to derive tweak")
}

// validScalar reports whether b interpreted as a big-endian integer is a valid secp256k1 scalar (0 < k < N).
// Matches secp256k1_ec_seckey_verify: rejects 0 and any value >= N.
func validScalar(b []byte) (*big.Int, bool) {
	k := new(big.Int).SetBytes(b)
	if k.Sign() <= 0 {
		return nil, false
	}
	if k.Cmp(btcec.S256().N) >= 0 {
		return nil, false
	}
	return k, true
}

// addScalarsModN computes (a + b) mod N and rejects a zero result, matching secp256k1_ec_seckey_tweak_add.
func addScalarsModN(a, b *big.Int) (*big.Int, error) {
	out := new(big.Int).Add(a, b)
	out.Mod(out, btcec.S256().N)
	if out.Sign() == 0 {
		return nil, errors.New("tweak produced zero scalar")
	}
	return out, nil
}

// scalarToCompressedPub returns the 33-byte compressed secp256k1 public key for scalar k.
func scalarToCompressedPub(k *big.Int) []byte {
	curve := btcec.S256()
	x, y := curve.ScalarBaseMult(k.Bytes())

	out := make([]byte, 33)
	if y.Bit(0) == 1 {
		out[0] = pubKeyOddPrefix
	} else {
		out[0] = pubKeyEvenPrefix
	}
	xBytes := x.Bytes()
	copy(out[33-len(xBytes):], xBytes)
	return out
}

// scalarToPrivKey wraps a 32-byte secp256k1 scalar as a stdlib ecdsa.PrivateKey.
func scalarToPrivKey(k *big.Int) *ecdsa.PrivateKey {
	curve := btcec.S256()
	priv := new(ecdsa.PrivateKey)
	priv.Curve = curve
	priv.D = new(big.Int).Set(k)
	priv.X, priv.Y = curve.ScalarBaseMult(k.Bytes())
	return priv
}
