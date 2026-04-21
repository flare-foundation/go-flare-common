package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVector256Single asserts a one-hash Vector256 payload equals the raw
// 32-byte hash concatenation.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json
//
//	accountState[2].json.Indexes = ["908D554A..."]
//
// source: rippled src/libxrpl/protocol/STVector256.cpp (Vector256 = packed hash256)
func TestVector256Single(t *testing.T) {
	const hashesJSON = `["908D554AA0D29F660716A3EE65C61DD886B744DDF60DE70E6B16EADB770635DB"]`
	const expected = "908D554AA0D29F660716A3EE65C61DD886B744DDF60DE70E6B16EADB770635DB"

	var hashes any
	require.NoError(t, json.Unmarshal([]byte(hashesJSON), &hashes))

	got, err := Vector256.ToBytes(hashes, false)
	require.NoError(t, err)

	want, err := hex.DecodeString(expected)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, 32, len(got))
}

// TestVector256Multi asserts two packed hashes land back-to-back with no
// separator and are recovered on decode.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/codec-fixtures.json
//
//	accountState[6].json.Indexes = ["A95EB289…", "52733E95…"]
func TestVector256Multi(t *testing.T) {
	const hashesJSON = `["A95EB2892EA15C8B7BCDAF6D1A8F1F21791192586EBD66B7DCBEC582BFAAA198",
		"52733E959FD0D25A72E188A26BC406768D91285883108AED061121408DAD4AF0"]`
	const expected = "A95EB2892EA15C8B7BCDAF6D1A8F1F21791192586EBD66B7DCBEC582BFAAA198" +
		"52733E959FD0D25A72E188A26BC406768D91285883108AED061121408DAD4AF0"

	var hashes []any
	require.NoError(t, json.Unmarshal([]byte(hashesJSON), &hashes))

	got, err := Vector256.ToBytes(hashes, false)
	require.NoError(t, err)

	want, err := hex.DecodeString(expected)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, 64, len(got))

	// Round-trip: length parameter is the total byte length.
	decoded, err := Vector256.ToJSON(bytes.NewBuffer(got), len(got))
	require.NoError(t, err)
	require.Equal(t, hashes, decoded)
}

// TestVector256EmptyRoundTrip verifies the zero-length vector path.
func TestVector256EmptyRoundTrip(t *testing.T) {
	got, err := Vector256.ToBytes([]any{}, false)
	require.NoError(t, err)
	require.Equal(t, 0, len(got))

	decoded, err := Vector256.ToJSON(bytes.NewBuffer(got), 0)
	require.NoError(t, err)
	require.Equal(t, []any{}, decoded)
}

// TestVector256InvalidLength asserts the decoder rejects a length that is not
// a multiple of 32 bytes.
//
// source: rippled src/libxrpl/protocol/STVector256.cpp (length must be multiple of 32)
func TestVector256InvalidLength(t *testing.T) {
	_, err := Vector256.ToJSON(bytes.NewBuffer(make([]byte, 31)), 31)
	require.Error(t, err)
}

// TestVector256InvalidInput rejects non-slice input and non-hex entries.
func TestVector256InvalidInput(t *testing.T) {
	_, err := Vector256.ToBytes("not a slice", false)
	require.Error(t, err)

	_, err = Vector256.ToBytes([]any{"not-hex"}, false)
	require.Error(t, err)

	// A hash of the wrong length is also invalid.
	_, err = Vector256.ToBytes([]any{"AABB"}, false)
	require.Error(t, err)
}
