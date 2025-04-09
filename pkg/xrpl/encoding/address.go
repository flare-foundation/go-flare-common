package encoding

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"

	//nolint
	"golang.org/x/crypto/ripemd160"
)

const (
	ed25519Prefix = 0xed
	accountPrefix = 0x00

	pubKeyEvenPrefix = 0x02
	pubKeyOddPrefix  = 0x03
)

func Secp256k1PrivateToAddress(prv *ecdsa.PrivateKey) string {
	accountID := Secp256k1PrivateToID(prv)
	return accountIDToAddress(accountID)
}

func Secp256k1PrivateToID(prv *ecdsa.PrivateKey) []byte {
	return sha256RipeMD160(secp256k1PrivateToPub(prv))
}

func Secp256k1PrivateToPub(prv *ecdsa.PrivateKey) string {
	return hex.EncodeToString(secp256k1PrivateToPub(prv))
}

func secp256k1PrivateToPub(prv *ecdsa.PrivateKey) []byte {
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

func Ed25519PrivateToAddress(prv []byte) (string, error) {
	if len(prv) != 64 {
		return "", fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	accountID := sha256RipeMD160(pub)

	return accountIDToAddress(accountID), nil
}

func Ed25519PrivateToID(prv []byte) ([]byte, error) {
	if len(prv) != 64 {
		return nil, fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	accountID := sha256RipeMD160(pub)

	return accountID, nil
}

func Ed25519PrivateToPub(prv []byte) (string, error) {
	if len(prv) != 64 {
		return "", fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return strings.ToUpper(hex.EncodeToString(pub)), nil
}

func accountIDToAddress(accountID []byte) string {
	augmented := make([]byte, 0, len(accountID)+1+4)

	augmented = append(augmented, accountPrefix)
	augmented = append(augmented, accountID...)

	cs := checksum(augmented)

	augmented = append(augmented, cs...)

	return base58.XRPLCoder.Encode(augmented)
}

func sha256RipeMD160(b []byte) []byte {
	ripe := ripemd160.New()
	sha := sha256.New()
	sha.Write(b)
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}

func checksum(input []byte) []byte {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	return h2[:4]
}
