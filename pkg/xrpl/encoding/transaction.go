package encoding

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"encoding/hex"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

func SetMultisig(prv *ecdsa.PrivateKey, sequence int, signer string) (string, error) {
	tx := map[string]any{
		"TransactionType": "SignerListSet",
		"Fee":             "120000",
		"SignerQuorum":    1,
		"SignerEntries": []map[string]map[string]any{
			{"SignerEntry": {
				"Account":      signer,
				"SignerWeight": 1,
			},
			},
		},
	}

	signed, err := SignTransactionSecpSingle(tx, uint32(sequence), prv)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signed), err
}

func PaymentMultisig(prv ed25519.PrivateKey, account string, sequence int) (string, error) {
	tx := map[string]any{
		"Account":         account,
		"Destination":     "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S",
		"TransactionType": "Payment",
		"Amount":          "19",
		"Fee":             "120000",
		"SigningPubKey":   "",
		"Sequence":        sequence,
	}

	signed, err := SignTransactionEDMultisig(tx, prv)
	if err != nil {
		return "", err
	}

	tx["Signers"] = []map[string]map[string]any{
		{"Signer": {
			"Account":       signed.Account,
			"TxnSignature":  signed.TxnSignature,
			"SigningPubKey": signed.SigningPubKey,
		}},
	}

	encoded, err := types.Encode(tx, false)

	return hex.EncodeToString(encoded), err
}
