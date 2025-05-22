// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeverification

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

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	RelayMessage       []byte
	TeeSignatures      []Signature
	CosignerSignatures []Signature
	Data               ITeeAvailabilityCheckResponse
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

// TeeVerificationMetaData contains all meta data concerning the TeeVerification contract.
var TeeVerificationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"CosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITeeAvailabilityCheck.TeeMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"}],\"name\":\"requestAvailabilityCheckAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"requestTeeAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITeeAvailabilityCheck.TeeMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAvailabilityCheckProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVerificationMetaData.ABI instead.
var TeeVerificationABI = TeeVerificationMetaData.ABI

// TeeVerification is an auto generated Go binding around an Ethereum contract.
type TeeVerification struct {
	TeeVerificationCaller     // Read-only binding to the contract
	TeeVerificationTransactor // Write-only binding to the contract
	TeeVerificationFilterer   // Log filterer for contract events
}

// TeeVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVerificationSession struct {
	Contract     *TeeVerification  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVerificationCallerSession struct {
	Contract *TeeVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeeVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVerificationTransactorSession struct {
	Contract     *TeeVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeeVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVerificationRaw struct {
	Contract *TeeVerification // Generic contract binding to access the raw methods on
}

// TeeVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVerificationCallerRaw struct {
	Contract *TeeVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVerificationTransactorRaw struct {
	Contract *TeeVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVerification creates a new instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerification(address common.Address, backend bind.ContractBackend) (*TeeVerification, error) {
	contract, err := bindTeeVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVerification{TeeVerificationCaller: TeeVerificationCaller{contract: contract}, TeeVerificationTransactor: TeeVerificationTransactor{contract: contract}, TeeVerificationFilterer: TeeVerificationFilterer{contract: contract}}, nil
}

// NewTeeVerificationCaller creates a new read-only instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationCaller(address common.Address, caller bind.ContractCaller) (*TeeVerificationCaller, error) {
	contract, err := bindTeeVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationCaller{contract: contract}, nil
}

// NewTeeVerificationTransactor creates a new write-only instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVerificationTransactor, error) {
	contract, err := bindTeeVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTransactor{contract: contract}, nil
}

// NewTeeVerificationFilterer creates a new log filterer instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVerificationFilterer, error) {
	contract, err := bindTeeVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationFilterer{contract: contract}, nil
}

