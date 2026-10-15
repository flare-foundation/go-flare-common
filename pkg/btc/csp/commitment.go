package csp

import (
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/walletpayments"
)

// BtcProposalCommitment is what a CSP proposal commits to on Bitcoin: the
// contract's IBtcAccounts.BtcProposalCommitment, as generated from the struct
// ABI.
//
// The attested type `CspProposalCheck` is chain-neutral: it carries who proposed,
// the score, the payment count and ONE `chainCommitmentHash`. The Bitcoin facts
// travel separately, in calldata to `finalizeProposal(proof, commitment)`, and the
// diamond binds them to the proof by this hash — `CommitmentMismatch` if they
// disagree. Nothing is trusted that was not trusted when these fields sat inside
// the response body; they are bound by hash instead of by inclusion, which is what
// lets a second CSP chain reuse the arbitration without registering a type of its
// own.
//
// THE PROPOSER AND THE VERIFIER MUST PRODUCE THE SAME BYTES, which is why this
// lives here rather than in either of them: the proposer builds the struct for
// calldata, the verifier hashes it into the response, and a disagreement would
// surface as `CommitmentMismatch` on a proposal that is in fact correct. The
// field order is the contract's, and it is the order `abi.encode` hashes.
type BtcProposalCommitment walletpayments.IBtcAccountsBtcProposalCommitment

// CommitmentHash returns `keccak256(abi.encode(commitment))`, the value the
// attested response carries as `chainCommitmentHash`.
func (c BtcProposalCommitment) CommitmentHash() [32]byte {
	packed, err := abicoder.Encode(walletpayments.BtcProposalCommitmentArgument, walletpayments.IBtcAccountsBtcProposalCommitment(c))
	if err != nil {
		// Every field is a static type, so encoding cannot fail on any value; a
		// failure here is a programming error.
		panic("csp: encoding BtcProposalCommitment: " + err.Error())
	}
	return [32]byte(crypto.Keccak256Hash(packed))
}

// CommitmentOf builds the commitment an envelope proposes. The successor anchor is
// the transaction's own output[0] — the batch grammar puts it there and every
// consumer parses it there.
func CommitmentOf(e Envelope) BtcProposalCommitment {
	txid := TxidOf(e)
	return BtcProposalCommitment{
		AnchorIndex:    e.AnchorIndex,
		Nonce:          e.Nonce,
		Txid:           txid,
		NextAnchorTxid: txid,
		NextAnchorVout: 0,
	}
}
