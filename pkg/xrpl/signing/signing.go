// Package signing provides XRPL multi-signature validation and assembly.
package signing

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/ed25519"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/secp256k1"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/utils"
)

// ValidateMultiSig checks that signer is valid for the transaction tx.
func ValidateMultiSig(tx map[string]any, s *signer.Signer) (bool, error) {
	id, err := types.AccountID.ToBytes(s.Account, false)
	if err != nil {
		return false, fmt.Errorf("getting accountID: %w", err)
	}

	txBlob, err := encoding.Encode(tx, true)
	if err != nil {
		return false, fmt.Errorf("encoding tx: %w", err)
	}

	msg, err := utils.Prepare(txBlob, true, id)
	if err != nil {
		return false, fmt.Errorf("preparing message: %w", err)
	}

	if len(s.SigningPubKey) < 2 {
		return false, fmt.Errorf("signing pub key too short: %d chars", len(s.SigningPubKey))
	}
	pubPrefix := s.SigningPubKey[0:2]

	sigBytes, err := hex.DecodeString(s.TxnSignature)
	if err != nil {
		return false, fmt.Errorf("decoding signature: %w", err)
	}

	switch strings.ToLower(pubPrefix) {
	case "ed":
		return ed25519.Validate(msg, sigBytes, s.SigningPubKey)
	case "02", "03":
		return secp256k1.Validate(msg, sigBytes, s.SigningPubKey)
	default:
		return false, fmt.Errorf("invalid public key %s", s.SigningPubKey)
	}
}

// JoinMultisig appends signers to the transaction and serializes the result.
//
// Every signer is checked: signature is verified against the for-signing form
// of tx (M6), and the slice must be in canonical address order per rippled's
// STTx::checkMultiSign (M9). Returns an error if any signer is invalid or
// the order is wrong.
func JoinMultisig(tx map[string]any, signers []*signer.Signer) ([]byte, error) {
	if err := validateSignersForJoin(tx, signers); err != nil {
		return nil, err
	}

	signersArray := make([]any, len(signers))
	for j, s := range signers {
		signersArray[j] = s.Format()
	}

	tx["Signers"] = signersArray

	return encoding.Encode(tx, false)
}

// JoinMultisigJSON appends signers to the transaction map and returns it.
//
// Same validation contract as JoinMultisig: signatures are checked and the
// slice must be in canonical address order. Returns an error if either check
// fails.
func JoinMultisigJSON(tx map[string]any, signers []*signer.Signer) (map[string]any, error) {
	if err := validateSignersForJoin(tx, signers); err != nil {
		return nil, err
	}

	signersArray := make([]any, len(signers))
	for j, s := range signers {
		signersArray[j] = s.Format()
	}

	tx["Signers"] = signersArray
	return tx, nil
}

// validateSignersForJoin verifies every signer's signature against tx and
// asserts the slice is in strictly-ascending canonical address order, which
// rippled requires for the multi-sig blob to validate.
func validateSignersForJoin(tx map[string]any, signers []*signer.Signer) error {
	for j, s := range signers {
		if _, err := s.Value(); err != nil {
			return fmt.Errorf("signer %d (%s): computing canonical value: %w", j, s.Account, err)
		}
		if j > 0 && signer.Compare(signers[j-1], s) >= 0 {
			return fmt.Errorf("signer %d not in canonical ascending order (account %s)", j, s.Account)
		}
		ok, err := ValidateMultiSig(tx, s)
		if err != nil {
			return fmt.Errorf("signer %d (%s): %w", j, s.Account, err)
		}
		if !ok {
			return fmt.Errorf("signer %d (%s): invalid signature", j, s.Account)
		}
	}
	return nil
}
