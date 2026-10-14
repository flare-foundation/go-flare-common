package batch

import (
	"bytes"
	"errors"
)

// Status mirrors the attestation's transactionStatus field.
type Status uint8

const (
	// StatusSuccess — the payment was delivered: value outputs to the expected
	// recipient sum to the expected amount.
	StatusSuccess Status = 0
	// StatusUndeliverable — the group carries no value to the expected recipient
	// (K=0). For a PMW redeem this means the amount was below the recipient's
	// dust threshold and was redirected to change. Whether that is legitimately
	// sub-dust is a dust-threshold check the caller may perform against the
	// instruction amount.
	StatusUndeliverable Status = 1
	// StatusNotFound — no confirmed transaction / group for this payment.
	// Returned by the caller (no group to Match), not here.
	StatusNotFound Status = 2
)

// Expected is the on-chain payment instruction a group is checked against. The
// recipient is a scriptPubKey — the caller decodes the instruction's
// recipientAddress to its script (keeping this package free of address/network
// concerns).
type Expected struct {
	Reference []byte // paymentReference
	Recipient []byte // recipient scriptPubKey
	Amount    int64  // satoshis the instruction says to send
}

// Match consistency errors (the on-chain group contradicts the instruction).
var (
	ErrReferenceMismatch = errors.New("btcbatch: group reference does not match the instruction")
	ErrAmountMismatch    = errors.New("btcbatch: value to recipient does not equal the instruction amount")
)

// Match evaluates a payment group against its expected instruction and returns
// the payment status and received amount (satoshis).
//
//   - The group's reference must equal the instruction's reference (otherwise
//     the caller mapped the wrong group / batch) → ErrReferenceMismatch.
//   - Value outputs paying the expected recipient are summed (outputs to any
//     other script — e.g. the trailing change — are ignored):
//   - sum > 0 and == amount → StatusSuccess, receivedAmount = sum.
//   - sum > 0 and != amount → ErrAmountMismatch (amount-balance violation).
//   - sum == 0 → StatusUndeliverable, receivedAmount 0.
//
// It does not compute dust thresholds or handle StatusNotFound — both are the
// caller's responsibility.
func Match(g Group, exp Expected) (Status, int64, error) {
	if !bytes.Equal(g.Reference, exp.Reference) {
		return StatusNotFound, 0, ErrReferenceMismatch
	}
	var received int64
	for _, o := range g.Outputs {
		if bytes.Equal(o.PkScript, exp.Recipient) {
			received += o.Value
		}
	}
	if received == 0 {
		return StatusUndeliverable, 0, nil
	}
	if received != exp.Amount {
		return StatusNotFound, 0, ErrAmountMismatch
	}
	return StatusSuccess, received, nil
}
