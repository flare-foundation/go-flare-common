// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletpayments

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBtcAccountsBtcAnchor is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcAnchor struct {
	GenesisTxid [32]byte
	GenesisVout uint32
	NextNonce   uint64
}

// IBtcAccountsBtcProposalCommitment is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcProposalCommitment struct {
	AnchorIndex    uint32
	Nonce          uint64
	NextAnchorVout uint32
	Txid           [32]byte
	NextAnchorTxid [32]byte
}

// IBtcEscrowsBtcEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type IBtcEscrowsBtcEscrowTerms struct {
	PreimageHash       [32]byte
	Amount             uint64
	ExpiresAt          uint64
	CounterpartyPubKey []byte
}

// ICspSettlementCspInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type ICspSettlementCspInstructionMessage struct {
	TeeIdKeyIdPairs     []TeeIdKeyIdPair
	WalletId            [32]byte
	SourceId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	FromPaymentId       uint64
	ToPaymentId         uint64
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	Proposer            common.Address
}

// ICspSettlementCspPaymentRecord is an auto generated low-level Go binding around an user-defined struct.
type ICspSettlementCspPaymentRecord struct {
	WalletId         [32]byte
	SourceId         [32]byte
	AccountAddress   string
	AccountIndex     uint32
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
	PaymentId        uint64
	BatchPaymentId   uint64
}

// IPaymentsPaymentInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type IPaymentsPaymentInstructionMessage struct {
	WalletId         [32]byte
	TeeIdKeyIdPairs  []TeeIdKeyIdPair
	SourceId         [32]byte
	SenderAddress    string
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	FeeSchedule      []byte
	PaymentReference [32]byte
	Nonce            uint64
	PaymentId        uint64
}

// IXrpEscrowsXrpEscrowCreateMessage is an auto generated low-level Go binding around an user-defined struct.
type IXrpEscrowsXrpEscrowCreateMessage struct {
	WalletId        [32]byte
	TeeIdKeyIdPairs []TeeIdKeyIdPair
	SourceId        [32]byte
	SenderAddress   string
	Destination     string
	Amount          *big.Int
	PreimageHash    [32]byte
	CancelAfter     uint64
	MaxFee          *big.Int
	Nonce           uint64
	PaymentId       uint64
	Nullified       bool
}

// IXrpEscrowsXrpEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type IXrpEscrowsXrpEscrowTerms struct {
	PreimageHash [32]byte
	Amount       *big.Int
	ExpiresAt    uint64
	Destination  string
	Nullified    bool
}

// PaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type PaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
}

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// WalletAccount is an auto generated low-level Go binding around an user-defined struct.
type WalletAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// WalletPaymentsMetaData contains all meta data concerning the WalletPayments contract.
var WalletPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIBtcAccounts.BtcAnchor\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcAnchorStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structIBtcEscrows.BtcEscrowTerms\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcEscrowTermsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcProposalCommitmentStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"internalType\":\"structICspSettlement.CspInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"cspInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"}],\"internalType\":\"structICspSettlement.CspPaymentRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"cspPaymentRecordStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeSchedule\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"}],\"internalType\":\"structIPayments.PaymentInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInstruction\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structReissueFeeParams\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"reissueFeeParamsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletAccountStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"cancelAfter\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"nullified\",\"type\":\"bool\"}],\"internalType\":\"structIXrpEscrows.XrpEscrowCreateMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"xrpEscrowCreateMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"nullified\",\"type\":\"bool\"}],\"internalType\":\"structIXrpEscrows.XrpEscrowTerms\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"xrpEscrowTermsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletPaymentsMetaData.ABI instead.
var WalletPaymentsABI = WalletPaymentsMetaData.ABI

// WalletPayments is an auto generated Go binding around an Ethereum contract.
type WalletPayments struct {
	WalletPaymentsCaller     // Read-only binding to the contract
	WalletPaymentsTransactor // Write-only binding to the contract
	WalletPaymentsFilterer   // Log filterer for contract events
}

// WalletPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletPaymentsSession struct {
	Contract     *WalletPayments   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletPaymentsCallerSession struct {
	Contract *WalletPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WalletPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletPaymentsTransactorSession struct {
	Contract     *WalletPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletPaymentsRaw struct {
	Contract *WalletPayments // Generic contract binding to access the raw methods on
}

// WalletPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletPaymentsCallerRaw struct {
	Contract *WalletPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// WalletPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletPaymentsTransactorRaw struct {
	Contract *WalletPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletPayments creates a new instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPayments(address common.Address, backend bind.ContractBackend) (*WalletPayments, error) {
	contract, err := bindWalletPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletPayments{WalletPaymentsCaller: WalletPaymentsCaller{contract: contract}, WalletPaymentsTransactor: WalletPaymentsTransactor{contract: contract}, WalletPaymentsFilterer: WalletPaymentsFilterer{contract: contract}}, nil
}

// NewWalletPaymentsCaller creates a new read-only instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsCaller(address common.Address, caller bind.ContractCaller) (*WalletPaymentsCaller, error) {
	contract, err := bindWalletPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsCaller{contract: contract}, nil
}

// NewWalletPaymentsTransactor creates a new write-only instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletPaymentsTransactor, error) {
	contract, err := bindWalletPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsTransactor{contract: contract}, nil
}

// NewWalletPaymentsFilterer creates a new log filterer instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletPaymentsFilterer, error) {
	contract, err := bindWalletPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsFilterer{contract: contract}, nil
}

// bindWalletPayments binds a generic wrapper to an already deployed contract.
func bindWalletPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletPaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletPayments *WalletPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletPayments.Contract.WalletPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletPayments *WalletPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletPayments *WalletPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletPayments *WalletPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletPayments *WalletPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletPayments *WalletPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletPayments.Contract.contract.Transact(opts, method, params...)
}

// BtcAnchorStruct is a paid mutator transaction binding the contract method 0x9ba77a6a.
//
// Solidity: function btcAnchorStruct((bytes32,uint32,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) BtcAnchorStruct(opts *bind.TransactOpts, arg0 IBtcAccountsBtcAnchor) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "btcAnchorStruct", arg0)
}

// BtcAnchorStruct is a paid mutator transaction binding the contract method 0x9ba77a6a.
//
// Solidity: function btcAnchorStruct((bytes32,uint32,uint64) ) returns()
func (_WalletPayments *WalletPaymentsSession) BtcAnchorStruct(arg0 IBtcAccountsBtcAnchor) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcAnchorStruct(&_WalletPayments.TransactOpts, arg0)
}

// BtcAnchorStruct is a paid mutator transaction binding the contract method 0x9ba77a6a.
//
// Solidity: function btcAnchorStruct((bytes32,uint32,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) BtcAnchorStruct(arg0 IBtcAccountsBtcAnchor) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcAnchorStruct(&_WalletPayments.TransactOpts, arg0)
}

// BtcEscrowTermsStruct is a paid mutator transaction binding the contract method 0xb8f4b596.
//
// Solidity: function btcEscrowTermsStruct((bytes32,uint64,uint64,bytes) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) BtcEscrowTermsStruct(opts *bind.TransactOpts, arg0 IBtcEscrowsBtcEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "btcEscrowTermsStruct", arg0)
}

// BtcEscrowTermsStruct is a paid mutator transaction binding the contract method 0xb8f4b596.
//
// Solidity: function btcEscrowTermsStruct((bytes32,uint64,uint64,bytes) ) returns()
func (_WalletPayments *WalletPaymentsSession) BtcEscrowTermsStruct(arg0 IBtcEscrowsBtcEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcEscrowTermsStruct(&_WalletPayments.TransactOpts, arg0)
}

// BtcEscrowTermsStruct is a paid mutator transaction binding the contract method 0xb8f4b596.
//
// Solidity: function btcEscrowTermsStruct((bytes32,uint64,uint64,bytes) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) BtcEscrowTermsStruct(arg0 IBtcEscrowsBtcEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcEscrowTermsStruct(&_WalletPayments.TransactOpts, arg0)
}

// BtcProposalCommitmentStruct is a paid mutator transaction binding the contract method 0x646116c4.
//
// Solidity: function btcProposalCommitmentStruct((uint32,uint64,uint32,bytes32,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) BtcProposalCommitmentStruct(opts *bind.TransactOpts, arg0 IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "btcProposalCommitmentStruct", arg0)
}

