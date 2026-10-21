package contracts_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bindv2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/contracts"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	relayv2 "github.com/flare-foundation/go-flare-common/pkg/contracts/v2/relay"
)

const testABI = `[
	{"type":"event","name":"Done","inputs":[]},
	{"type":"function","name":"run","inputs":[],"outputs":[]},
	{"type":"function","name":"run","inputs":[{"name":"n","type":"uint256"}],"outputs":[]},
	{"type":"error","name":"Failed","inputs":[]}
]`

func TestMustABI(t *testing.T) {
	tests := []struct {
		name string
		load func(json string) *abi.ABI
	}{
		{name: "v1", load: func(json string) *abi.ABI { return contracts.MustABI(&bind.MetaData{ABI: json}) }},
		{name: "v2", load: func(json string) *abi.ABI { return contracts.MustABI(&bindv2.MetaData{ABI: json}) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Contains(t, tc.load(testABI).Events, "Done")
			require.PanicsWithError(t, "loading ABI: invalid character 'b' looking for beginning of value", func() { tc.load("bad") })
		})
	}

	require.PanicsWithError(t, "nil metadata", func() { contracts.MustABI((*bind.MetaData)(nil)) })
	require.PanicsWithError(t, "nil metadata", func() { contracts.MustABI((*bindv2.MetaData)(nil)) })
}

// TestMustABIBindings reads one contract's ABI from both binding trees.
func TestMustABIBindings(t *testing.T) {
	v1 := contracts.MustABI(relay.RelayMetaData)
	require.Same(t, v1, contracts.MustABI(relay.RelayMetaData), "the binding's cached parse")

	v2 := contracts.MustABI(&relayv2.RelayMetaData)
	require.Equal(t, contracts.MustEvent(v1, "SigningPolicyInitialized").ID, contracts.MustEvent(v2, "SigningPolicyInitialized").ID)
}

func TestMustLookups(t *testing.T) {
	a := contracts.MustABI(&bind.MetaData{ABI: testABI})

	tests := []struct {
		name    string
		lookup  func(a *abi.ABI) string
		want    string
		wantErr string
	}{
		{name: "event", lookup: func(a *abi.ABI) string { return contracts.MustEvent(a, "Done").Sig }, want: "Done()"},
		{name: "method", lookup: func(a *abi.ABI) string { return contracts.MustMethod(a, "run").Sig }, want: "run()"},
		{name: "overloaded method", lookup: func(a *abi.ABI) string { return contracts.MustMethod(a, "run0").Sig }, want: "run(uint256)"},
		{name: "error", lookup: func(a *abi.ABI) string { return contracts.MustError(a, "Failed").Sig }, want: "Failed()"},
		{name: "missing event", lookup: func(a *abi.ABI) string { return contracts.MustEvent(a, "run").Sig }, wantErr: "event run not found in ABI"},
		{name: "missing method", lookup: func(a *abi.ABI) string { return contracts.MustMethod(a, "Done").Sig }, wantErr: "method Done not found in ABI"},
		{name: "missing error", lookup: func(a *abi.ABI) string { return contracts.MustError(a, "Done").Sig }, wantErr: "error Done not found in ABI"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr != "" {
				require.PanicsWithError(t, tc.wantErr, func() { tc.lookup(a) })
				return
			}
			require.Equal(t, tc.want, tc.lookup(a))
			require.PanicsWithError(t, "nil ABI", func() { tc.lookup(nil) })
		})
	}
}

func ExampleMustEvent() {
	relayABI := contracts.MustABI(relay.RelayMetaData)
	fmt.Println(contracts.MustEvent(relayABI, "SigningPolicyInitialized").Sig)
	// Output: SigningPolicyInitialized(uint24,uint32,uint16,uint256,address[],uint16[],bytes,uint64)
}
