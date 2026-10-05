package fdc2requestfeeconfigurations_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	feev1 "github.com/flare-foundation/go-flare-common/pkg/contracts/fdc2/fdc2requestfeeconfigurations"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/v2/fdc2/fdc2requestfeeconfigurations"
)

// The v2 ABI renames _type, which must not change the encoding.
func TestPackMatchesV1(t *testing.T) {
	v1, err := feev1.Fdc2RequestFeeConfigurationsMetaData.GetAbi()
	require.NoError(t, err)
	f := fdc2requestfeeconfigurations.NewFdc2RequestFeeConfigurations()

	typ := [32]byte{0xaa}
	source := [32]byte{0xbb}
	fee := big.NewInt(500)

	tests := []struct {
		method string
		got    []byte
		args   []any
	}{
		{"getTypeAndSourceFee", f.PackGetTypeAndSourceFee(typ, source), []any{typ, source}},
		{"removeTypeAndSourceFee", f.PackRemoveTypeAndSourceFee(typ, source), []any{typ, source}},
		{"setTypeAndSourceFee", f.PackSetTypeAndSourceFee(typ, source, fee), []any{typ, source, fee}},
	}

	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			want, err := v1.Pack(tc.method, tc.args...)
			require.NoError(t, err)
			require.Equal(t, want, tc.got)
		})
	}
}
