package aggregator

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
)

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

		// todo: check that pubKey matches address

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

func (a *Account) Finalize(id common.Hash) ([]byte, error) {
	tx, ok := a.txs[id]
	if !ok {
		return nil, fmt.Errorf("no transaction with id %v to finalize", id)
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
