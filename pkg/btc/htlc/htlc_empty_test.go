package htlc

import (
	"crypto/sha256"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
)

// TestBuildScript_RejectsEmptyMultisigKeys pins that BuildScript
// refuses an empty multisigKeys slice. Without an explicit assertion the
// only protection is the implicit "threshold > len(multisigKeys)" check
// (threshold ≥ 1 by spec invariant, len == 0 ⇒ rejected) — but a future
// refactor that allowed threshold=0 or skipped the range check would
// silently emit an unspendable timeout branch.
func TestBuildScript_RejectsEmptyMultisigKeys(t *testing.T) {
	// Produce a real custodian pubkey so we hit the multisigKeys path,
	// not the custodian-validation path.
	priv, err := btcec.NewPrivateKey()
	if err != nil {
		t.Fatalf("priv key: %v", err)
	}
	custodian := priv.PubKey().SerializeCompressed()
	preimage := [32]byte{0x42}
	preimageHash := sha256.Sum256(preimage[:])

	cases := []struct {
		name      string
		keys      [][]byte
		threshold int
		want      string
	}{
		{"nil_keys_threshold_1", nil, 1, "threshold 1 out of range for n=0"},
		{"empty_keys_threshold_1", [][]byte{}, 1, "threshold 1 out of range for n=0"},
		{"nil_keys_threshold_0", nil, 0, "threshold 0 out of range"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := BuildScript(preimageHash, custodian, 1_000_000, c.keys, c.threshold)
			if err == nil {
				t.Fatalf("expected reject for %s, got nil", c.name)
			}
			if !strings.Contains(err.Error(), c.want) {
				t.Errorf("error %q does not contain %q", err, c.want)
			}
		})
	}
}
