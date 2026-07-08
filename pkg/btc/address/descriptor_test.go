package address

import (
	"strings"
	"testing"
)

func TestComputeDescriptorChecksum_KnownVectors(t *testing.T) {
	// Vectors verified against Bitcoin Core's getdescriptorinfo RPC and
	// rust-miniscript. These are widely cited in BIP-380 tooling docs.
	cases := []struct {
		expr     string
		expected string
	}{
		// Minimal raw() descriptor — the canonical first vector cited in
		// most BIP-380 tooling docs.
		{"raw(deadbeef)", "89f8spxm"},
		// pkh() with the fixed test pubkey from Bitcoin Core's
		// src/test/descriptor_tests.cpp; checksum is documented as
		// "8fhd9pwu" there.
		{"pkh(02c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5)", "8fhd9pwu"},
	}
	for _, c := range cases {
		got, err := ComputeDescriptorChecksum(c.expr)
		if err != nil {
			t.Errorf("ComputeDescriptorChecksum(%q): unexpected error %v", c.expr, err)
			continue
		}
		if got != c.expected {
			t.Errorf("ComputeDescriptorChecksum(%q) = %q, want %q", c.expr, got, c.expected)
		}
	}
}

func TestVerifyDescriptorChecksum_RoundTrip(t *testing.T) {
	// A realistic wsh(sortedmulti(...)) expression — the shape used by
	// BtcAccountConfigured.scriptDescriptor for our wallets.
	expr := "wsh(sortedmulti(2," +
		"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz/0/*," +
		"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz/0/*," +
		"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz/0/*))"
	csum, err := ComputeDescriptorChecksum(expr)
	if err != nil {
		t.Fatalf("ComputeDescriptorChecksum: %v", err)
	}
	if len(csum) != 8 {
		t.Fatalf("checksum length = %d, want 8", len(csum))
	}
	full := expr + "#" + csum
	got, err := VerifyDescriptorChecksum(full)
	if err != nil {
		t.Fatalf("VerifyDescriptorChecksum round-trip: %v", err)
	}
	if got != expr {
		t.Errorf("VerifyDescriptorChecksum returned expression = %q, want %q", got, expr)
	}
}

func TestVerifyDescriptorChecksum_Rejections(t *testing.T) {
	const baseExpr = "raw(deadbeef)"
	csum, err := ComputeDescriptorChecksum(baseExpr)
	if err != nil {
		t.Fatalf("compute: %v", err)
	}

	// Flip one character of the checksum to a different valid bech32 char.
	tampered := []byte(csum)
	for i := range len(descriptorChecksumCharset) {
		if descriptorChecksumCharset[i] != tampered[0] {
			tampered[0] = descriptorChecksumCharset[i]
			break
		}
	}

	cases := []struct {
		name string
		full string
		want string
	}{
		{"missing separator", baseExpr, "missing '#'"},
		{"empty checksum", baseExpr + "#", "checksum length 0"},
		{"short checksum", baseExpr + "#qpzry9x", "checksum length 7"},
		{"long checksum", baseExpr + "#qpzry9x88", "checksum length 9"},
		{"non-bech32 char", baseExpr + "#qpzry9xb", "outside bech32 alphabet"}, // 'b' isn't in alphabet
		{"checksum mismatch", baseExpr + "#" + string(tampered), "checksum mismatch"},
		{"invalid expression char", "raw(deadbeef)\x01#" + csum, "invalid character"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := VerifyDescriptorChecksum(c.full)
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", c.want)
			}
			if !strings.Contains(err.Error(), c.want) {
				t.Errorf("error %q does not contain %q", err.Error(), c.want)
			}
		})
	}
}

func TestComputeDescriptorChecksum_AllowedAlphabet(t *testing.T) {
	// Every character in descriptorInputCharset is valid; passing all of
	// them should not error.
	_, err := ComputeDescriptorChecksum(descriptorInputCharset)
	if err != nil {
		t.Errorf("alphabet itself should be a valid expression for checksum computation, got %v", err)
	}
}
