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
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

type Signer struct {
	Account       string
	TxnSignature  string
	SigningPubKey string
	value         *big.Int
}

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

func (s *Signer) Format() types.ArrayObject {
	return map[string]map[string]any{
		"Signer": {
			"Account":       s.Account,
			"TxnSignature":  s.TxnSignature,
			"SigningPubKey": s.SigningPubKey,
		},
	}
}

func SignTransactionEDSingle(tx map[string]any, sequence uint32, prv ed25519.PrivateKey) ([]byte, error) {
	tx["Sequence"] = sequence

	pub, err := Ed25519PrivateToPub(prv)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}
	tx["SigningPubKey"] = pub

	addr, err := Ed25519PrivateToAddress(prv)
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

	accID, err := Ed25519PrivateToID(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get account id: %v", err)
	}

	msg := MessageToSign(encoded, true, accID)
	signature := ed25519.Sign(prv, msg)

	add, err := Ed25519PrivateToAddress(prv)
	if err != nil {
		return nil, fmt.Errorf("cannot get address %v", err)
	}

	pub, err := Ed25519PrivateToPub(prv)
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

	pub := Secp256k1PrivateToPub(prv)
	tx["SigningPubKey"] = pub

	addr := Secp256k1PrivateToAddress(prv)
	tx["Account"] = addr

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	id := Secp256k1PrivateToID(prv)

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

	accID := Secp256k1PrivateToID(prv)

	msg := MessageToSign(encoded, true, accID)
	signature := Secp256k1Sign(msg, prv)

	pub := Secp256k1PrivateToPub(prv)
	addr := Secp256k1PrivateToAddress(prv)

	return &Signer{
		Account:       addr,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

func Secp256k1Sign(message []byte, privKey *ecdsa.PrivateKey) []byte {
	hash := Sha512Half(message)
	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, hash)

	return sig2.Serialize()
}

func JoinMultisig(tx map[string]any, signers []*Signer) ([]byte, error) {
	slices.SortFunc(signers, compare)

	signersArray := make([]map[string]map[string]any, len(signers))

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

func ValidateMultiSig(tx map[string]any, signer *Signer) (bool, error) {
	accountID, err := (&types.AccountID{}).ToBytes(signer.Account, false)
	if err != nil {
		return false, fmt.Errorf("cannot get accountID: %v", err)
	}

	fmt.Printf("accountID: %v\n", accountID)

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return false, fmt.Errorf("cannot encode tx: %v", err)
	}

	signed := MessageToSign(encoded, true, accountID)

	pubPrefix := signer.SigningPubKey[0:2]

	signatureBytes, err := hex.DecodeString(signer.TxnSignature)
	if err != nil {
		return false, fmt.Errorf("cannot decode signature: %v", err)
	}

	fmt.Printf("pubPrefix: %v\n", pubPrefix)

	switch pubPrefix {
	case "ED", "ed":
		return validateEd25519(signed, signatureBytes, signer.SigningPubKey[2:])
	case "02", "03":
		return validateSecp256k1(signed, signatureBytes, signer.SigningPubKey[:])
	default:
		return false, fmt.Errorf("invalid public key %s", signer.SigningPubKey)
	}
}

func validateSecp256k1(msg []byte, sig []byte, pub string) (bool, error) {
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

	ok := sigParse.Verify(Sha512Half(msg), pubPar)

	return ok, nil
}

func validateEd25519(signature, msg []byte, pub string) (bool, error) {
	pubBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false, fmt.Errorf("cannot read pub: %v", err)
	}

	if len(pubBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid pubKey length (require 32 bytes): %s", pub)
	}

	if len(signature) != ed25519.SignatureSize {
		return false, fmt.Errorf("invalid signature length (require 64 bytes): %s", pub)
	}

	return ed25519.Verify(pubBytes, msg, signature), nil
}

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
