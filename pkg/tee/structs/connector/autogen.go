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

// IFtdcHubFtdcProve is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcProve struct {
	TeeIds             []common.Address
	ThresholdBIPS      uint16
	Cosigners          []common.Address
	CosignersThreshold uint64
	AttestationRequest []byte
}

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	RelayMessage       []byte
	TeeSignatures      []Signature
	CosignerSignatures []Signature
	Data               ITeeAvailabilityCheckResponse
}

// ITeeAvailabilityCheckRequest is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequest struct {
	AttestationType [32]byte
	SourceId        [32]byte
	RequestBody     ITeeAvailabilityCheckRequestBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId             common.Address
	InitialTeeId      common.Address
	Url               string
	CodeHash          [32]byte
	Platform          [32]byte
	TeeGovernanceHash [32]byte
	RewardEpochId     *big.Int
	Challenge         *big.Int
}

// ITeeAvailabilityCheckResponse is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponse struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	Timestamp          uint64
	Cosigners          []common.Address
	CosignersThreshold uint64
	RequestBody        ITeeAvailabilityCheckRequestBody
	ResponseBody       ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponseBody struct {
	Status        uint8
	MachineStatus uint8
	TeeTimestamp  uint64
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ConnectorMetaData contains all meta data concerning the Connector contract.
var ConnectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITeeAvailabilityCheck.TeeMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckRequestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITeeAvailabilityCheck.TeeMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckResponseStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"teeIds\",\"type\":\"address[]\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"attestationRequest\",\"type\":\"bytes\"}],\"internalType\":\"structIFtdcHub.FtdcProve\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ftdcProveStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x4a636bf6.
//
// Solidity: function availabilityCheckProofStruct((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckProofStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckProofStruct", arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x4a636bf6.
//
// Solidity: function availabilityCheckProofStruct((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckProofStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x4a636bf6.
//
// Solidity: function availabilityCheckProofStruct((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckProofStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0x017d17db.
//
// Solidity: function availabilityCheckRequestStruct((bytes32,bytes32,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256)) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckRequestStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckRequestStruct", arg0)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0x017d17db.
//
// Solidity: function availabilityCheckRequestStruct((bytes32,bytes32,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256)) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckRequestStruct(arg0 ITeeAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckRequestStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0x017d17db.
//
// Solidity: function availabilityCheckRequestStruct((bytes32,bytes32,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256)) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckRequestStruct(arg0 ITeeAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckRequestStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x5051e552.
//
// Solidity: function availabilityCheckResponseStruct((bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64)) ) returns()
func (_Connector *ConnectorTransactor) AvailabilityCheckResponseStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "availabilityCheckResponseStruct", arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x5051e552.
//
// Solidity: function availabilityCheckResponseStruct((bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64)) ) returns()
func (_Connector *ConnectorSession) AvailabilityCheckResponseStruct(arg0 ITeeAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckResponseStruct(&_Connector.TransactOpts, arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x5051e552.
//
// Solidity: function availabilityCheckResponseStruct((bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64)) ) returns()
func (_Connector *ConnectorTransactorSession) AvailabilityCheckResponseStruct(arg0 ITeeAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Connector.Contract.AvailabilityCheckResponseStruct(&_Connector.TransactOpts, arg0)
}

// FtdcProveStruct is a paid mutator transaction binding the contract method 0x74cee18f.
//
// Solidity: function ftdcProveStruct((address[],uint16,address[],uint64,bytes) ) returns()
func (_Connector *ConnectorTransactor) FtdcProveStruct(opts *bind.TransactOpts, arg0 IFtdcHubFtdcProve) (*types.Transaction, error) {
	return _Connector.contract.Transact(opts, "ftdcProveStruct", arg0)
}

// FtdcProveStruct is a paid mutator transaction binding the contract method 0x74cee18f.
//
// Solidity: function ftdcProveStruct((address[],uint16,address[],uint64,bytes) ) returns()
func (_Connector *ConnectorSession) FtdcProveStruct(arg0 IFtdcHubFtdcProve) (*types.Transaction, error) {
	return _Connector.Contract.FtdcProveStruct(&_Connector.TransactOpts, arg0)
}

// FtdcProveStruct is a paid mutator transaction binding the contract method 0x74cee18f.
//
// Solidity: function ftdcProveStruct((address[],uint16,address[],uint64,bytes) ) returns()
func (_Connector *ConnectorTransactorSession) FtdcProveStruct(arg0 IFtdcHubFtdcProve) (*types.Transaction, error) {
	return _Connector.Contract.FtdcProveStruct(&_Connector.TransactOpts, arg0)
}