// BtcProposalCommitmentStruct is a paid mutator transaction binding the contract method 0x646116c4.
//
// Solidity: function btcProposalCommitmentStruct((uint32,uint64,uint32,bytes32,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsSession) BtcProposalCommitmentStruct(arg0 IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcProposalCommitmentStruct(&_WalletPayments.TransactOpts, arg0)
}

// BtcProposalCommitmentStruct is a paid mutator transaction binding the contract method 0x646116c4.
//
// Solidity: function btcProposalCommitmentStruct((uint32,uint64,uint32,bytes32,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) BtcProposalCommitmentStruct(arg0 IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.Contract.BtcProposalCommitmentStruct(&_WalletPayments.TransactOpts, arg0)
}

// CspInstructionMessageStruct is a paid mutator transaction binding the contract method 0x10f47dab.
//
// Solidity: function cspInstructionMessageStruct(((address,uint64)[],bytes32,bytes32,uint32,uint64,uint32,uint64,uint64,bytes32,bytes32,address) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) CspInstructionMessageStruct(opts *bind.TransactOpts, arg0 ICspSettlementCspInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "cspInstructionMessageStruct", arg0)
}

// CspInstructionMessageStruct is a paid mutator transaction binding the contract method 0x10f47dab.
//
// Solidity: function cspInstructionMessageStruct(((address,uint64)[],bytes32,bytes32,uint32,uint64,uint32,uint64,uint64,bytes32,bytes32,address) ) returns()
func (_WalletPayments *WalletPaymentsSession) CspInstructionMessageStruct(arg0 ICspSettlementCspInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.CspInstructionMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// CspInstructionMessageStruct is a paid mutator transaction binding the contract method 0x10f47dab.
//
// Solidity: function cspInstructionMessageStruct(((address,uint64)[],bytes32,bytes32,uint32,uint64,uint32,uint64,uint64,bytes32,bytes32,address) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) CspInstructionMessageStruct(arg0 ICspSettlementCspInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.CspInstructionMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// CspPaymentRecordStruct is a paid mutator transaction binding the contract method 0x73ee5ab8.
//
// Solidity: function cspPaymentRecordStruct((bytes32,bytes32,string,uint32,string,bytes,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) CspPaymentRecordStruct(opts *bind.TransactOpts, arg0 ICspSettlementCspPaymentRecord) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "cspPaymentRecordStruct", arg0)
}

// CspPaymentRecordStruct is a paid mutator transaction binding the contract method 0x73ee5ab8.
//
// Solidity: function cspPaymentRecordStruct((bytes32,bytes32,string,uint32,string,bytes,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsSession) CspPaymentRecordStruct(arg0 ICspSettlementCspPaymentRecord) (*types.Transaction, error) {
	return _WalletPayments.Contract.CspPaymentRecordStruct(&_WalletPayments.TransactOpts, arg0)
}

