// Package aggregator collects and validates XRPL multi-signatures for quorum-based transaction finalization.
package aggregator

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
)

// Account represents XRPL multisig account.
type Account struct {
	Address    string
	SignerList map[string]bool
	Quorum     uint
	txs        map[common.Hash]*transaction
}

type transaction struct {
	transaction   map[string]any
	signers       map[string]*signer.Signer
	quorumReached bool
	id            common.Hash
}

// AddSignatures accepts an encoded transaction with a Signers field.
// If the transaction's "Account" field matches the account, signers are extracted, validated, and stored.
// The boolean is true the first time the quorum is reached.
func (a *Account) AddSignatures(blob []byte) (*transaction, bool, error) {
	txJSON, err := encoding.Decode(blob)
	if err != nil {
		return nil, false, fmt.Errorf("invalid tx bloc: %v", err)
	}

	en, err := encoding.Encode(txJSON, true)
	if err != nil {
		return nil, false, fmt.Errorf("encoding: %v", err)
	}

	identifier := crypto.Keccak256Hash(en)

	var txExists bool
	var tx *transaction
	tx, txExists = a.txs[identifier]
	if !txExists {
		add, exists := txJSON["Account"]
		if !exists {
			return nil, false, errors.New("missing Account field")
		}
		addStr, ok := add.(string)
		if !ok {
			return nil, false, fmt.Errorf("invalid Account, %v", add)
		}

		if a.Address != addStr {
			return nil, false, fmt.Errorf("expected address %s got %s", a.Address, addStr)
		}

		signingKey, exists := txJSON["SigningPubKey"]
		if !exists {
			return nil, false, errors.New("missing SigningPubKey field")
		}
		if signingKey != "" {
			return nil, false, errors.New("signingPubKey should be empty string")
		}

		tx = new(transaction)
		tx.signers = make(map[string]*signer.Signer)
		tx.id = identifier

		decoded, err := encoding.Decode(en)
		if err != nil {
			return nil, false, fmt.Errorf("decoding: %v", err)
		}
		tx.transaction = decoded
	}

	s, exists := txJSON["Signers"]
	if !exists {
		return nil, false, errors.New("missing Signers field")
	}

	si, ok := s.([]any)
	if !ok {
		return nil, false, fmt.Errorf("invalid Signers, %v", s)
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
		return nil, false, errors.New("no new valid signatures added")
	}

	a.txs[identifier] = tx // add transaction if not yet added

	qr := len(tx.signers) >= int(a.Quorum)
	if qr {
		defer func() { tx.quorumReached = true }()
	}

	return tx, qr && !tx.quorumReached, nil
}

// Finalize checks that enough signatures are collected for a transaction with id and returns an encoded transaction ready to be submitted.
func (a *Account) Finalize(id common.Hash) ([]byte, error) {
	tx, ok := a.txs[id]
	if !ok {
		return nil, fmt.Errorf("no transaction with id %v to finalize", id)
	}

	if !tx.quorumReached {
		return nil, fmt.Errorf("quorum not yet reached for %v", id)
	}

	s, err := sort(tx.signers, int(a.Quorum))
	if err != nil {
		return nil, err
	}

	blob, err := signing.JoinMultisig(tx.transaction, s)
	if err != nil {
		return nil, errors.New("joining signatures")
	}

	return blob, nil
}

// sort takes quorum of Signers and sorts according to the numerical value of their addresses.
func sort[T comparable](sig map[T]*signer.Signer, quorum int) ([]*signer.Signer, error) {
	out := make([]*signer.Signer, quorum)
	i := 0
	for j := range sig {
		out[i] = sig[j]
		i++
		if i == quorum {
			break
		}
	}
	if i != quorum {
		return nil, errors.New("quorum not reached")
	}

	slices.SortFunc(out, signer.Compare)

	return out, nil
}
