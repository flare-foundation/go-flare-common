// Package safe implements the Safe (Gnosis Safe) v1.3.0 EIP-712 transaction
// encoding offline (encodeTransactionData), so a verifier can check that a
// threshold of known owners signed a specific execTransaction — e.g. the
// approveMachinePathList call a TEE governance Safe makes to approve a machine
// path list — without trusting any on-chain flags.
//
// It needs only the execTransaction ABI (selector + argument types), which is
// defined directly here rather than generated from an ABI/abigen: the Safe
// contract is external and only this one method is used. The definitions match
// Safe v1.3.0 (execTransaction selector 0x6a761202).
package safe

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// OperationCall is the Safe operation value for a plain CALL.
const OperationCall uint8 = 0

// SignatureLength is the length of one static owner-signature entry
// (r ‖ s ‖ v) in a Safe signatures blob.
const SignatureLength = 65

var (
	// domainSeparatorTypehash = keccak256("EIP712Domain(uint256 chainId,address verifyingContract)")
	domainSeparatorTypehash = crypto.Keccak256Hash([]byte("EIP712Domain(uint256 chainId,address verifyingContract)"))
	// safeTxTypehash = keccak256("SafeTx(address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)")
	safeTxTypehash = crypto.Keccak256Hash([]byte("SafeTx(address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)"))

	domainArgs abi.Arguments
	structArgs abi.Arguments

	// execTransactionSignature is the canonical Safe v1.3.0 execTransaction
	// signature; its keccak-256's first 4 bytes are ExecTransactionSelector
	// (0x6a761202).
	execTransactionSignature = "execTransaction(address,uint256,bytes,uint8,uint256,uint256,uint256,address,address,bytes)"

	// ExecTransactionSelector is the 4-byte selector of execTransaction.
	ExecTransactionSelector [4]byte

	// execTransactionArgs are the execTransaction argument types, used to
	// decode/encode the calldata of executed Safe transactions. The argument
	// names match ExecTransactionInputs fields so Unpack/Copy maps directly.
	execTransactionArgs abi.Arguments
)

func init() {
	copy(ExecTransactionSelector[:], crypto.Keccak256([]byte(execTransactionSignature))[:4])

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		panic(err)
	}
	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		panic(err)
	}
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}
	uint8Ty, err := abi.NewType("uint8", "", nil)
	if err != nil {
		panic(err)
	}
	bytesTy, err := abi.NewType("bytes", "", nil)
	if err != nil {
		panic(err)
	}

	// execTransaction(address to, uint256 value, bytes data, uint8 operation,
	// uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken,
	// address refundReceiver, bytes signatures)
	execTransactionArgs = abi.Arguments{
		{Name: "to", Type: addressTy},
		{Name: "value", Type: uint256Ty},
		{Name: "data", Type: bytesTy},
		{Name: "operation", Type: uint8Ty},
		{Name: "safeTxGas", Type: uint256Ty},
		{Name: "baseGas", Type: uint256Ty},
		{Name: "gasPrice", Type: uint256Ty},
		{Name: "gasToken", Type: addressTy},
		{Name: "refundReceiver", Type: addressTy},
		{Name: "signatures", Type: bytesTy},
	}

	// keccak256(abi.encode(DOMAIN_SEPARATOR_TYPEHASH, chainId, verifyingContract))
	domainArgs = abi.Arguments{{Type: bytes32Ty}, {Type: uint256Ty}, {Type: addressTy}}
	// keccak256(abi.encode(SAFE_TX_TYPEHASH, to, value, keccak256(data), operation,
	// safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce))
	structArgs = abi.Arguments{
		{Type: bytes32Ty}, {Type: addressTy}, {Type: uint256Ty}, {Type: bytes32Ty},
		{Type: uint8Ty}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: addressTy}, {Type: addressTy}, {Type: uint256Ty},
	}
}

// Tx holds the Safe transaction fields that feed the EIP-712 SafeTx struct
// hash. They mirror the execTransaction parameters plus the Safe nonce the
// transaction was executed at.
type Tx struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Nonce          *big.Int
}

// DomainSeparator returns the Safe's EIP-712 domain separator for the given
// chain: keccak256(abi.encode(DOMAIN_SEPARATOR_TYPEHASH, chainId, safe)).
func DomainSeparator(chainID uint64, safeAddress common.Address) (common.Hash, error) {
	enc, err := domainArgs.Pack(
		[32]byte(domainSeparatorTypehash),
		new(big.Int).SetUint64(chainID),
		safeAddress,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("packing domain separator: %w", err)
	}
	return crypto.Keccak256Hash(enc), nil
}