// CspPaymentRecordStruct is a paid mutator transaction binding the contract method 0x73ee5ab8.
//
// Solidity: function cspPaymentRecordStruct((bytes32,bytes32,string,uint32,string,bytes,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) CspPaymentRecordStruct(arg0 ICspSettlementCspPaymentRecord) (*types.Transaction, error) {
	return _WalletPayments.Contract.CspPaymentRecordStruct(&_WalletPayments.TransactOpts, arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) PaymentInstructionMessageStruct(opts *bind.TransactOpts, arg0 IPaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "paymentInstructionMessageStruct", arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsSession) PaymentInstructionMessageStruct(arg0 IPaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.PaymentInstructionMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) PaymentInstructionMessageStruct(arg0 IPaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.PaymentInstructionMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) PaymentInstructionStruct(opts *bind.TransactOpts, arg0 PaymentInstruction) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "paymentInstructionStruct", arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsSession) PaymentInstructionStruct(arg0 PaymentInstruction) (*types.Transaction, error) {
	return _WalletPayments.Contract.PaymentInstructionStruct(&_WalletPayments.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) PaymentInstructionStruct(arg0 PaymentInstruction) (*types.Transaction, error) {
	return _WalletPayments.Contract.PaymentInstructionStruct(&_WalletPayments.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) ReissueFeeParamsStruct(opts *bind.TransactOpts, arg0 ReissueFeeParams) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "reissueFeeParamsStruct", arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_WalletPayments *WalletPaymentsSession) ReissueFeeParamsStruct(arg0 ReissueFeeParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReissueFeeParamsStruct(&_WalletPayments.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) ReissueFeeParamsStruct(arg0 ReissueFeeParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReissueFeeParamsStruct(&_WalletPayments.TransactOpts, arg0)
}

// WalletAccountStruct is a paid mutator transaction binding the contract method 0x7fea4de9.
//
// Solidity: function walletAccountStruct((bytes32,string) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) WalletAccountStruct(opts *bind.TransactOpts, arg0 WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "walletAccountStruct", arg0)
}

// WalletAccountStruct is a paid mutator transaction binding the contract method 0x7fea4de9.
//
// Solidity: function walletAccountStruct((bytes32,string) ) returns()
func (_WalletPayments *WalletPaymentsSession) WalletAccountStruct(arg0 WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletAccountStruct(&_WalletPayments.TransactOpts, arg0)
}

// WalletAccountStruct is a paid mutator transaction binding the contract method 0x7fea4de9.
//
// Solidity: function walletAccountStruct((bytes32,string) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) WalletAccountStruct(arg0 WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletAccountStruct(&_WalletPayments.TransactOpts, arg0)
}

// XrpEscrowCreateMessageStruct is a paid mutator transaction binding the contract method 0x76c2575b.
//
// Solidity: function xrpEscrowCreateMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,uint256,bytes32,uint64,uint256,uint64,uint64,bool) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) XrpEscrowCreateMessageStruct(opts *bind.TransactOpts, arg0 IXrpEscrowsXrpEscrowCreateMessage) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "xrpEscrowCreateMessageStruct", arg0)
}

// XrpEscrowCreateMessageStruct is a paid mutator transaction binding the contract method 0x76c2575b.
//
// Solidity: function xrpEscrowCreateMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,uint256,bytes32,uint64,uint256,uint64,uint64,bool) ) returns()
func (_WalletPayments *WalletPaymentsSession) XrpEscrowCreateMessageStruct(arg0 IXrpEscrowsXrpEscrowCreateMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.XrpEscrowCreateMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// XrpEscrowCreateMessageStruct is a paid mutator transaction binding the contract method 0x76c2575b.
//
// Solidity: function xrpEscrowCreateMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,uint256,bytes32,uint64,uint256,uint64,uint64,bool) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) XrpEscrowCreateMessageStruct(arg0 IXrpEscrowsXrpEscrowCreateMessage) (*types.Transaction, error) {
	return _WalletPayments.Contract.XrpEscrowCreateMessageStruct(&_WalletPayments.TransactOpts, arg0)
}

// XrpEscrowTermsStruct is a paid mutator transaction binding the contract method 0xaf6d2f4c.
//
// Solidity: function xrpEscrowTermsStruct((bytes32,uint256,uint64,string,bool) ) returns()
func (_WalletPayments *WalletPaymentsTransactor) XrpEscrowTermsStruct(opts *bind.TransactOpts, arg0 IXrpEscrowsXrpEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "xrpEscrowTermsStruct", arg0)
}

// XrpEscrowTermsStruct is a paid mutator transaction binding the contract method 0xaf6d2f4c.
//
// Solidity: function xrpEscrowTermsStruct((bytes32,uint256,uint64,string,bool) ) returns()
func (_WalletPayments *WalletPaymentsSession) XrpEscrowTermsStruct(arg0 IXrpEscrowsXrpEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.Contract.XrpEscrowTermsStruct(&_WalletPayments.TransactOpts, arg0)
}

// XrpEscrowTermsStruct is a paid mutator transaction binding the contract method 0xaf6d2f4c.
//
// Solidity: function xrpEscrowTermsStruct((bytes32,uint256,uint64,string,bool) ) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) XrpEscrowTermsStruct(arg0 IXrpEscrowsXrpEscrowTerms) (*types.Transaction, error) {
	return _WalletPayments.Contract.XrpEscrowTermsStruct(&_WalletPayments.TransactOpts, arg0)
}
