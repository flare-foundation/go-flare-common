package defs

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFieldIDBranches verifies the four-way branch of the field ID encoding
// (type<16/name<16 single byte; type<16/name>=16; type>=16/name<16; type>=16/name>=16).
// Vectors sourced from the xrpl.js binary-codec data-driven tests, which are
// generated from rippled's SField definitions.
//
// source: xrpl.js ripple-binary-codec/test/fixtures/data-driven-tests.json fields_tests
// source: rippled src/libxrpl/protocol/Serializer.cpp:106-142 (Serializer::addFieldID)
func TestFieldIDBranches(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		expected  string
	}{
		// type<16, name<16 — common type / common name (1 byte)
		{"UInt16 LedgerEntryType t=1 n=1", "LedgerEntryType", "11"},
		{"UInt16 TransactionType t=1 n=2", "TransactionType", "12"},
		{"UInt64 IndexNext t=3 n=1", "IndexNext", "31"},
		{"Hash128 EmailHash t=4 n=1", "EmailHash", "41"},
		{"Hash256 LedgerHash t=5 n=1", "LedgerHash", "51"},
		{"Blob PublicKey t=7 n=1", "PublicKey", "71"},
		{"AccountID Account t=8 n=1", "Account", "81"},
		{"STObject ObjectEndMarker t=14 n=1", "ObjectEndMarker", "E1"},
		{"STArray ArrayEndMarker t=15 n=1", "ArrayEndMarker", "F1"},

		// type<16, name>=16 — common type / uncommon name (2 bytes)
		{"UInt32 HighQualityIn t=2 n=16", "HighQualityIn", "2010"},

		// type>=16, name<16 — uncommon type / common name (2 bytes)
		{"UInt8 CloseResolution t=16 n=1", "CloseResolution", "0110"},
		{"Hash160 TakerPaysCurrency t=17 n=1", "TakerPaysCurrency", "0111"},
		{"PathSet Paths t=18 n=1", "Paths", "0112"},
		{"Vector256 Indexes t=19 n=1", "Indexes", "0113"},

		// type>=16, name>=16 — uncommon type / uncommon name (3 bytes)
		{"UInt8 TickSize t=16 n=16", "TickSize", "001010"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			field, ok := NameToField[test.fieldName]
			require.True(t, ok, "field %s must exist in NameToField", test.fieldName)

			id, err := field.ID()
			require.NoError(t, err)

			expected, err := hex.DecodeString(test.expected)
			require.NoError(t, err)
			require.Equal(t, expected, id)

			// Round-trip through ReadID.
			pair, err := ReadID(bytes.NewBuffer(id))
			require.NoError(t, err)
			require.Equal(t, field.Nth, pair.F)
			require.Equal(t, field.Type, pair.T)
		})
	}
}

// TestObjectAndArrayEndMarkers pins the serialized values 0xE1 / 0xF1, which
// the codec relies on as STObject / STArray terminators.
//
// source: rippled include/xrpl/protocol/SField.h:59-60 (STI_OBJECT=14, STI_ARRAY=15)
// source: xrpl.js ripple-binary-codec/test/fixtures/data-driven-tests.json fields_tests
//
//	"ObjectEndMarker" -> "E1", "ArrayEndMarker" -> "F1"
func TestObjectAndArrayEndMarkers(t *testing.T) {
	for _, c := range []struct {
		name   string
		expect byte
	}{
		{"ObjectEndMarker", 0xE1},
		{"ArrayEndMarker", 0xF1},
	} {
		t.Run(c.name, func(t *testing.T) {
			field, ok := NameToField[c.name]
			require.True(t, ok)
			id, err := field.ID()
			require.NoError(t, err)
			require.Equal(t, []byte{c.expect}, id)
		})
	}
}

// TestFieldOrdering verifies Less sorts first by type code then by field code,
// matching the XRPL canonical serialization order.
//
// source: rippled src/libxrpl/protocol/STObject.cpp canonical field ordering
// (fields are emitted in ascending (type, field) order)
func TestFieldOrdering(t *testing.T) {
	// TransactionType (t=1 n=2) precedes Account (t=8 n=1) — different types.
	ttype := NameToField["TransactionType"]
	acct := NameToField["Account"]
	require.Equal(t, -1, ttype.Less(&acct))
	require.Equal(t, 1, acct.Less(&ttype))

	// Within UInt32 (t=2): Flags (n=2) precedes Sequence (n=4) precedes HighQualityIn (n=16).
	flags := NameToField["Flags"]
	seq := NameToField["Sequence"]
	hqi := NameToField["HighQualityIn"]
	require.Equal(t, -1, flags.Less(&seq))
	require.Equal(t, -1, seq.Less(&hqi))
	require.Equal(t, -1, flags.Less(&hqi))
	require.Equal(t, 1, hqi.Less(&flags))
}
