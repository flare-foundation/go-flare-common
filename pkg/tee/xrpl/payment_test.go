package xrpl

import (
	"encoding/hex"
	"maps"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payment"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentTransactionMultisig(t *testing.T) {
	// using xrpl.js
	expectedBlobStr := "120000240000000A61400000000000000A68400000000000000A73008114AA8114C65DA8EA1BE40849F974685289CC145CCF8314F667B0CA50CC7709A220B0561B85E53A48461FA8F9EA7D209C22FF5F21F0B81B113E63F7DB6DA94FEDEF11B2119B4088B89664FB9A3CB658E1F1"
	expectedBlob, err := hex.DecodeString(expectedBlobStr)
	require.NoError(t, err)

	instruction := payment.ITeePaymentsPaymentInstructionMessage{
		WalletId:         [32]byte{1},
		TeeIdKeyIdPairs:  []payment.TeeIdKeyIdPair{},
		SenderAddress:    "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		RecipientAddress: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:           big.NewInt(10),
		PaymentReference: crypto.Keccak256Hash([]byte("test")),
		Nonce:            10,
		SubNonce:         0,
		MaxFee:           big.NewInt(10),
		FeeSchedule:      []byte{0x27, 0x10, 0, 1},
		BatchEndTs:       0,
	}

	tx, err := PaymentTxFromInstruction(instruction, 0)
	require.NoError(t, err)

	err = CheckNativePayment(tx)
	require.NoError(t, err)

	blob, err := types.Encode(tx, true)
	require.NoError(t, err)
	require.Equal(t, expectedBlob, blob)
}

// Source: rippled include/xrpl/protocol/detail/transactions.macro Payment common fields
// (TxFormats.cpp lines 16-33) and Nullify's intentional divergence — AccountSet instead of Payment.
// Verifies the nullify branch encodes into a valid signing-mode AccountSet blob that round-trips.
func TestPaymentTxFromInstructionNullify(t *testing.T) {
	// feeBIPS = -0x2710 = -10000 → Nullify path (parseScheduledFee).
	instruction := payment.ITeePaymentsPaymentInstructionMessage{
		WalletId:         [32]byte{2},
		TeeIdKeyIdPairs:  []payment.TeeIdKeyIdPair{},
		SenderAddress:    "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		RecipientAddress: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:           big.NewInt(1000),
		PaymentReference: crypto.Keccak256Hash([]byte("nullify")),
		Nonce:            5,
		SubNonce:         0,
		MaxFee:           big.NewInt(100),
		FeeSchedule:      []byte{0xD8, 0xF0, 0, 1}, // -10000 bips, 1s delay
		BatchEndTs:       0,
	}

	tx, err := PaymentTxFromInstruction(instruction, 0)
	require.NoError(t, err)

	require.Equal(t, "AccountSet", tx["TransactionType"])
	require.Equal(t, instruction.SenderAddress, tx["Account"])
	require.Equal(t, "", tx["SigningPubKey"])
	_, hasDestination := tx["Destination"]
	require.False(t, hasDestination, "nullify must not have Destination")
	_, hasAmount := tx["Amount"]
	require.False(t, hasAmount, "nullify must not have Amount")

	// CheckNativePayment rejects this (it's not a Payment).
	require.Error(t, CheckNativePayment(tx))

	blob, err := types.Encode(tx, true)
	require.NoError(t, err)
	require.NotEmpty(t, blob)

	decoded, err := types.Decode(blob)
	require.NoError(t, err)
	require.Equal(t, "AccountSet", decoded["TransactionType"])
}

// TestCheckNativePaymentRejectsIOUAmount verifies that CheckNativePayment rejects a Payment
// whose Amount is an IOU issued-currency object rather than an XRP drops string.
func TestCheckNativePaymentRejectsIOUAmount(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		"Amount": map[string]any{
			"currency": "USD",
			"value":    "10",
			"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		},
		"Fee":           "10",
		"Sequence":      uint32(1),
		"SigningPubKey": "",
	}
	require.Error(t, CheckNativePayment(tx))
}

