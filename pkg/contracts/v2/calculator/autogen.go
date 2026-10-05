// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package calculator

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

// CalculatorMetaData contains all meta data concerning the Calculator contract.
var CalculatorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_wNatCapPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignNonPunishableDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignNonPunishableDurationBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignNoRewardsDurationBlocks\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"delegationFeeBIPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wNatWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wNatCappedWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nodeWeights\",\"type\":\"uint256[]\"}],\"name\":\"VoterRegistrationInfo\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"calculateBurnFactorPPM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_votePowerBlockNumber\",\"type\":\"uint256\"}],\"name\":\"calculateRegistrationWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_registrationWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enablePChainStakeMirror\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityManager\",\"outputs\":[{\"internalType\":\"contractIIEntityManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainStakeMirror\",\"outputs\":[{\"internalType\":\"contractIPChainStakeMirror\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainStakeMirrorEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_wNatCapPPM\",\"type\":\"uint24\"}],\"name\":\"setWNatCapPPM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicySignNoRewardsDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicySignNonPunishableDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicySignNonPunishableDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNat\",\"outputs\":[{\"internalType\":\"contractIWNat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNatCapPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNatDelegationFee\",\"outputs\":[{\"internalType\":\"contractIWNatDelegationFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "Calculator",
}

// Calculator is an auto generated Go binding around an Ethereum contract.
type Calculator struct {
	abi abi.ABI
}

// NewCalculator creates a new instance of Calculator.
func NewCalculator() *Calculator {
	parsed, err := CalculatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Calculator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Calculator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _wNatCapPPM, uint64 _signingPolicySignNonPunishableDurationSeconds, uint64 _signingPolicySignNonPunishableDurationBlocks, uint64 _signingPolicySignNoRewardsDurationBlocks) returns()
func (calculator *Calculator) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _wNatCapPPM *big.Int, _signingPolicySignNonPunishableDurationSeconds uint64, _signingPolicySignNonPunishableDurationBlocks uint64, _signingPolicySignNoRewardsDurationBlocks uint64) []byte {
	enc, err := calculator.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _wNatCapPPM, _signingPolicySignNonPunishableDurationSeconds, _signingPolicySignNonPunishableDurationBlocks, _signingPolicySignNoRewardsDurationBlocks)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCalculateBurnFactorPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9350f57c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateBurnFactorPPM(uint24 _rewardEpochId, address _voter) view returns(uint256)
