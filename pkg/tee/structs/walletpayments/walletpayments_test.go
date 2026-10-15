package walletpayments_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/walletpayments"
)

var pairs = []walletpayments.TeeIdKeyIdPair{{TeeId: common.Address{1}, KeyId: 2}}

// roundTrips encodes v under arg and decodes it back into its own type, which
// abicoder refuses unless the type's shape is the argument's.
func roundTrips[T any](t *testing.T, arg abi.Argument, v T) {
	t.Helper()
	b, err := abicoder.Encode(arg, v)
	require.NoError(t, err)
	got, err := abicoder.Decode[T](arg, b)
	require.NoError(t, err)
	assert.Equal(t, v, got)
}

// TestMessageArguments pins each command of the account-based model to the
// struct the contract emits with it.
func TestMessageArguments(t *testing.T) {
	require.Len(t, walletpayments.MessageArguments, 3)

	payment := walletpayments.IPaymentsPaymentInstructionMessage{
		WalletId: [32]byte{1}, TeeIdKeyIdPairs: pairs, RecipientAddress: "r",
		TokenId: []byte{}, Amount: big.NewInt(10), MaxFee: big.NewInt(1), FeeSchedule: []byte{}, PaymentId: 3,
	}
	roundTrips(t, walletpayments.MessageArguments[op.Pay], payment)
	roundTrips(t, walletpayments.MessageArguments[op.Reissue], payment)

	// A struct of another shape does not decode, so a wrong mapping fails here.
	b, err := abicoder.Encode(walletpayments.MessageArguments[op.Pay], payment)
	require.NoError(t, err)
	_, err = abicoder.Decode[walletpayments.IXrpEscrowsXrpEscrowCreateMessage](walletpayments.MessageArguments[op.Pay], b)
	require.Error(t, err)

	roundTrips(t, walletpayments.MessageArguments[op.EscrowCreate], walletpayments.IXrpEscrowsXrpEscrowCreateMessage{
		WalletId: [32]byte{1}, TeeIdKeyIdPairs: pairs, Destination: "d",
		Amount: big.NewInt(10), MaxFee: big.NewInt(1), CancelAfter: 100, PaymentId: 3,
	})
}

// TestBtcProposalCommitmentArgument pins the commitment a Bitcoin CSP proposal
// is bound by.
func TestBtcProposalCommitmentArgument(t *testing.T) {
	roundTrips(t, walletpayments.BtcProposalCommitmentArgument, walletpayments.IBtcAccountsBtcProposalCommitment{
		AnchorIndex: 1, Nonce: 2, NextAnchorVout: 3, Txid: [32]byte{4}, NextAnchorTxid: [32]byte{5},
	})
}

// TestBtcConfirmedAttemptCommitmentArgument pins the commitment a
// confirmed-attempt proposal is bound by.
func TestBtcConfirmedAttemptCommitmentArgument(t *testing.T) {
	roundTrips(t, walletpayments.BtcConfirmedAttemptCommitmentArgument, walletpayments.IBtcAccountsBtcConfirmedAttemptCommitment{
		Attempt: 1, Txid: [32]byte{2},
	})
}

// TestCspMessageArgument pins the message CSP settlement emits.
func TestCspMessageArgument(t *testing.T) {
	roundTrips(t, walletpayments.CspMessageArgument, walletpayments.ICspSettlementCspInstructionMessage{
		TeeIdKeyIdPairs: pairs, WalletId: [32]byte{1}, SequencePosition: 4, Attempt: 1,
		PackageHash: [32]byte{5}, ChainCommitmentHash: [32]byte{6}, Proposer: common.Address{7},
	})
}