// TestCheckNativePaymentMissingFields exercises every required-field rejection.
func TestCheckNativePaymentMissingFields(t *testing.T) {
	base := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		"Amount":          "10",
		"Fee":             "10",
		"Sequence":        uint32(1),
		"SigningPubKey":   "",
	}

	cases := []string{"TransactionType", "Account", "Destination", "Amount", "Fee", "Sequence", "SigningPubKey"}

	for _, field := range cases {
		t.Run("missing "+field, func(t *testing.T) {
			tx := make(map[string]any, len(base))
			maps.Copy(tx, base)
			delete(tx, field)
			require.Error(t, CheckNativePayment(tx))
		})
	}
}

// TestCheckNativePaymentRejectsZeroAmount documents that amount must be strictly positive.
func TestCheckNativePaymentRejectsZeroAmount(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		"Amount":          "0",
		"Fee":             "10",
		"Sequence":        uint32(1),
		"SigningPubKey":   "",
	}
	require.Error(t, CheckNativePayment(tx))
}

// TestCheckNativePaymentRejectsNonEmptySigningPubKey documents multisig convention.
func TestCheckNativePaymentRejectsNonEmptySigningPubKey(t *testing.T) {
	tx := map[string]any{
		"TransactionType": "Payment",
		"Account":         "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		"Destination":     "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		"Amount":          "10",
		"Fee":             "10",
		"Sequence":        uint32(1),
		"SigningPubKey":   "034AADB09CFF4A4804073701EC53C3510CDC95917C2BB0150FB742D0C66E6CEE9E",
	}
	require.Error(t, CheckNativePayment(tx))
}

func TestParseScheduledFeeBounds(t *testing.T) {
	cases := []struct {
		name    string
		bytes   []byte
		want    ScheduledFee
		wantErr bool
	}{
		{
			name:  "max positive bips",
			bytes: []byte{0x27, 0x10, 0x00, 0x00},
			want:  ScheduledFee{FeeBIPS: 10000, Delay: 0, Nullify: false},
		},
		{
			name:  "max negative bips nullify",
			bytes: []byte{0xD8, 0xF0, 0x00, 0x01},
			want:  ScheduledFee{FeeBIPS: 10000, Delay: 1 * time.Second, Nullify: true},
		},
		{
			name:  "small positive bips",
			bytes: []byte{0x00, 0x01, 0x00, 0x0A},
			want:  ScheduledFee{FeeBIPS: 1, Delay: 10 * time.Second, Nullify: false},
		},
		{
			name:    "zero bips rejected",
			bytes:   []byte{0x00, 0x00, 0x00, 0x00},
			wantErr: true,
		},
		{
			name:    "bips below -10000 rejected",
			bytes:   []byte{0xD8, 0xEF, 0x00, 0x00},
			wantErr: true,
		},
		{
			name:    "bips above +10000 rejected",
			bytes:   []byte{0x27, 0x11, 0x00, 0x00},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := parseScheduledFee(c.bytes)
			if c.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestParseFeeEntriesMultiple(t *testing.T) {
	// Two entries: (1 bips, 0s) then (-5000 bips / nullify, 60s).
	schedule := []byte{0x00, 0x01, 0x00, 0x00, 0xEC, 0x78, 0x00, 0x3C}
	got, err := ParseFeeEntries(schedule)
	require.NoError(t, err)
	require.Len(t, got, 2)

	assert.Equal(t, ScheduledFee{FeeBIPS: 1, Delay: 0, Nullify: false}, got[0])
	assert.Equal(t, ScheduledFee{FeeBIPS: 5000, Delay: 60 * time.Second, Nullify: true}, got[1])

	// Individual Try access
	first, err := ParseFeeEntry(schedule, 0)
	require.NoError(t, err)
	require.Equal(t, got[0], first)

	second, err := ParseFeeEntry(schedule, 1)
	require.NoError(t, err)
	require.Equal(t, got[1], second)

	_, err = ParseFeeEntry(schedule, 2)
	require.Error(t, err)

	_, err = ParseFeeEntries([]byte{0x00, 0x01, 0x00})
	require.Error(t, err)
}

func TestScheduledFeeComputation(t *testing.T) {
	s := ScheduledFee{FeeBIPS: 2500} // 25%
	got := s.Fee(big.NewInt(1000))
	require.Equal(t, "250", got)

	// Max fee 0 → 0 fee.
	got = s.Fee(big.NewInt(0))
	require.Equal(t, "0", got)
}
