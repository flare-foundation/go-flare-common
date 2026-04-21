package types

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// fixtureTx is a transaction vector copied from xrpl.js codec-fixtures, which
// is the XRPL binary-codec round-trip reference suite.
type fixtureTx struct {
	name   string
	source string
	binary string
	tx     string
}

// fixtureTxs enumerates transactions from xrpl.js codec-fixtures that exercise
// the in-scope encoding types (AccountID length prefix, UInt16 TransactionType
// string mapping, UInt32, field ordering).
var fixtureTxs = []fixtureTx{
	{
		name:   "OracleDelete",
		source: "xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json transactions[26]",
		binary: "1200342033000004D2811401476926B590BA3245F63C829116A0A3AF7F382D",
		tx: `{
			"TransactionType": "OracleDelete",
			"Account": "rfmDuhDyLGgx94qiwf3YF8BUV5j6KSvE8",
			"OracleDocumentID": 1234
		}`,
	},
}

// TestCodecFixtures round-trips the xrpl.js codec-fixtures transactions
// through Encode/Decode. It also asserts the AccountID "Account" field is
// serialized as 0x81 (field ID) followed by the VL prefix 0x14 (=20) then
// the 20-byte AccountID — exercising task item 9.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json
// source: rippled src/libxrpl/protocol/STAccount.cpp (AccountID is VL-prefixed 20 bytes)
func TestCodecFixtures(t *testing.T) {
	for _, fx := range fixtureTxs {
		t.Run(fx.name, func(t *testing.T) {
			blob, err := hex.DecodeString(fx.binary)
			require.NoError(t, err)

			var parsed map[string]any
			require.NoError(t, json.Unmarshal([]byte(fx.tx), &parsed))

			encoded, err := Encode(parsed, false)
			require.NoError(t, err)
			require.Equal(t, blob, encoded, "encode mismatch for %s\n  source: %s", fx.name, fx.source)

			decoded, err := Decode(blob)
			require.NoError(t, err)

			reEncoded, err := Encode(decoded, false)
			require.NoError(t, err)
			require.Equal(t, blob, reEncoded, "decode+re-encode mismatch for %s", fx.name)
		})
	}
}

// TestAccountIDVLPrefix verifies the on-wire form of an Account field: the
// field ID (0x81) is followed by the VL prefix 0x14 (=20) and then the
// AccountID raw bytes. The fixture below is the OracleDelete trailer from
// codec-fixtures transactions[26].
//
// source: xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json transactions[26]
// source: rippled include/xrpl/protocol/SField.h (sfAccount type=8, nth=1 -> ID 0x81)
// source: rippled src/libxrpl/protocol/Serializer.cpp:178-187 (addVL)
func TestAccountIDVLPrefix(t *testing.T) {
	// "Account": "rfmDuhDyLGgx94qiwf3YF8BUV5j6KSvE8"
	// Expected wire bytes for this field: 81 14 01476926B590BA3245F63C829116A0A3AF7F382D
	const expected = "811401476926B590BA3245F63C829116A0A3AF7F382D"

	obj := map[string]any{
		"Account": "rfmDuhDyLGgx94qiwf3YF8BUV5j6KSvE8",
	}
	got, err := Encode(obj, false)
	require.NoError(t, err)

	require.Equal(t, expected, strings.ToUpper(hex.EncodeToString(got)))
}
