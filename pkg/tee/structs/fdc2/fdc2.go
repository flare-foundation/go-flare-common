//go:generate  abigen --abi=fdc2.abi --pkg=fdc2 --type=Fdc2 --out=autogen.go
package fdc2

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

const (
	reqStruct   string = "RequestBodyStruct"
	resStruct   string = "ResponseBodyStruct"
	proofStruct string = "ProofStruct"
)

var opCommands = []op.Command{
	op.Prove,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the type of message emitted with i-th element of opCommands.
var methods = []string{
	"fdc2AttestationRequestStruct",
}

var MessageArguments map[op.Command]abi.Argument

type AttestationType string

const (
	AvailabilityCheck            AttestationType = "TeeAvailabilityCheck"
	WalletPaymentStatus          AttestationType = "WalletPaymentStatus"
	NativeNonceAccountConfigured AttestationType = "NativeNonceAccountConfigured"
	// WalletSettlement proves a range of an account's instructions settled on
	// the source chain — firstSequencePosition and the count after it on a CSP
	// account — and reports two totals for it: what the confirmed transactions
	// cost the wallet (on Bitcoin the miner fee plus the 330-sat P2A output) and
	// what the confirmed attempts were authorised to spend, plus every position
	// that took more than one attempt. The diamond's verifyWalletSettlement
	// moves the account's confirmed position with it (SequenceConfirmed).
	WalletSettlement     AttestationType = "WalletSettlement"
	BtcAccountConfigured AttestationType = "BtcAccountConfigured"
	// CspProposalCheck is CSP's proposal predicate: the data providers
	// decide, before any signature exists, which of several competing
	// transactions fills a batch.
	CspProposalCheck AttestationType = "CspProposalCheck"
	// WalletDeposit proves that one OUTPUT of one confirmed transaction paid
	// receivingAddress — on Bitcoin an address a wallet derives at a reserved
	// index — in the asset tokenId names (empty: the chain's native asset, the
	// only one accepted so far). It is the deposit counterpart of
	// WalletPaymentStatus: that type proves a payment a protocol-managed wallet
	// MADE, this one proves a payment it RECEIVED, and the receiving address is
	// the identity — Bitcoin has no destination tag, so the index the address
	// derives from is what says who to credit.
	//
	// It also REPORTS one input, on request. `inputIndex` names an input and the
	// response carries that input's prevout address plus the distinct sighash
	// bytes on its signatures; 0xFFFFFFFF means "do not look" and leaves the input
	// fields empty (on XRPL both indices are 0). The sighash set is in the response rather than assumed
	// because only SIGHASH_ALL commits the signer to the outputs — a consumer
	// meaning "that address approved THIS transaction" must require exactly
	// [0x01], and an empty set (nothing parsable) must fail that test. There is
	// deliberately no 0x00 sentinel: 0x00 is a valid sighash byte under BIP-341.
	WalletDeposit AttestationType = "WalletDeposit"
	// BtcWalletAddress proves the Bitcoin address that
	// (walletId, accountIndex, derivationIndex) derives to, and is the ONLY
	// attestation type that reads wallet keys from a registry.
	//
	// The deposit type used to carry that triple and derive the address on every
	// deposit, which made a Bitcoin observation depend on a registry, an RPC and
	// a key read to answer a question whose answer never changes. An address
	// derives from keys and an index and neither moves, so it is proven once
	// here and stored; a deposit then names the address itself and reads no
	// Flare at all.
	//
	// The address comes back as a STRING so a depositor can read where to pay
	// from the contract that stored the binding, rather than deriving it
	// off-chain and hoping they matched.
	BtcWalletAddress AttestationType = "BtcWalletAddress"
	// WalletOperationStatus proves the outcome of one OPERATION of a wallet
	// account — a consolidation, an escrow create or an escrow reclaim on
	// Bitcoin, an escrow on XRPL — named by its operation id (the sequence
	// position on a Bitcoin CSP account, the payment id on XRPL): its kind,
	// the escrowed asset and amount, whether it was nullified, its transaction
	// and exact fee, and for an escrow create, given the transaction that spent
	// the escrow, how the escrow ended (rescued by the counterparty or
	// reclaimed by the wallet).
	WalletOperationStatus AttestationType = "WalletOperationStatus"
)

var attestationTypes = []AttestationType{
	AvailabilityCheck,
	WalletPaymentStatus,
	NativeNonceAccountConfigured,
	WalletSettlement,
	BtcAccountConfigured,
	CspProposalCheck,
	WalletDeposit,
	BtcWalletAddress,
	WalletOperationStatus,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the proof type of i-th attestation type.
var attestationTypeMethods = []string{
	"availabilityCheck",
	"walletPaymentStatus",
	"nativeNonceAccountConfigured",
	"walletSettlement",
	"btcAccountConfigured",
	"cspProposalCheck",
	"walletDeposit",
	"btcWalletAddress",
	"walletOperationStatus",
}

type AttestationArguments struct {
	Request  abi.Argument
	Response abi.Argument
	Proof    abi.Argument
}

var RequestHeaderArg abi.Argument
var ResponseHeaderArg abi.Argument
var AttestationRequestArg abi.Argument

var AttestationTypeArguments map[AttestationType]AttestationArguments

func init() {
	fdc2ABI, err := Fdc2MetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting tee data connector abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := fdc2ABI.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	if len(attestationTypes) != len(attestationTypeMethods) {
		panic("attestationTypes, attestationTypeMethods miss match")
	}

	AttestationTypeArguments = make(map[AttestationType]AttestationArguments)

	for j := range attestationTypes {
		request := attestationTypeMethods[j] + reqStruct
		method, ok := fdc2ABI.Methods[request]
		if !ok {
			panic(fmt.Sprintf("missing method %s", request))
		}
		requestABI := method.Inputs[0]

		response := attestationTypeMethods[j] + resStruct
		method, ok = fdc2ABI.Methods[response]
		if !ok {
			panic(fmt.Sprintf("missing method %s", response))
		}
		responseABI := method.Inputs[0]

		proof := attestationTypeMethods[j] + proofStruct
		method, ok = fdc2ABI.Methods[proof]
		if !ok {
			panic(fmt.Sprintf("missing method %s", proof))
		}
		proofABI := method.Inputs[0]

		AttestationTypeArguments[attestationTypes[j]] = AttestationArguments{
			Request:  requestABI,
			Response: responseABI,
			Proof:    proofABI,
		}
	}

	method, ok := fdc2ABI.Methods["fdc2RequestHeaderStruct"]
	if !ok {
		panic("missing method fdc2RequestHeaderStruct")
	}
	RequestHeaderArg = method.Inputs[0]

	method, ok = fdc2ABI.Methods["fdc2ResponseHeaderStruct"]
	if !ok {
		panic("missing method fdc2ResponseHeaderStruct")
	}
	ResponseHeaderArg = method.Inputs[0]

	method, ok = fdc2ABI.Methods["fdc2AttestationRequestStruct"]
	if !ok {
		panic("missing method fdc2AttestationRequestStruct")
	}
	AttestationRequestArg = method.Inputs[0]
}
