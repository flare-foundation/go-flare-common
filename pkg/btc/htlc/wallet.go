package htlc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/flare-foundation/go-flare-common/pkg/btc/address"
)

// Terms are the escrow half of a proposal: everything about an HTLC that the
// wallet does not already know.
//
// Deliberately not the multisig keys. Those come from the wallet's own
// provisioned xpubs at (Chain, Index), so a proposal cannot name a signer set —
// which is what keeps a signer from being an oracle. A proposal that lied about
// the keys would produce a script the wallet never funded, and the spend simply
// would not verify.
type Terms struct {
	PreimageHash    [32]byte
	CustodianPubKey []byte
	Timeout         int64
	Chain           address.Chain
	Index           uint32
}

// ForWallet builds an escrow's witness script and address from a wallet's
// PUBLISHED PARENT xpubs and the terms a proposal carries.
//
// This is the one function a signer, a verifier and a proposer all call, so
// that "the escrow at these terms" means one script to all three. Deriving it
// twice from prose is how the three drift.
//
// It takes PARENT keys and applies the account step itself, deliberately. The
// published key is at the wallet level and the account is a non-hardened child,
// so a caller holding parent keys and a caller holding account keys derive
// different addresses from the same terms — and both look correct in isolation.
// Taking the account index as an argument removes the choice.
func ForWallet(
	parentXpubs []string,
	accountIndex uint32,
	threshold int,
	terms Terms,
	params *chaincfg.Params,
) (witnessScript []byte, addr btcutil.Address, err error) {
	if threshold < 1 || threshold > len(parentXpubs) {
		return nil, nil, fmt.Errorf("threshold %d is not a k of %d", threshold, len(parentXpubs))
	}
	accountXpubs, err := address.DeriveAccountXpubs(parentXpubs, accountIndex, params)
	if err != nil {
		return nil, nil, fmt.Errorf("deriving account xpubs: %w", err)
	}
	// The same derivation the wallet's ordinary addresses use, so the timeout
	// branch is spendable by exactly the quorum that governs the wallet.
	_, _, sortedPubs, err := address.Derive(accountXpubs, threshold, terms.Chain, terms.Index, params)
	if err != nil {
		return nil, nil, fmt.Errorf("deriving the escrow's multisig keys: %w", err)
	}

	witnessScript, err = BuildScript(terms.PreimageHash, terms.CustodianPubKey, terms.Timeout, sortedPubs, threshold)
	if err != nil {
		return nil, nil, err
	}
	addr, err = Address(witnessScript, params)
	if err != nil {
		return nil, nil, err
	}
	return witnessScript, addr, nil
}
