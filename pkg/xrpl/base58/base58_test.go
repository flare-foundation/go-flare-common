package base58

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct{ input []byte }{
		{
			[]byte{},
		},
		{
			[]byte("random"),
		},
		{
			[]byte{0, 0, 0},
		},
		{
			[]byte{0, 1, 0},
		},
	}

	for _, test := range tests {
		encoded := XRPLCoder.Encode(test.input)

		decoded, err := XRPLCoder.Decode(encoded)
		require.NoError(t, err)

		require.Equal(t, test.input, decoded)
	}
}

func TestDecodeEncode(t *testing.T) {
	tests := []struct {
		input string
		ok    bool
	}{
		{
			"r",
			true,
		},
		{
			"rMdpzP6SSRcNDSgHQGGx4hV3qeRiuRoqQM",
			true,
		},
		{
			"rrrr",
			true,
		},
		{
			"rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz",
			true,
		},
		{
			"l",
			false,
		},
		{
			"日本語",
			false,
		},
	}

	for _, test := range tests {
		decoded, err := XRPLCoder.Decode(test.input)
		if test.ok {
			require.NoError(t, err)
			encoded := XRPLCoder.Encode(decoded)

			require.Equal(t, test.input, encoded)
		} else {
			require.Error(t, err)
		}
	}
}

func TestNewCoderFail(t *testing.T) {
	alphabets := []string{
		"",
		"rpshnaf39wBUDNEGHJKLM4PQRST7V XYZ2bcdeCg65jkm8oFqi1tuvAxyz",      // whitespace
		"rpshnaf39wBUDNEGHJKLM4PQRST7V\u0000eXYZ2bcdeCg65jkm8oFqi1tuvAxy", // whitespace
		"rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyr",      // duplicate
		"rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyzl",     // too long
		"rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxœ",       // non standard character
	}

	for _, alphabet := range alphabets {
		_, err := NewCoder(alphabet)
		require.Error(t, err)
	}
}

// TestAlphabetMatchesRippled pins the XRPL alphabet to the value defined in
// rippled. Any divergence would silently break all base58 interop.
func TestAlphabetMatchesRippled(t *testing.T) {
	// source: rippled src/libxrpl/protocol/tokens.cpp:126 (alphabetForward)
	const rippledAlphabet = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
	require.Equal(t, rippledAlphabet, AlphabetXRPL)
}

// TestDecodePrimaryVectors checks that raw base58 decoding of canonical XRPL
// tokens yields the documented <version><payload> prefix, leaving only the
// 4-byte trailing sha256d checksum unchecked. Values come from rippled and
// xrpl.js primary-source test fixtures.
func TestDecodePrimaryVectors(t *testing.T) {
	tests := []struct {
		name       string
		encoded    string
		prefixHex  string
		payloadHex string
	}{
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:48
			name:       "accountID",
			encoded:    "rJrRMgiRgrU6hDF4pgu5DXQdWyPbY35ErN",
			prefixHex:  "00",
			payloadHex: "BA8E78626EE42C41B46D46C3048DF3A1C3C87072",
		},
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:55
			name:       "nodePublic",
			encoded:    "n9MXXueo837zYH36DvMc13BwHcqtfAWNJY5czWVbp7uYTj7x17TH",
			prefixHex:  "1C",
			payloadHex: "0388E5BA87A000CB807240DF8C848EB0B5FFA5C8E5A521BC8E105C0F0A44217828",
		},
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:62
			name:       "accountPublic",
			encoded:    "aB44YfzW24VDEJQ2UuLPV2PvqcPCSoLnL7y5M1EzhdW4LnK5xMS3",
			prefixHex:  "23",
			payloadHex: "023693F15967AE357D0327974AD46FE3C127113B1110D6044FD41E723689F81CC6",
		},
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:72
			name:       "secp256k1Seed",
			encoded:    "sn259rEFXrQrWyx3Q7XneWcwV6dfL",
			prefixHex:  "21",
			payloadHex: "CF2DE378FBDD7E2EE87D486DFB5A7BFF",
		},
		{
			// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:67
			name:       "ed25519Seed",
			encoded:    "sEdTM1uX8pu2do5XvTnutH6HsouMaM2",
			prefixHex:  "01E14B",
			payloadHex: "4C3A1D213FBDFB14C7C28D609469B341",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prefix, err := hex.DecodeString(test.prefixHex)
			require.NoError(t, err)
			payload, err := hex.DecodeString(test.payloadHex)
			require.NoError(t, err)

			got, err := XRPLCoder.Decode(test.encoded)
			require.NoError(t, err)

			// Layout: <prefix><payload><4-byte checksum>.
			require.Len(t, got, len(prefix)+len(payload)+4)
			assert.Equal(t, prefix, got[:len(prefix)])
			assert.Equal(t, payload, got[len(prefix):len(prefix)+len(payload)])

			// And the reverse direction must round-trip.
			assert.Equal(t, test.encoded, XRPLCoder.Encode(got))
		})
	}
}

