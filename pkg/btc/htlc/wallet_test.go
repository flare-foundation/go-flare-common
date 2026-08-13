package htlc_test

import (
	"testing"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/btc/address"
	"github.com/flare-foundation/go-flare-common/pkg/btc/htlc"
)

// parentXpubs is n published wallet-level keys, the form ForWallet takes.
func parentXpubs(t *testing.T, n int) []string {
	t.Helper()
	params := &chaincfg.RegressionNetParams
	out := make([]string, 0, n)
	for i := range n {
		seed := make([]byte, 32)
		for j := range seed {
			seed[j] = byte(i + 1)
		}
		master, err := hdkeychain.NewMaster(seed, params)
		require.NoError(t, err)
		pub, err := master.Neuter()
		require.NoError(t, err)
		out = append(out, pub.String())
	}
	return out
}

// TestForWalletAppliesTheAccountStep is the property the signature exists for.
//
// The published key is at the WALLET level and the account is a non-hardened
// child, so a caller holding parent keys and one holding account keys derive
// different addresses from identical terms — and both look correct alone. That
// is exactly how a proposer and a verifier came to disagree about where an
// escrow was: the proposal paid one address and the predicate looked at
// another, and the only symptom was "the transaction pays nothing to the
// escrow".
func TestForWalletAppliesTheAccountStep(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	parents := parentXpubs(t, 3)
	terms := htlc.Terms{
		PreimageHash:    [32]byte{0xAB},
		CustodianPubKey: makeKey(t),
		Timeout:         1_800_000_000,
		Chain:           address.External,
		Index:           32,
	}

	_, viaParents, err := htlc.ForWallet(parents, 0, 2, terms, params)
	require.NoError(t, err)

	// The same terms, but with the account step already applied by the caller
	// and then applied AGAIN inside — a different address, which is what a
	// mismatched caller would produce.
	accountLevel, err := address.DeriveAccountXpubs(parents, 0, params)
	require.NoError(t, err)
	_, viaAccount, err := htlc.ForWallet(accountLevel, 0, 2, terms, params)
	require.NoError(t, err)

	require.NotEqual(t, viaParents.EncodeAddress(), viaAccount.EncodeAddress(),
		"if these agreed, the account step would not be happening at all")

	// And a different account index is a different escrow, so the step is not
	// merely present but keyed on the index it is given.
	_, other, err := htlc.ForWallet(parents, 1, 2, terms, params)
	require.NoError(t, err)
	require.NotEqual(t, viaParents.EncodeAddress(), other.EncodeAddress())
}

func makeKey(t *testing.T) []byte {
	t.Helper()
	params := &chaincfg.RegressionNetParams
	seed := make([]byte, 32)
	for i := range seed {
		seed[i] = 0xC0
	}
	master, err := hdkeychain.NewMaster(seed, params)
	require.NoError(t, err)
	pub, err := master.ECPubKey()
	require.NoError(t, err)
	return pub.SerializeCompressed()
}