// bindTeeVerification binds a generic wrapper to an already deployed contract.
func bindTeeVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVerification *TeeVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVerification.Contract.TeeVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVerification *TeeVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVerification.Contract.TeeVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVerification *TeeVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVerification.Contract.TeeVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVerification *TeeVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVerification *TeeVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVerification *TeeVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVerification.Contract.contract.Transact(opts, method, params...)
}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationCaller) GetCosigners(opts *bind.CallOpts) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getCosigners")

	outstruct := new(struct {
		Cosigners          []common.Address
		CosignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationSession) GetCosigners() (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeVerification.Contract.GetCosigners(&_TeeVerification.CallOpts)
}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationCallerSession) GetCosigners() (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeVerification.Contract.GetCosigners(&_TeeVerification.CallOpts)
}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationCaller) GetSettings(opts *bind.CallOpts) (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getSettings")

	outstruct := new(struct {
		AvailabilityCheckValidityDurationSeconds *big.Int
		ChallengeValidityDurationSeconds         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailabilityCheckValidityDurationSeconds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChallengeValidityDurationSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationSession) GetSettings() (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	return _TeeVerification.Contract.GetSettings(&_TeeVerification.CallOpts)
}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationCallerSession) GetSettings() (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	return _TeeVerification.Contract.GetSettings(&_TeeVerification.CallOpts)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x910c99ec.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactor) ConfirmAvailability(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "confirmAvailability", _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x910c99ec.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x910c99ec.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactorSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestAvailabilityCheckAttestation(opts *bind.TransactOpts, _teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestAvailabilityCheckAttestation", _teeId, _testOnTeeId)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _testOnTeeId)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _testOnTeeId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestTeeAttestation(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestTeeAttestation", _teeId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestTeeAttestation(_teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestTeeAttestation(_teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0x7f1cbcdf.
//
// Solidity: function verifyAvailabilityCheckProof((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactor) VerifyAvailabilityCheckProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "verifyAvailabilityCheckProof", _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0x7f1cbcdf.
//
// Solidity: function verifyAvailabilityCheckProof((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0x7f1cbcdf.
//
// Solidity: function verifyAvailabilityCheckProof((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,address[],uint64,(address,address,string,bytes32,bytes32,bytes32,uint24,uint256),(uint8,uint8,uint64))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactorSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// TeeVerificationAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the TeeVerification contract.
type TeeVerificationAvailabilityCheckValidityExtendedIterator struct {
	Event *TeeVerificationAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationAvailabilityCheckValidityExtended)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeVerificationAvailabilityCheckValidityExtended)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the TeeVerification contract.
type TeeVerificationAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address, owner []common.Address) (*TeeVerificationAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationAvailabilityCheckValidityExtendedIterator{contract: _TeeVerification.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *TeeVerificationAvailabilityCheckValidityExtended, teeId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationAvailabilityCheckValidityExtended)
				if err := _TeeVerification.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAvailabilityCheckValidityExtended is a log parse operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*TeeVerificationAvailabilityCheckValidityExtended, error) {
	event := new(TeeVerificationAvailabilityCheckValidityExtended)
	if err := _TeeVerification.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationCosignersSetIterator is returned from FilterCosignersSet and is used to iterate over the raw logs and unpacked data for CosignersSet events raised by the TeeVerification contract.
type TeeVerificationCosignersSetIterator struct {
	Event *TeeVerificationCosignersSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeVerificationCosignersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationCosignersSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeVerificationCosignersSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeVerificationCosignersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationCosignersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationCosignersSet represents a CosignersSet event raised by the TeeVerification contract.
type TeeVerificationCosignersSet struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCosignersSet is a free log retrieval operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) FilterCosignersSet(opts *bind.FilterOpts) (*TeeVerificationCosignersSetIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "CosignersSet")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationCosignersSetIterator{contract: _TeeVerification.contract, event: "CosignersSet", logs: logs, sub: sub}, nil
}

// WatchCosignersSet is a free log subscription operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) WatchCosignersSet(opts *bind.WatchOpts, sink chan<- *TeeVerificationCosignersSet) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "CosignersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationCosignersSet)
				if err := _TeeVerification.contract.UnpackLog(event, "CosignersSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCosignersSet is a log parse operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) ParseCosignersSet(log types.Log) (*TeeVerificationCosignersSet, error) {
	event := new(TeeVerificationCosignersSet)
	if err := _TeeVerification.contract.UnpackLog(event, "CosignersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationSettingsUpdatedIterator is returned from FilterSettingsUpdated and is used to iterate over the raw logs and unpacked data for SettingsUpdated events raised by the TeeVerification contract.
type TeeVerificationSettingsUpdatedIterator struct {
	Event *TeeVerificationSettingsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeVerificationSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationSettingsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeVerificationSettingsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeVerificationSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationSettingsUpdated represents a SettingsUpdated event raised by the TeeVerification contract.
type TeeVerificationSettingsUpdated struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
	Raw                                      types.Log // Blockchain specific contextual infos
}

// FilterSettingsUpdated is a free log retrieval operation binding the contract event 0x290c7608baac88950f6f2bf919aee31e5caf8c7bf387f0e489159253ad2e39d5.
//
// Solidity: event SettingsUpdated(uint256 availabilityCheckValidityDurationSeconds, uint256 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) FilterSettingsUpdated(opts *bind.FilterOpts) (*TeeVerificationSettingsUpdatedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationSettingsUpdatedIterator{contract: _TeeVerification.contract, event: "SettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettingsUpdated is a free log subscription operation binding the contract event 0x290c7608baac88950f6f2bf919aee31e5caf8c7bf387f0e489159253ad2e39d5.
//
// Solidity: event SettingsUpdated(uint256 availabilityCheckValidityDurationSeconds, uint256 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) WatchSettingsUpdated(opts *bind.WatchOpts, sink chan<- *TeeVerificationSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationSettingsUpdated)
				if err := _TeeVerification.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettingsUpdated is a log parse operation binding the contract event 0x290c7608baac88950f6f2bf919aee31e5caf8c7bf387f0e489159253ad2e39d5.
//
// Solidity: event SettingsUpdated(uint256 availabilityCheckValidityDurationSeconds, uint256 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) ParseSettingsUpdated(log types.Log) (*TeeVerificationSettingsUpdated, error) {
	event := new(TeeVerificationSettingsUpdated)
	if err := _TeeVerification.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationTeeAttestationRequestedIterator is returned from FilterTeeAttestationRequested and is used to iterate over the raw logs and unpacked data for TeeAttestationRequested events raised by the TeeVerification contract.
type TeeVerificationTeeAttestationRequestedIterator struct {
	Event *TeeVerificationTeeAttestationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeVerificationTeeAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationTeeAttestationRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeVerificationTeeAttestationRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeVerificationTeeAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationTeeAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationTeeAttestationRequested represents a TeeAttestationRequested event raised by the TeeVerification contract.
type TeeVerificationTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeAttestationRequested is a free log retrieval operation binding the contract event 0x1943691de6eb6b645aef3a008925b3da6a5b3d33ceebd3cd042c2b152cf2f5fb.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, uint256 challenge)
func (_TeeVerification *TeeVerificationFilterer) FilterTeeAttestationRequested(opts *bind.FilterOpts, teeId []common.Address) (*TeeVerificationTeeAttestationRequestedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTeeAttestationRequestedIterator{contract: _TeeVerification.contract, event: "TeeAttestationRequested", logs: logs, sub: sub}, nil
}

// WatchTeeAttestationRequested is a free log subscription operation binding the contract event 0x1943691de6eb6b645aef3a008925b3da6a5b3d33ceebd3cd042c2b152cf2f5fb.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, uint256 challenge)
func (_TeeVerification *TeeVerificationFilterer) WatchTeeAttestationRequested(opts *bind.WatchOpts, sink chan<- *TeeVerificationTeeAttestationRequested, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationTeeAttestationRequested)
				if err := _TeeVerification.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeAttestationRequested is a log parse operation binding the contract event 0x1943691de6eb6b645aef3a008925b3da6a5b3d33ceebd3cd042c2b152cf2f5fb.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, uint256 challenge)
func (_TeeVerification *TeeVerificationFilterer) ParseTeeAttestationRequested(log types.Log) (*TeeVerificationTeeAttestationRequested, error) {
	event := new(TeeVerificationTeeAttestationRequested)
	if err := _TeeVerification.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
