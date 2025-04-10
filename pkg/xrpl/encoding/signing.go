package encoding

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"

	"github.com/btcsuite/btcd/btcec/v2"
	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

type Signer struct {
	Account       string
	TxnSignature  string
	SigningPubKey string
	value         *big.Int
}

// Value of a signer is the number represented bt the account ID.
//
// At the first call the value is computed, stored, and returned. At subsequent calls, the stored value is returned.
func (s *Signer) Value() (*big.Int, error) {
	if s.value != nil {
		return s.value, nil
	}

	value, err := base58.XRPLCoder.Decode(s.Account)
	if err != nil {
		return nil, err
	}

	s.value = new(big.Int).SetBytes(value)

	return s.value, nil
}

// Format returns an array object that is ready to be serialized.
func (s *Signer) Format() types.ArrayObject {
	return types.ArrayObject{
		"Signer": types.Object{
			"Account":       s.Account,
			"TxnSignature":  s.TxnSignature,
			"SigningPubKey": s.SigningPubKey,
		},
	}
}

func SignTransactionEDSingle(tx map[string]any, sequence uint32, prv ed25519.PrivateKey) ([]byte, error) {
	tx["Sequence"] = sequence

	pub, err := Ed25519PrvToPub(prv)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}
	tx["SigningPubKey"] = pub

	addr, err := Ed25519PrvToAddress(prv)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}
	tx["Account"] = addr

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	msg := MessageToSign(encoded, false, nil)

	signature := ed25519.Sign(prv, msg)

	return signature, nil
}

func SignTransactionEDMultisig(tx map[string]any, prv ed25519.PrivateKey) (*Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	accID, err := Ed25519PrvToID(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get account id: %v", err)
	}

	msg := MessageToSign(encoded, true, accID)
	signature := ed25519.Sign(prv, msg)

	add, err := Ed25519PrvToAddress(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get address %v", err)
	}

	pub, err := Ed25519PrvToPub(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get address %v", err)
	}

	return &Signer{
		Account:       add,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

func SignTransactionSecpSingle(tx map[string]any, sequence uint32, prv *ecdsa.PrivateKey) ([]byte, error) {
	tx["Sequence"] = float64(sequence)

	pub := Secp256k1PrvToPub(prv)
	tx["SigningPubKey"] = pub

	addr := Secp256k1PrvToAddress(prv)
	tx["Account"] = addr

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	id := Secp256k1PrvToID(prv)

	msg := MessageToSign(encoded, false, id)
	signature := Secp256k1Sign(msg, prv)

	tx["TxnSignature"] = hex.EncodeToString(signature)

	signed, err := types.Encode(tx, false)
	if err != nil {
		return nil, fmt.Errorf("cannot encode signed tx: %v", err)
	}

	return signed, nil
}

func SignTransactionSecpMultisig(tx map[string]any, prv *ecdsa.PrivateKey) (*Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	accID := Secp256k1PrvToID(prv)

	msg := MessageToSign(encoded, true, accID)
	signature := Secp256k1Sign(msg, prv)

	pub := Secp256k1PrvToPub(prv)
	addr := Secp256k1PrvToAddress(prv)

	return &Signer{
		Account:       addr,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// Secp256k1Sign computes Secp256k1 signature of the message and returns in DER format.
func Secp256k1Sign(message []byte, privKey *ecdsa.PrivateKey) []byte {
	hash := hash.Sha512Half(message)
	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, hash)

	return sig2.Serialize()
}

// JoinMultisig appends signers to transactions and serializes it.
//
// If any of the signers is invalid, error is returned.
func JoinMultisig(tx map[string]any, signers []*Signer) ([]byte, error) {
	slices.SortFunc(signers, compare)

	signersArray := make([]any, len(signers))

	for j, signer := range signers {
		ok, err := ValidateMultiSig(tx, signer)

		if err != nil {
			return nil, fmt.Errorf("invalid signer %v", signer)
		} else if !ok {
			return nil, fmt.Errorf("invalid signature  %v", signer)
		}

		signersArray[j] = signer.Format()
	}

	tx["Signers"] = signersArray

	return types.Encode(tx, false)
}

// ValidateMultiSig checks that signer is valid for the transaction tx.
func ValidateMultiSig(tx map[string]any, signer *Signer) (bool, error) {
	id, err := types.AccountID.ToBytes(signer.Account, false)
	if err != nil {
		return false, fmt.Errorf("cannot get accountID: %v", err)
	}

	txBlob, err := types.Encode(tx, true)
	if err != nil {
		return false, fmt.Errorf("cannot encode tx: %v", err)
	}

	msg := MessageToSign(txBlob, true, id)

	pubPrefix := signer.SigningPubKey[0:2]

	sigBytes, err := hex.DecodeString(signer.TxnSignature)
	if err != nil {
		return false, fmt.Errorf("cannot decode signature: %v", err)
	}

	switch pubPrefix {
	case "ED", "ed":
		return validateEd25519(msg, sigBytes, signer.SigningPubKey[2:])
	case "02", "03":
		return validateSecp256k1(msg, sigBytes, signer.SigningPubKey)
	default:
		return false, fmt.Errorf("invalid public key %s", signer.SigningPubKey)
	}
}

// validateSecp256k1 checks that the sig is a valid Secp256k1 signature of msg by pub.
func validateSecp256k1(msg, sig []byte, pub string) (bool, error) {
	sigParse, err := btcecdsa.ParseDERSignature(sig)
	if err != nil {
		return false, err
	}

	pubBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false, err
	}

	pubPar, err := btcec.ParsePubKey(pubBytes)
	if err != nil {
		return false, err
	}

	ok := sigParse.Verify(hash.Sha512Half(msg), pubPar)

	return ok, nil
}

// validateEd25519 checks that the sig is a valid Ed25519 signature of msg by pub.
func validateEd25519(msg, sig []byte, pub string) (bool, error) {
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

// compare compares values signers
//   - -1 if s1 is lesser than s2
//   - 0 if s1 is equal to s2
//   - 1 if s1 is greater than s2
func compare(s1, s2 *Signer) int {
	if s1 != nil || s2 != nil {
		return 0
	}

	val1, err := s1.Value()
	if err != nil {
		return 0
	}

	val2, err := s2.Value()
	if err != nil {
		return 0
	}

	return val1.Cmp(val2)
}
