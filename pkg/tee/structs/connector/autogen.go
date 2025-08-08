// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connector

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

// IFtdcHubFtdcAttestationRequest is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcAttestationRequest struct {
	Header      IFtdcHubFtdcRequestHeader
	RequestBody []byte
}

// IFtdcHubFtdcRequestHeader is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcRequestHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// IFtdcHubFtdcResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFtdcVerificationFtdcSignatures is an auto generated low-level Go binding around an user-defined struct.
type IFtdcVerificationFtdcSignatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IPMWMultisigAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  IPMWMultisigAccountConfiguredRequestBody
	ResponseBody IPMWMultisigAccountConfiguredResponseBody
}

// IPMWMultisigAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredRequestBody struct {
	WalletAddress string
	PublicKeys    [][]byte
	Threshold     uint64
	OpType        [32]byte
}

// IPMWMultisigAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// IPMWPaymentStatusProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  IPMWPaymentStatusRequestBody
	ResponseBody IPMWPaymentStatusResponseBody
}

// IPMWPaymentStatusRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusRequestBody struct {
	WalletId [32]byte
	Nonce    uint64
	SubNonce uint64
}

// IPMWPaymentStatusResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusResponseBody struct {
	SenderAddress     string
	RecipientAddress  string
	Amount            *big.Int
	Fee               *big.Int
	PaymentReference  [32]byte
	TransactionStatus uint8
	RevertReason      string
	ReceivedAmount    *big.Int
	TransactionFee    *big.Int
	TransactionId     [32]byte
	BlockNumber       uint64
	BlockTimestamp    uint64
}

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  ITeeAvailabilityCheckRequestBody
	ResponseBody ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId     common.Address
	Url       string
	Challenge [32]byte
}

// ITeeAvailabilityCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponseBody struct {
	Status                 uint8
	TeeTimestamp           uint64
	CodeHash               [32]byte
	Platform               [32]byte
	InitialSigningPolicyId uint32
	LastSigningPolicyId    uint32
	State                  ITeeAvailabilityCheckTeeState
}

// ITeeAvailabilityCheckTeeState is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckTeeState struct {
	SystemState        []byte
	SystemStateVersion [32]byte
	State              []byte
	StateVersion       [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ConnectorMetaData contains all meta data concerning the Connector contract.
var ConnectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckTeeStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcRequestHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"}],\"internalType\":\"structIFtdcHub.FtdcAttestationRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ftdcAttestationRequestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcRequestHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ftdcRequestHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ftdcResponseHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWPaymentStatus.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConnectorABI is the input ABI used to generate the binding from.
// Deprecated: Use ConnectorMetaData.ABI instead.
var ConnectorABI = ConnectorMetaData.ABI

// Connector is an auto generated Go binding around an Ethereum contract.
type Connector struct {
	ConnectorCaller     // Read-only binding to the contract
	ConnectorTransactor // Write-only binding to the contract
	ConnectorFilterer   // Log filterer for contract events
}

// ConnectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConnectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConnectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConnectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConnectorSession struct {
	Contract     *Connector        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConnectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConnectorCallerSession struct {
	Contract *ConnectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConnectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConnectorTransactorSession struct {
	Contract     *ConnectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConnectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConnectorRaw struct {
	Contract *Connector // Generic contract binding to access the raw methods on
}

// ConnectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConnectorCallerRaw struct {
	Contract *ConnectorCaller // Generic read-only contract binding to access the raw methods on
}

// ConnectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConnectorTransactorRaw struct {
	Contract *ConnectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConnector creates a new instance of Connector, bound to a specific deployed contract.
func NewConnector(address common.Address, backend bind.ContractBackend) (*Connector, error) {
	contract, err := bindConnector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Connector{ConnectorCaller: ConnectorCaller{contract: contract}, ConnectorTransactor: ConnectorTransactor{contract: contract}, ConnectorFilterer: ConnectorFilterer{contract: contract}}, nil
}

// NewConnectorCaller creates a new read-only instance of Connector, bound to a specific deployed contract.
func NewConnectorCaller(address common.Address, caller bind.ContractCaller) (*ConnectorCaller, error) {
	contract, err := bindConnector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorCaller{contract: contract}, nil
}

// NewConnectorTransactor creates a new write-only instance of Connector, bound to a specific deployed contract.
func NewConnectorTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectorTransactor, error) {
	contract, err := bindConnector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorTransactor{contract: contract}, nil
}

// NewConnectorFilterer creates a new log filterer instance of Connector, bound to a specific deployed contract.
func NewConnectorFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectorFilterer, error) {
	contract, err := bindConnector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectorFilterer{contract: contract}, nil
}

