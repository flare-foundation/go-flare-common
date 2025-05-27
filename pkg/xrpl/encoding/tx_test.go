package encoding_test

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	ed "github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/ed25519"
)

type ed25519key struct {
	priv ed25519.PrivateKey
}

const privateKey = "965E359115DD65D08FD9CF020969BBE0D2A5D263F4211C46AC4A21B77CE720AD"
const privateKey2 = "111B4877EB4ACBD4943E7CCE0908DB89A26BDB77ED0F0DEB029235FEF54B7870"

// Returns first 32 bytes of a SHA512 of the input bytes.
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

func TestMake(t *testing.T) {
	prvBytes, err := hex.DecodeString(privateKey)
	require.NoError(t, err)

	privED := ed25519.NewKeyFromSeed(prvBytes)
	require.NoError(t, err)

	addrED, err := ed.PrvToAddress(privED)
	require.NoError(t, err)

	pubED, err := ed.PrvToPub(privED)
	require.NoError(t, err)

	fmt.Printf("pubED: %v\n", pubED)
	fmt.Printf("addrED: %v\n", addrED)

	tx := map[string]any{
		"TransactionType": "AccountSet",
		"Account":         addrED,
		"Fee":             "120000",
		"Sequence":        7168324,
		"SigningPubKey":   pubED,
		"SetFlag":         13,
		"Flags":           1048576 | 65536,
	}

	signature, err := ed.SignTx(tx, 7168331, privED)
	require.NoError(t, err)

	tx["TxnSignature"] = hex.EncodeToString(signature)

	txBlob, err := types.Encode(tx, false)
	require.NoError(t, err)

	fmt.Printf("toSend: %v\n", hex.EncodeToString(txBlob))
}

func TestMake2(t *testing.T) {
	prvBytes, err := hex.DecodeString(privateKey2)
	require.NoError(t, err)

	privED := ed25519.NewKeyFromSeed(prvBytes)
	require.NoError(t, err)

	addrED, err := ed.PrvToAddress(privED)
	require.NoError(t, err)

	pubED, err := ed.PrvToPub(privED)
	require.NoError(t, err)

	fmt.Printf("pubED: %v\n", pubED)
	fmt.Printf("addrED: %v\n", addrED)

	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         addrED,
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "1200000",
		"Amount":          "1",
		"Sequence":        7168498,
		"SigningPubKey":   pubED,
		"DestinationTag":  3,
	}

	signature, err := ed.SignTx(tx, 7168498, privED)
	require.NoError(t, err)

	tx["TxnSignature"] = hex.EncodeToString(signature)

	txBlob, err := types.Encode(tx, false)
	require.NoError(t, err)

	fmt.Printf("toSend: %v\n", hex.EncodeToString(txBlob))
}

// func TestMake2(t *testing.T) {
// 	prvBytes, err := hex.DecodeString(privateKey)
// 	require.NoError(t, err)

// 	privED := ed25519.NewKeyFromSeed(prvBytes)
// 	privSecp, err := crypto.ToECDSA(prvBytes)
// 	require.NoError(t, err)

// 	addrED, err := address.Ed25519PrvToAddress(privED)
// 	require.NoError(t, err)

// 	pubED, err := address.Ed25519PrvToPub(privED)
// 	require.NoError(t, err)

// 	fmt.Printf("pubED: %v\n", pubED)
// 	fmt.Printf("addrED: %v\n", addrED)

// 	addrSecp := address.Secp256k1PrvToAddress(privSecp)

// 	pubSecp := address.Secp256k1PrvToPub(privSecp)

// 	fmt.Printf("addrSep: %v\n", addrSecp)
// 	fmt.Printf("pubSecp: %v\n", pubSecp)

// 	tx := map[string]any{
// 		"TransactionType": "Payment",
// 		"Account":         addrSecp,
// 		"Destination":     addrED,
// 		"Amount":          "100",
// 		"Fee":             "1200000",
// 		"Sequence":        6303019,
// 		"SigningPubKey":   pubSecp,
// 	}

// 	signature, err := encoding.SignTransactionSecp256k1Single(tx, privSecp)
// 	require.NoError(t, err)

// 	fmt.Printf("signature: %v\n", signature)

// 	tx["TxnSignature"] = hex.EncodeToString(signature)

// 	toSend, err := types.Encode(tx, false)
// 	require.NoError(t, err)

// 	fmt.Printf("toSend: %v\n", hex.EncodeToString(toSend))
// }
