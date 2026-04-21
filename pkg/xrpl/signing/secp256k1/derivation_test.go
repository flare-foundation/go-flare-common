package secp256k1

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

// source: rippled src/test/protocol/SecretKey_test.cpp secp256k1TestVectors[0..2] (lines 356-374).
func TestPrivKeyFromSeedRippledVectors(t *testing.T) {
	tests := []struct {
		seedHex string
		pubHex  string
		secHex  string
		addr    string
	}{
		{
			seedHex: "DEDCE9CE67B451D852FD4E846FCDE31C",
			pubHex:  "0330E7FC9D56BB25D6893BA3F317AE5BCF33B3291BD63DB32654A313222F7FD020",
			secHex:  "1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7",
			addr:    "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		},
		{
			seedHex: "F75C48FEC46D4D64928B795F3FBABBA0",
			pubHex:  "03AF53E8011E85B36664F1710890501C3E86FC2C6658C2EE83CA580DC9972541B1",
			secHex:  "5B8AB0E7CDAF48874D5D9934BF3E7B2CB06BC4C7EAAAF762682ED8D0A31E3C70",
			addr:    "r9ZERztesFu3ZBs7zsWTeCvBg14GQ9zWF7",
		},
		{
			seedHex: "7AEA88C148A8C4A89069F98A3733167B",
			pubHex:  "037AD930B31B0F81036D65A62D4E1131EF605F738E7D7F95465BBECBCBFFA7076E",
			secHex:  "3DDC5FE16378AE8503E974233017791CC07C1077653AB723C8F65B17B9980DA1",
			addr:    "rfmZ1iAuSzB2gxjoJgBub6iJPCk92D7YfC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			seed, err := hex.DecodeString(tt.seedHex)
			require.NoError(t, err)

			priv, err := PrivKeyFromSeed(seed)
			require.NoError(t, err)

			gotSec := strings.ToUpper(hex.EncodeToString(paddedScalar(priv.D, 32)))
			assert.Equal(t, tt.secHex, gotSec, "private scalar")

			assert.Equal(t, tt.pubHex, PrvToPub(priv), "compressed pubkey")
			assert.Equal(t, tt.addr, PrvToAddress(priv), "account address")

			// Address via the public-key path must agree.
			viaPub, err := address.PubToAddress(tt.pubHex)
			require.NoError(t, err)
			assert.Equal(t, tt.addr, viaPub)
		})
	}
}

// source: rippled src/test/protocol/Seed_test.cpp testPassphrase masterpassphrase line 201-210.
// generateKeyPair(secp256k1, generateSeed("masterpassphrase")) → rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh.
func TestPrivKeyFromFamilySeedMasterpassphrase(t *testing.T) {
	priv, err := PrivKeyFromFamilySeed("snoPBrXtMeMyMHUVTgbuqAfg1SUTb")
	require.NoError(t, err)
	assert.Equal(t, "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", PrvToAddress(priv))
}

func TestPrivKeyFromSeedRejectsBadLength(t *testing.T) {
	_, err := PrivKeyFromSeed(nil)
	assert.Error(t, err)
	_, err = PrivKeyFromSeed(make([]byte, 15))
	assert.Error(t, err)
	_, err = PrivKeyFromSeed(make([]byte, 17))
	assert.Error(t, err)
}

func paddedScalar(d interface{ Bytes() []byte }, n int) []byte {
	b := d.Bytes()
	if len(b) >= n {
		return b
	}
	out := make([]byte, n)
	copy(out[n-len(b):], b)
	return out
}
