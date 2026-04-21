package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToInt64(t *testing.T) {
	tests := []struct {
		input  any
		output int64
	}{
		{10, 10}, {uint8(10), 10}, {uint16(10), 10}, {uint32(10), 10}, {uint64(10), 10}, {uint(10), 10},
		{-10, -10}, {int8(-10), -10}, {int32(-10), -10}, {int64(-10), -10}, {int(-10), -10},
		{10.0, 10}, {float32(10.0), 10}, {float64(10.0), 10},
	}

	for j, test := range tests {
		output, err := convertInt64(test.input, fmt.Sprintf("test %d", j))
		require.NoError(t, err)

		require.Equal(t, test.output, output)
	}
}

func TestToInt64Error(t *testing.T) {
	tests := []any{"1", []byte("neki"), 1.1, float32(1.1), float64(1.1), uint64(0xffffffffffffffff)}

	for j := range tests {
		_, err := convertInt64(tests[j], fmt.Sprintf("test %d", j))
		require.Error(t, err)
	}
}

func TestUintEncode(t *testing.T) {
	tests := []struct {
		t      Coder
		inputs []any
		output string
		err    bool
	}{
		{
			t:      &UInt8{},
			inputs: []any{uint(8), uint8(8), uint16(8), uint32(8), uint64(8), int(8), int8(8), int16(8), int32(8), int64(8), float32(8), float64(8)},
			output: "08",
			err:    false,
		},
		{
			t:      &UInt8{},
			inputs: []any{255},
			output: "ff",
			err:    false,
		},
		{
			t:      &UInt8{},
			inputs: []any{8.1, -8, 256, "10"},
			output: "",
			err:    true,
		},
		{
			t:      &UInt16{},
			inputs: []any{uint(8), uint8(8), uint16(8), uint32(8), uint64(8), int(8), int8(8), int16(8), int32(8), int64(8), float32(8), float64(8), "OfferCancel"},
			output: "0008",
			err:    false,
		},
		{
			t:      &UInt16{},
			inputs: []any{65535, 65535.0, uint16(65535)},
			output: "ffff",
			err:    false,
		},
		{
			t:      &UInt16{},
			inputs: []any{8.1, -8, 65536, "10"},
			output: "",
			err:    true,
		},
		{
			t:      &UInt32{},
			inputs: []any{uint(8), uint8(8), uint16(8), uint32(8), uint64(8), int(8), int8(8), int16(8), int32(8), int64(8), float32(8), float64(8)},
			output: "00000008",
			err:    false,
		},
		{
			t:      &UInt32{},
			inputs: []any{4294967295, 4294967295, 4294967295.0, uint32(4294967295)},
			output: "ffffffff",
			err:    false,
		},
		{
			t:      &UInt32{},
			inputs: []any{8.1, -8, 4294967296, "10"},
			output: "",
			err:    true,
		},
		{
			t:      &UInt64{},
			inputs: []any{uint(8), uint8(8), uint16(8), uint32(8), uint64(8), int(8), int8(8), int16(8), int32(8), int64(8), float32(8), float64(8), "08"},
			output: "0000000000000008",
			err:    false,
		},
		{
			t:      &UInt64{},
			inputs: []any{uint(10), uint8(10), uint16(10), uint32(10), uint64(10), int(10), int8(10), int16(10), int32(10), int64(10), float32(10), float64(10), "10"},
			output: "000000000000000a",
			err:    false,
		},
		{
			t:      &UInt64{},
			inputs: []any{uint64(0xffffffffffffffff)},
			output: "ffffffffffffffff",
			err:    false,
		},
		{
			t:      &UInt64{},
			inputs: []any{8.1, -8, "-10", "gg", "aa", "10,1", "10.0", "18446744073709551616"},
			output: "",
			err:    true,
		},
	}

	for _, test := range tests {
		var outputBytes []byte

		if !test.err {
			var err error
			outputBytes, err = hex.DecodeString(test.output)
			require.NoError(t, err)
		}

		for _, input := range test.inputs {
			output, err := test.t.ToBytes(input, false)

			if test.err {
				require.Error(t, err, input)
			} else {
				require.NoError(t, err)
				require.Equal(t, outputBytes, output, reflect.TypeOf(input))
			}
		}
	}
}

