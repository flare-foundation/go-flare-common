package fdc2

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
// ICspProposalCheckRequestBody without the matching rename in fdc2.abi, and
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
		WalletPaymentStatus: {
			IWalletPaymentStatusRequestBody{},
			IWalletPaymentStatusResponseBody{
				Amount:         big.NewInt(0),
				MaxFee:         big.NewInt(0),
				ReceivedAmount: big.NewInt(0),
				TransactionFee: big.NewInt(0),
			},
		},
		NativeNonceAccountConfigured: {
			INativeNonceAccountConfiguredRequestBody{},
			INativeNonceAccountConfiguredResponseBody{},
		},
		WalletSettlement: {
			IWalletSettlementRequestBody{},
			IWalletSettlementResponseBody{
				ActualFee:     big.NewInt(0),
				AuthorisedFee: big.NewInt(0),
			},
		},
		BtcAccountConfigured: {
			IBtcAccountConfiguredRequestBody{},
			IBtcAccountConfiguredResponseBody{},
		},
		CspProposalCheck: {
			ICspProposalCheckRequestBody{},
			ICspProposalCheckResponseBody{},
		},
		WalletDeposit: {
			IWalletDepositRequestBody{},
			IWalletDepositResponseBody{Amount: big.NewInt(0)},
		},
		BtcWalletAddress: {
			IBtcWalletAddressRequestBody{},
			IBtcWalletAddressResponseBody{},
		},
		WalletOperationStatus: {
			IWalletOperationStatusRequestBody{},
			IWalletOperationStatusResponseBody{Amount: big.NewInt(0), TransactionFee: big.NewInt(0)},
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

// TestWalletDepositBodiesRoundTrip goes further than encoding for the deposit type
// specifically: every field is given a distinct value and required to survive
// the trip. Packing alone proves the shapes agree; this proves no field is
// silently dropped or transposed with its neighbour — which is what an addition
// in the wrong position looks like, since tuple encoding is positional and two
// adjacent fields of the same type swap without complaint.
func TestWalletDepositBodiesRoundTrip(t *testing.T) {
	args, ok := AttestationTypeArguments[WalletDeposit]
	require.True(t, ok, "WalletDeposit is not registered")

	t.Run("request", func(t *testing.T) {
		want := IWalletDepositRequestBody{
			TransactionId: [32]byte{4},
			// Both indices are 32 bits, as an outpoint's is on the wire: a value
			// past 65535 has to survive the round trip.
			OutputIndex:      70_000,
			ReceivingAddress: "bcrt1qdeposit",
			// Empty is the chain's native asset, the only one accepted so far; a
			// non-empty value must still survive, since it is what a proof names.
			TokenId:          []byte{0xab},
			MinConfirmations: 6,
			// 0xFFFFFFFF is "do not look at the inputs", which is what every caller
			// wanting only the deposit fact sends; a real index is the value the
			// round trip has to preserve.
			InputIndex: 65_536,
		}
		require.Equal(t, want, roundTrip[IWalletDepositRequestBody](t, args.Request, want))
	})

	t.Run("response", func(t *testing.T) {
		want := IWalletDepositResponseBody{
			Amount:           big.NewInt(9),
			Memo:             []byte{14, 15},
			BlockNumber:      11,
			BlockTimestamp:   12,
			InputAddress:     "bcrt1qpayer",
			InputAddressHash: [32]byte{16},
			// Two distinct sighash bytes: the mixed-signature case that a single
			// reported byte would hide.
			InputSighashTypes: []byte{0x01, 0x02},
		}
		require.Equal(t, want, roundTrip[IWalletDepositResponseBody](t, args.Response, want))
	})
}

// TestBodiesMatchTheSolidityEncoding pins the bodies whose shape changed with
// flare-smart-contracts-v2 csp-utxo-payments e4d6d26c (and 728b30bf) to bytes produced
// outside this package: each golden is `cast abi-encode` over a tuple type
// written by hand from the Solidity struct (userInterfaces/fdc2/I<Type>.sol),
// not from fdc2.abi. So a field the ABI file and the Go struct both lost, or
// both reordered, still fails here — the round trip above cannot see that,
// because both of its ends are read from the same ABI. The FDC2 digest is
// taken over exactly these bytes.
//
// The goldens are reproducible with
//
//	b(){ printf '0x%s' $(printf "$1%.0s" $(seq 32)); }   # b 11 = 0x1111…11
//	cast abi-encode 'f((bytes32,string,uint64,uint64))' "($(b 11),bc1qsender,7,3)"
//	cast abi-encode 'f((uint256,uint256,(uint64,uint32)[]))' "(12345,67890,[(8,2),(9,3)])"
//	cast abi-encode 'f((bytes32,string,uint64,bytes32))' "($(b 22),bc1qvault,42,$(b 33))"
//	cast abi-encode 'f((uint8,bytes,uint256,bool,uint256,bytes32,uint64,uint64,uint8))' \
//	    "(1,0xab,500000,true,1234,$(b 44),100,1700000000,2)"
//	cast abi-encode 'f((string,uint64,uint32,uint64,bytes32))' "(bc1qacct,5,2,9,$(b 55))"
//	cast abi-encode 'f((bytes32,string,uint64))' "($(b 66),bc1qpmw,77)"
//	cast abi-encode 'f((bytes32,uint32,string,bytes,uint16,uint32))' "($(b 04),70000,bcrt1qdeposit,0x,6,65536)"
func TestBodiesMatchTheSolidityEncoding(t *testing.T) {
	b := func(v byte) (out [32]byte) {
		for i := range out {
			out[i] = v
		}
		return out
	}

	cases := []struct {
		name   string
		typ    AttestationType
		arg    func(AttestationArguments) abi.Argument
		body   any
		golden string
	}{
		{
			name: "WalletSettlement request", typ: WalletSettlement, arg: request,
			body: IWalletSettlementRequestBody{
				OpType: b(0x11), SenderAddress: "bc1qsender", FirstSequencePosition: 7, Count: 3,
			},
			golden: "0x00000000000000000000000000000000000000000000000000000000000000201111111111111111111111111111111111111111111111111111111111111111000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000a6263317173656e64657200000000000000000000000000000000000000000000",
		},
		{
			name: "WalletSettlement response", typ: WalletSettlement, arg: response,
			body: IWalletSettlementResponseBody{
				ActualFee: big.NewInt(12345), AuthorisedFee: big.NewInt(67890),
				MultiAttemptPositions: []IWalletSettlementPositionAttempts{
					{SequencePosition: 8, AttemptCount: 2},
					{SequencePosition: 9, AttemptCount: 3},
				},
			},
			golden: "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000010932000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000003",
		},
		{
			name: "WalletOperationStatus request", typ: WalletOperationStatus, arg: request,
			body: IWalletOperationStatusRequestBody{
				OpType: b(0x22), SenderAddress: "bc1qvault", OperationId: 42, ResolutionTransactionId: b(0x33),
			},
			golden: "0x000000000000000000000000000000000000000000000000000000000000002022222222222222222222222222222222222222222222222222222222222222220000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000002a33333333333333333333333333333333333333333333333333333333333333330000000000000000000000000000000000000000000000000000000000000009626331717661756c740000000000000000000000000000000000000000000000",
		},
		{
			name: "WalletOperationStatus response", typ: WalletOperationStatus, arg: response,
			body: IWalletOperationStatusResponseBody{
				Kind:           1, // ESCROW_CREATE
				TokenId:        []byte{0xab},
				Amount:         big.NewInt(500_000),
				Nullified:      true,
				TransactionFee: big.NewInt(1234),
				TransactionId:  b(0x44),
				BlockNumber:    100,
				BlockTimestamp: 1_700_000_000,
				// Only the outcome: the owners' 728b30bf dropped the resolution
				// block and time, as the resolution transaction is not the wallet's.
				Resolution: 2, // ESCROW_RECLAIMED
			},
			golden: "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000007a120000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000004d244444444444444444444444444444444444444444444444444444444444444440000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000006553f10000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001ab00000000000000000000000000000000000000000000000000000000000000",
		},
		{
			name: "CspProposalCheck request", typ: CspProposalCheck, arg: request,
			body: ICspProposalCheckRequestBody{
				SenderAddress: "bc1qacct", SequencePosition: 5, Attempt: 2, EligibleGeneration: 9, PackageHash: b(0x55),
			},
			golden: "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000009555555555555555555555555555555555555555555555555555555555555555500000000000000000000000000000000000000000000000000000000000000086263317161636374000000000000000000000000000000000000000000000000",
		},
		{
			name: "WalletPaymentStatus request", typ: WalletPaymentStatus, arg: request,
			body:   IWalletPaymentStatusRequestBody{OpType: b(0x66), SenderAddress: "bc1qpmw", PaymentId: 77},
			golden: "0x000000000000000000000000000000000000000000000000000000000000002066666666666666666666666666666666666666666666666666666666666666660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000004d000000000000000000000000000000000000000000000000000000000000000762633171706d7700000000000000000000000000000000000000000000000000",
		},
		{
			name: "WalletDeposit request", typ: WalletDeposit, arg: request,
			body: IWalletDepositRequestBody{
				TransactionId: b(0x04), OutputIndex: 70_000, ReceivingAddress: "bcrt1qdeposit",
				TokenId: []byte{}, MinConfirmations: 6, InputIndex: 65_536,
			},
			golden: "0x00000000000000000000000000000000000000000000000000000000000000200404040404040404040404040404040404040404040404040404040404040404000000000000000000000000000000000000000000000000000000000001117000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000d6263727431716465706f736974000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			args, ok := AttestationTypeArguments[c.typ]
			require.True(t, ok, "%s is not registered", c.typ)
			packed, err := abi.Arguments{c.arg(args)}.Pack(c.body)
			require.NoError(t, err)
			require.Equal(t, c.golden, hexutil.Encode(packed),
				"the Go encoding differs from the Solidity struct's — the digest would not verify")
		})
	}
}

func request(a AttestationArguments) abi.Argument  { return a.Request }
func response(a AttestationArguments) abi.Argument { return a.Response }

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
