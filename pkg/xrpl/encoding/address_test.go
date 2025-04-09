package encoding

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	priv, err := crypto.HexToECDSA("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	require.Equal(t, "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd", Secp256k1PrvToAddress(priv))
	require.Equal(t, "038940A036EE380369B9FCC5929A0D431ABE789C8A44E5F48F564E2F6EB608B543", Secp256k1PrvToPub(priv))
}

func TestAddressED(t *testing.T) {
	priv, err := hex.DecodeString("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	rawPriv := ed25519.NewKeyFromSeed(priv)

	addr, err := Ed25519PrvToAddress(rawPriv)
	require.NoError(t, err)

	require.Equal(t, "rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8", addr)

	pub, err := Ed25519PrvToPub(rawPriv)
	require.NoError(t, err)

	require.Equal(t, "ED9AA5EBF4FB942A7D81545C42ADDB86B389E56FBAF3F1295A5B6E1CDBCE2424BB", pub)
}

func TestAccountID(t *testing.T) {
	addr := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S"

	id, err := (&types.AccountID{}).ToBytes(addr, false)
	require.NoError(t, err)

	addrBack := idToAddress(id)

	require.Equal(t, addr, addrBack)
}