func TestUintFail(t *testing.T) {
	tests := []struct {
		t      Coder
		inputs []string
	}{
		{
			t:      &UInt8{},
			inputs: []string{""},
		},
		{
			t:      &UInt16{},
			inputs: []string{"", "00"},
		},
		{
			t:      &UInt32{},
			inputs: []string{"", "00", "000000"},
		},
		{
			t:      &UInt64{},
			inputs: []string{"", "00", "000000", "00000000000000"},
		},
	}

	for j, test := range tests {
		for _, input := range test.inputs {
			inputBytes, err := hex.DecodeString(input)
			require.NoError(t, err, j, input)
			b := bytes.NewBuffer(inputBytes)
			_, err = test.t.ToJSON(b, 0)
			require.Error(t, err, j, input)
		}
	}
}

func TestUint32(t *testing.T) {
	inputs := []uint32{0, 1, 1752791}

	for j, input := range inputs {
		serialized, err := (&UInt32{}).ToBytes(input, false)
		require.NoError(t, err, j)

		b := bytes.NewBuffer(serialized)
		deserialized, err := (&UInt32{}).ToJSON(b, 0)
		require.NoError(t, err, j)

		deserializedU, ok := deserialized.(uint32)
		require.True(t, ok, j)

		require.Equal(t, input, deserializedU)
	}
}

func TestUint64(t *testing.T) {
	inputs := []string{"10", "0", "1234567890"}

	for j, input := range inputs {
		serialized, err := (&UInt64{}).ToBytes(input, false)
		require.NoError(t, err, j)

		b := bytes.NewBuffer(serialized)
		deserialized, err := (&UInt64{}).ToJSON(b, 0)
		require.NoError(t, err, j)

		deserializedStr, ok := deserialized.(string)
		require.True(t, ok, j)

		require.Equal(t, input, deserializedStr)
	}
}

func TestUint16ToTxType(t *testing.T) {
	tests := []struct {
		output string
		input  uint16
	}{
		{"XChainModifyBridge", 47},
		{"AMMVote", 38},
		{"AMMCreate", 35},
		{"Batch", 71},
		{"LedgerStateFix", 53},
		{"NFTokenCreateOffer", 27},
		{"SetFee", 101},
		{"VaultDeposit", 68},
		{"XChainAccountCreateCommit", 44},
	}

	for _, test := range tests {
		output, err := Uint16ToTxType(test.input)
		require.NoError(t, err)
		require.Equal(t, test.output, output)
	}

	fails := []any{
		uint32(45),
		uint16(10000),
		"12",
	}

	for _, fail := range fails {
		_, err := Uint16ToTxType(fail)
		require.Error(t, err)
	}
}

// TestUInt16TxTypeMapping asserts TransactionType string values encode to the
// exact numeric transaction type codes enumerated by xrpl.js binary-codec.
// The canonical list lives in
// xrpl.js ripple-binary-codec/src/enums/definitions.json (TRANSACTION_TYPES)
// and is sourced from rippled include/xrpl/protocol/TxFormats.h.
//
// source: xrpl.js ripple-binary-codec/src/enums/definitions.json TRANSACTION_TYPES
func TestUInt16TxTypeMapping(t *testing.T) {
	tests := []struct {
		name    string
		code    uint16
		wantHex string
	}{
		{"Payment", 0, "0000"},
		{"EscrowCreate", 1, "0001"},
		{"EscrowFinish", 2, "0002"},
		{"AccountSet", 3, "0003"},
		{"EscrowCancel", 4, "0004"},
		{"OfferCreate", 7, "0007"},
		{"OfferCancel", 8, "0008"},
		{"TicketCreate", 10, "000A"},
		{"SignerListSet", 12, "000C"},
		{"NFTokenMint", 25, "0019"},
		{"NFTokenBurn", 26, "001A"},
		{"Clawback", 30, "001E"},
		{"AMMCreate", 35, "0023"},
		{"OracleDelete", 52, "0034"},
		{"Batch", 71, "0047"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := (&UInt16{}).ToBytes(tc.name, false)
			require.NoError(t, err)
			require.Equal(t, tc.wantHex, strings.ToUpper(hex.EncodeToString(got)))

			// Round-trip: decode back to uint16 then to name.
			decoded, err := (&UInt16{}).ToJSON(bytes.NewBuffer(got), 0)
			require.NoError(t, err)
			require.Equal(t, tc.code, decoded)

			name, err := Uint16ToTxType(decoded)
			require.NoError(t, err)
			require.Equal(t, tc.name, name)
		})
	}
}

// TestUInt16UnknownTxTypeString asserts that an unknown transaction-type
// string is rejected rather than silently mapped.
func TestUInt16UnknownTxTypeString(t *testing.T) {
	_, err := (&UInt16{}).ToBytes("NotARealTransactionType", false)
	require.Error(t, err)
}
