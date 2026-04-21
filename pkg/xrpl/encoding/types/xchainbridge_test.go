package types

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestXChainBridgeEncode asserts the full 82-byte serialized form of
// XChainBridge for the XChainCreateBridge transaction in xrpl.js codec-fixtures.
// Layout per rippled include/xrpl/protocol/XChainBridge.h and
// src/libxrpl/protocol/XChainBridge.cpp:
//
//	LockingChainDoor  (length-prefixed AccountID: 0x14 + 20 bytes)
//	LockingChainIssue (20-byte Issue; XRP=all zeros)
//	IssuingChainDoor  (length-prefixed AccountID: 0x14 + 20 bytes)
//	IssuingChainIssue (20-byte Issue; XRP=all zeros)
//
// source: xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json transactions[1]
func TestXChainBridgeEncode(t *testing.T) {
	const bridgeJSON = `{
		"LockingChainDoor": "rGzx83BVoqTYbGn7tiVAnFw7cbxjin13jL",
		"LockingChainIssue": {"currency": "XRP"},
		"IssuingChainDoor": "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
		"IssuingChainIssue": {"currency": "XRP"}
	}`
	const expected = "14AF80285F637EE4AF3C20378F9DFB12511ACB8D27" +
		"0000000000000000000000000000000000000000" +
		"14550FC62003E785DC231A1058A05E56E3F09CF4E6" +
		"0000000000000000000000000000000000000000"

	var parsed map[string]any
	require.NoError(t, json.Unmarshal([]byte(bridgeJSON), &parsed))

	got, err := (&XChainBridge{}).ToBytes(parsed, false)
	require.NoError(t, err)

	want, err := hex.DecodeString(expected)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, 82, len(got))
}

// TestXChainBridgeFieldOrdering verifies the XChainBridge encoder emits fields
// in the canonical order LockingChainDoor, LockingChainIssue, IssuingChainDoor,
// IssuingChainIssue regardless of JSON input order.
//
// source: rippled src/libxrpl/protocol/XChainBridge.cpp (field order in
//
//	add())
func TestXChainBridgeFieldOrdering(t *testing.T) {
	bridge := map[string]any{
		// order deliberately shuffled
		"IssuingChainIssue": map[string]any{"currency": "XRP"},
		"IssuingChainDoor":  "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
		"LockingChainIssue": map[string]any{"currency": "XRP"},
		"LockingChainDoor":  "rGzx83BVoqTYbGn7tiVAnFw7cbxjin13jL",
	}

	got, err := (&XChainBridge{}).ToBytes(bridge, false)
	require.NoError(t, err)

	// First door after its length byte must be rGzx83…, i.e. LockingChainDoor.
	require.Equal(t, byte(0x14), got[0])
	const lockingDoorID = "AF80285F637EE4AF3C20378F9DFB12511ACB8D27"
	require.Equal(t, lockingDoorID, strings.ToUpper(hex.EncodeToString(got[1:21])))

	// Second door (after the first XRP issue + its length byte) must be IssuingChainDoor.
	require.Equal(t, byte(0x14), got[41])
	const issuingDoorID = "550FC62003E785DC231A1058A05E56E3F09CF4E6"
	require.Equal(t, issuingDoorID, strings.ToUpper(hex.EncodeToString(got[42:62])))
}

// TestXChainBridgeErrors exercises the structural guards in the encoder.
func TestXChainBridgeErrors(t *testing.T) {
	tests := []any{
		// wrong type
		"not-a-map",
		// missing LockingChainDoor
		map[string]any{
			"LockingChainIssue": map[string]any{"currency": "XRP"},
			"IssuingChainDoor":  "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
			"IssuingChainIssue": map[string]any{"currency": "XRP"},
		},
		// missing LockingChainIssue
		map[string]any{
			"LockingChainDoor":  "rGzx83BVoqTYbGn7tiVAnFw7cbxjin13jL",
			"IssuingChainDoor":  "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
			"IssuingChainIssue": map[string]any{"currency": "XRP"},
		},
		// missing IssuingChainIssue
		map[string]any{
			"LockingChainDoor":  "rGzx83BVoqTYbGn7tiVAnFw7cbxjin13jL",
			"LockingChainIssue": map[string]any{"currency": "XRP"},
			"IssuingChainDoor":  "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
		},
		// invalid door address
		map[string]any{
			"LockingChainDoor":  "rXXXXXXXXXXXXX",
			"LockingChainIssue": map[string]any{"currency": "XRP"},
			"IssuingChainDoor":  "r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
			"IssuingChainIssue": map[string]any{"currency": "XRP"},
		},
	}
	for _, in := range tests {
		_, err := (&XChainBridge{}).ToBytes(in, false)
		require.Error(t, err)
	}
}
