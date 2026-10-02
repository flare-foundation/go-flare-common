package htlc

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func TestBuildScript_ProducesValidScript(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	sortedPubs := walletKeys(t, 3)

	// Create a fake custodian pubkey (33 bytes compressed)
	custodianSeed := bytes.Repeat([]byte{0xFF}, 32)
	custodianMaster, err := hdkeychain.NewMaster(custodianSeed, params)
	if err != nil {
		t.Fatalf("custodian master: %v", err)
	}
	custodianPub, err := custodianMaster.ECPubKey()
	if err != nil {
		t.Fatalf("custodian pubkey: %v", err)
	}
	custodianPubBytes := custodianPub.SerializeCompressed()

	preimage := [32]byte{0xAA, 0xBB, 0xCC}
	preimageHash := PreimageHash(preimage)
	timeout := int64(1700000000)

	script, err := BuildScript(preimageHash, custodianPubBytes, timeout, sortedPubs, 2)
	if err != nil {
		t.Fatalf("BuildScript: %v", err)
	}

	// Script should be non-empty
	if len(script) == 0 {
		t.Fatal("empty HTLC script")
	}

	// Script should parse without errors
	_, err = txscript.DisasmString(script)
	if err != nil {
		t.Fatalf("script disasm failed: %v", err)
	}

	// Address should produce a valid P2WSH address
	addr, err := Address(script, params)
	if err != nil {
		t.Fatalf("Address: %v", err)
	}
	addrStr := addr.EncodeAddress()
	if addrStr[:5] != "bcrt1" {
		t.Errorf("expected bcrt1 prefix, got %s", addrStr[:5])
	}

	// PreimageHash should be deterministic
	hash2 := PreimageHash(preimage)
	if hash2 != preimageHash {
		t.Error("PreimageHash not deterministic")
	}
}

func TestBuildScript_ValidationErrors(t *testing.T) {
	pubkeys := [][]byte{
		bytes.Repeat([]byte{0x02}, 33),
		bytes.Repeat([]byte{0x03}, 33),
	}
	preimageHash := [32]byte{}
	custodian := bytes.Repeat([]byte{0x02}, 33)
	validTimeout := int64(1_700_000_000) // valid UNIX timestamp

	// threshold too low
	_, err := BuildScript(preimageHash, custodian, validTimeout, pubkeys, 0)
	if err == nil {
		t.Error("expected error for threshold 0")
	}

	// threshold too high
	_, err = BuildScript(preimageHash, custodian, validTimeout, pubkeys, 3)
	if err == nil {
		t.Error("expected error for threshold > n")
	}

	// wrong custodian pubkey length
	_, err = BuildScript(preimageHash, []byte{0x01, 0x02}, validTimeout, pubkeys, 1)
	if err == nil {
		t.Error("expected error for wrong custodian pubkey length")
	}
}

// TestBuildScript_TimeoutBoundary covers the CLTV time-vs-height boundary.
// Values < 500,000,000 are block heights under BIP-65; CV escrows must use
// UNIX timestamps (≥ 500,000,000). Any sub-threshold value must be rejected
// at build time.
func TestBuildScript_TimeoutBoundary(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	sortedPubs := walletKeys(t, 3)
	custodianSeed := bytes.Repeat([]byte{0xFF}, 32)
	custodianMaster, _ := hdkeychain.NewMaster(custodianSeed, params)
	custodianPub, _ := custodianMaster.ECPubKey()
	custodianPubBytes := custodianPub.SerializeCompressed()
	preimageHash := [32]byte{}

	cases := []struct {
		name    string
		timeout int64
		wantErr bool
	}{
		// Height-mode values (must be rejected)
		{"zero", 0, true},
		{"negative", -1, true},
		{"height_1", 1, true},
		{"height_499999999", 499_999_999, true},
		// Time-mode boundary
		{"time_boundary_500000000", 500_000_000, false},
		{"time_current_era", 1_700_000_000, false},
		{"time_int32max", maxTimeout, false},
		// Above int32 max (must be rejected)
		{"above_int32max", maxTimeout + 1, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := BuildScript(preimageHash, custodianPubBytes, tc.timeout, sortedPubs, 2)
			if tc.wantErr && err == nil {
				t.Fatalf("timeout %d: expected error, got nil", tc.timeout)
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("timeout %d: unexpected error: %v", tc.timeout, err)
			}
		})
	}
}
