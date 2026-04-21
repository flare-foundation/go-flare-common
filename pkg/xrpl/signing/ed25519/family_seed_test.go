package ed25519

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

// source: rippled src/test/protocol/SecretKey_test.cpp ed25519TestVectors[0..2] (lines 929-947).
func TestPrivKeyFromSeedRippledVectors(t *testing.T) {
	tests := []struct {
		seedHex string
		pubHex  string
		secHex  string
		addr    string
	}{
		{
			seedHex: "AF41FF66F75EBD3A6B18FB7A1DF61C97",
			pubHex:  "ED48CBBBE0EE7B8686A7DE9F0A0159734E65F9C369947F2E2696232B461E553213",
			secHex:  "1A1097FCD9CE4E1DA24666B698879766E1757547D1D4E364B64355F7C84BA0F3",
			addr:    "rVAEQBhWT6nZ4woEifdN3TMMdUZaxeXnR",
		},
		{
			seedHex: "140C1D081319339C799DC6A165951BE1",
			pubHex:  "ED3BC82EF45F8909CC00F8B7AAF05931681411758C117124875066C28398FE156D",
			secHex:  "FE3E5A82B80DD82E915F7638942A332CE3068879740C7E90E220A4FB0B37CEC8",
			addr:    "rK57dJ9533WtoY8NNwVWGY7ffuAc8WCcPE",
		},
		{
			seedHex: "865323B6E45AD6EEFA279FA584852ED4",
			pubHex:  "EDE83422EBE075707312661E4B033A29BA86386230500CF28CF165F3C21E906D00",
			secHex:  "A3C4E243A4644E738A247A59AABB5C89E4090D1B7302F245826487C838DA6989",
			addr:    "rfZiEDieHSHsQJ1UNfv2jYDuQawdRSBFwz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			seed, err := hex.DecodeString(tt.seedHex)
			require.NoError(t, err)

			priv, err := PrivKeyFromSeed(seed)
			require.NoError(t, err)

			gotSec := strings.ToUpper(hex.EncodeToString(priv[:32]))
			assert.Equal(t, tt.secHex, gotSec, "private scalar (sha512Half(seed))")

			gotPub, err := PrvToPub(priv)
			require.NoError(t, err)
			assert.Equal(t, tt.pubHex, gotPub, "compressed pubkey")

			gotAddr, err := PrvToAddress(priv)
			require.NoError(t, err)
			assert.Equal(t, tt.addr, gotAddr, "account address")

			viaPub, err := address.PubToAddress(tt.pubHex)
			require.NoError(t, err)
			assert.Equal(t, tt.addr, viaPub)
		})
	}
}

// source: rippled src/test/protocol/Seed_test.cpp:242-243.
// generateKeyPair(ed25519, generateSeed("masterpassphrase")) → rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf.
func TestPrivKeyFromFamilySeedMasterpassphrase(t *testing.T) {
	priv, err := PrivKeyFromFamilySeed("snoPBrXtMeMyMHUVTgbuqAfg1SUTb")
	require.NoError(t, err)
	gotAddr, err := PrvToAddress(priv)
	require.NoError(t, err)
	assert.Equal(t, "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", gotAddr)
}

func TestPrivKeyFromSeedRejectsBadLength(t *testing.T) {
	_, err := PrivKeyFromSeed(nil)
	assert.Error(t, err)
	_, err = PrivKeyFromSeed(make([]byte, 15))
	assert.Error(t, err)
	_, err = PrivKeyFromSeed(make([]byte, 17))
	assert.Error(t, err)
}
