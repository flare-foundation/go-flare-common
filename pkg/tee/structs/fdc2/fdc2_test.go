package fdc2

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

// TestEveryRegisteredTypeEncodes guards the one thing that can silently break
// this package: fdc2.abi and the generated Go struct drifting apart.
//
// Nothing enforces the correspondence at compile time. The ABI is data, the
// struct is code, and a field renamed in one and not the other still builds and
// still passes every test that does not encode. It then fails at Pack time,
// inside a verifier, on a request that looks perfectly ordinary — and since the
// signed FDC2 digest is taken over the ENCODED bodies, a type that cannot
// encode cannot produce a proof at all.
//
// That is not hypothetical: `packageHash` was renamed on
// IPMWUtxoProposalCheckRequestBody without the matching rename in fdc2.abi, and
// for weeks every attempt to encode that request body failed with "field
// proposalHash for tuple not found in the given struct".
//
// Packing a value of each registered type's request and response body is the
// cheapest statement that both directions still agree.
func TestEveryRegisteredTypeEncodes(t *testing.T) {
	// One entry per registered attestation type. A new type with no entry fails
	// the completeness check below rather than being quietly unguarded.
	bodies := map[AttestationType]struct{ request, response any }{
		AvailabilityCheck: {
			ITeeAvailabilityCheckRequestBody{},
			ITeeAvailabilityCheckResponseBody{},
		},
		PMWPaymentStatus: {
			IPMWPaymentStatusRequestBody{},
			IPMWPaymentStatusResponseBody{
				Amount:         big.NewInt(0),
				MaxFee:         big.NewInt(0),
				ReceivedAmount: big.NewInt(0),
				TransactionFee: big.NewInt(0),
			},
		},
		PMWMultisigAccountConfigured: {
			IPMWMultisigAccountConfiguredRequestBody{},
			IPMWMultisigAccountConfiguredResponseBody{},
		},
		PMWFeeProof: {
			IPMWFeeProofRequestBody{},
			IPMWFeeProofResponseBody{ActualFee: big.NewInt(0), EstimatedFee: big.NewInt(0)},
		},
		PMWMultisigUtxoConfigured: {
			IPMWMultisigUtxoConfiguredRequestBody{},
			IPMWMultisigUtxoConfiguredResponseBody{},
		},
		PMWUtxoProposalCheck: {
			IPMWUtxoProposalCheckRequestBody{},
			IPMWUtxoProposalCheckResponseBody{},
		},
		BtcDeposit: {
			IBtcDepositRequestBody{},
			IBtcDepositResponseBody{Amount: big.NewInt(0)},
		},
		BtcPayment: {
			IBtcPaymentRequestBody{},
			IBtcPaymentResponseBody{
				ReceivedAmountSats: big.NewInt(0),
				SpentAmountSats:    big.NewInt(0),
			},
		},
	}

	require.Len(t, bodies, len(attestationTypes),
		"a registered attestation type has no entry here — add one, or drift in it goes unnoticed")

	for _, at := range attestationTypes {
		body, ok := bodies[at]
		require.True(t, ok, "no bodies listed for %s", at)

		args, ok := AttestationTypeArguments[at]
		require.True(t, ok, "%s is not registered", at)

		t.Run(string(at), func(t *testing.T) {
			_, err := abi.Arguments{args.Request}.Pack(body.request)
			require.NoError(t, err, "the request body and fdc2.abi disagree")

			_, err = abi.Arguments{args.Response}.Pack(body.response)
			require.NoError(t, err, "the response body and fdc2.abi disagree")
		})
	}
}

// TestBtcDepositBodiesRoundTrip goes further than encoding for the deposit type
// specifically: every field is given a distinct value and required to survive
// the trip. Packing alone proves the shapes agree; this proves no field is
// silently dropped or transposed with its neighbour — which is what an addition
// in the wrong position looks like, since tuple encoding is positional and two
// adjacent fields of the same type swap without complaint.
func TestBtcDepositBodiesRoundTrip(t *testing.T) {
	args, ok := AttestationTypeArguments[BtcDeposit]
	require.True(t, ok, "BtcDeposit is not registered")

	t.Run("request", func(t *testing.T) {
		want := IBtcDepositRequestBody{
			WalletId:         [32]byte{1},
			AccountIndex:     2,
			DerivationIndex:  3,
			TransactionId:    [32]byte{4},
			OutputIndex:      5,
			MinConfirmations: 6,
			// 0xFFFF is "do not look at the inputs", which is what every caller
			// wanting only the deposit fact sends; a real index is the value the
			// round trip has to preserve.
			InputIndex: 7,
		}
		require.Equal(t, want, roundTrip[IBtcDepositRequestBody](t, args.Request, want))
	})

	t.Run("response", func(t *testing.T) {
		want := IBtcDepositResponseBody{
			Status:               0,
			ReceivingAddress:     "bcrt1qrecipient",
			ReceivingAddressHash: [32]byte{8},
			Amount:               big.NewInt(9),
			TransactionId:        [32]byte{10},
			BlockNumber:          11,
			BlockTimestamp:       12,
			Confirmations:        13,
			Memo:                 []byte{14, 15},
			InputAddress:         "bcrt1qpayer",
			InputAddressHash:     [32]byte{16},
			// Two distinct sighash bytes: the mixed-signature case that a single
			// reported byte would hide.
			InputSighashTypes: []byte{0x01, 0x02},
		}
		require.Equal(t, want, roundTrip[IBtcDepositResponseBody](t, args.Response, want))
	})
}

func roundTrip[T any](t *testing.T, arg abi.Argument, v T) T {
	t.Helper()

	packed, err := abi.Arguments{arg}.Pack(v)
	require.NoError(t, err, "packing failed — the struct has a field the ABI does not")

	unpacked, err := abi.Arguments{arg}.Unpack(packed)
	require.NoError(t, err)
	require.Len(t, unpacked, 1)

	// The decoder builds an anonymous struct from the ABI's own components, so
	// the conversion below succeeds only if that shape matches T field for
	// field — which is the drift being guarded against.
	out, ok := abi.ConvertType(unpacked[0], new(T)).(*T)
	require.True(t, ok, "the decoded value is not a %T", *new(T))
	return *out
}
