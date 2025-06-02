package signing

import (
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/ed25519"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/secp256k1"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/utils"
)

// ValidateMultiSig checks that signer is valid for the transaction tx.
func ValidateMultiSig(tx map[string]any, s *utils.Signer) (bool, error) {
	id, err := types.AccountID.ToBytes(s.Account, false)
	if err != nil {
		return false, fmt.Errorf("cannot get accountID: %v", err)
	}

	txBlob, err := types.Encode(tx, true)
	if err != nil {
		return false, fmt.Errorf("cannot encode tx: %v", err)
	}

	msg := utils.Prepare(txBlob, true, id)

	pubPrefix := s.SigningPubKey[0:2]

	sigBytes, err := hex.DecodeString(s.TxnSignature)
	if err != nil {
		return false, fmt.Errorf("cannot decode signature: %v", err)
	}

	switch pubPrefix {
	case "ED", "ed":
		return ed25519.Validate(msg, sigBytes, s.SigningPubKey[2:])
	case "02", "03":
		return secp256k1.Validate(msg, sigBytes, s.SigningPubKey)
	default:
		return false, fmt.Errorf("invalid public key %s", s.SigningPubKey)
	}
}

// JoinMultisig appends signers to transactions and serializes it.
//
// If any of the signers is invalid, error is returned.
func JoinMultisig(tx map[string]any, signers []*utils.Signer) ([]byte, error) {
	slices.SortFunc(signers, utils.Compare)

	signersArray := make([]any, len(signers))

	for j, signer := range signers {
		ok, err := ValidateMultiSig(tx, signer)

		if err != nil {
			return nil, fmt.Errorf("invalid signer %v", signer)
		} else if !ok {
			return nil, fmt.Errorf("invalid signature  %v", signer)
		}

		signersArray[j] = signer.Format()
	}

	tx["Signers"] = signersArray

	return types.Encode(tx, false)
}
