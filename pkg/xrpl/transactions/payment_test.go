package transactions

import (
	"maps"
	"testing"

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
