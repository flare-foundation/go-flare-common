package types

import (
	"fmt"
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
