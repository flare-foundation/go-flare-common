package secp256k1

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
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

// SignTxMultisig returns a signature of a transaction needed for multisig.
func SignTxMultisig(tx map[string]any, prv *ecdsa.PrivateKey) (*utils.Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	accID := PrvToID(prv)

	msg := utils.Prepare(encoded, true, accID)
	signature, err := Sign(msg, prv)
	if err != nil {
		return nil, fmt.Errorf("signing %v,", err)
	}

	pub := PrvToPub(prv)
	addr := PrvToAddress(prv)

	return &utils.Signer{
		Account:       addr,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// Sign computes Secp256k1 signature of the message and returns it in DER format.
func Sign(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	h := hash.Sha512Half(message)

	sig, err := crypto.Sign(h, privKey)
	if err != nil {
		return nil, err
	}
	return secToDER(sig), nil
}

// Sign computes Secp256k1 signature of the message and returns it in DER format.
func sign(message []byte, privKey *ecdsa.PrivateKey) (*SignatureWithRecovery, error) {
	h := hash.Sha512Half(message)

	sig, err := crypto.Sign(h, privKey)
	if err != nil {
		return nil, err
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

// secToDER converts SEC (Standards for Efficient Cryptography) encoded signature r||s(||v) to DER (Distinguished Encoding Rules) encoded signature.
func secToDER(sig []byte) []byte {
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])

	out := bytes.NewBuffer(nil)

	rb := r.Bytes()
	sb := s.Bytes()

	out.WriteByte(0x30)

	out.WriteByte(0) // to be amended at the end

	out.WriteByte(0x02)

	if rb[0] > 0x7f {
		out.WriteByte(1 + uint8(len(rb)))
		out.WriteByte(0)
	} else {
		out.WriteByte(uint8(len(rb)))
	}
	out.Write(rb)

	out.WriteByte(0x02)
	if sb[0] > 0x7f {
		out.WriteByte(1 + uint8(len(sb)))
		out.WriteByte(0)
	} else {
		out.WriteByte(uint8(len(sb)))
	}
	out.Write(sb)

	o := out.Bytes()

	o[1] = byte(out.Len() - 2)

	return o
}

// DERToSEC converts DER (Distinguished Encoding Rules) encoded signature to  SEC (Standards for Efficient Cryptography) encoded signature r||s.
func DERToSEC(sig []byte) []byte {
	rLen := int(sig[3])

	rStart := 4
	rEnd := rStart + rLen

	if sig[rStart] == 0 {
		rStart++
		rLen--
	}

	out := make([]byte, 64)

	copy(out[32-rLen:32], sig[rStart:rEnd])

	sLen := int(sig[rEnd+1])
	sStart := rEnd + 2
	sEnd := sStart + sLen

	if sig[sStart] == 0 {
		sStart++
		sLen--
	}

	// parity := sig[sEnd-1] % 2

	copy(out[64-sLen:64], sig[sStart:sEnd])

	return out
}
