package csp_test

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcd/wire"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
)

// batch builds a transaction in the batch grammar: the three fixed outputs,
// then `groups` payment groups of `perGroup` value outputs each, then change.
func batch(t *testing.T, groups, perGroup int, change bool) csp.Envelope {
	t.Helper()
	tx := wire.NewMsgTx(3)
	tx.AddTxIn(wire.NewTxIn(&wire.OutPoint{}, nil, nil))
	tx.AddTxOut(wire.NewTxOut(10_000, []byte{0x00, 0x20})) // anchor
	tx.AddTxOut(wire.NewTxOut(330, []byte{0x51, 0x02}))    // P2A
	tx.AddTxOut(wire.NewTxOut(0, []byte{0x6a, 0x0c}))      // nonce OP_RETURN
	for range groups {
		tx.AddTxOut(wire.NewTxOut(0, []byte{0x6a, 0x20})) // reference OP_RETURN opens the group
		for range perGroup {
			tx.AddTxOut(wire.NewTxOut(1000, []byte{0x00, 0x14}))
		}
	}
	if change {
		tx.AddTxOut(wire.NewTxOut(5000, []byte{0x00, 0x20}))
	}
	var buf bytes.Buffer
	if err := tx.SerializeNoWitness(&buf); err != nil {
		t.Fatalf("serializing: %v", err)
	}
	return csp.Envelope{RawUnsignedTx: buf.Bytes(), PaymentCount: uint32(groups)}
}

func TestPaymentGroupsCountsGroupsNotOutputs(t *testing.T) {
	// A group may pay several outputs (a CV top-up spreads across
	// denominations); the payment is the GROUP, so K must not inflate the count.
	for _, perGroup := range []int{1, 3} {
		got, err := csp.PaymentGroups(batch(t, 4, perGroup, true))
		if err != nil {
			t.Fatalf("K=%d: %v", perGroup, err)
		}
		if got != 4 {
			t.Errorf("K=%d: counted %d payments, want 4", perGroup, got)
		}
	}
}

func TestPaymentGroupsIgnoresChange(t *testing.T) {
	with, err := csp.PaymentGroups(batch(t, 2, 1, true))
	if err != nil {
		t.Fatal(err)
	}
	without, err := csp.PaymentGroups(batch(t, 2, 1, false))
	if err != nil {
		t.Fatal(err)
	}
	if with != without || with != 2 {
		t.Errorf("change changed the count: %d with, %d without", with, without)
	}
}

func TestPaymentGroupsRejectsATruncatedBatch(t *testing.T) {
	// Fewer than the three fixed outputs is not a batch at all, and returning
	// zero would read as "an empty batch" rather than as malformed.
	tx := wire.NewMsgTx(3)
	tx.AddTxOut(wire.NewTxOut(10_000, []byte{0x00, 0x20}))
	var buf bytes.Buffer
	if err := tx.SerializeNoWitness(&buf); err != nil {
		t.Fatal(err)
	}
	if _, err := csp.PaymentGroups(csp.Envelope{RawUnsignedTx: buf.Bytes()}); err == nil {
		t.Error("a one-output transaction was accepted as a batch")
	}
}
