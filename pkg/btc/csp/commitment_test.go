package csp

import (
	"encoding/hex"
	"testing"
)

// The commitment is `abi.encode` of five STATIC types, so the preimage is exactly
// five 32-byte words in declaration order — the CONTRACT's declaration order,
// in which nextAnchorVout sits third so that it shares a storage slot with
// anchorIndex and nonce. This asserts the layout directly rather
// than against our own packer, because the value has to match what Solidity
// computes — a disagreement would surface only as `CommitmentMismatch` on a
// proposal that is in fact correct.
func TestCommitmentPreimageIsFiveStaticWords(t *testing.T) {
	c := BtcProposalCommitment{
		AnchorIndex:    1,
		Nonce:          2,
		Txid:           [32]byte{0xaa},
		NextAnchorTxid: [32]byte{0xbb},
		NextAnchorVout: 3,
	}
	packed, err := commitmentArgs.Pack(c.AnchorIndex, c.Nonce, c.NextAnchorVout, c.Txid, c.NextAnchorTxid)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	if len(packed) != 5*32 {
		t.Fatalf("preimage is %d bytes, want 160 (five static words)", len(packed))
	}
	want := "" +
		"0000000000000000000000000000000000000000000000000000000000000001" + // anchorIndex
		"0000000000000000000000000000000000000000000000000000000000000002" + // nonce
		"0000000000000000000000000000000000000000000000000000000000000003" + // nextAnchorVout
		"aa00000000000000000000000000000000000000000000000000000000000000" + // txid, left-aligned
		"bb00000000000000000000000000000000000000000000000000000000000000" //   nextAnchorTxid
	if got := hex.EncodeToString(packed); got != want {
		t.Fatalf("preimage layout\n got %s\nwant %s", got, want)
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
