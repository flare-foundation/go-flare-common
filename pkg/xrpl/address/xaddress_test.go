package address

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
)

func tagPtr(v uint32) *uint32 { return &v }

// Golden vectors generated with the reference implementation ripple-address-codec.
var xVectors = []struct {
	name    string
	classic string
	tag     *uint32
	test    bool
	x       string
}{
	{"no tag", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", nil, false, "XVLhHMPHU98es4dbozjVtdWzVrDjtV5fdx1mHp98tDMoQXb"},
	{"tag 0", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", tagPtr(0), false, "XVLhHMPHU98es4dbozjVtdWzVrDjtV8AqEL4xcZj5whKbmc"},
	{"tag 13371337", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", tagPtr(13371337), false, "XVLhHMPHU98es4dbozjVtdWzVrDjtVijiy6dVhdE4mRdMDU"},
	{"tag max", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", tagPtr(4294967295), false, "XVLhHMPHU98es4dbozjVtdWzVrDjtV18pX8yuPT7y4xaEHi"},
	{"testnet no tag", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", nil, true, "TVE26TYGhfLC7tQDno7G8dGtxSkYQn49b3qD26PK7FcGSKE"},
	{"other account", "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", tagPtr(42), false, "XVPcpSm47b1CZkf5AkKM9a84dQHe3mTCLZc5ZAoh11sd5nY"},
}

func TestClassicToX(t *testing.T) {
	for _, tc := range xVectors {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ClassicToX(tc.classic, tc.tag, tc.test)
			require.NoError(t, err)
			assert.Equal(t, tc.x, got)
		})
	}
}

func TestXToClassic(t *testing.T) {
	for _, tc := range xVectors {
		t.Run(tc.name, func(t *testing.T) {
			classic, tag, test, err := XToClassic(tc.x)
			require.NoError(t, err)
			assert.Equal(t, tc.classic, classic)
			assert.Equal(t, tc.test, test)
			if tc.tag == nil {
				assert.Nil(t, tag)
			} else {
				require.NotNil(t, tag)
				assert.Equal(t, *tc.tag, *tag)
			}
		})
	}
}

func TestXAddressRoundTrip(t *testing.T) {
	id, err := ID("rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf")
	require.NoError(t, err)

	for _, tc := range xVectors {
		t.Run(tc.name, func(t *testing.T) {
			x, err := XAddress(id, tc.tag, tc.test)
			require.NoError(t, err)

			gotID, gotTag, gotTest, err := XID(x)
			require.NoError(t, err)
			assert.Equal(t, id, gotID)
			assert.Equal(t, tc.test, gotTest)
			if tc.tag == nil {
				assert.Nil(t, gotTag)
			} else {
				require.NotNil(t, gotTag)
				assert.Equal(t, *tc.tag, *gotTag)
			}
		})
	}
}

func TestXAddressWrongIDLength(t *testing.T) {
	_, err := XAddress([]byte{1, 2, 3}, nil, false)
	require.Error(t, err)
}

func TestXIDFail(t *testing.T) {
	valid := "XVLhHMPHU98es4dbozjVtdWzVrDjtV5fdx1mHp98tDMoQXb"
	cases := []struct {
		name string
		x    string
	}{
		{"bad checksum", valid[:len(valid)-1] + "c"},
		{"classic address", "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh"},
		{"truncated", valid[:40]},
		{"too long", strings.Repeat("r", 49)},
		{"invalid base58 char", valid[:len(valid)-1] + "0"},
		{"empty", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, err := XID(tc.x)
			require.Error(t, err)
		})
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		classic string
		tag     *uint32
	}{
		{"classic passthrough", "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", nil},
		{"x tagless", "XVLhHMPHU98es4dbozjVtdWzVrDjtV5fdx1mHp98tDMoQXb", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", nil},
		{"x tagged", "XVLhHMPHU98es4dbozjVtdWzVrDjtVijiy6dVhdE4mRdMDU", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", tagPtr(13371337)},
		{"x testnet", "TVE26TYGhfLC7tQDno7G8dGtxSkYQn49b3qD26PK7FcGSKE", "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf", nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			classic, tag, err := Normalize(tc.in)
			require.NoError(t, err)
			assert.Equal(t, tc.classic, classic)
			if tc.tag == nil {
				assert.Nil(t, tag)
			} else {
				require.NotNil(t, tag)
				assert.Equal(t, *tc.tag, *tag)
			}
		})
	}
}

func TestNormalizeFail(t *testing.T) {
	cases := []string{
		"",
		"pHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", // unrecognized prefix
		"rHb9CJAWyB4rj91VRWn96DkukG4bwdtyT0", // invalid classic
		"XVLhHMPHU98es4dbozjVtdWzVrDjtV5fdx1mHp98tDMoQXc", // invalid x-address checksum
	}
	for _, in := range cases {
		_, _, err := Normalize(in)
		require.Error(t, err)
	}
}

func TestXIDRejectsUnsupportedFlag(t *testing.T) {
	id, err := ID("rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf")
	require.NoError(t, err)

	build := func(flag byte, tagBytes []byte) string {
		payload := append([]byte{0x05, 0x44}, id...)
		payload = append(payload, flag)
		payload = append(payload, tagBytes...)
		return base58.XRPLCoder.Encode(append(payload, hash.Checksum(payload)...))
	}

	// flag 2 (64-bit tag) is reserved and unsupported.
	_, _, _, err = XID(build(0x02, make([]byte, 8)))
	require.Error(t, err)

	// no-tag flag with non-zero tag bytes.
	_, _, _, err = XID(build(flagNoTag, []byte{1, 0, 0, 0, 0, 0, 0, 0}))
	require.Error(t, err)

	// 32-bit tag flag with the high tag bytes set.
	_, _, _, err = XID(build(flagTag, []byte{0, 0, 0, 0, 1, 0, 0, 0}))
	require.Error(t, err)
}