// TestDecodeEncodeRippledStrings round-trips additional canonical tokens.
func TestDecodeEncodeRippledStrings(t *testing.T) {
	tests := []string{
		// source: rippled src/test/protocol/Seed_test.cpp:69
		"snMKnVku798EnBwUfxeSD8953sLYA",
		// source: rippled src/test/protocol/Seed_test.cpp:71
		"sspUXGrmjQhq6mgc24jiRuevZiwKT",
		// source: rippled src/test/protocol/Seed_test.cpp:122 (NodePublic, secp256k1 from masterpassphrase)
		"n94a1u4jAz288pZLtw6yFWVbi89YamiC6JBXPVUj5zmExe5fTVg9",
		// source: rippled src/test/protocol/Seed_test.cpp:125 (NodePrivate, secp256k1 from masterpassphrase)
		"pnen77YEeUd4fFKG7iycBWcwKpTaeFRkW2WFostaATy1DSupwXe",
		// source: rippled src/test/protocol/Seed_test.cpp:164 (NodePublic, ed25519 from masterpassphrase)
		"nHUeeJCSY2dM71oxM8Cgjouf5ekTuev2mwDpc374aLMxzDLXNmjf",
		// source: rippled src/test/protocol/Seed_test.cpp:203 (AccountID, secp256k1 from masterpassphrase)
		"rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		// source: rippled src/test/protocol/Seed_test.cpp:205 (AccountPublic, secp256k1 from masterpassphrase)
		"aBQG8RQAzjs1eTKFEAQXr2gS4utcDiEC9wmi7pfUPTi27VCahwgw",
		// source: rippled src/test/protocol/Seed_test.cpp:242 (AccountID, ed25519 from masterpassphrase)
		"rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
		// source: rippled src/test/protocol/Seed_test.cpp:245 (AccountPublic, ed25519 from masterpassphrase)
		"aKGheSBjmCsKJVuLNKRAKpZXT6wpk2FCuEZAXJupXgdAxX5THCqR",
		// source: xrpl.js packages/ripple-address-codec/test/xrp-codec.test.ts:215
		// (codec.encode(stringToBytes('123456789'), {versions:[0]}))
		"rnaC7gW34M77Kneb78s",
	}

	for _, s := range tests {
		decoded, err := XRPLCoder.Decode(s)
		require.NoError(t, err, "decode %s", s)
		assert.Equal(t, s, XRPLCoder.Encode(decoded))
	}
}

// TestDecodeRejectsRippledInvalidChars covers the characters absent from
// rippled's alphabet (0, O, I, l) that Seed_test.cpp uses to probe rejection.
func TestDecodeRejectsRippledInvalidChars(t *testing.T) {
	tests := []string{
		// source: rippled src/test/protocol/Seed_test.cpp:88 (contains 'O')
		"sspOXGrmjQhq6mgc24jiRuevZiwKT",
		// source: rippled src/test/protocol/Seed_test.cpp:89 (contains '/')
		"ssp/XGrmjQhq6mgc24jiRuevZiwKT",
	}

	for _, s := range tests {
		_, err := XRPLCoder.Decode(s)
		assert.Error(t, err, "expected decode to fail: %s", s)
	}
}
