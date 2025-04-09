package encoding_test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/require"

	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

type ed25519key struct {
	priv ed25519.PrivateKey
}

const privateKey = "76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872"

// Returns first 32 bytes of a SHA512 of the input bytes
func Sha512Half(b []byte) []byte {
	hasher := sha512.New()
	hasher.Write(b)
	return hasher.Sum(nil)[:32]
}

func NewEd25519Key(seed []byte) (*ed25519key, error) {
	r := rand.Reader
	if seed != nil {
		r = bytes.NewReader(Sha512Half(seed))
	}
	_, priv, err := ed25519.GenerateKey(r)
	if err != nil {
		return nil, err
	}
	return &ed25519key{priv: priv}, nil
}

func XrpSign(txHash []byte, privKey *ecdsa.PrivateKey) []byte {

	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, txHash)

	return sig2.Serialize()
}

func TestMakeAndSign(t *testing.T) {

	// pub := "ED9768D77765A383FEAD745F716DBBE5BF63A8CDFC6BEB41C943EFCCB746554FB4"
	// tx["SigningPubKey"] = pub

	// toSign := encoding.HashToSign(encoded, false)

	// // secret := "sEdSQL8ZUkRdeaMV2ECakSrmzcR83py"

	prvStr := "76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872"

	seedBytes, err := hex.DecodeString(prvStr)
	require.NoError(t, err)

	rawPriv := ed25519.NewKeyFromSeed(seedBytes)

	signer, err := encoding.Ed25519PrivateToAddress(rawPriv)
	require.NoError(t, err)

	fmt.Printf("signer: %v\n", signer)

	// fmt.Println(hex.DecodeString(pub))

	privSecp, err := crypto.ToECDSA(seedBytes)
	require.NoError(t, err)

	addSec := encoding.Secp256k1PrivateToAddress(privSecp)

	tx, err := encoding.SetMultisig(privSecp, 6303020, signer)
	require.NoError(t, err)

	fmt.Printf("lowercase: %v\n", tx)

	multi, err := encoding.PaymentMultisig(rawPriv, addSec, 6303021)
	require.NoError(t, err)

	fmt.Printf("lowercase: %v\n", multi)
}

func TestMake(t *testing.T) {
	prvBytes, err := hex.DecodeString(privateKey)
	require.NoError(t, err)

	privED := ed25519.NewKeyFromSeed(prvBytes)
	privSecp, err := crypto.ToECDSA(prvBytes)
	require.NoError(t, err)

	addrED, err := encoding.Ed25519PrivateToAddress(privED)
	require.NoError(t, err)

	pubED, err := encoding.Ed25519PrivateToPub(privED)
	require.NoError(t, err)

	fmt.Printf("pubED: %v\n", pubED)
	fmt.Printf("addrED: %v\n", addrED)

	addrSecp := encoding.Secp256k1PrivateToAddress(privSecp)

	pubSecp := encoding.Secp256k1PrivateToPub(privSecp)

	fmt.Printf("addrSep: %v\n", addrSecp)
	fmt.Printf("pubSecp: %v\n", pubSecp)

	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         addrED,
		"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		"Amount":          "100",
		"Fee":             "1200000",
		"Sequence":        6302691,
		"SigningPubKey":   "",
	}

	signature, err := encoding.SignTransactionSecpMultisig(tx, privSecp)
	require.NoError(t, err)

	signers := []*encoding.Signer{signature}

	toSend, err := encoding.JoinMultisig(tx, signers)
	require.NoError(t, err)

	fmt.Printf("toSend: %v\n", hex.EncodeToString(toSend))
}

func TestMake2(t *testing.T) {
	prvBytes, err := hex.DecodeString(privateKey)
	require.NoError(t, err)

	privED := ed25519.NewKeyFromSeed(prvBytes)
	privSecp, err := crypto.ToECDSA(prvBytes)
	require.NoError(t, err)

	addrED, err := encoding.Ed25519PrivateToAddress(privED)
	require.NoError(t, err)

	pubED, err := encoding.Ed25519PrivateToPub(privED)
	require.NoError(t, err)

	fmt.Printf("pubED: %v\n", pubED)
	fmt.Printf("addrED: %v\n", addrED)

	addrSecp := encoding.Secp256k1PrivateToAddress(privSecp)

	pubSecp := encoding.Secp256k1PrivateToPub(privSecp)

	fmt.Printf("addrSep: %v\n", addrSecp)
	fmt.Printf("pubSecp: %v\n", pubSecp)

	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         addrSecp,
		"Destination":     addrED,
		"Amount":          "100",
		"Fee":             "1200000",
		"Sequence":        6303019,
		"SigningPubKey":   pubSecp,
	}

	signature, err := encoding.SignTransactionSecpSingle(tx, 6303019, privSecp)
	require.NoError(t, err)

	fmt.Printf("signature: %v\n", signature)

	tx["TxnSignature"] = hex.EncodeToString(signature)

	toSend, err := types.Encode(tx, false)
	require.NoError(t, err)

	fmt.Printf("toSend: %v\n", hex.EncodeToString(toSend))
}
