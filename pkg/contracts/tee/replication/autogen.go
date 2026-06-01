// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package replication

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

// IFdc2HubFdc2ResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2ResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	ProofOwner         common.Address
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFdc2VerificationFdc2Signatures is an auto generated low-level Go binding around an user-defined struct.
type IFdc2VerificationFdc2Signatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  ITeeAvailabilityCheckRequestBody
	ResponseBody ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId         common.Address
	TeeProxyId    common.Address
	Url           string
	Challenge     [32]byte
	InstructionId [32]byte
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

// ReplicationMetaData contains all meta data concerning the Replication contract.
var ReplicationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMachinePath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSystemStateVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveMachinePathList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"NotReplicationCapable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMachineOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReplicationNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"PauseBeforeUpgradeMinDurationSecondsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"TeeMachinePausedForUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTeeId\",\"type\":\"address\"}],\"name\":\"TeeMachineReplicationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTeeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"machinePathListNonce\",\"type\":\"uint256\"}],\"name\":\"TeeMachineReplicationTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmReplicate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"}],\"name\":\"getReplicatingTeeId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"replicateFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"setPauseBeforeUpgradeMinDurationSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"toPauseForUpgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ReplicationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReplicationMetaData.ABI instead.
var ReplicationABI = ReplicationMetaData.ABI

// Replication is an auto generated Go binding around an Ethereum contract.
type Replication struct {
	ReplicationCaller     // Read-only binding to the contract
	ReplicationTransactor // Write-only binding to the contract
	ReplicationFilterer   // Log filterer for contract events
}

// ReplicationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReplicationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReplicationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReplicationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReplicationSession struct {
	Contract     *Replication      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReplicationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReplicationCallerSession struct {
	Contract *ReplicationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReplicationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReplicationTransactorSession struct {
	Contract     *ReplicationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReplicationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReplicationRaw struct {
	Contract *Replication // Generic contract binding to access the raw methods on
}

// ReplicationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReplicationCallerRaw struct {
	Contract *ReplicationCaller // Generic read-only contract binding to access the raw methods on
}

// ReplicationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReplicationTransactorRaw struct {
	Contract *ReplicationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReplication creates a new instance of Replication, bound to a specific deployed contract.
func NewReplication(address common.Address, backend bind.ContractBackend) (*Replication, error) {
	contract, err := bindReplication(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Replication{ReplicationCaller: ReplicationCaller{contract: contract}, ReplicationTransactor: ReplicationTransactor{contract: contract}, ReplicationFilterer: ReplicationFilterer{contract: contract}}, nil
}

// NewReplicationCaller creates a new read-only instance of Replication, bound to a specific deployed contract.
func NewReplicationCaller(address common.Address, caller bind.ContractCaller) (*ReplicationCaller, error) {
	contract, err := bindReplication(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicationCaller{contract: contract}, nil
}

// NewReplicationTransactor creates a new write-only instance of Replication, bound to a specific deployed contract.
func NewReplicationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReplicationTransactor, error) {
	contract, err := bindReplication(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicationTransactor{contract: contract}, nil
}

// NewReplicationFilterer creates a new log filterer instance of Replication, bound to a specific deployed contract.
func NewReplicationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReplicationFilterer, error) {
	contract, err := bindReplication(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReplicationFilterer{contract: contract}, nil
}

// bindReplication binds a generic wrapper to an already deployed contract.
func bindReplication(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReplicationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Replication *ReplicationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Replication.Contract.ReplicationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Replication *ReplicationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Replication.Contract.ReplicationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Replication *ReplicationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Replication.Contract.ReplicationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Replication *ReplicationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Replication.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Replication *ReplicationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Replication.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Replication *ReplicationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Replication.Contract.contract.Transact(opts, method, params...)
}

// GetReplicatingTeeId is a free data retrieval call binding the contract method 0x62fe8ad7.
//
// Solidity: function getReplicatingTeeId(address _oldTeeId) view returns(address)
func (_Replication *ReplicationCaller) GetReplicatingTeeId(opts *bind.CallOpts, _oldTeeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _Replication.contract.Call(opts, &out, "getReplicatingTeeId", _oldTeeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReplicatingTeeId is a free data retrieval call binding the contract method 0x62fe8ad7.
//
// Solidity: function getReplicatingTeeId(address _oldTeeId) view returns(address)
func (_Replication *ReplicationSession) GetReplicatingTeeId(_oldTeeId common.Address) (common.Address, error) {
	return _Replication.Contract.GetReplicatingTeeId(&_Replication.CallOpts, _oldTeeId)
}

// GetReplicatingTeeId is a free data retrieval call binding the contract method 0x62fe8ad7.
//
// Solidity: function getReplicatingTeeId(address _oldTeeId) view returns(address)
func (_Replication *ReplicationCallerSession) GetReplicatingTeeId(_oldTeeId common.Address) (common.Address, error) {
	return _Replication.Contract.GetReplicatingTeeId(&_Replication.CallOpts, _oldTeeId)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x250c4322.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_Replication *ReplicationTransactor) ConfirmReplicate(opts *bind.TransactOpts, _newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Replication.contract.Transact(opts, "confirmReplicate", _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x250c4322.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_Replication *ReplicationSession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Replication.Contract.ConfirmReplicate(&_Replication.TransactOpts, _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x250c4322.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_Replication *ReplicationTransactorSession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Replication.Contract.ConfirmReplicate(&_Replication.TransactOpts, _newTeeId, _proof)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x79848111.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof, address _claimBackAddress) payable returns()
func (_Replication *ReplicationTransactor) ReplicateFrom(opts *bind.TransactOpts, _oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.contract.Transact(opts, "replicateFrom", _oldTeeId, _proof, _claimBackAddress)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x79848111.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof, address _claimBackAddress) payable returns()
func (_Replication *ReplicationSession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.Contract.ReplicateFrom(&_Replication.TransactOpts, _oldTeeId, _proof, _claimBackAddress)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x79848111.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof, address _claimBackAddress) payable returns()
func (_Replication *ReplicationTransactorSession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.Contract.ReplicateFrom(&_Replication.TransactOpts, _oldTeeId, _proof, _claimBackAddress)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_Replication *ReplicationTransactor) SetPauseBeforeUpgradeMinDurationSeconds(opts *bind.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _Replication.contract.Transact(opts, "setPauseBeforeUpgradeMinDurationSeconds", _pauseBeforeUpgradeMinDurationSeconds)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_Replication *ReplicationSession) SetPauseBeforeUpgradeMinDurationSeconds(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _Replication.Contract.SetPauseBeforeUpgradeMinDurationSeconds(&_Replication.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_Replication *ReplicationTransactorSession) SetPauseBeforeUpgradeMinDurationSeconds(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _Replication.Contract.SetPauseBeforeUpgradeMinDurationSeconds(&_Replication.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xf030c99c.
//
// Solidity: function toPauseForUpgrade(address _teeId, address _claimBackAddress) payable returns()
func (_Replication *ReplicationTransactor) ToPauseForUpgrade(opts *bind.TransactOpts, _teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.contract.Transact(opts, "toPauseForUpgrade", _teeId, _claimBackAddress)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xf030c99c.
//
// Solidity: function toPauseForUpgrade(address _teeId, address _claimBackAddress) payable returns()
func (_Replication *ReplicationSession) ToPauseForUpgrade(_teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.Contract.ToPauseForUpgrade(&_Replication.TransactOpts, _teeId, _claimBackAddress)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xf030c99c.
//
// Solidity: function toPauseForUpgrade(address _teeId, address _claimBackAddress) payable returns()
func (_Replication *ReplicationTransactorSession) ToPauseForUpgrade(_teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Replication.Contract.ToPauseForUpgrade(&_Replication.TransactOpts, _teeId, _claimBackAddress)
}

// ReplicationAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the Replication contract.
type ReplicationAvailabilityCheckValidityExtendedIterator struct {
	Event *ReplicationAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

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
func (it *ReplicationAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationAvailabilityCheckValidityExtended)
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
		it.Event = new(ReplicationAvailabilityCheckValidityExtended)
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
func (it *ReplicationAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the Replication contract.
type ReplicationAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_Replication *ReplicationFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address, owner []common.Address) (*ReplicationAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationAvailabilityCheckValidityExtendedIterator{contract: _Replication.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_Replication *ReplicationFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *ReplicationAvailabilityCheckValidityExtended, teeId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationAvailabilityCheckValidityExtended)
				if err := _Replication.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
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
func (_Replication *ReplicationFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*ReplicationAvailabilityCheckValidityExtended, error) {
	event := new(ReplicationAvailabilityCheckValidityExtended)
	if err := _Replication.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Replication contract.
type ReplicationGovernanceCallTimelockedIterator struct {
	Event *ReplicationGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *ReplicationGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationGovernanceCallTimelocked)
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
		it.Event = new(ReplicationGovernanceCallTimelocked)
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
func (it *ReplicationGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Replication contract.
type ReplicationGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_Replication *ReplicationFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*ReplicationGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Replication.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &ReplicationGovernanceCallTimelockedIterator{contract: _Replication.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_Replication *ReplicationFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *ReplicationGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Replication.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationGovernanceCallTimelocked)
				if err := _Replication.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_Replication *ReplicationFilterer) ParseGovernanceCallTimelocked(log types.Log) (*ReplicationGovernanceCallTimelocked, error) {
	event := new(ReplicationGovernanceCallTimelocked)
	if err := _Replication.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Replication contract.
type ReplicationInitializedIterator struct {
	Event *ReplicationInitialized // Event containing the contract specifics and raw log

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
func (it *ReplicationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationInitialized)
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
		it.Event = new(ReplicationInitialized)
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
func (it *ReplicationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationInitialized represents a Initialized event raised by the Replication contract.
type ReplicationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Replication *ReplicationFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReplicationInitializedIterator, error) {

	logs, sub, err := _Replication.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReplicationInitializedIterator{contract: _Replication.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Replication *ReplicationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReplicationInitialized) (event.Subscription, error) {

	logs, sub, err := _Replication.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationInitialized)
				if err := _Replication.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Replication *ReplicationFilterer) ParseInitialized(log types.Log) (*ReplicationInitialized, error) {
	event := new(ReplicationInitialized)
	if err := _Replication.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator is returned from FilterPauseBeforeUpgradeMinDurationSecondsSet and is used to iterate over the raw logs and unpacked data for PauseBeforeUpgradeMinDurationSecondsSet events raised by the Replication contract.
type ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator struct {
	Event *ReplicationPauseBeforeUpgradeMinDurationSecondsSet // Event containing the contract specifics and raw log

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
func (it *ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationPauseBeforeUpgradeMinDurationSecondsSet)
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
		it.Event = new(ReplicationPauseBeforeUpgradeMinDurationSecondsSet)
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
func (it *ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationPauseBeforeUpgradeMinDurationSecondsSet represents a PauseBeforeUpgradeMinDurationSecondsSet event raised by the Replication contract.
type ReplicationPauseBeforeUpgradeMinDurationSecondsSet struct {
	PauseBeforeUpgradeMinDurationSeconds *big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterPauseBeforeUpgradeMinDurationSecondsSet is a free log retrieval operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_Replication *ReplicationFilterer) FilterPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.FilterOpts) (*ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator, error) {

	logs, sub, err := _Replication.contract.FilterLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return &ReplicationPauseBeforeUpgradeMinDurationSecondsSetIterator{contract: _Replication.contract, event: "PauseBeforeUpgradeMinDurationSecondsSet", logs: logs, sub: sub}, nil
}

// WatchPauseBeforeUpgradeMinDurationSecondsSet is a free log subscription operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_Replication *ReplicationFilterer) WatchPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.WatchOpts, sink chan<- *ReplicationPauseBeforeUpgradeMinDurationSecondsSet) (event.Subscription, error) {

	logs, sub, err := _Replication.contract.WatchLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationPauseBeforeUpgradeMinDurationSecondsSet)
				if err := _Replication.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
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

// ParsePauseBeforeUpgradeMinDurationSecondsSet is a log parse operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_Replication *ReplicationFilterer) ParsePauseBeforeUpgradeMinDurationSecondsSet(log types.Log) (*ReplicationPauseBeforeUpgradeMinDurationSecondsSet, error) {
	event := new(ReplicationPauseBeforeUpgradeMinDurationSecondsSet)
	if err := _Replication.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the Replication contract.
type ReplicationTeeInstructionsSentIterator struct {
	Event *ReplicationTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *ReplicationTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationTeeInstructionsSent)
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
		it.Event = new(ReplicationTeeInstructionsSent)
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
func (it *ReplicationTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationTeeInstructionsSent represents a TeeInstructionsSent event raised by the Replication contract.
type ReplicationTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_Replication *ReplicationFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*ReplicationTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationTeeInstructionsSentIterator{contract: _Replication.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_Replication *ReplicationFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *ReplicationTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationTeeInstructionsSent)
				if err := _Replication.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_Replication *ReplicationFilterer) ParseTeeInstructionsSent(log types.Log) (*ReplicationTeeInstructionsSent, error) {
	event := new(ReplicationTeeInstructionsSent)
	if err := _Replication.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationTeeMachinePausedForUpgradeIterator is returned from FilterTeeMachinePausedForUpgrade and is used to iterate over the raw logs and unpacked data for TeeMachinePausedForUpgrade events raised by the Replication contract.
type ReplicationTeeMachinePausedForUpgradeIterator struct {
	Event *ReplicationTeeMachinePausedForUpgrade // Event containing the contract specifics and raw log

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
func (it *ReplicationTeeMachinePausedForUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationTeeMachinePausedForUpgrade)
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
		it.Event = new(ReplicationTeeMachinePausedForUpgrade)
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
func (it *ReplicationTeeMachinePausedForUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationTeeMachinePausedForUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationTeeMachinePausedForUpgrade represents a TeeMachinePausedForUpgrade event raised by the Replication contract.
type ReplicationTeeMachinePausedForUpgrade struct {
	TeeId common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTeeMachinePausedForUpgrade is a free log retrieval operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_Replication *ReplicationFilterer) FilterTeeMachinePausedForUpgrade(opts *bind.FilterOpts, teeId []common.Address) (*ReplicationTeeMachinePausedForUpgradeIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "TeeMachinePausedForUpgrade", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationTeeMachinePausedForUpgradeIterator{contract: _Replication.contract, event: "TeeMachinePausedForUpgrade", logs: logs, sub: sub}, nil
}

// WatchTeeMachinePausedForUpgrade is a free log subscription operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_Replication *ReplicationFilterer) WatchTeeMachinePausedForUpgrade(opts *bind.WatchOpts, sink chan<- *ReplicationTeeMachinePausedForUpgrade, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "TeeMachinePausedForUpgrade", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationTeeMachinePausedForUpgrade)
				if err := _Replication.contract.UnpackLog(event, "TeeMachinePausedForUpgrade", log); err != nil {
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

// ParseTeeMachinePausedForUpgrade is a log parse operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_Replication *ReplicationFilterer) ParseTeeMachinePausedForUpgrade(log types.Log) (*ReplicationTeeMachinePausedForUpgrade, error) {
	event := new(ReplicationTeeMachinePausedForUpgrade)
	if err := _Replication.contract.UnpackLog(event, "TeeMachinePausedForUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationTeeMachineReplicationConfirmedIterator is returned from FilterTeeMachineReplicationConfirmed and is used to iterate over the raw logs and unpacked data for TeeMachineReplicationConfirmed events raised by the Replication contract.
type ReplicationTeeMachineReplicationConfirmedIterator struct {
	Event *ReplicationTeeMachineReplicationConfirmed // Event containing the contract specifics and raw log

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
func (it *ReplicationTeeMachineReplicationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationTeeMachineReplicationConfirmed)
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
		it.Event = new(ReplicationTeeMachineReplicationConfirmed)
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
func (it *ReplicationTeeMachineReplicationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationTeeMachineReplicationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationTeeMachineReplicationConfirmed represents a TeeMachineReplicationConfirmed event raised by the Replication contract.
type ReplicationTeeMachineReplicationConfirmed struct {
	OldTeeId common.Address
	NewTeeId common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineReplicationConfirmed is a free log retrieval operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_Replication *ReplicationFilterer) FilterTeeMachineReplicationConfirmed(opts *bind.FilterOpts, oldTeeId []common.Address, newTeeId []common.Address) (*ReplicationTeeMachineReplicationConfirmedIterator, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "TeeMachineReplicationConfirmed", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationTeeMachineReplicationConfirmedIterator{contract: _Replication.contract, event: "TeeMachineReplicationConfirmed", logs: logs, sub: sub}, nil
}

// WatchTeeMachineReplicationConfirmed is a free log subscription operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_Replication *ReplicationFilterer) WatchTeeMachineReplicationConfirmed(opts *bind.WatchOpts, sink chan<- *ReplicationTeeMachineReplicationConfirmed, oldTeeId []common.Address, newTeeId []common.Address) (event.Subscription, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "TeeMachineReplicationConfirmed", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationTeeMachineReplicationConfirmed)
				if err := _Replication.contract.UnpackLog(event, "TeeMachineReplicationConfirmed", log); err != nil {
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

// ParseTeeMachineReplicationConfirmed is a log parse operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_Replication *ReplicationFilterer) ParseTeeMachineReplicationConfirmed(log types.Log) (*ReplicationTeeMachineReplicationConfirmed, error) {
	event := new(ReplicationTeeMachineReplicationConfirmed)
	if err := _Replication.contract.UnpackLog(event, "TeeMachineReplicationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationTeeMachineReplicationTriggeredIterator is returned from FilterTeeMachineReplicationTriggered and is used to iterate over the raw logs and unpacked data for TeeMachineReplicationTriggered events raised by the Replication contract.
type ReplicationTeeMachineReplicationTriggeredIterator struct {
	Event *ReplicationTeeMachineReplicationTriggered // Event containing the contract specifics and raw log

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
func (it *ReplicationTeeMachineReplicationTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationTeeMachineReplicationTriggered)
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
		it.Event = new(ReplicationTeeMachineReplicationTriggered)
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
func (it *ReplicationTeeMachineReplicationTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationTeeMachineReplicationTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationTeeMachineReplicationTriggered represents a TeeMachineReplicationTriggered event raised by the Replication contract.
type ReplicationTeeMachineReplicationTriggered struct {
	OldTeeId             common.Address
	NewTeeId             common.Address
	MachinePathListNonce *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineReplicationTriggered is a free log retrieval operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 machinePathListNonce)
func (_Replication *ReplicationFilterer) FilterTeeMachineReplicationTriggered(opts *bind.FilterOpts, oldTeeId []common.Address, newTeeId []common.Address) (*ReplicationTeeMachineReplicationTriggeredIterator, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "TeeMachineReplicationTriggered", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationTeeMachineReplicationTriggeredIterator{contract: _Replication.contract, event: "TeeMachineReplicationTriggered", logs: logs, sub: sub}, nil
}

// WatchTeeMachineReplicationTriggered is a free log subscription operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 machinePathListNonce)
func (_Replication *ReplicationFilterer) WatchTeeMachineReplicationTriggered(opts *bind.WatchOpts, sink chan<- *ReplicationTeeMachineReplicationTriggered, oldTeeId []common.Address, newTeeId []common.Address) (event.Subscription, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "TeeMachineReplicationTriggered", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationTeeMachineReplicationTriggered)
				if err := _Replication.contract.UnpackLog(event, "TeeMachineReplicationTriggered", log); err != nil {
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

// ParseTeeMachineReplicationTriggered is a log parse operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 machinePathListNonce)
func (_Replication *ReplicationFilterer) ParseTeeMachineReplicationTriggered(log types.Log) (*ReplicationTeeMachineReplicationTriggered, error) {
	event := new(ReplicationTeeMachineReplicationTriggered)
	if err := _Replication.contract.UnpackLog(event, "TeeMachineReplicationTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicationTeeMachineStatusChangedIterator is returned from FilterTeeMachineStatusChanged and is used to iterate over the raw logs and unpacked data for TeeMachineStatusChanged events raised by the Replication contract.
type ReplicationTeeMachineStatusChangedIterator struct {
	Event *ReplicationTeeMachineStatusChanged // Event containing the contract specifics and raw log

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
func (it *ReplicationTeeMachineStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationTeeMachineStatusChanged)
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
		it.Event = new(ReplicationTeeMachineStatusChanged)
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
func (it *ReplicationTeeMachineStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationTeeMachineStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationTeeMachineStatusChanged represents a TeeMachineStatusChanged event raised by the Replication contract.
type ReplicationTeeMachineStatusChanged struct {
	TeeId     common.Address
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineStatusChanged is a free log retrieval operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_Replication *ReplicationFilterer) FilterTeeMachineStatusChanged(opts *bind.FilterOpts, teeId []common.Address, newStatus []uint8) (*ReplicationTeeMachineStatusChangedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _Replication.contract.FilterLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return &ReplicationTeeMachineStatusChangedIterator{contract: _Replication.contract, event: "TeeMachineStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTeeMachineStatusChanged is a free log subscription operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_Replication *ReplicationFilterer) WatchTeeMachineStatusChanged(opts *bind.WatchOpts, sink chan<- *ReplicationTeeMachineStatusChanged, teeId []common.Address, newStatus []uint8) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _Replication.contract.WatchLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationTeeMachineStatusChanged)
				if err := _Replication.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
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

// ParseTeeMachineStatusChanged is a log parse operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_Replication *ReplicationFilterer) ParseTeeMachineStatusChanged(log types.Log) (*ReplicationTeeMachineStatusChanged, error) {
	event := new(ReplicationTeeMachineStatusChanged)
	if err := _Replication.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
