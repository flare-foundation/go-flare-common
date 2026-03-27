package convert

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBigToUint64(t *testing.T) {
	tests := []struct {
		name      string
		input     *big.Int
		expected  uint64
		shouldErr bool
	}{
		{
			name:      "zero",
			input:     big.NewInt(0),
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "small positive",
			input:     big.NewInt(123),
			expected:  123,
			shouldErr: false,
		},
		{
			name:      "max uint64",
			input:     new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF),
			expected:  0xFFFFFFFFFFFFFFFF,
			shouldErr: false,
		},
		{
			name:      "negative number",
			input:     big.NewInt(-1),
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "overflow uint64",
			input:     new(big.Int).Add(new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF), big.NewInt(1)),
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "large positive overflow",
			input:     new(big.Int).Lsh(big.NewInt(1), 65), // 2^65
			expected:  0,
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BigToUint64Safe(test.input)

			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, result)
			}
		})
	}
}

func TestBigToUint32(t *testing.T) {
	tests := []struct {
		name      string
		input     *big.Int
		expected  uint32
		shouldErr bool
	}{
		{
			name:      "zero",
			input:     big.NewInt(0),
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "small positive",
			input:     big.NewInt(123),
			expected:  123,
			shouldErr: false,
		},
		{
			name:      "max uint32",
			input:     new(big.Int).SetUint64(0xFFFFFFFF),
			expected:  0xFFFFFFFF,
			shouldErr: false,
		},
		{
			name:      "negative number",
			input:     big.NewInt(-1),
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "overflow uint32",
			input:     new(big.Int).Add(new(big.Int).SetUint64(0xFFFFFFFF), big.NewInt(1)),
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "large positive overflow",
			input:     new(big.Int).Lsh(big.NewInt(1), 33), // 2^33
			expected:  0,
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BigToUint32Safe(test.input)

			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, result)
			}
		})
	}
}