// Hash returns the Safe transaction hash the owners sign:
// keccak256(0x19 ‖ 0x01 ‖ domainSeparator ‖ safeTxStructHash).
func (t Tx) Hash(chainID uint64, safeAddress common.Address) (common.Hash, error) {
	domain, err := DomainSeparator(chainID, safeAddress)
	if err != nil {
		return common.Hash{}, err
	}

	enc, err := structArgs.Pack(
		[32]byte(safeTxTypehash),
		t.To,
		bigOrZero(t.Value),
		[32]byte(crypto.Keccak256Hash(t.Data)),
		t.Operation,
		bigOrZero(t.SafeTxGas),
		bigOrZero(t.BaseGas),
		bigOrZero(t.GasPrice),
		t.GasToken,
		t.RefundReceiver,
		bigOrZero(t.Nonce),
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("packing safe tx struct: %w", err)
	}
	structHash := crypto.Keccak256Hash(enc)

	return crypto.Keccak256Hash([]byte{0x19, 0x01}, domain[:], structHash[:]), nil
}

// RecoverSigners parses a Safe signatures blob and returns the addresses
// recovered from its ECDSA entries. The blob is a concatenation of 65-byte
// (r ‖ s ‖ v) entries; per the Safe signature encoding:
//   - v ∈ {27, 28}: an ECDSA signature over the Safe transaction hash;
//   - v ∈ {31, 32}: an eth_sign signature — ECDSA over the EIP-191 prefixed
//     Safe transaction hash, with v offset by 4;
//   - v == 1 (approved hash) and v == 0 (EIP-1271 contract signature) cannot
//     be verified from the blob alone and are skipped, as are entries that
//     fail recovery. Trailing dynamic bytes of contract signatures may be
//     misparsed as entries; they recover to arbitrary addresses that callers
//     filter against a known signer set, so they cannot create false
//     positives.
func RecoverSigners(txHash common.Hash, signatures []byte) []common.Address {
	n := len(signatures) / SignatureLength
	out := make([]common.Address, 0, n)
	for i := range n {
		entry := signatures[i*SignatureLength : (i+1)*SignatureLength]
		v := entry[64]

		var digest []byte
		var recID byte
		switch v {
		case 27, 28:
			digest = txHash[:]
			recID = v - 27
		case 31, 32:
			digest = accounts.TextHash(txHash[:])
			recID = v - 31
		default:
			continue
		}

		sig := make([]byte, SignatureLength)
		copy(sig, entry[:64])
		sig[64] = recID

		pubKey, err := crypto.SigToPub(digest, sig)
		if err != nil || pubKey == nil {
			continue
		}
		out = append(out, crypto.PubkeyToAddress(*pubKey))
	}
	return out
}

// ErrThresholdNotMet is returned by Approval.Verify when, at the given nonce,
// fewer than threshold distinct signer-set addresses are recovered.
var ErrThresholdNotMet = errors.New("safe approval signatures do not meet the owner threshold at the given nonce")

// countSigners returns how many distinct addresses recovered from the
// signatures blob (over txHash) belong to signerSet.
func countSigners(txHash common.Hash, signatures []byte, signerSet map[common.Address]struct{}) int {
	distinct := make(map[common.Address]struct{})
	for _, addr := range RecoverSigners(txHash, signatures) {
		if _, ok := signerSet[addr]; ok {
			distinct[addr] = struct{}{}
		}
	}
	return len(distinct)
}

// signerSetOf builds a lookup set from a signer slice.
func signerSetOf(signers []common.Address) map[common.Address]struct{} {
	set := make(map[common.Address]struct{}, len(signers))
	for _, s := range signers {
		set[s] = struct{}{}
	}
	return set
}

// Approval is the evidence for one executed Safe transaction: the
// execTransaction inputs (as decoded from calldata) plus the Safe nonce the
// transaction executed at. The nonce is not part of execTransaction calldata,
// so it is carried alongside — read from the MachinePathListApproved event's
// safeNonce field — so a downstream verifier need not scan. It carries no
// authorization weight: under a wrong nonce the owner signatures fail to
// recover to the signer set and Verify rejects.
type Approval struct {
	ExecTransaction ExecTransactionInputs `json:"execTransaction"`
	Nonce           uint64                `json:"nonce"`
}

