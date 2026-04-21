package transactions

import (
	"encoding/hex"
	"encoding/json"
	"maps"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/require"
)

func TestCheckAndEncodePayment(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "1",
		"Amount":          "1",
	}

	_, err := CheckAndEncodePayment(tx, true)

	require.NoError(t, err)
}

func TestCheckAndEncodePaymentNonNative(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "12",
		"Amount": map[string]any{
			"currency": "USD",
			"value":    "10",
			"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		},
	}

	_, err := CheckAndEncodePayment(tx, false)
	require.NoError(t, err)
}

func TestCheckAndEncodePaymentErrors(t *testing.T) {
	base := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "1",
		"Amount":          "1",
	}

	copyTx := func() map[string]any {
		tx := make(map[string]any, len(base))
		maps.Copy(tx, base)
		return tx
	}

	tests := []struct {
		name   string
		modify func(map[string]any)
		native bool
	}{
		{
			name: "wrong transaction type",
			modify: func(tx map[string]any) {
				tx["TransactionType"] = "TrustSet"
			},
			native: true,
		},
		{
			name: "self payment",
			modify: func(tx map[string]any) {
				tx["Destination"] = tx["Account"]
			},
			native: true,
		},
		{
			name: "zero fee",
			modify: func(tx map[string]any) {
				tx["Fee"] = "0"
			},
			native: true,
		},
		{
			name: "zero native amount",
			modify: func(tx map[string]any) {
				tx["Amount"] = "0"
			},
			native: true,
		},
		{
			name: "non-string amount when native",
			modify: func(tx map[string]any) {
				tx["Amount"] = map[string]any{
					"currency": "USD",
					"value":    "10",
					"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				}
			},
			native: true,
		},
		{
			name: "missing Account",
			modify: func(tx map[string]any) {
				delete(tx, "Account")
			},
			native: true,
		},
		{
			name: "missing Destination",
			modify: func(tx map[string]any) {
				delete(tx, "Destination")
			},
			native: true,
		},
		{
			name: "missing Fee",
			modify: func(tx map[string]any) {
				delete(tx, "Fee")
			},
			native: true,
		},
		{
			name: "missing Amount",
			modify: func(tx map[string]any) {
				delete(tx, "Amount")
			},
			native: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tx := copyTx()
			test.modify(tx)
			_, err := CheckAndEncodePayment(tx, test.native)
			require.Error(t, err)
		})
	}
}

// Source: xrpl.js packages/ripple-binary-codec/test/fixtures/codec-fixtures.json line 4496 Payment.
func TestCheckAndEncodePaymentXrplFixture(t *testing.T) {
	raw := `{
		"Account": "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
		"Destination": "rLQBHVhFnaC5gLEkgr6HgBJJ3bgeZHg9cj",
		"TransactionType": "Payment",
		"SigningPubKey": "034AADB09CFF4A4804073701EC53C3510CDC95917C2BB0150FB742D0C66E6CEE9E",
		"Amount": "10000000000",
		"Fee": "10",
		"Flags": 0,
		"Sequence": 62
	}`

	var tx map[string]any
	require.NoError(t, json.Unmarshal([]byte(raw), &tx))

	blob, err := CheckAndEncodePayment(tx, true)
	require.NoError(t, err)

	decoded, err := types.Decode(blob)
	require.NoError(t, err)
	require.Equal(t, "Payment", decoded["TransactionType"])
	require.Equal(t, "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV", decoded["Account"])
	require.Equal(t, "rLQBHVhFnaC5gLEkgr6HgBJJ3bgeZHg9cj", decoded["Destination"])
	require.Equal(t, "10000000000", decoded["Amount"])
	require.Equal(t, "10", decoded["Fee"])
}

// Source: rippled src/test/protocol/STTx_test.cpp getPayment() lines 1390-1402 minimum viable Payment.
func TestCheckAndEncodePaymentSTTxTestMinimal(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Amount":          "10000000000",
		"Fee":             "10",
		"Sequence":        uint32(1),
		"SigningPubKey":   "",
	}

	blob, err := CheckAndEncodePayment(tx, true)
	require.NoError(t, err)

	decoded, err := types.Decode(blob)
	require.NoError(t, err)
	require.Equal(t, "10000000000", decoded["Amount"])
}

// Source: rippled include/xrpl/protocol/detail/transactions.macro lines 27-41 — Payment optional fields.
func TestCheckAndEncodePaymentOptionalFields(t *testing.T) {
	cases := []struct {
		name  string
		extra map[string]any
	}{
		{
			name:  "DestinationTag",
			extra: map[string]any{"DestinationTag": uint32(12345)},
		},
		{
			name:  "InvoiceID",
			extra: map[string]any{"InvoiceID": "342B8D16BEE494D169034AFF0908FDE35874A38E548D4CEC8DFC5C49E9A33B76"},
		},
		{
			name:  "LastLedgerSequence",
			extra: map[string]any{"LastLedgerSequence": uint32(65953073)},
		},
		{
			name:  "Flags",
			extra: map[string]any{"Flags": uint32(2147483648)},
		},
		{
			name:  "SendMax",
			extra: map[string]any{"SendMax": "15000"},
		},
		{
			name: "Memos",
			extra: map[string]any{"Memos": []any{
				map[string]any{"Memo": map[string]any{"MemoData": hex.EncodeToString([]byte("hi"))}},
			}},
		},
	}

	base := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "10",
		"Amount":          "1",
		"Sequence":        uint32(1),
		"SigningPubKey":   "",
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tx := make(map[string]any, len(base)+len(c.extra))
			maps.Copy(tx, base)
			maps.Copy(tx, c.extra)
			_, err := CheckAndEncodePayment(tx, true)
			require.NoError(t, err)
		})
	}
}

// Deviation: rippled TxFormats.cpp common fields mark Sequence and SigningPubKey as soeREQUIRED (lines 19, 27),
// but CheckAndEncodePayment does not enforce their presence on the decoded tx.
// This test documents that current behavior (passes without these fields when types.Encode tolerates omission).
func TestCheckAndEncodePaymentMissingCommonRequiredFields(t *testing.T) {
	cases := []struct {
		name   string
		omit   string
		accept bool
	}{
		{name: "missing Sequence", omit: "Sequence", accept: true},
		{name: "missing SigningPubKey", omit: "SigningPubKey", accept: true},
	}

	base := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rw16SLQGtfnjQJxgp1RCfkxsMCk8G7PaCJ",
		"Destination":     "rfNgpzQecR231M6d9rRZKsMCmrykK5bYLB",
		"Fee":             "10",
		"Amount":          "1",
		"Sequence":        uint32(1),
		"SigningPubKey":   "",
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tx := make(map[string]any, len(base))
			maps.Copy(tx, base)
			delete(tx, c.omit)
			_, err := CheckAndEncodePayment(tx, true)
			if c.accept {
				require.NoError(t, err, "documents deviation from rippled TxFormats.cpp required fields")
			} else {
				require.Error(t, err)
			}
		})
	}
}
