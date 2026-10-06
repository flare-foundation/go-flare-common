// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package submission

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// SubmissionMetaData contains all meta data concerning the Submission contract.
var SubmissionMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_submit3MethodEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewVotingRoundInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandomWithQuality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isSecureRandom\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandomWithQualityAndTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isSecureRandom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_randomTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_submit1Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submit2Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submit3Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_submitSignaturesAddresses\",\"type\":\"address[]\"}],\"name\":\"initNewVotingRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSubmit3MethodEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAndPassContract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_submitAndPassSelector\",\"type\":\"bytes4\"}],\"name\":\"setSubmitAndPassData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3MethodEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitAndPass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitAndPassContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitAndPassSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Submission",
}

// Submission is an auto generated Go binding around an Ethereum contract.
type Submission struct {
	abi abi.ABI
}

// NewSubmission creates a new instance of Submission.
func NewSubmission() *Submission {
	parsed, err := SubmissionMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Submission{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Submission) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, bool _submit3MethodEnabled) returns()
func (submission *Submission) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _submit3MethodEnabled bool) []byte {
	enc, err := submission.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _submit3MethodEnabled)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (submission *Submission) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := submission.abi.Pack("cancelGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (submission *Submission) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return submission.abi.Pack("cancelGovernanceCall", selector)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (submission *Submission) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := submission.abi.Pack("executeGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (submission *Submission) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return submission.abi.Pack("executeGovernanceCall", selector)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (submission *Submission) PackFlareSystemsManager() []byte {
	enc, err := submission.abi.Pack("flareSystemsManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareSystemsManager() view returns(address)
func (submission *Submission) TryPackFlareSystemsManager() ([]byte, error) {
	return submission.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (submission *Submission) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (submission *Submission) PackGetAddressUpdater() []byte {
	enc, err := submission.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (submission *Submission) TryPackGetAddressUpdater() ([]byte, error) {
	return submission.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (submission *Submission) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetCurrentRandom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd89601fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (submission *Submission) PackGetCurrentRandom() []byte {
	enc, err := submission.abi.Pack("getCurrentRandom")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentRandom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd89601fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (submission *Submission) TryPackGetCurrentRandom() ([]byte, error) {
	return submission.abi.Pack("getCurrentRandom")
}

// UnpackGetCurrentRandom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256 _currentRandom)
func (submission *Submission) UnpackGetCurrentRandom(data []byte) (*big.Int, error) {
	out, err := submission.abi.Unpack("getCurrentRandom", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentRandomWithQuality is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa978fb6b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _isSecureRandom)
func (submission *Submission) PackGetCurrentRandomWithQuality() []byte {
	enc, err := submission.abi.Pack("getCurrentRandomWithQuality")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentRandomWithQuality is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa978fb6b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _isSecureRandom)
func (submission *Submission) TryPackGetCurrentRandomWithQuality() ([]byte, error) {
	return submission.abi.Pack("getCurrentRandomWithQuality")
}

// GetCurrentRandomWithQualityOutput serves as a container for the return parameters of contract
// method GetCurrentRandomWithQuality.
type GetCurrentRandomWithQualityOutput struct {
	CurrentRandom  *big.Int
	IsSecureRandom bool
}

// UnpackGetCurrentRandomWithQuality is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa978fb6b.
//
// Solidity: function getCurrentRandomWithQuality() view returns(uint256 _currentRandom, bool _isSecureRandom)
func (submission *Submission) UnpackGetCurrentRandomWithQuality(data []byte) (GetCurrentRandomWithQualityOutput, error) {
	out, err := submission.abi.Unpack("getCurrentRandomWithQuality", data)
	outstruct := new(GetCurrentRandomWithQualityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CurrentRandom = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsSecureRandom = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetCurrentRandomWithQualityAndTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf9fbc3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentRandomWithQualityAndTimestamp() view returns(uint256 _currentRandom, bool _isSecureRandom, uint256 _randomTimestamp)
func (submission *Submission) PackGetCurrentRandomWithQualityAndTimestamp() []byte {
	enc, err := submission.abi.Pack("getCurrentRandomWithQualityAndTimestamp")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentRandomWithQualityAndTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf9fbc3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentRandomWithQualityAndTimestamp() view returns(uint256 _currentRandom, bool _isSecureRandom, uint256 _randomTimestamp)
func (submission *Submission) TryPackGetCurrentRandomWithQualityAndTimestamp() ([]byte, error) {
	return submission.abi.Pack("getCurrentRandomWithQualityAndTimestamp")
}

// GetCurrentRandomWithQualityAndTimestampOutput serves as a container for the return parameters of contract
// method GetCurrentRandomWithQualityAndTimestamp.
type GetCurrentRandomWithQualityAndTimestampOutput struct {
	CurrentRandom   *big.Int
	IsSecureRandom  bool
	RandomTimestamp *big.Int
}

// UnpackGetCurrentRandomWithQualityAndTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf9fbc3e.
//
// Solidity: function getCurrentRandomWithQualityAndTimestamp() view returns(uint256 _currentRandom, bool _isSecureRandom, uint256 _randomTimestamp)
func (submission *Submission) UnpackGetCurrentRandomWithQualityAndTimestamp(data []byte) (GetCurrentRandomWithQualityAndTimestampOutput, error) {
	out, err := submission.abi.Unpack("getCurrentRandomWithQualityAndTimestamp", data)
	outstruct := new(GetCurrentRandomWithQualityAndTimestampOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CurrentRandom = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsSecureRandom = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomTimestamp = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (submission *Submission) PackGovernance() []byte {
	enc, err := submission.abi.Pack("governance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governance() view returns(address)
func (submission *Submission) TryPackGovernance() ([]byte, error) {
	return submission.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (submission *Submission) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("governance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governanceSettings() view returns(address)
func (submission *Submission) PackGovernanceSettings() []byte {
	enc, err := submission.abi.Pack("governanceSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governanceSettings() view returns(address)
func (submission *Submission) TryPackGovernanceSettings() ([]byte, error) {
	return submission.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (submission *Submission) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitNewVotingRound is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ae8a2f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initNewVotingRound(address[] _submit1Addresses, address[] _submit2Addresses, address[] _submit3Addresses, address[] _submitSignaturesAddresses) returns()
func (submission *Submission) PackInitNewVotingRound(submit1Addresses []common.Address, submit2Addresses []common.Address, submit3Addresses []common.Address, submitSignaturesAddresses []common.Address) []byte {
	enc, err := submission.abi.Pack("initNewVotingRound", submit1Addresses, submit2Addresses, submit3Addresses, submitSignaturesAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitNewVotingRound is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ae8a2f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initNewVotingRound(address[] _submit1Addresses, address[] _submit2Addresses, address[] _submit3Addresses, address[] _submitSignaturesAddresses) returns()
func (submission *Submission) TryPackInitNewVotingRound(submit1Addresses []common.Address, submit2Addresses []common.Address, submit3Addresses []common.Address, submitSignaturesAddresses []common.Address) ([]byte, error) {
	return submission.abi.Pack("initNewVotingRound", submit1Addresses, submit2Addresses, submit3Addresses, submitSignaturesAddresses)
}

// PackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (submission *Submission) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := submission.abi.Pack("initialise", governanceSettings, initialGovernance)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (submission *Submission) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return submission.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (submission *Submission) PackIsExecutor(address common.Address) []byte {
	enc, err := submission.abi.Pack("isExecutor", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (submission *Submission) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return submission.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (submission *Submission) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (submission *Submission) PackProductionMode() []byte {
	enc, err := submission.abi.Pack("productionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function productionMode() view returns(bool)
func (submission *Submission) TryPackProductionMode() ([]byte, error) {
	return submission.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (submission *Submission) UnpackProductionMode(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay() view returns(address)
func (submission *Submission) PackRelay() []byte {
	enc, err := submission.abi.Pack("relay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay() view returns(address)
func (submission *Submission) TryPackRelay() ([]byte, error) {
	return submission.abi.Pack("relay")
}

// UnpackRelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (submission *Submission) UnpackRelay(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("relay", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetSubmit3MethodEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x941877d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubmit3MethodEnabled(bool _enabled) returns()
func (submission *Submission) PackSetSubmit3MethodEnabled(enabled bool) []byte {
	enc, err := submission.abi.Pack("setSubmit3MethodEnabled", enabled)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSubmit3MethodEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x941877d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSubmit3MethodEnabled(bool _enabled) returns()
func (submission *Submission) TryPackSetSubmit3MethodEnabled(enabled bool) ([]byte, error) {
	return submission.abi.Pack("setSubmit3MethodEnabled", enabled)
}

// PackSetSubmitAndPassData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ee7fe4d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubmitAndPassData(address _submitAndPassContract, bytes4 _submitAndPassSelector) returns()
func (submission *Submission) PackSetSubmitAndPassData(submitAndPassContract common.Address, submitAndPassSelector [4]byte) []byte {
	enc, err := submission.abi.Pack("setSubmitAndPassData", submitAndPassContract, submitAndPassSelector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSubmitAndPassData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ee7fe4d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSubmitAndPassData(address _submitAndPassContract, bytes4 _submitAndPassSelector) returns()
func (submission *Submission) TryPackSetSubmitAndPassData(submitAndPassContract common.Address, submitAndPassSelector [4]byte) ([]byte, error) {
	return submission.abi.Pack("setSubmitAndPassData", submitAndPassContract, submitAndPassSelector)
}

// PackSubmit1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c532fae.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submit1() returns(bool)
func (submission *Submission) PackSubmit1() []byte {
	enc, err := submission.abi.Pack("submit1")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmit1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c532fae.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submit1() returns(bool)
func (submission *Submission) TryPackSubmit1() ([]byte, error) {
	return submission.abi.Pack("submit1")
}

// UnpackSubmit1 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6c532fae.
//
// Solidity: function submit1() returns(bool)
func (submission *Submission) UnpackSubmit1(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submit1", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d00c9fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submit2() returns(bool)
func (submission *Submission) PackSubmit2() []byte {
	enc, err := submission.abi.Pack("submit2")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d00c9fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submit2() returns(bool)
func (submission *Submission) TryPackSubmit2() ([]byte, error) {
	return submission.abi.Pack("submit2")
}

// UnpackSubmit2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9d00c9fd.
//
// Solidity: function submit2() returns(bool)
func (submission *Submission) UnpackSubmit2(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submit2", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmit3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1b157e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submit3() returns(bool)
func (submission *Submission) PackSubmit3() []byte {
	enc, err := submission.abi.Pack("submit3")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmit3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1b157e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submit3() returns(bool)
func (submission *Submission) TryPackSubmit3() ([]byte, error) {
	return submission.abi.Pack("submit3")
}

// UnpackSubmit3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1b157e7.
//
// Solidity: function submit3() returns(bool)
func (submission *Submission) UnpackSubmit3(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submit3", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmit3MethodEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32de7a9f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (submission *Submission) PackSubmit3MethodEnabled() []byte {
	enc, err := submission.abi.Pack("submit3MethodEnabled")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmit3MethodEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32de7a9f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (submission *Submission) TryPackSubmit3MethodEnabled() ([]byte, error) {
	return submission.abi.Pack("submit3MethodEnabled")
}

// UnpackSubmit3MethodEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32de7a9f.
//
// Solidity: function submit3MethodEnabled() view returns(bool)
func (submission *Submission) UnpackSubmit3MethodEnabled(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submit3MethodEnabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmitAndPass is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x833bf6c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (submission *Submission) PackSubmitAndPass(data []byte) []byte {
	enc, err := submission.abi.Pack("submitAndPass", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitAndPass is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x833bf6c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (submission *Submission) TryPackSubmitAndPass(data []byte) ([]byte, error) {
	return submission.abi.Pack("submitAndPass", data)
}

// UnpackSubmitAndPass is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x833bf6c0.
//
// Solidity: function submitAndPass(bytes _data) returns(bool)
func (submission *Submission) UnpackSubmitAndPass(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submitAndPass", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmitAndPassContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93953af1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitAndPassContract() view returns(address)
func (submission *Submission) PackSubmitAndPassContract() []byte {
	enc, err := submission.abi.Pack("submitAndPassContract")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitAndPassContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93953af1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitAndPassContract() view returns(address)
func (submission *Submission) TryPackSubmitAndPassContract() ([]byte, error) {
	return submission.abi.Pack("submitAndPassContract")
}

// UnpackSubmitAndPassContract is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x93953af1.
//
// Solidity: function submitAndPassContract() view returns(address)
func (submission *Submission) UnpackSubmitAndPassContract(data []byte) (common.Address, error) {
	out, err := submission.abi.Unpack("submitAndPassContract", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSubmitAndPassSelector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xafd7f821.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (submission *Submission) PackSubmitAndPassSelector() []byte {
	enc, err := submission.abi.Pack("submitAndPassSelector")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitAndPassSelector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xafd7f821.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (submission *Submission) TryPackSubmitAndPassSelector() ([]byte, error) {
	return submission.abi.Pack("submitAndPassSelector")
}

// UnpackSubmitAndPassSelector is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xafd7f821.
//
// Solidity: function submitAndPassSelector() view returns(bytes4)
func (submission *Submission) UnpackSubmitAndPassSelector(data []byte) ([4]byte, error) {
	out, err := submission.abi.Unpack("submitAndPassSelector", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackSubmitSignatures is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57eed580.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitSignatures() returns(bool)
func (submission *Submission) PackSubmitSignatures() []byte {
	enc, err := submission.abi.Pack("submitSignatures")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitSignatures is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57eed580.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitSignatures() returns(bool)
func (submission *Submission) TryPackSubmitSignatures() ([]byte, error) {
	return submission.abi.Pack("submitSignatures")
}

// UnpackSubmitSignatures is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57eed580.
//
// Solidity: function submitSignatures() returns(bool)
func (submission *Submission) UnpackSubmitSignatures(data []byte) (bool, error) {
	out, err := submission.abi.Unpack("submitSignatures", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (submission *Submission) PackSwitchToProductionMode() []byte {
	enc, err := submission.abi.Pack("switchToProductionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToProductionMode() returns()
func (submission *Submission) TryPackSwitchToProductionMode() ([]byte, error) {
	return submission.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (submission *Submission) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := submission.abi.Pack("timelockedCalls", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (submission *Submission) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return submission.abi.Pack("timelockedCalls", selector)
}

// TimelockedCallsOutput serves as a container for the return parameters of contract
// method TimelockedCalls.
type TimelockedCallsOutput struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}

// UnpackTimelockedCalls is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (submission *Submission) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := submission.abi.Unpack("timelockedCalls", data)
	outstruct := new(TimelockedCallsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AllowedAfterTimestamp = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (submission *Submission) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := submission.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (submission *Submission) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return submission.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// SubmissionGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Submission contract.
type SubmissionGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const SubmissionGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (SubmissionGovernanceCallTimelocked) ContractEventName() string {
	return SubmissionGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (submission *Submission) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*SubmissionGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SubmissionGovernanceInitialised represents a GovernanceInitialised event raised by the Submission contract.
type SubmissionGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const SubmissionGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (SubmissionGovernanceInitialised) ContractEventName() string {
	return SubmissionGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (submission *Submission) UnpackGovernanceInitialisedEvent(log *types.Log) (*SubmissionGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SubmissionGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Submission contract.
type SubmissionGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const SubmissionGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (SubmissionGovernedProductionModeEntered) ContractEventName() string {
	return SubmissionGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (submission *Submission) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*SubmissionGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SubmissionNewVotingRoundInitiated represents a NewVotingRoundInitiated event raised by the Submission contract.
type SubmissionNewVotingRoundInitiated struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const SubmissionNewVotingRoundInitiatedEventName = "NewVotingRoundInitiated"

// ContractEventName returns the user-defined event name.
func (SubmissionNewVotingRoundInitiated) ContractEventName() string {
	return SubmissionNewVotingRoundInitiatedEventName
}

// UnpackNewVotingRoundInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewVotingRoundInitiated()
func (submission *Submission) UnpackNewVotingRoundInitiatedEvent(log *types.Log) (*SubmissionNewVotingRoundInitiated, error) {
	event := "NewVotingRoundInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionNewVotingRoundInitiated)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SubmissionTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Submission contract.
type SubmissionTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SubmissionTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (SubmissionTimelockedGovernanceCallCanceled) ContractEventName() string {
	return SubmissionTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (submission *Submission) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*SubmissionTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SubmissionTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Submission contract.
type SubmissionTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SubmissionTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (SubmissionTimelockedGovernanceCallExecuted) ContractEventName() string {
	return SubmissionTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (submission *Submission) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*SubmissionTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != submission.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SubmissionTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := submission.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range submission.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
