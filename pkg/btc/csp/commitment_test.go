package csp

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

// The commitment is `abi.encode` of five STATIC types, so the preimage is exactly
// five 32-byte words in the CONTRACT's declaration order, in which
// nextAnchorVout sits third so that it shares a storage slot with anchorIndex
// and nonce. The hash is checked against those words written out by hand,
// because the value has to match what Solidity computes — a disagreement would
// surface only as `CommitmentMismatch` on a proposal that is in fact correct.
func TestCommitmentPreimageIsFiveStaticWords(t *testing.T) {
	c := BtcProposalCommitment{
		AnchorIndex:    1,
		Nonce:          2,
		Txid:           [32]byte{0xaa},
		NextAnchorTxid: [32]byte{0xbb},
		NextAnchorVout: 3,
	}
	preimage, err := hex.DecodeString("" +
		"0000000000000000000000000000000000000000000000000000000000000001" + // anchorIndex
		"0000000000000000000000000000000000000000000000000000000000000002" + // nonce
		"0000000000000000000000000000000000000000000000000000000000000003" + // nextAnchorVout
		"aa00000000000000000000000000000000000000000000000000000000000000" + // txid, left-aligned
		"bb00000000000000000000000000000000000000000000000000000000000000") //  nextAnchorTxid
	if err != nil {
		t.Fatalf("decoding the preimage: %v", err)
	}
	if got, want := c.CommitmentHash(), [32]byte(crypto.Keccak256Hash(preimage)); got != want {
		t.Fatalf("commitment hash %x, want keccak256 of the five words %x", got, want)
	}
}

// A commitment that differs in ANY field must hash differently — otherwise the
// binding would admit a proposal that named a different transaction.
func TestCommitmentHashIsSensitiveToEveryField(t *testing.T) {
	base := BtcProposalCommitment{AnchorIndex: 1, Nonce: 2, Txid: [32]byte{3}, NextAnchorTxid: [32]byte{4}, NextAnchorVout: 5}
	seen := map[[32]byte]string{base.CommitmentHash(): "base"}
	for name, m := range map[string]func(*BtcProposalCommitment){
		"anchorIndex":    func(c *BtcProposalCommitment) { c.AnchorIndex++ },
		"nonce":          func(c *BtcProposalCommitment) { c.Nonce++ },
		"txid":           func(c *BtcProposalCommitment) { c.Txid[31] ^= 1 },
		"nextAnchorTxid": func(c *BtcProposalCommitment) { c.NextAnchorTxid[31] ^= 1 },
		"nextAnchorVout": func(c *BtcProposalCommitment) { c.NextAnchorVout++ },
	} {
		v := base
		m(&v)
		h := v.CommitmentHash()
		if other, dup := seen[h]; dup {
			t.Fatalf("changing %s collides with %s", name, other)
		}
		seen[h] = name
	}
}
