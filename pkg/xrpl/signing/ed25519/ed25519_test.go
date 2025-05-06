package ed25519

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressED(t *testing.T) {
	priv, err := hex.DecodeString("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	rawPriv := ed25519.NewKeyFromSeed(priv)

	addr, err := PrvToAddress(rawPriv)
	require.NoError(t, err)

	require.Equal(t, "rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8", addr)

	pub, err := PrvToPub(rawPriv)
	require.NoError(t, err)

	require.Equal(t, "ED9AA5EBF4FB942A7D81545C42ADDB86B389E56FBAF3F1295A5B6E1CDBCE2424BB", pub)
}
