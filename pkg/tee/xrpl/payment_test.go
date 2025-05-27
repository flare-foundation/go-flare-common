package xrpl

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payment"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/require"
)

func TestPaymentTransactionMultisig(t *testing.T) {
	// using xrpl.js
	expectedBlobStr := "120000240000000A61400000000000000A68400000000000000A73008114AA8114C65DA8EA1BE40849F974685289CC145CCF8314F667B0CA50CC7709A220B0561B85E53A48461FA8F9EA7D209C22FF5F21F0B81B113E63F7DB6DA94FEDEF11B2119B4088B89664FB9A3CB658E1F1"
	expectedBlob, err := hex.DecodeString(expectedBlobStr)
	require.NoError(t, err)

	instruction := payment.ITeePaymentsPaymentInstructionMessage{
		WalletId:         [32]byte{1},
		TeeIdKeyIdPairs:  []payment.TeeIdKeyIdPair{},
		SenderAddress:    "rGYYWKxT1XgNipUJouCq4cKiyAdq8xBoE9",
		RecipientAddress: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:           big.NewInt(10),
		PaymentReference: crypto.Keccak256Hash([]byte("test")),
		Nonce:            10,
		SubNonce:         0,
		Fee:              big.NewInt(10),
		BatchEndTs:       0,
	}

	tx := PaymentTransactionMultisig(instruction)

	err = CheckNativePayment(tx)
	require.NoError(t, err)

	blob, err := types.Encode(tx, true)
	require.NoError(t, err)
	require.Equal(t, expectedBlob, blob)
}