// bindConnector binds a generic wrapper to an already deployed contract.
func bindConnector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConnectorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Connector *ConnectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Connector.Contract.ConnectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Connector *ConnectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Connector.Contract.ConnectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Connector *ConnectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Connector.Contract.ConnectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Connector *ConnectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Connector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Connector *ConnectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Connector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Connector *ConnectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Connector.Contract.contract.Transact(opts, method, params...)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0xf6c6bdb0.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckProofStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckProofStruct", arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0xf6c6bdb0.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckProofStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0xf6c6bdb0.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckProofStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xbdf0dadc.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,string,bytes32) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckRequestBodyStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckRequestBodyStruct", arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xbdf0dadc.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,string,bytes32) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckRequestBodyStruct(arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xbdf0dadc.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,string,bytes32) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckRequestBodyStruct(arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckResponseBodyStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckResponseBodyStruct", arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckResponseBodyStruct(arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckResponseBodyStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckResponseBodyStruct(arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckResponseBodyStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckTeeStateStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckTeeStateStruct", arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckTeeStateStruct(arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckTeeStateStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckTeeStateStruct(arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckTeeStateStruct(&_Connector.TransactOpts, arg0)
}

// FtdcAttestationRequestStruct is a paid mutator transaction binding the contract method 0x627d482e.
//
// Solidity: function ftdcAttestationRequestStruct(((bytes32,bytes32,uint16,address[],uint64),bytes) ) returns()
func (_Connector *ConnectorTransactor) FtdcAttestationRequestStruct(opts *bind.TransactOpts, arg0 IFtdcHubFtdcAttestationRequest) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "ftdcAttestationRequestStruct", arg0)
}

// FtdcAttestationRequestStruct is a paid mutator transaction binding the contract method 0x627d482e.
//
// Solidity: function ftdcAttestationRequestStruct(((bytes32,bytes32,uint16,address[],uint64),bytes) ) returns()
func (_Connector *ConnectorSession) FtdcAttestationRequestStruct(arg0 IFtdcHubFtdcAttestationRequest) (*types.Transaction, error) {
	return _Connector.Contract.FtdcAttestationRequestStruct(&_Connector.TransactOpts, arg0)
}

// FtdcAttestationRequestStruct is a paid mutator transaction binding the contract method 0x627d482e.
//
// Solidity: function ftdcAttestationRequestStruct(((bytes32,bytes32,uint16,address[],uint64),bytes) ) returns()
func (_Connector *ConnectorTransactorSession) FtdcAttestationRequestStruct(arg0 IFtdcHubFtdcAttestationRequest) (*types.Transaction, error) {
	return _Connector.Contract.FtdcAttestationRequestStruct(&_Connector.TransactOpts, arg0)
}

// FtdcRequestHeaderStruct is a paid mutator transaction binding the contract method 0x5b421a8b.
//
// Solidity: function ftdcRequestHeaderStruct((bytes32,bytes32,uint16,address[],uint64) ) returns()
func (_Connector *ConnectorTransactor) FtdcRequestHeaderStruct(opts *bind.TransactOpts, arg0 IFtdcHubFtdcRequestHeader) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "ftdcRequestHeaderStruct", arg0)
}

// FtdcRequestHeaderStruct is a paid mutator transaction binding the contract method 0x5b421a8b.
//
// Solidity: function ftdcRequestHeaderStruct((bytes32,bytes32,uint16,address[],uint64) ) returns()
func (_Connector *ConnectorSession) FtdcRequestHeaderStruct(arg0 IFtdcHubFtdcRequestHeader) (*types.Transaction, error) {
	return _Connector.Contract.FtdcRequestHeaderStruct(&_Connector.TransactOpts, arg0)
}

// FtdcRequestHeaderStruct is a paid mutator transaction binding the contract method 0x5b421a8b.
//
// Solidity: function ftdcRequestHeaderStruct((bytes32,bytes32,uint16,address[],uint64) ) returns()
func (_Connector *ConnectorTransactorSession) FtdcRequestHeaderStruct(arg0 IFtdcHubFtdcRequestHeader) (*types.Transaction, error) {
	return _Connector.Contract.FtdcRequestHeaderStruct(&_Connector.TransactOpts, arg0)
}