// Verify checks that the approval's execTransaction, reconstructed at the
// recorded Nonce, was signed by at least threshold distinct addresses from
// signers. The nonce is taken as given (no scan). It validates only the owner
// signatures and threshold; callers are responsible for binding the
// execTransaction target/operation/calldata to the action being authorized.
func (a Approval) Verify(chainID uint64, safeAddress common.Address, signers []common.Address, threshold uint64) error {
	if threshold == 0 {
		return errors.New("threshold must be positive")
	}
	tx := a.ExecTransaction.SafeTx()
	tx.Nonce = new(big.Int).SetUint64(a.Nonce)
	txHash, err := tx.Hash(chainID, safeAddress)
	if err != nil {
		return err
	}
	if uint64(countSigners(txHash, a.ExecTransaction.Signatures, signerSetOf(signers))) < threshold {
		return ErrThresholdNotMet
	}
	return nil
}

// ExecTransactionInputs mirrors the execTransaction parameter list; field
// names match the ABI argument names so the argument set unpacks directly
// into the struct. It doubles as the wire form of a Safe approval: the full
// parameter set of an executed Safe transaction, including the owners'
// signatures blob, from which a verifier reconstructs the Safe transaction
// hash and recovers the owner signatures.
type ExecTransactionInputs struct {
	To             common.Address `json:"to"`
	Value          *big.Int       `json:"value"`
	Data           []byte         `json:"data"`
	Operation      uint8          `json:"operation"`
	SafeTxGas      *big.Int       `json:"safeTxGas"`
	BaseGas        *big.Int       `json:"baseGas"`
	GasPrice       *big.Int       `json:"gasPrice"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
	Signatures     []byte         `json:"signatures"`
}

// SafeTx maps the inputs onto the EIP-712 SafeTx fields. The Safe nonce is
// not part of execTransaction calldata, so Nonce is left unset; callers set it
// from the approval (which carries the event's safeNonce).
func (in ExecTransactionInputs) SafeTx() Tx {
	return Tx{
		To:             in.To,
		Value:          in.Value,
		Data:           in.Data,
		Operation:      in.Operation,
		SafeTxGas:      in.SafeTxGas,
		BaseGas:        in.BaseGas,
		GasPrice:       in.GasPrice,
		GasToken:       in.GasToken,
		RefundReceiver: in.RefundReceiver,
	}
}

// ErrInvalidExecTransactionCalldata is returned when calldata is too short or
// does not carry the execTransaction selector.
var ErrInvalidExecTransactionCalldata = errors.New("invalid execTransaction calldata")

// DecodeExecTransactionCalldata unpacks execTransaction calldata (selector
// included) into its inputs.
func DecodeExecTransactionCalldata(input []byte) (ExecTransactionInputs, error) {
	if len(input) < 4 || [4]byte(input[:4]) != ExecTransactionSelector {
		return ExecTransactionInputs{}, ErrInvalidExecTransactionCalldata
	}

	values, err := execTransactionArgs.Unpack(input[4:])
	if err != nil {
		return ExecTransactionInputs{}, err
	}

	var out ExecTransactionInputs
	if err := execTransactionArgs.Copy(&out, values); err != nil {
		return ExecTransactionInputs{}, err
	}
	return out, nil
}

// EncodeExecTransactionCalldata packs the inputs into execTransaction
// calldata (selector included) — the inverse of DecodeExecTransactionCalldata,
// used by tests and Safe clients.
func EncodeExecTransactionCalldata(in ExecTransactionInputs) ([]byte, error) {
	packed, err := execTransactionArgs.Pack(
		in.To, bigOrZero(in.Value), in.Data, in.Operation, bigOrZero(in.SafeTxGas),
		bigOrZero(in.BaseGas), bigOrZero(in.GasPrice), in.GasToken, in.RefundReceiver, in.Signatures,
	)
	if err != nil {
		return nil, err
	}
	return append(ExecTransactionSelector[:], packed...), nil
}

// bigOrZero returns b, or a zero big.Int when b is nil, so optional numeric
// fields encode as zero instead of panicking inside the ABI encoder.
func bigOrZero(b *big.Int) *big.Int {
	if b == nil {
		return new(big.Int)
	}
	return b
}
