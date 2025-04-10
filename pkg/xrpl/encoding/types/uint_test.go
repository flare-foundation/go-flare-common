package types

import (
	"encoding/hex"
	"fmt"
	"reflect"
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

func TestUint(t *testing.T) {
	tests := []struct {
		t      Encoder
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
