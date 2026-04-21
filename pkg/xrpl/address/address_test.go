package address

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountID(t *testing.T) {
	addr := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S"

	id, err := ID(addr)
	require.NoError(t, err)

	addrBack, err := Address(id)
	require.NoError(t, err)

	require.Equal(t, addr, addrBack)
}

func TestAccountIDFail(t *testing.T) {
	failAddrLengthAlphabet := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1l"
	_, err := ID(failAddrLengthAlphabet)
	require.Error(t, err)

	failAddrLength := "rKTyuZsNuQbL"
	_, err = ID(failAddrLength)
	require.Error(t, err)

	failAddrLeading := "hhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S"
	_, err = ID(failAddrLeading)
	require.Error(t, err)

	failAddrChecksum := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1h"
	_, err = ID(failAddrChecksum)
	require.Error(t, err)

	invalidID0 := []byte{}
	invalidID1 := make([]byte, 21)

	_, err = Address(invalidID0)
	require.Error(t, err)

	_, err = Address(invalidID1)
	require.Error(t, err)
}

func TestPubToAddress(t *testing.T) {
	tests := []struct {
		pub string
		add string
	}{
		{
			pub: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
			add: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		},
		{
			pub: "ED457192C53CC901123BB8C6E75C8636979BD88D46C32F6D915F89D66CE8862635",
			add: "rMuZNV2kjCKs8v8rd8QFizAaPdvCDYTPc7",
		},
		{
			pub: "032315B578DC026DD182F2FC89723E4E43AF812D7D3C86635B80D3608FBC76C224",
			add: "rf27pmNLFaXFwQfZPbTKe2Z665e1V76oC5",
		},
		{
			pub: "02707A7AE05A8DACDB89CC93429949CDA26F68200D9CE8753D4DCB04D6F80CFCB7",
			add: "rN5N6fJbc8xyViPDeQFMQMpYfVHuxSGV2G",
		},
		{
			pub: "035DB05B1CEA82785FB8B3F7E68E5C7429A1B00BE47CC6B1A651AA12C7B8D9592C",
			add: "rJQesZZEQzW9J3Eb1X1Snc7E6YGk7kTMoK",
		},
		{
			pub: "030C141E3E131B1D25B7EA85B10283E315F77F27A2FA085A45D65D47393DB9219F",
			add: "r9cvJhquqeExszdWZSw2rrFP98fsVFLdPe",
		},
	}

	for _, test := range tests {
		addFromPub, err := PubToAddress(test.pub)
		require.NoError(t, err)

		require.Equal(t, test.add, addFromPub)
	}
}

// TestAccountIDPrimaryVectors round-trips known <address, hex-account-id> pairs
// from xrpl.js primary-source fixtures.
func TestAccountIDPrimaryVectors(t *testing.T) {
	tests := []struct {
		name    string
		address string
		idHex   string
	}{
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:48
			name:    "xrpCodecAccountID",
			address: "rJrRMgiRgrU6hDF4pgu5DXQdWyPbY35ErN",
			idHex:   "BA8E78626EE42C41B46D46C3048DF3A1C3C87072",
		},
		{
			// source: rippled src/test/protocol/Seed_test.cpp:203 (masterpassphrase, secp256k1)
			name:    "masterPassphraseSecp256k1",
			address: "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			idHex:   "B5F762798A53D543A014CAF8B297CFF8F2F937E8",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want, err := hex.DecodeString(test.idHex)
			require.NoError(t, err)

			id, err := ID(test.address)
			require.NoError(t, err)
			assert.Equal(t, want, id)

			addr, err := Address(id)
			require.NoError(t, err)
			assert.Equal(t, test.address, addr)
		})
	}
}

// TestPubToAddressRippledVectors derives canonical XRPL addresses from public
// keys emitted by xrpl.js primary-source fixtures.
func TestPubToAddressRippledVectors(t *testing.T) {
	tests := []struct {
		name    string
		pub     string
		address string
	}{
		{
			// source: xrpl.js packages/ripple-keypairs/test/fixtures/api.json (secp256k1)
			name:    "apiSecp256k1",
			pub:     "030D58EB48B4420B1F7B9DF55087E0E29FEF0E8468F9A6825B01CA2C361042D435",
			address: "rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1",
		},
		{
			// source: xrpl.js packages/ripple-keypairs/test/fixtures/api.json (ed25519)
			name:    "apiEd25519",
			pub:     "ED01FA53FA5A7E77798F882ECE20B1ABC00BB358A9E55A202D0D0676BD0CE37A63",
			address: "rLUEXYuLiQptky37CqLcm9USQpPiz5rkpD",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addr, err := PubToAddress(test.pub)
			require.NoError(t, err)
			assert.Equal(t, test.address, addr)
		})
	}
}

// TestClassicAddressValidity mirrors xrpl.js isValidClassicAddress behaviour:
// ID must succeed on a canonical address and must reject the tampered variant.
func TestClassicAddressValidity(t *testing.T) {
	// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:86
	_, err := ID("rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1")
	require.NoError(t, err)

	// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:90
	_, err = ID("rLUEXYuLiQptky37CqLcm9USQpPiz5rkpD")
	require.NoError(t, err)

	// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:94
	// Last character flipped: checksum must fail.
	_, err = ID("rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw2")
	assert.Error(t, err)
}

func TestPubToAddressFail(t *testing.T) {
	tests := []string{
		"",        // empty string
		"invalid", // not hex
		"028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25",    // too short
		"028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25BAA", // too long
		"FF8FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",   // invalid first byte
		"048FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",   // invalid first byte
	}

	for _, test := range tests {
		_, err := PubToAddress(test)
		require.Error(t, err)
	}
}