// FtdcResponseHeaderStruct is a paid mutator transaction binding the contract method 0x72d0e906.
//
// Solidity: function ftdcResponseHeaderStruct((bytes32,bytes32,uint16,address[],uint64,uint64) ) returns()
func (_Connector *ConnectorTransactor) FtdcResponseHeaderStruct(opts *bind.TransactOpts, arg0 IFtdcHubFtdcResponseHeader) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "ftdcResponseHeaderStruct", arg0)
}

// FtdcResponseHeaderStruct is a paid mutator transaction binding the contract method 0x72d0e906.
//
// Solidity: function ftdcResponseHeaderStruct((bytes32,bytes32,uint16,address[],uint64,uint64) ) returns()
func (_Connector *ConnectorSession) FtdcResponseHeaderStruct(arg0 IFtdcHubFtdcResponseHeader) (*types.Transaction, error) {
	return _Connector.Contract.FtdcResponseHeaderStruct(&_Connector.TransactOpts, arg0)
}

// FtdcResponseHeaderStruct is a paid mutator transaction binding the contract method 0x72d0e906.
//
// Solidity: function ftdcResponseHeaderStruct((bytes32,bytes32,uint16,address[],uint64,uint64) ) returns()
func (_Connector *ConnectorTransactorSession) FtdcResponseHeaderStruct(arg0 IFtdcHubFtdcResponseHeader) (*types.Transaction, error) {
	return _Connector.Contract.FtdcResponseHeaderStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x288d8e74.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) ) returns()
func (_Connector *ConnectorTransactor) PmwMultisigAccountConfiguredProofStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwMultisigAccountConfiguredProofStruct", arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x288d8e74.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) ) returns()
func (_Connector *ConnectorSession) PmwMultisigAccountConfiguredProofStruct(arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredProofStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x288d8e74.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) ) returns()
func (_Connector *ConnectorTransactorSession) PmwMultisigAccountConfiguredProofStruct(arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredProofStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xb41fcc74.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64,bytes32) ) returns()
func (_Connector *ConnectorTransactor) PmwMultisigAccountConfiguredRequestBodyStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwMultisigAccountConfiguredRequestBodyStruct", arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xb41fcc74.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64,bytes32) ) returns()
func (_Connector *ConnectorSession) PmwMultisigAccountConfiguredRequestBodyStruct(arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xb41fcc74.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64,bytes32) ) returns()
func (_Connector *ConnectorTransactorSession) PmwMultisigAccountConfiguredRequestBodyStruct(arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Connector *ConnectorTransactor) PmwMultisigAccountConfiguredResponseBodyStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwMultisigAccountConfiguredResponseBodyStruct", arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Connector *ConnectorSession) PmwMultisigAccountConfiguredResponseBodyStruct(arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredResponseBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Connector *ConnectorTransactorSession) PmwMultisigAccountConfiguredResponseBodyStruct(arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwMultisigAccountConfiguredResponseBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xed0bf61c.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(bytes32,uint64,uint64),(string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Connector *ConnectorTransactor) PmwPaymentStatusProofStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwPaymentStatusProofStruct", arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xed0bf61c.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(bytes32,uint64,uint64),(string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Connector *ConnectorSession) PmwPaymentStatusProofStruct(arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusProofStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xed0bf61c.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(bytes32,uint64,uint64),(string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Connector *ConnectorTransactorSession) PmwPaymentStatusProofStruct(arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusProofStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0xd64b83c7.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorTransactor) PmwPaymentStatusRequestBodyStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwPaymentStatusRequestBodyStruct", arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0xd64b83c7.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorSession) PmwPaymentStatusRequestBodyStruct(arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0xd64b83c7.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorTransactorSession) PmwPaymentStatusRequestBodyStruct(arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusRequestBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x2aba206e.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorTransactor) PmwPaymentStatusResponseBodyStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "pmwPaymentStatusResponseBodyStruct", arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x2aba206e.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorSession) PmwPaymentStatusResponseBodyStruct(arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusResponseBodyStruct(&_Connector.TransactOpts, arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x2aba206e.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,string,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Connector *ConnectorTransactorSession) PmwPaymentStatusResponseBodyStruct(arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Connector.Contract.PmwPaymentStatusResponseBodyStruct(&_Connector.TransactOpts, arg0)
}
