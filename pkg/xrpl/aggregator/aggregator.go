// Package aggregator collects and validates XRPL multi-signatures for quorum-based transaction finalization.
package aggregator

import (
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
)

// Account represents XRPL multisig account.
//
// SignerList maps each authorized signer's account to its SignerWeight, mirroring
// rippled's on-ledger SignerEntries (Account, SignerWeight). Quorum is the threshold
// against the SUMMED weight of submitted signers — not their count — per
// rippled STTx::checkMultiSign. The same model degenerates to "N of M with weight 1"
// when every weight is set to 1.
//
// AddSignatures and Finalize are safe to call concurrently from multiple
// goroutines (mu serializes the internal state). The exported fields
// (Address, SignerList, Quorum) must be set once at construction and not
// mutated afterwards; the mutex does not cover them.
type Account struct {
	Address    string
	SignerList map[string]uint16
	Quorum     uint

	mu  sync.Mutex
	txs map[common.Hash]*transaction
}

// NewAccount constructs an Account ready for concurrent use. Prefer this
// over struct-literal initialization: it initializes the internal tx map,
// which AddSignatures otherwise lazy-allocates under the lock.
func NewAccount(address string, signerList map[string]uint16, quorum uint) *Account {
	return &Account{
		Address:    address,
		SignerList: signerList,
		Quorum:     quorum,
		txs:        make(map[common.Hash]*transaction),
	}
}

type transaction struct {
	transaction   map[string]any
	signers       map[string]*signer.Signer
	quorumReached bool
}

// AddSignatures accepts an encoded transaction with a Signers field.
// If the transaction's "Account" field matches the account, signers are
// extracted, validated, and stored. The returned common.Hash identifies the
// transaction and can be passed to Finalize once quorum is reached. The
// boolean is true the first time the quorum is reached.
//
// Audit M10: previously this returned *transaction whose useful field was
// unexported, so external callers could not finalize. The id is the public
// handle.
func (a *Account) AddSignatures(blob []byte) (common.Hash, bool, error) {
	txJSON, err := encoding.Decode(blob)
	if err != nil {
		return common.Hash{}, false, fmt.Errorf("invalid tx blob: %w", err)
	}

	en, err := encoding.Encode(txJSON, true)
	if err != nil {
		return common.Hash{}, false, fmt.Errorf("encoding: %w", err)
	}

	identifier := crypto.Keccak256Hash(en)

	a.mu.Lock()
	defer a.mu.Unlock()

	// Lazy-init for callers that used a struct literal without going through
	// NewAccount; protected by the mutex so the check-and-init is race-free.
	if a.txs == nil {
		a.txs = make(map[common.Hash]*transaction)
	}

	var txExists bool
	var tx *transaction
	tx, txExists = a.txs[identifier]
	if !txExists {
		add, exists := txJSON["Account"]
		if !exists {
			return common.Hash{}, false, errors.New("missing Account field")
		}
		addStr, ok := add.(string)
		if !ok {
			return common.Hash{}, false, fmt.Errorf("invalid Account, %v", add)
		}

		if a.Address != addStr {
			return common.Hash{}, false, fmt.Errorf("expected address %s got %s", a.Address, addStr)
		}

		signingKey, exists := txJSON["SigningPubKey"]
		if !exists {
			return common.Hash{}, false, errors.New("missing SigningPubKey field")
		}
		if signingKey != "" {
			return common.Hash{}, false, errors.New("signingPubKey should be empty string")
		}

		tx = new(transaction)
		tx.signers = make(map[string]*signer.Signer)

		decoded, err := encoding.Decode(en)
		if err != nil {
			return common.Hash{}, false, fmt.Errorf("decoding: %w", err)
		}
		tx.transaction = decoded
	}

	s, exists := txJSON["Signers"]
	if !exists {
		return common.Hash{}, false, errors.New("missing Signers field")
	}

	si, ok := s.([]any)
	if !ok {
		return common.Hash{}, false, fmt.Errorf("invalid Signers, %v", s)
	}

	somethingAdded := false
	for _, sig := range si {
		sMap, ok := sig.(map[string]any)
		if !ok {
			continue
		}

		s, err := signer.Parse(sMap)
		if err != nil {
			continue
		}

		_, ok = tx.signers[s.Account]
		if ok { // already added
			continue
		}

		_, ok = a.SignerList[s.Account]
		if !ok { // not among signers
			continue
		}

		ok, err = signing.ValidateMultiSig(tx.transaction, s)
		if err != nil || !ok {
			continue
		}

		addrFromPub, err := address.PubToAddress(s.SigningPubKey)
		if err != nil {
			continue
		}
		if addrFromPub != s.Account {
			continue
		}

		tx.signers[s.Account] = s
		somethingAdded = true
	}

	if !somethingAdded {
		return common.Hash{}, false, errors.New("no new valid signatures added")
	}

	a.txs[identifier] = tx // add transaction if not yet added

	qr := a.collectedWeight(tx) >= a.Quorum
	if qr {
		defer func() { tx.quorumReached = true }()
	}

	return identifier, qr && !tx.quorumReached, nil
}

// collectedWeight sums the SignerWeights of every signer that has submitted a
// valid signature for tx. Mirrors the weighted check in rippled STTx::checkMultiSign.
func (a *Account) collectedWeight(tx *transaction) uint {
	var sum uint
	for acct := range tx.signers {
		sum += uint(a.SignerList[acct])
	}
	return sum
}

// Finalize checks that enough signatures are collected for a transaction with id and returns an encoded transaction ready to be submitted.
func (a *Account) Finalize(id common.Hash) ([]byte, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	tx, ok := a.txs[id]
	if !ok {
		return nil, fmt.Errorf("no transaction with id %v to finalize", id)
	}

	if !tx.quorumReached {
		return nil, fmt.Errorf("quorum not yet reached for %v", id)
	}

	if len(tx.signers) == 0 {
		return nil, fmt.Errorf("no signers collected for %v", id)
	}

	blob, err := signing.JoinMultisig(tx.transaction, sortSigners(tx.signers))
	if err != nil {
		return nil, fmt.Errorf("joining signatures: %w", err)
	}

	return blob, nil
}

// sortSigners returns every collected signer ordered by the numeric value of
// their address.
//
// XRPL accepts any subset of authorized signers whose summed weight meets the
// account's Quorum (rippled STTx::checkMultiSign), so emitting all collected
// valid signers is always correct. Doing so also makes the produced blob
// deterministic: previously the helper truncated to int(Quorum) entries by
// iterating a Go map, which silently broke for weighted multisig (Quorum is a
// weight threshold, not a count — a single weight-2 signer satisfies Quorum=2
// but len(signers) == 1) and picked a non-deterministic subset when more
// signers had been collected than the count threshold demanded.
func sortSigners[T comparable](sig map[T]*signer.Signer) []*signer.Signer {
	out := make([]*signer.Signer, 0, len(sig))
	for _, s := range sig {
		out = append(out, s)
	}
	slices.SortFunc(out, signer.Compare)
	return out
}
