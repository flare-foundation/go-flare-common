// Package secp256k1 provides XRPL secp256k1 key management and transaction signing.
package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/utils"
)

const (
	pubKeyEvenPrefix = 0x02
	pubKeyOddPrefix  = 0x03
)

// // SignTx signs and encodes a transaction with prv as a master key of the account.
// func SignTx(tx map[string]any, prv *ecdsa.PrivateKey) ([]byte, error) {
// 	pub := PrvToPub(prv)
// 	tx["SigningPubKey"] = pub

// 	addr := PrvToAddress(prv)
// 	tx["Account"] = addr

// 	encoded, err := types.Encode(tx, true)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot encode tx: %v", err)
// 	}

// 	id := PrvToID(prv)

// 	msg := signing.MessageToSign(encoded, false, id)
// 	signature, err := Sign(msg, prv)
// 	if err != nil {
// 		return nil, fmt.Errorf("signing %v,", err)
// 	}

// 	tx["TxnSignature"] = hex.EncodeToString(signature)

// 	signed, err := types.Encode(tx, false)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot encode signed tx: %v", err)
// 	}

// 	return signed, nil
// }

// SignTxMultisig signs a transaction for multi-signing using a secp256k1 private key.
func SignTxMultisig(tx map[string]any, prv *ecdsa.PrivateKey) (*signer.Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := encoding.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %w", err)
	}

	accID := PrvToID(prv)

	msg, err := utils.Prepare(encoded, true, accID)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare message: %w", err)
	}

	signature, err := SignXRPL(msg, prv)
	if err != nil {
		return nil, fmt.Errorf("signing %w,", err)
	}

	pub := PrvToPub(prv)
	addr := PrvToAddress(prv)

	return &signer.Signer{
		Account:       addr,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// SignXRPL computes Secp256k1 signature of the message and returns it in DER format
// as needed for signing of an XRPL transaction.
func SignXRPL(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	h := hash.Sha512Half(message)

	sig, err := sign(h, privKey)
	if err != nil {
		return nil, fmt.Errorf("signing: %w", err)
	}

	return sig.DER(), nil
}

// sign computes Secp256k1 signature of the hash.
func sign(hash []byte, privKey *ecdsa.PrivateKey) (*SignatureWithRecovery, error) {
	sig, err := crypto.Sign(hash, privKey)
	if err != nil {
		return nil, fmt.Errorf("crypto sign: %w", err)
	}
	return MarshalRecID(sig)
}

// PrvToAddress calculates XRPL address for ECDSA private key.
func PrvToAddress(prv *ecdsa.PrivateKey) string {
	accountID := PrvToID(prv)
	return address.IDToAddress(accountID)
}

// PrvToID calculates account ID for ECDSA private key.
func PrvToID(prv *ecdsa.PrivateKey) []byte {
	return hash.Sha256RipeMD160(secp256k1PrvToPub(prv))
}

// PrvToPub returns public Key for ECDSA private key in hex string.
func PrvToPub(prv *ecdsa.PrivateKey) string {
	return strings.ToUpper(hex.EncodeToString(secp256k1PrvToPub(prv)))
}

// secp256k1PrivateToPub returns compressed public Key for ECDSA private key in byte slice.
func secp256k1PrvToPub(prv *ecdsa.PrivateKey) []byte {
	return toBytesCompressed(&prv.PublicKey)
}
