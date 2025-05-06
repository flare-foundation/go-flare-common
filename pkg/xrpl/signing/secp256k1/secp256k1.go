package secp256k1

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
)

const (
	pubKeyEvenPrefix = 0x02
	pubKeyOddPrefix  = 0x03
)

// SignTx signs a tra
func SignTx(tx map[string]any, prv *ecdsa.PrivateKey) ([]byte, error) {
	pub := PrvToPub(prv)
	tx["SigningPubKey"] = pub

	addr := PrvToAddress(prv)
	tx["Account"] = addr

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	id := PrvToID(prv)

	msg := signing.MessageToSign(encoded, false, id)
	signature := Sign(msg, prv)

	tx["TxnSignature"] = hex.EncodeToString(signature)

	signed, err := types.Encode(tx, false)
	if err != nil {
		return nil, fmt.Errorf("cannot encode signed tx: %v", err)
	}

	return signed, nil
}

func SignTxMultisig(tx map[string]any, prv *ecdsa.PrivateKey) (*signing.Signer, error) {
	tx["SigningPubKey"] = ""

	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("cannot encode tx: %v", err)
	}

	accID := PrvToID(prv)

	msg := signing.MessageToSign(encoded, true, accID)
	signature := Sign(msg, prv)

	pub := PrvToPub(prv)
	addr := PrvToAddress(prv)

	return &signing.Signer{
		Account:       addr,
		TxnSignature:  hex.EncodeToString(signature),
		SigningPubKey: pub,
	}, nil
}

// Sign computes Secp256k1 signature of the message and returns in DER format.
func Sign(message []byte, privKey *ecdsa.PrivateKey) []byte {
	hash := hash.Sha512Half(message)
	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, hash)

	return sig2.Serialize()
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

// PrvToPub return public Key for ECDSA private key in hex string.
func PrvToPub(prv *ecdsa.PrivateKey) string {
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

func SECToDER(sig []byte) []byte {
	if sig[len(sig)-1] != 0 && sig[len(sig)-1] != 1 {
		fmt.Printf("sig[len(sig)-1]: %v\n", sig[len(sig)-1])
	}

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
