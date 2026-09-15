package csp

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// BtcProposalCommitment is what a CSP proposal commits to on Bitcoin.
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
// surface as `CommitmentMismatch` on a proposal that is in fact correct.
type BtcProposalCommitment struct {
	AnchorIndex    uint32
	Nonce          uint64
	Txid           [32]byte
	NextAnchorTxid [32]byte
	NextAnchorVout uint32
}

// commitmentArgs mirrors `abi.encode(BtcProposalCommitment)` — five static words,
// in declaration order. Built once; the types cannot fail to parse.
var commitmentArgs = func() abi.Arguments {
	u32, _ := abi.NewType("uint32", "", nil)
	u64, _ := abi.NewType("uint64", "", nil)
	b32, _ := abi.NewType("bytes32", "", nil)
	return abi.Arguments{
		{Type: u32}, {Type: u64}, {Type: b32}, {Type: b32}, {Type: u32},
	}
}()

// CommitmentHash returns `keccak256(abi.encode(commitment))`, the value the
// attested response carries as `chainCommitmentHash`.
func (c BtcProposalCommitment) CommitmentHash() [32]byte {
	packed, err := commitmentArgs.Pack(
		c.AnchorIndex, c.Nonce, c.Txid, c.NextAnchorTxid, c.NextAnchorVout,
	)
	if err != nil {
		// Every field is a static type built from constants above, so packing
		// cannot fail on well-formed input; a failure here is a programming error.
		panic("csp: packing BtcProposalCommitment: " + err.Error())
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
