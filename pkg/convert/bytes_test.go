package convert

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBytesToUint64(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  uint64
		shouldErr bool
	}{
		{
			name:      "empty slice",
			input:     []byte{},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "prefixed zero byte",
			input:     []byte{0x00},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "single byte",
			input:     []byte{0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "single byte with prefix zero",
			input:     []byte{0x00, 0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "two bytes",
			input:     []byte{0x01, 0x02},
			expected:  0x0102,
			shouldErr: false,
		},
		{
			name:      "four bytes",
			input:     []byte{0x01, 0x02, 0x03, 0x04},
			expected:  0x01020304,
			shouldErr: false,
		},
		{
			name:      "eight bytes",
			input:     []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			expected:  0x0102030405060708,
			shouldErr: false,
		},
		{
			name:      "max uint64",
			input:     []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			expected:  0xFFFFFFFFFFFFFFFF, // max uint64
			shouldErr: false,
		},
		{
			name:      "too many bytes",
			input:     []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "too many zero bytes",
			input:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected:  0,
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BytesToUint64(test.input)

			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, test.expected, result)
			}
		})
	}
}

func TestBytesToUint32(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  uint32
		shouldErr bool
	}{
		{
			name:      "empty slice",
			input:     []byte{},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "prefixed zero byte",
			input:     []byte{0x00},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "single byte",
			input:     []byte{0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "single byte with prefix zero",
			input:     []byte{0x00, 0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "two bytes",
			input:     []byte{0x01, 0x02},
			expected:  0x0102,
			shouldErr: false,
		},
		{
			name:      "four bytes",
			input:     []byte{0x01, 0x02, 0x03, 0x04},
			expected:  0x01020304,
			shouldErr: false,
		},
		{
			name:      "max uint32",
			input:     []byte{0xFF, 0xFF, 0xFF, 0xFF},
			expected:  0xFFFFFFFF,
			shouldErr: false,
		},
		{
			name:      "too many bytes",
			input:     []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "too many zero bytes",
			input:     []byte{0x00, 0x00, 0x00, 0x00, 0x00},
			expected:  0,
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BytesToUint32(test.input)

			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, result)
			}
		})
	}
}

func TestBytesToUint16(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  uint16
		shouldErr bool
	}{
		{
			name:      "empty slice",
			input:     []byte{},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "prefixed zero byte",
			input:     []byte{0x00},
			expected:  0,
			shouldErr: false,
		},
		{
			name:      "single byte",
			input:     []byte{0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "single byte with prefix zero",
			input:     []byte{0x00, 0x01},
			expected:  1,
			shouldErr: false,
		},
		{
			name:      "two bytes",
			input:     []byte{0x01, 0x02},
			expected:  0x0102,
			shouldErr: false,
		},
		{
			name:      "max uint16",
			input:     []byte{0xFF, 0xFF},
			expected:  0xFFFF,
			shouldErr: false,
		},
		{
			name:      "too many bytes",
			input:     []byte{0x01, 0x02, 0x03},
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "too many zero bytes",
			input:     []byte{0x00, 0x00, 0x00},
			expected:  0,
			shouldErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BytesToUint16(test.input)
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, result)
			}
		})
	}
}

func TestUint64ToHash(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want common.Hash
	}{
		{
			name: "max_uint64",
			n:    uint64(0xFFFFFFFFFFFFFFFF),
			want: common.HexToHash("0x000000000000000000000000ffffffffffffffff"),
		},
		{
			name: "min_uint64",
			n:    uint64(0),
			want: common.HexToHash("0x0000000000000000000000000000000000000000"),
		},
		{
			name: "value1",
			n:    uint64(1),
			want: common.HexToHash("0x0000000000000000000000000000000000000001"),
		},
		{
			name: "value2",
			n:    uint64(0x12345678),
			want: common.HexToHash("0x0000000000000000000000000000000012345678"),
		},
		{
			name: "value2",
			n:    uint64(0x1234567890123456),
			want: common.HexToHash("0x0000000000000000000000001234567890123456"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Uint64ToHash(test.n)
			require.Equal(t, test.want, got)
		})
	}
}

func TestUint32ToHash(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want common.Hash
	}{
		{
			name: "max_uint32",
			n:    uint32(0xFFFFFFFF),
			want: common.HexToHash("0x000000000000000000000000ffffffff"),
		},
		{
			name: "min_uint32",
			n:    uint32(0),
			want: common.HexToHash("0x00000000000000000000000000000000"),
		},
		{
			name: "value1",
			n:    uint32(1),
			want: common.HexToHash("0x00000000000000000000000000000001"),
		},
		{
			name: "value2",
			n:    uint32(0x12345678),
			want: common.HexToHash("0x00000000000000000000000012345678"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Uint32ToHash(test.n)
			require.Equal(t, test.want, got)
		})
	}
}

func TestUint16ToHash(t *testing.T) {
	tests := []struct {
		name string
		n    uint16
		want common.Hash
	}{
		{
			name: "max_uint16",
			n:    uint16(0xFFFF),
			want: common.HexToHash("0x000000000000000000000000ffff"),
		},
		{
			name: "min_uint16",
			n:    uint16(0),
			want: common.HexToHash("0x0000000000000000000000000000"),
		},
		{
			name: "value1",
			n:    uint16(1),
			want: common.HexToHash("0x0000000000000000000000000001"),
		},
		{
			name: "value2",
			n:    uint16(0x1234),
			want: common.HexToHash("0x0000000000000000000000001234"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Uint16ToHash(test.n)
			require.Equal(t, test.want, got)
		})
	}
}
