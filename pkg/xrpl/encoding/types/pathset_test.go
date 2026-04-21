package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPathSetAccountStep encodes a single-path, single-step PathSet that uses
// the 0x01 (account) flag and asserts the exact wire form, then round-trips.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/data-driven-tests.json
//
//	whole_objects[4].fields[...] (Paths: [[{account: razqQK…}]])
//	Binary: 01 41C8BE2C0A6AA17471B9F6D0AF92AAB1C94D5A25 00
//
// source: rippled include/xrpl/protocol/STPathSet.h:29-34
//
//	typeNone=0x00, typeAccount=0x01, typeBoundary=0xFF
func TestPathSetAccountStep(t *testing.T) {
	const pathsJSON = `[[{"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"}]]`
	const expected = "0141C8BE2C0A6AA17471B9F6D0AF92AAB1C94D5A2500"

	var paths any
	require.NoError(t, json.Unmarshal([]byte(pathsJSON), &paths))

	got, err := (&PathSet{}).ToBytes(paths, false)
	require.NoError(t, err)

	want, err := hex.DecodeString(expected)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Round-trip.
	decoded, err := (&PathSet{}).ToJSON(bytes.NewBuffer(got), 0)
	require.NoError(t, err)
	require.Equal(t, paths, decoded)
}

// TestPathSetCurrencyIssuerStep encodes a single-path step that combines
// currency (0x10) and issuer (0x20) flags, yielding 0x30 as the step flag,
// followed by the 20-byte currency and 20-byte issuer.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/data-driven-tests.json
//
//	whole_objects[10].fields[Paths]
//	Binary: 30 0000000000000000000000004254430000000000
//	        DD39C650A96EDA48334E70CC4A85B8B2E8502CD3 00
//
// source: rippled include/xrpl/protocol/STPathSet.h:30-32 (0x10 currency, 0x20 issuer)
func TestPathSetCurrencyIssuerStep(t *testing.T) {
	const pathsJSON = `[[{"currency": "BTC", "issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"}]]`
	const expected = "300000000000000000000000004254430000000000DD39C650A96EDA48334E70CC4A85B8B2E8502CD300"

	var paths any
	require.NoError(t, json.Unmarshal([]byte(pathsJSON), &paths))

	got, err := (&PathSet{}).ToBytes(paths, false)
	require.NoError(t, err)

	want, err := hex.DecodeString(expected)
	require.NoError(t, err)
	require.Equal(t, want, got)

	decoded, err := (&PathSet{}).ToJSON(bytes.NewBuffer(got), 0)
	require.NoError(t, err)
	require.Equal(t, paths, decoded)
}

// TestPathSetBoundaryByteBetweenPaths asserts that multi-path sets use
// typeBoundary=0xFF between paths and typeNone=0x00 after the last path.
//
// source: rippled include/xrpl/protocol/STPathSet.h:29,34 (typeNone/typeBoundary)
// source: rippled src/libxrpl/protocol/STPathSet.cpp:227 (s.add8(typeBoundary))
func TestPathSetBoundaryByteBetweenPaths(t *testing.T) {
	const pathsJSON = `[
		[{"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"}],
		[{"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"}]
	]`
	var paths any
	require.NoError(t, json.Unmarshal([]byte(pathsJSON), &paths))

	got, err := (&PathSet{}).ToBytes(paths, false)
	require.NoError(t, err)

	// Each path is 1 flag byte + 20 bytes account = 21 bytes, plus the
	// trailing boundary / end flag => 22 bytes per path.
	require.Equal(t, 44, len(got))
	require.Equal(t, byte(0xFF), got[21], "expected typeBoundary between paths")
	require.Equal(t, byte(0x00), got[43], "expected typeNone terminator")
}

// TestPathSetRejectEmpty asserts the encoder refuses a zero-length path set.
//
// source: rippled src/libxrpl/protocol/STPathSet.cpp path-set size constraints
func TestPathSetRejectEmpty(t *testing.T) {
	_, err := (&PathSet{}).ToBytes([]any{}, false)
	require.Error(t, err)
}

// TestPathSetRejectTooLarge asserts the encoder rejects >6 paths (xrpl path
// limit) and >8 steps per path (XRPL_MAX_PATH_LENGTH).
//
// source: rippled src/libxrpl/protocol/Protocol.cpp:26-27
//
//	maxPathCount = 6, maxPathLength = 8
func TestPathSetRejectTooLarge(t *testing.T) {
	step := map[string]any{"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"}

	// 7 paths > max 6.
	manyPaths := make([]any, 7)
	for i := range manyPaths {
		manyPaths[i] = []any{step}
	}
	_, err := (&PathSet{}).ToBytes(manyPaths, false)
	require.Error(t, err)

	// 9 steps > max 8 in one path.
	longPath := make([]any, 9)
	for i := range longPath {
		longPath[i] = step
	}
	_, err = (&PathSet{}).ToBytes([]any{longPath}, false)
	require.Error(t, err)
}

// TestPathSetAccountOnlyRejectExtraFields verifies that account cannot coexist
// with currency or issuer in the same step — XRPL path steps use either
// account-only or currency/issuer combinations.
//
// source: rippled src/libxrpl/protocol/STPathSet.cpp:179-184 (account XOR currency)
func TestPathSetAccountOnlyRejectExtraFields(t *testing.T) {
	bad := []any{[]any{map[string]any{
		"account":  "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA",
		"currency": "USD",
	}}}
	_, err := (&PathSet{}).ToBytes(bad, false)
	require.Error(t, err)
}

// TestPathSetUnknownFlag asserts the decoder rejects illegal step flags.
//
// source: rippled include/xrpl/protocol/STPathSet.h typeAll = 0x71 (account|currency|issuer|mpt)
func TestPathSetUnknownFlag(t *testing.T) {
	// 0x02 is not a valid path step flag.
	b := bytes.NewBuffer([]byte{0x02, 0x00})
	_, err := (&PathSet{}).ToJSON(b, 0)
	require.Error(t, err)
}
