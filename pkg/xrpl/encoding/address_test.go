package encoding

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	priv, err := crypto.HexToECDSA("E53D66E77BE631BE5BF1583864C1E126E14276D9E3F747E3733E6809FC49D09B")
	require.NoError(t, err)

	fmt.Println(Secp256k1PrivateToAddress(priv))

	fmt.Printf("Secp256k1PrivateToPub(priv): %v\n", Secp256k1PrivateToPub(priv))
}

func TestAddressED(t *testing.T) {
	priv, err := hex.DecodeString("695CFDC7148F58705FEFD622812D8016EB2F6D14C4BDD8C8F5D4BCDDAB3D5DE2")
	require.NoError(t, err)

	rawPriv := ed25519.NewKeyFromSeed(priv)

	fmt.Println(Ed25519PrivateToAddress(rawPriv))

	fmt.Println(Ed25519PrivateToPub(rawPriv))

	// fmt.Printf("Secp256k1PrivateToPub(priv): %v\n", Ed25519PrivateToPub(priv))
}
