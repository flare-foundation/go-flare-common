package ed25519

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/utils"
)

const (
	ed25519Prefix = 0xed
)

// func SignTx(tx map[string]any, sequence uint32, prv ed25519.PrivateKey) ([]byte, error) {
// 	tx["Sequence"] = sequence

// 	pub, err := PrvToPub(prv)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid private key: %v", err)
// 	}
// 	tx["SigningPubKey"] = pub

// 	addr, err := PrvToAddress(prv)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid private key: %v", err)
// 	}
// 	tx["Account"] = addr

// 	encoded, err := types.Encode(tx, true)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot encode tx: %v", err)
// 	}

// 	msg := signing.MessageToSign(encoded, false, nil)

// 	signature := ed25519.Sign(prv, msg)

// 	return signature, nil
// }

func SignTxMultisig(tx map[string]any, prv ed25519.PrivateKey) (*utils.Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	accID, err := PrvToID(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get account id: %v", err)
	}

	msg := utils.Prepare(encoded, true, accID)
	signature := ed25519.Sign(prv, msg)

	add, err := PrvToAddress(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get address %v", err)
	}

	pub, err := PrvToPub(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get address %v", err)
	}

	return &utils.Signer{
		Account:       add,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// PrvToAddress calculates xrpl address for Ed25519 private key.
func PrvToAddress(prv []byte) (string, error) {
	accountID, err := PrvToID(prv)
	if err != nil {
		return "", fmt.Errorf("private key to address: %v", err)
	}

	add, err := address.Address(accountID)
	if err != nil {
		return "", fmt.Errorf("private key to address: %v", err)
	}

	return add, nil
}

// PrvToID calculates xrpl account ID for Ed25519 private key.
func PrvToID(prv []byte) ([]byte, error) {
	if len(prv) != 64 {
		return nil, fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return hash.Sha256RipeMD160(pub), nil
}

// PrvToPub calculates public key in hex string for Ed25519 private key.
func PrvToPub(prv []byte) (string, error) {
	if len(prv) != 64 {
		return "", fmt.Errorf("wrong length prv is %d bytes long, should be 64", len(prv))
	}

	pub := make([]byte, 0, 33)
	pub = append(pub, ed25519Prefix)
	pub = append(pub, prv[32:]...)

	return strings.ToUpper(hex.EncodeToString(pub)), nil
}

// validate checks that the sig is a valid Ed25519 signature of msg by pub.
func Validate(msg, sig []byte, pub string) (bool, error) {
	pubBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false, fmt.Errorf("cannot read pub: %v", err)
	}

	if len(pubBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid pubKey length (require 32 bytes): %s", pub)
	}

	if len(sig) != ed25519.SignatureSize {
		return false, fmt.Errorf("invalid signature length (require 64 bytes): %s", pub)
	}

	return ed25519.Verify(pubBytes, msg, sig), nil
}
