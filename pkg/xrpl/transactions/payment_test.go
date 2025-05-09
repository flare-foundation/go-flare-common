package transactions

import (
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
