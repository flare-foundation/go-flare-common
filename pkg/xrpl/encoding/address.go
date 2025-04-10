package encoding

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/hash"
	//nolint
)

const (
	ed25519Prefix = 0xed
	accountPrefix = 0x00

	pubKeyEvenPrefix = 0x02
	pubKeyOddPrefix  = 0x03
)

// Secp256k1PrvToAddress calculates XRPL address for ECDSA private key.
func Secp256k1PrvToAddress(prv *ecdsa.PrivateKey) string {
	accountID := Secp256k1PrvToID(prv)
	return idToAddress(accountID)
}

// Secp256k1PrvToID calculates account ID for ECDSA private key.
func Secp256k1PrvToID(prv *ecdsa.PrivateKey) []byte {
	return hash.Sha256RipeMD160(secp256k1PrvToPub(prv))
}

// Secp256k1PrvToPub return public Key for ECDSA private key in hex string.
func Secp256k1PrvToPub(prv *ecdsa.PrivateKey) string {
	return strings.ToUpper(hex.EncodeToString(secp256k1PrvToPub(prv)))
}

// Secp256k1PrivateToPub return public Key for ECDSA private key in byte slice.
func secp256k1PrvToPub(prv *ecdsa.PrivateKey) []byte {
	xrpPub := make([]byte, 0, 33)

	pub := prv.PublicKey

	if pub.Y.Bit(0) == 0 {
		xrpPub = append(xrpPub, pubKeyEvenPrefix)
	} else {
		xrpPub = append(xrpPub, pubKeyOddPrefix)
	}

	xrpPub = append(xrpPub, pub.X.Bytes()...)

	return xrpPub
}

// Ed25519PrvToAddress calculates xrpl address for Ed25519 private key.
func Ed25519PrvToAddress(prv []byte) (string, error) {
	accountID, err := Ed25519PrvToID(prv)
	if err != nil {
		return "", err
	}

	return idToAddress(accountID), nil
}

// Ed25519PrvToID calculates xrpl account ID for Ed25519 private key.
func Ed25519PrvToID(prv []byte) ([]byte, error) {
	if len(prv) != 64 {
		return nil, fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return hash.Sha256RipeMD160(pub), nil
}

// Ed25519PrvToPub calculates public key in hex string for Ed25519 private key.
func Ed25519PrvToPub(prv []byte) (string, error) {
	if len(prv) != 64 {
		return "", fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return strings.ToUpper(hex.EncodeToString(pub)), nil
}

// idToAddress calculates XRPL address from accountID.
// It does not check that id is 20 bytes long.
func idToAddress(id []byte) string {
	augmented := make([]byte, 0, len(id)+1+4)

	augmented = append(augmented, accountPrefix) // account prefix
	augmented = append(augmented, id...)

	cs := hash.Checksum(augmented) // checksum

	augmented = append(augmented, cs...)

	return base58.XRPLCoder.Encode(augmented)
}
