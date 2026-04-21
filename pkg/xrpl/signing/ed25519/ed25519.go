// Package ed25519 provides XRPL Ed25519 key management and transaction signing.
package ed25519

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/seed"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/utils"
)

const (
	ed25519Prefix = 0xed
)

// SignTxMultisig signs a transaction for multi-signing using an Ed25519 private key.
func SignTxMultisig(tx map[string]any, prv ed25519.PrivateKey) (*signer.Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := encoding.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("encoding tx: %w", err)
	}

	accID, err := PrvToID(prv)
	if err != nil {
		return nil, fmt.Errorf("getting account id: %w", err)
	}

	msg, err := utils.Prepare(encoded, true, accID)
	if err != nil {
		return nil, fmt.Errorf("preparing message: %w", err)
	}

	signature := ed25519.Sign(prv, msg)

	add, err := PrvToAddress(prv)
	if err != nil {
		return nil, fmt.Errorf("getting address: %w", err)
	}

	pub, err := PrvToPub(prv)
	if err != nil {
		return nil, fmt.Errorf("getting address: %w", err)
	}

	return &signer.Signer{
		Account:       add,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// PrivKeyFromSecret converts a xrpl secret string to an Ed25519 private key.
// The secret should be a base58-encoded XRPL secret starting with 'sEd'.
func PrivKeyFromSecret(secret string) (ed25519.PrivateKey, error) {
	// xrpl.js Ed25519 seed is 23 decoded bytes → at most ~32 chars encoded; cap at 40 with slack.
	const encodedSecretCap = 40
	if len(secret) > encodedSecretCap {
		return nil, fmt.Errorf("secret too long, got %d chars", len(secret))
	}
	secretBytes, err := base58.XRPLCoder.Decode(secret)
	if err != nil {
		return nil, fmt.Errorf("decoding secret: %w", err)
	}

	if len(secretBytes) < 7 { // 3 prefix bytes + at least 4 checksum bytes
		return nil, errors.New("invalid secret length")
	}

	cs := hash.Checksum(secretBytes[:len(secretBytes)-4])

	if !bytes.Equal(cs, secretBytes[len(secretBytes)-4:]) {
		return nil, errors.New("invalid checksum")
	}

	if !bytes.Equal(secretBytes[:3], []byte{0x01, 0xe1, 0x4b}) {
		return nil, errors.New("invalid secret start")
	}

	stripped := secretBytes[3 : len(secretBytes)-4] // strip checksum
	rawPriv := hash.Sha512Half(stripped)

	prv := ed25519.NewKeyFromSeed(rawPriv)

	return prv, nil
}

// PrivKeyFromSeed derives an Ed25519 private key from a 16-byte rippled family seed.
// Mirrors rippled generateSecretKey(KeyType::ed25519, seed) in SecretKey.cpp:283-289: key = sha512Half(seed).
func PrivKeyFromSeed(familySeed []byte) (ed25519.PrivateKey, error) {
	if len(familySeed) != seed.Length {
		return nil, fmt.Errorf("invalid seed length %d, want %d", len(familySeed), seed.Length)
	}
	return ed25519.NewKeyFromSeed(hash.Sha512Half(familySeed)), nil
}

// PrivKeyFromFamilySeed derives an Ed25519 private key from a rippled-native base58 family seed string (TokenType::FamilySeed, prefix 0x21).
func PrivKeyFromFamilySeed(s string) (ed25519.PrivateKey, error) {
	payload, err := seed.DecodeFamilySeed(s)
	if err != nil {
		return nil, fmt.Errorf("decoding family seed: %w", err)
	}
	return PrivKeyFromSeed(payload)
}

// PrvToAddress calculates xrpl address for Ed25519 private key.
func PrvToAddress(prv ed25519.PrivateKey) (string, error) {
	accountID, err := PrvToID(prv)
	if err != nil {
		return "", fmt.Errorf("private key to address: %w", err)
	}

	add, err := address.Address(accountID)
	if err != nil {
		return "", fmt.Errorf("private key to address: %w", err)
	}

	return add, nil
}

// PrvToID calculates xrpl account ID for Ed25519 private key.
func PrvToID(prv ed25519.PrivateKey) ([]byte, error) {
	if len(prv) != 64 {
		return nil, fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return hash.Sha256RipeMD160(pub), nil
}

// PrvToPub calculates (xrpl) public key in hex string for Ed25519 private key.
// Public key is ED prefixed.
func PrvToPub(prv ed25519.PrivateKey) (string, error) {
	if len(prv) != 64 {
		return "", fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return strings.ToUpper(hex.EncodeToString(pub)), nil
}

// Validate checks that the sig is a valid Ed25519 signature of msg by pub.
// The public key should hex string be prefixed with ED and 33 bytes (including the prefix).
func Validate(msg, sig []byte, pub string) (bool, error) {
	pubBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false, fmt.Errorf("reading pub: %w", err)
	}

	if pubBytes[0] != 0xed {
		return false, errors.New("pub key should ED (or ed) prefixed")
	}

	if len(pubBytes) != ed25519.PublicKeySize+1 {
		return false, fmt.Errorf("invalid pubKey length (require %d bytes)", ed25519.PublicKeySize+1)
	}

	if len(sig) != ed25519.SignatureSize {
		return false, fmt.Errorf("invalid signature length (require 64 bytes): %s", pub)
	}

	return ed25519.Verify(pubBytes[1:], msg, sig), nil
}