func (calculator *Calculator) PackCalculateBurnFactorPPM(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := calculator.abi.Pack("calculateBurnFactorPPM", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateBurnFactorPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9350f57c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateBurnFactorPPM(uint24 _rewardEpochId, address _voter) view returns(uint256)
func (calculator *Calculator) TryPackCalculateBurnFactorPPM(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return calculator.abi.Pack("calculateBurnFactorPPM", rewardEpochId, voter)
}

// UnpackCalculateBurnFactorPPM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9350f57c.
//
// Solidity: function calculateBurnFactorPPM(uint24 _rewardEpochId, address _voter) view returns(uint256)
func (calculator *Calculator) UnpackCalculateBurnFactorPPM(data []byte) (*big.Int, error) {
	out, err := calculator.abi.Unpack("calculateBurnFactorPPM", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCalculateRegistrationWeight is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb65185d6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateRegistrationWeight(address _voter, uint24 _rewardEpochId, uint256 _votePowerBlockNumber) returns(uint256 _registrationWeight)
func (calculator *Calculator) PackCalculateRegistrationWeight(voter common.Address, rewardEpochId *big.Int, votePowerBlockNumber *big.Int) []byte {
	enc, err := calculator.abi.Pack("calculateRegistrationWeight", voter, rewardEpochId, votePowerBlockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateRegistrationWeight is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb65185d6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateRegistrationWeight(address _voter, uint24 _rewardEpochId, uint256 _votePowerBlockNumber) returns(uint256 _registrationWeight)
func (calculator *Calculator) TryPackCalculateRegistrationWeight(voter common.Address, rewardEpochId *big.Int, votePowerBlockNumber *big.Int) ([]byte, error) {
	return calculator.abi.Pack("calculateRegistrationWeight", voter, rewardEpochId, votePowerBlockNumber)
}

// UnpackCalculateRegistrationWeight is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb65185d6.
//
// Solidity: function calculateRegistrationWeight(address _voter, uint24 _rewardEpochId, uint256 _votePowerBlockNumber) returns(uint256 _registrationWeight)
func (calculator *Calculator) UnpackCalculateRegistrationWeight(data []byte) (*big.Int, error) {
	out, err := calculator.abi.Unpack("calculateRegistrationWeight", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (calculator *Calculator) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := calculator.abi.Pack("cancelGovernanceCall", selector)
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
func (calculator *Calculator) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return calculator.abi.Pack("cancelGovernanceCall", selector)
}

// PackEnablePChainStakeMirror is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb006b4e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enablePChainStakeMirror() returns()
func (calculator *Calculator) PackEnablePChainStakeMirror() []byte {
	enc, err := calculator.abi.Pack("enablePChainStakeMirror")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnablePChainStakeMirror is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb006b4e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enablePChainStakeMirror() returns()
func (calculator *Calculator) TryPackEnablePChainStakeMirror() ([]byte, error) {
	return calculator.abi.Pack("enablePChainStakeMirror")
}

// PackEntityManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50b1d61b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function entityManager() view returns(address)
func (calculator *Calculator) PackEntityManager() []byte {
	enc, err := calculator.abi.Pack("entityManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEntityManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50b1d61b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function entityManager() view returns(address)
func (calculator *Calculator) TryPackEntityManager() ([]byte, error) {
	return calculator.abi.Pack("entityManager")
}

// UnpackEntityManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (calculator *Calculator) UnpackEntityManager(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("entityManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (calculator *Calculator) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := calculator.abi.Pack("executeGovernanceCall", selector)
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
func (calculator *Calculator) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return calculator.abi.Pack("executeGovernanceCall", selector)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (calculator *Calculator) PackFlareSystemsManager() []byte {
	enc, err := calculator.abi.Pack("flareSystemsManager")
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
func (calculator *Calculator) TryPackFlareSystemsManager() ([]byte, error) {
	return calculator.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (calculator *Calculator) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("flareSystemsManager", data)
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
func (calculator *Calculator) PackGetAddressUpdater() []byte {
	enc, err := calculator.abi.Pack("getAddressUpdater")
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
func (calculator *Calculator) TryPackGetAddressUpdater() ([]byte, error) {
	return calculator.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (calculator *Calculator) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (calculator *Calculator) PackGovernance() []byte {
	enc, err := calculator.abi.Pack("governance")
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
func (calculator *Calculator) TryPackGovernance() ([]byte, error) {
	return calculator.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (calculator *Calculator) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("governance", data)
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
func (calculator *Calculator) PackGovernanceSettings() []byte {
	enc, err := calculator.abi.Pack("governanceSettings")
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
func (calculator *Calculator) TryPackGovernanceSettings() ([]byte, error) {
	return calculator.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (calculator *Calculator) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (calculator *Calculator) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := calculator.abi.Pack("initialise", governanceSettings, initialGovernance)
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
func (calculator *Calculator) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return calculator.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (calculator *Calculator) PackIsExecutor(address common.Address) []byte {
	enc, err := calculator.abi.Pack("isExecutor", address)
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
func (calculator *Calculator) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return calculator.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (calculator *Calculator) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := calculator.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPChainStakeMirror is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62d9c89a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (calculator *Calculator) PackPChainStakeMirror() []byte {
	enc, err := calculator.abi.Pack("pChainStakeMirror")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPChainStakeMirror is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62d9c89a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (calculator *Calculator) TryPackPChainStakeMirror() ([]byte, error) {
	return calculator.abi.Pack("pChainStakeMirror")
}

// UnpackPChainStakeMirror is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62d9c89a.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (calculator *Calculator) UnpackPChainStakeMirror(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("pChainStakeMirror", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPChainStakeMirrorEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bf756c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (calculator *Calculator) PackPChainStakeMirrorEnabled() []byte {
	enc, err := calculator.abi.Pack("pChainStakeMirrorEnabled")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPChainStakeMirrorEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bf756c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (calculator *Calculator) TryPackPChainStakeMirrorEnabled() ([]byte, error) {
	return calculator.abi.Pack("pChainStakeMirrorEnabled")
}

// UnpackPChainStakeMirrorEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7bf756c9.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (calculator *Calculator) UnpackPChainStakeMirrorEnabled(data []byte) (bool, error) {
	out, err := calculator.abi.Unpack("pChainStakeMirrorEnabled", data)
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
func (calculator *Calculator) PackProductionMode() []byte {
	enc, err := calculator.abi.Pack("productionMode")
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
func (calculator *Calculator) TryPackProductionMode() ([]byte, error) {
	return calculator.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (calculator *Calculator) UnpackProductionMode(data []byte) (bool, error) {
	out, err := calculator.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetWNatCapPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d7cf608.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setWNatCapPPM(uint24 _wNatCapPPM) returns()
func (calculator *Calculator) PackSetWNatCapPPM(wNatCapPPM *big.Int) []byte {
	enc, err := calculator.abi.Pack("setWNatCapPPM", wNatCapPPM)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetWNatCapPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d7cf608.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setWNatCapPPM(uint24 _wNatCapPPM) returns()
func (calculator *Calculator) TryPackSetWNatCapPPM(wNatCapPPM *big.Int) ([]byte, error) {
	return calculator.abi.Pack("setWNatCapPPM", wNatCapPPM)
}

// PackSigningPolicySignNoRewardsDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dd1018c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicySignNoRewardsDurationBlocks() view returns(uint64)
func (calculator *Calculator) PackSigningPolicySignNoRewardsDurationBlocks() []byte {
	enc, err := calculator.abi.Pack("signingPolicySignNoRewardsDurationBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicySignNoRewardsDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dd1018c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicySignNoRewardsDurationBlocks() view returns(uint64)
func (calculator *Calculator) TryPackSigningPolicySignNoRewardsDurationBlocks() ([]byte, error) {
	return calculator.abi.Pack("signingPolicySignNoRewardsDurationBlocks")
}

// UnpackSigningPolicySignNoRewardsDurationBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9dd1018c.
//
// Solidity: function signingPolicySignNoRewardsDurationBlocks() view returns(uint64)
func (calculator *Calculator) UnpackSigningPolicySignNoRewardsDurationBlocks(data []byte) (uint64, error) {
	out, err := calculator.abi.Unpack("signingPolicySignNoRewardsDurationBlocks", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSigningPolicySignNonPunishableDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87fd1ba1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicySignNonPunishableDurationBlocks() view returns(uint64)
func (calculator *Calculator) PackSigningPolicySignNonPunishableDurationBlocks() []byte {
	enc, err := calculator.abi.Pack("signingPolicySignNonPunishableDurationBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicySignNonPunishableDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87fd1ba1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicySignNonPunishableDurationBlocks() view returns(uint64)
func (calculator *Calculator) TryPackSigningPolicySignNonPunishableDurationBlocks() ([]byte, error) {
	return calculator.abi.Pack("signingPolicySignNonPunishableDurationBlocks")
}

// UnpackSigningPolicySignNonPunishableDurationBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87fd1ba1.
//
// Solidity: function signingPolicySignNonPunishableDurationBlocks() view returns(uint64)
func (calculator *Calculator) UnpackSigningPolicySignNonPunishableDurationBlocks(data []byte) (uint64, error) {
	out, err := calculator.abi.Unpack("signingPolicySignNonPunishableDurationBlocks", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSigningPolicySignNonPunishableDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x96ca1472.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicySignNonPunishableDurationSeconds() view returns(uint64)
func (calculator *Calculator) PackSigningPolicySignNonPunishableDurationSeconds() []byte {
	enc, err := calculator.abi.Pack("signingPolicySignNonPunishableDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicySignNonPunishableDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x96ca1472.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicySignNonPunishableDurationSeconds() view returns(uint64)
func (calculator *Calculator) TryPackSigningPolicySignNonPunishableDurationSeconds() ([]byte, error) {
	return calculator.abi.Pack("signingPolicySignNonPunishableDurationSeconds")
}

// UnpackSigningPolicySignNonPunishableDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x96ca1472.
//
// Solidity: function signingPolicySignNonPunishableDurationSeconds() view returns(uint64)
func (calculator *Calculator) UnpackSigningPolicySignNonPunishableDurationSeconds(data []byte) (uint64, error) {
	out, err := calculator.abi.Unpack("signingPolicySignNonPunishableDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSqrt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x677342ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sqrt(uint256 _x) pure returns(uint128)
func (calculator *Calculator) PackSqrt(x *big.Int) []byte {
	enc, err := calculator.abi.Pack("sqrt", x)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSqrt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x677342ce.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sqrt(uint256 _x) pure returns(uint128)
func (calculator *Calculator) TryPackSqrt(x *big.Int) ([]byte, error) {
	return calculator.abi.Pack("sqrt", x)
}

// UnpackSqrt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x677342ce.
//
// Solidity: function sqrt(uint256 _x) pure returns(uint128)
func (calculator *Calculator) UnpackSqrt(data []byte) (*big.Int, error) {
	out, err := calculator.abi.Unpack("sqrt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (calculator *Calculator) PackSwitchToProductionMode() []byte {
	enc, err := calculator.abi.Pack("switchToProductionMode")
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
func (calculator *Calculator) TryPackSwitchToProductionMode() ([]byte, error) {
	return calculator.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (calculator *Calculator) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := calculator.abi.Pack("timelockedCalls", selector)
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
func (calculator *Calculator) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return calculator.abi.Pack("timelockedCalls", selector)
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
func (calculator *Calculator) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := calculator.abi.Unpack("timelockedCalls", data)
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
func (calculator *Calculator) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := calculator.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (calculator *Calculator) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return calculator.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackVoterRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe60040e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistry() view returns(address)
func (calculator *Calculator) PackVoterRegistry() []byte {
	enc, err := calculator.abi.Pack("voterRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoterRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe60040e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voterRegistry() view returns(address)
func (calculator *Calculator) TryPackVoterRegistry() ([]byte, error) {
	return calculator.abi.Pack("voterRegistry")
}

// UnpackVoterRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (calculator *Calculator) UnpackVoterRegistry(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("voterRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWNat is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9edbf007.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function wNat() view returns(address)
func (calculator *Calculator) PackWNat() []byte {
	enc, err := calculator.abi.Pack("wNat")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWNat is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9edbf007.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function wNat() view returns(address)
func (calculator *Calculator) TryPackWNat() ([]byte, error) {
	return calculator.abi.Pack("wNat")
}

// UnpackWNat is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (calculator *Calculator) UnpackWNat(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("wNat", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWNatCapPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5edf7596.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function wNatCapPPM() view returns(uint24)
func (calculator *Calculator) PackWNatCapPPM() []byte {
	enc, err := calculator.abi.Pack("wNatCapPPM")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWNatCapPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5edf7596.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function wNatCapPPM() view returns(uint24)
func (calculator *Calculator) TryPackWNatCapPPM() ([]byte, error) {
	return calculator.abi.Pack("wNatCapPPM")
}

// UnpackWNatCapPPM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5edf7596.
//
// Solidity: function wNatCapPPM() view returns(uint24)
func (calculator *Calculator) UnpackWNatCapPPM(data []byte) (*big.Int, error) {
	out, err := calculator.abi.Unpack("wNatCapPPM", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackWNatDelegationFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87c5ab51.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function wNatDelegationFee() view returns(address)
func (calculator *Calculator) PackWNatDelegationFee() []byte {
	enc, err := calculator.abi.Pack("wNatDelegationFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWNatDelegationFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87c5ab51.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function wNatDelegationFee() view returns(address)
func (calculator *Calculator) TryPackWNatDelegationFee() ([]byte, error) {
	return calculator.abi.Pack("wNatDelegationFee")
}

// UnpackWNatDelegationFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87c5ab51.
//
// Solidity: function wNatDelegationFee() view returns(address)
func (calculator *Calculator) UnpackWNatDelegationFee(data []byte) (common.Address, error) {
	out, err := calculator.abi.Unpack("wNatDelegationFee", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// CalculatorGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Calculator contract.
type CalculatorGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const CalculatorGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (CalculatorGovernanceCallTimelocked) ContractEventName() string {
	return CalculatorGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (calculator *Calculator) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*CalculatorGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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

// CalculatorGovernanceInitialised represents a GovernanceInitialised event raised by the Calculator contract.
type CalculatorGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const CalculatorGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (CalculatorGovernanceInitialised) ContractEventName() string {
	return CalculatorGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (calculator *Calculator) UnpackGovernanceInitialisedEvent(log *types.Log) (*CalculatorGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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

// CalculatorGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Calculator contract.
type CalculatorGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const CalculatorGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (CalculatorGovernedProductionModeEntered) ContractEventName() string {
	return CalculatorGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (calculator *Calculator) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*CalculatorGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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

// CalculatorTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Calculator contract.
type CalculatorTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const CalculatorTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (CalculatorTimelockedGovernanceCallCanceled) ContractEventName() string {
	return CalculatorTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (calculator *Calculator) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*CalculatorTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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

// CalculatorTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Calculator contract.
type CalculatorTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const CalculatorTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (CalculatorTimelockedGovernanceCallExecuted) ContractEventName() string {
	return CalculatorTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (calculator *Calculator) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*CalculatorTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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

// CalculatorVoterRegistrationInfo represents a VoterRegistrationInfo event raised by the Calculator contract.
type CalculatorVoterRegistrationInfo struct {
	Voter             common.Address
	RewardEpochId     *big.Int
	DelegationAddress common.Address
	DelegationFeeBIPS uint16
	WNatWeight        *big.Int
	WNatCappedWeight  *big.Int
	NodeIds           [][20]byte
	NodeWeights       []*big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const CalculatorVoterRegistrationInfoEventName = "VoterRegistrationInfo"

// ContractEventName returns the user-defined event name.
func (CalculatorVoterRegistrationInfo) ContractEventName() string {
	return CalculatorVoterRegistrationInfoEventName
}

// UnpackVoterRegistrationInfoEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoterRegistrationInfo(address indexed voter, uint24 indexed rewardEpochId, address delegationAddress, uint16 delegationFeeBIPS, uint256 wNatWeight, uint256 wNatCappedWeight, bytes20[] nodeIds, uint256[] nodeWeights)
func (calculator *Calculator) UnpackVoterRegistrationInfoEvent(log *types.Log) (*CalculatorVoterRegistrationInfo, error) {
	event := "VoterRegistrationInfo"
	if len(log.Topics) == 0 || log.Topics[0] != calculator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CalculatorVoterRegistrationInfo)
	if len(log.Data) > 0 {
		if err := calculator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range calculator.abi.Events[event].Inputs {
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
