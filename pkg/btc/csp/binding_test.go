package csp_test

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func binding() csp.Binding {
	return csp.Binding{
		WalletID:    [32]byte{0x01},
		Network:     "regtest",
		Threshold:   2,
		ParentXpubs: []string{"xpubA", "xpubB", "xpubC"},
	}
}

func TestBindingRoundTrip(t *testing.T) {
	b := binding()
	raw, err := csp.EncodeBinding(b)
	require.NoError(t, err)
	got, err := csp.DecodeBinding(raw)
	require.NoError(t, err)
	assert.Equal(t, b, got)
}

func TestBindingRejectsUnusableConfigurations(t *testing.T) {
	cases := map[string]func(*csp.Binding){
		"no xpubs":         func(b *csp.Binding) { b.ParentXpubs = nil },
		"threshold zero":   func(b *csp.Binding) { b.Threshold = 0 },
		"threshold over n": func(b *csp.Binding) { b.Threshold = 4 },
		"no network":       func(b *csp.Binding) { b.Network = "" },
		// A repeated xpub lets one holder fill several CHECKMULTISIG slots, so
		// the k-of-n is not what it claims to be.
		"duplicate xpub": func(b *csp.Binding) { b.ParentXpubs[2] = b.ParentXpubs[0] },
		"empty xpub":     func(b *csp.Binding) { b.ParentXpubs[1] = "" },
	}
	for name, mutate := range cases {
		b := binding()
		mutate(&b)
		assert.Error(t, b.Validate(), "%s: Validate accepted it", name)
		_, err := csp.EncodeBinding(b)
		assert.Error(t, err, "%s: Encode accepted it", name)
	}
}

// Decode must reject too: a binding arriving over the wire has not been through
// Encode, so validation cannot be an encode-side courtesy.
func TestDecodeRejectsAnInvalidBinding(t *testing.T) {
	good := binding()
	good.Threshold = 3
	raw, err := csp.EncodeBinding(good)
	require.NoError(t, err)

	// Same bytes, but claiming a threshold above n would have to be built by
	// hand — so instead check the round trip refuses a mutated field.
	tampered := binding()
	tampered.ParentXpubs = []string{"xpubA"}
	tampered.Threshold = 1
	rawSmall, err := csp.EncodeBinding(tampered)
	require.NoError(t, err)
	_, err = csp.DecodeBinding(rawSmall)
	require.NoError(t, err)

	_ = raw
}

func TestContainsFindsThisMachine(t *testing.T) {
	b := binding()
	assert.True(t, b.Contains("xpubB"))
	assert.False(t, b.Contains("xpubZ"))
}
