package seed

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// source: rippled src/test/protocol/SecretKey_test.cpp secp256k1TestVectors[0] (line 357)
// raw 16-byte seed, encoded then decoded to confirm the round-trip.
func TestDecodeFamilySeedRoundTrip(t *testing.T) {
	raw, err := hex.DecodeString("DEDCE9CE67B451D852FD4E846FCDE31C")
	require.NoError(t, err)

	encoded, err := EncodeFamilySeed(raw)
	require.NoError(t, err)

	got, err := DecodeFamilySeed(encoded)
	require.NoError(t, err)
	assert.Equal(t, raw, got)
}

// source: rippled src/libxrpl/protocol/Seed.cpp parseBase58<Seed> + masterpassphrase keypair
// generation (Seed_test.cpp testKeypairGenerationAndSigning line 203). The canonical seed string
// for generateSeed("masterpassphrase") is "snoPBrXtMeMyMHUVTgbuqAfg1SUTb".
func TestDecodeFamilySeedMasterpassphrase(t *testing.T) {
	const encoded = "snoPBrXtMeMyMHUVTgbuqAfg1SUTb"
	const wantHex = "dedce9ce67b451d852fd4e846fcde31c"

	got, err := DecodeFamilySeed(encoded)
	require.NoError(t, err)
	assert.Equal(t, wantHex, hex.EncodeToString(got))
}

// TestDecodeFamilySeedErrorDoesNotLeakInput covers audit finding F-SIGN-1:
// base58.Decode's error embeds its raw input, and a wrapped %w propagated
// that input — which is private-key-equivalent — into caller logs. The
// returned error must not contain the seed string in any form.
func TestDecodeFamilySeedErrorDoesNotLeakInput(t *testing.T) {
	// A clearly invalid base58 string with a distinctive sentinel substring.
	const sentinel = "LEAK_SENTINEL_!!!"
	_, err := DecodeFamilySeed(sentinel)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), sentinel)
}

func TestDecodeFamilySeedRejectsBadInputs(t *testing.T) {
	raw, err := hex.DecodeString("DEDCE9CE67B451D852FD4E846FCDE31C")
	require.NoError(t, err)
	encoded, err := EncodeFamilySeed(raw)
	require.NoError(t, err)

	// flip a body byte → checksum mismatch
	bad := []byte(encoded)
	bad[len(bad)/2]++
	_, err = DecodeFamilySeed(string(bad))
	assert.Error(t, err)

	// empty
	_, err = DecodeFamilySeed("")
	assert.Error(t, err)

	// oversized
	_, err = DecodeFamilySeed("sssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
	assert.Error(t, err)
}
