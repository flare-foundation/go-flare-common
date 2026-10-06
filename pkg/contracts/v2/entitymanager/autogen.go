// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entitymanager

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

// IEntityManagerVoterAddresses is an auto generated low-level Go binding around an user-defined struct.
type IEntityManagerVoterAddresses struct {
	SubmitAddress           common.Address
	SubmitSignaturesAddress common.Address
	SigningPolicyAddress    common.Address
}

// EntityManagerMetaData contains all meta data concerning the EntityManager contract.
var EntityManagerMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"}],\"name\":\"DelegationAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"}],\"name\":\"DelegationAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxNodeIdsPerEntity\",\"type\":\"uint256\"}],\"name\":\"MaxNodeIdsPerEntitySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"}],\"name\":\"NodeIdRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"}],\"name\":\"NodeIdUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part2\",\"type\":\"bytes32\"}],\"name\":\"PublicKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part2\",\"type\":\"bytes32\"}],\"name\":\"PublicKeyUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"SigningPolicyAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"SigningPolicyAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"}],\"name\":\"SubmitAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"}],\"name\":\"SubmitAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"SubmitSignaturesAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"SubmitSignaturesAddressRegistrationConfirmed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmDelegationAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSigningPolicyAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSubmitAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSubmitSignaturesAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getDelegationAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getDelegationAddressOfAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getNodeIdsOf\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getNodeIdsOfAt\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getPublicKeyOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPublicKeyOfAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"internalType\":\"structIEntityManager.VoterAddresses\",\"name\":\"_addresses\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterAddressesAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"internalType\":\"structIEntityManager.VoterAddresses\",\"name\":\"_addresses\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForDelegationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForNodeId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForPublicKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSigningPolicyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSubmitAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSubmitSignaturesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationAddress\",\"type\":\"address\"}],\"name\":\"proposeDelegationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"proposeSigningPolicyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"name\":\"proposeSubmitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"proposeSubmitSignaturesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"_certificateRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerNodeId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_verificationData\",\"type\":\"bytes\"}],\"name\":\"registerPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"}],\"name\":\"unregisterNodeId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "EntityManager",
}

// EntityManager is an auto generated Go binding around an Ethereum contract.
type EntityManager struct {
	abi abi.ABI
}

// NewEntityManager creates a new instance of EntityManager.
func NewEntityManager() *EntityManager {
	parsed, err := EntityManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EntityManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EntityManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConfirmDelegationAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5360006.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmDelegationAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) PackConfirmDelegationAddressRegistration(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("confirmDelegationAddressRegistration", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmDelegationAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5360006.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmDelegationAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) TryPackConfirmDelegationAddressRegistration(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("confirmDelegationAddressRegistration", voter)
}

// PackConfirmSigningPolicyAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb8d5125.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmSigningPolicyAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) PackConfirmSigningPolicyAddressRegistration(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("confirmSigningPolicyAddressRegistration", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmSigningPolicyAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb8d5125.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmSigningPolicyAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) TryPackConfirmSigningPolicyAddressRegistration(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("confirmSigningPolicyAddressRegistration", voter)
}

// PackConfirmSubmitAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f1fb56a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmSubmitAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) PackConfirmSubmitAddressRegistration(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("confirmSubmitAddressRegistration", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmSubmitAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f1fb56a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmSubmitAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) TryPackConfirmSubmitAddressRegistration(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("confirmSubmitAddressRegistration", voter)
}

// PackConfirmSubmitSignaturesAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbda1a84e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmSubmitSignaturesAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) PackConfirmSubmitSignaturesAddressRegistration(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("confirmSubmitSignaturesAddressRegistration", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmSubmitSignaturesAddressRegistration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbda1a84e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmSubmitSignaturesAddressRegistration(address _voter) returns()
func (entityManager *EntityManager) TryPackConfirmSubmitSignaturesAddressRegistration(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("confirmSubmitSignaturesAddressRegistration", voter)
}

// PackGetDelegationAddressOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5d70de0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (entityManager *EntityManager) PackGetDelegationAddressOf(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("getDelegationAddressOf", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDelegationAddressOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5d70de0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (entityManager *EntityManager) TryPackGetDelegationAddressOf(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("getDelegationAddressOf", voter)
}

// UnpackGetDelegationAddressOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5d70de0.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (entityManager *EntityManager) UnpackGetDelegationAddressOf(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getDelegationAddressOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetDelegationAddressOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bedbf7a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (entityManager *EntityManager) PackGetDelegationAddressOfAt(voter common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getDelegationAddressOfAt", voter, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDelegationAddressOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bedbf7a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (entityManager *EntityManager) TryPackGetDelegationAddressOfAt(voter common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getDelegationAddressOfAt", voter, blockNumber)
}

// UnpackGetDelegationAddressOfAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2bedbf7a.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (entityManager *EntityManager) UnpackGetDelegationAddressOfAt(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getDelegationAddressOfAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeIdsOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6450b97.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (entityManager *EntityManager) PackGetNodeIdsOf(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("getNodeIdsOf", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeIdsOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6450b97.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (entityManager *EntityManager) TryPackGetNodeIdsOf(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("getNodeIdsOf", voter)
}

// UnpackGetNodeIdsOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe6450b97.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (entityManager *EntityManager) UnpackGetNodeIdsOf(data []byte) ([][20]byte, error) {
	out, err := entityManager.abi.Unpack("getNodeIdsOf", data)
	if err != nil {
		return *new([][20]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)
	return out0, nil
}

// PackGetNodeIdsOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b4be5f2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (entityManager *EntityManager) PackGetNodeIdsOfAt(voter common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getNodeIdsOfAt", voter, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeIdsOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b4be5f2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (entityManager *EntityManager) TryPackGetNodeIdsOfAt(voter common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getNodeIdsOfAt", voter, blockNumber)
}

// UnpackGetNodeIdsOfAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b4be5f2.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (entityManager *EntityManager) UnpackGetNodeIdsOfAt(data []byte) ([][20]byte, error) {
	out, err := entityManager.abi.Unpack("getNodeIdsOfAt", data)
	if err != nil {
		return *new([][20]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)
	return out0, nil
}

// PackGetPublicKeyOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75e68605.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (entityManager *EntityManager) PackGetPublicKeyOf(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("getPublicKeyOf", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPublicKeyOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75e68605.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (entityManager *EntityManager) TryPackGetPublicKeyOf(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("getPublicKeyOf", voter)
}

// GetPublicKeyOfOutput serves as a container for the return parameters of contract
// method GetPublicKeyOf.
type GetPublicKeyOfOutput struct {
	Arg0 [32]byte
	Arg1 [32]byte
}

// UnpackGetPublicKeyOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x75e68605.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (entityManager *EntityManager) UnpackGetPublicKeyOf(data []byte) (GetPublicKeyOfOutput, error) {
	out, err := entityManager.abi.Unpack("getPublicKeyOf", data)
	outstruct := new(GetPublicKeyOfOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackGetPublicKeyOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11104561.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (entityManager *EntityManager) PackGetPublicKeyOfAt(voter common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getPublicKeyOfAt", voter, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPublicKeyOfAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11104561.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (entityManager *EntityManager) TryPackGetPublicKeyOfAt(voter common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getPublicKeyOfAt", voter, blockNumber)
}

// GetPublicKeyOfAtOutput serves as a container for the return parameters of contract
// method GetPublicKeyOfAt.
type GetPublicKeyOfAtOutput struct {
	Arg0 [32]byte
	Arg1 [32]byte
}

// UnpackGetPublicKeyOfAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11104561.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (entityManager *EntityManager) UnpackGetPublicKeyOfAt(data []byte) (GetPublicKeyOfAtOutput, error) {
	out, err := entityManager.abi.Unpack("getPublicKeyOfAt", data)
	outstruct := new(GetPublicKeyOfAtOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackGetVoterAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5771dbc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) PackGetVoterAddresses(voter common.Address) []byte {
	enc, err := entityManager.abi.Pack("getVoterAddresses", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5771dbc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) TryPackGetVoterAddresses(voter common.Address) ([]byte, error) {
	return entityManager.abi.Pack("getVoterAddresses", voter)
}

// UnpackGetVoterAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe5771dbc.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) UnpackGetVoterAddresses(data []byte) (IEntityManagerVoterAddresses, error) {
	out, err := entityManager.abi.Unpack("getVoterAddresses", data)
	if err != nil {
		return *new(IEntityManagerVoterAddresses), err
	}
	out0 := *abi.ConvertType(out[0], new(IEntityManagerVoterAddresses)).(*IEntityManagerVoterAddresses)
	return out0, nil
}

// PackGetVoterAddressesAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2519ac8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) PackGetVoterAddressesAt(voter common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterAddressesAt", voter, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterAddressesAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2519ac8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) TryPackGetVoterAddressesAt(voter common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterAddressesAt", voter, blockNumber)
}

// UnpackGetVoterAddressesAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb2519ac8.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (entityManager *EntityManager) UnpackGetVoterAddressesAt(data []byte) (IEntityManagerVoterAddresses, error) {
	out, err := entityManager.abi.Unpack("getVoterAddressesAt", data)
	if err != nil {
		return *new(IEntityManagerVoterAddresses), err
	}
	out0 := *abi.ConvertType(out[0], new(IEntityManagerVoterAddresses)).(*IEntityManagerVoterAddresses)
	return out0, nil
}

// PackGetVoterForDelegationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc8c4a5f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForDelegationAddress(delegationAddress common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForDelegationAddress", delegationAddress, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForDelegationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc8c4a5f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForDelegationAddress(delegationAddress common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForDelegationAddress", delegationAddress, blockNumber)
}

// UnpackGetVoterForDelegationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc8c4a5f.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForDelegationAddress(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForDelegationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVoterForNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a812fff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForNodeId(nodeId [20]byte, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForNodeId", nodeId, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a812fff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForNodeId(nodeId [20]byte, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForNodeId", nodeId, blockNumber)
}

// UnpackGetVoterForNodeId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0a812fff.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForNodeId(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForNodeId", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVoterForPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09c80fa6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForPublicKey(part1 [32]byte, part2 [32]byte, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForPublicKey", part1, part2, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09c80fa6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForPublicKey(part1 [32]byte, part2 [32]byte, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForPublicKey", part1, part2, blockNumber)
}

// UnpackGetVoterForPublicKey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09c80fa6.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForPublicKey(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForPublicKey", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVoterForSigningPolicyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe178ec4c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForSigningPolicyAddress(signingPolicyAddress common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForSigningPolicyAddress", signingPolicyAddress, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForSigningPolicyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe178ec4c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForSigningPolicyAddress(signingPolicyAddress common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForSigningPolicyAddress", signingPolicyAddress, blockNumber)
}

// UnpackGetVoterForSigningPolicyAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe178ec4c.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForSigningPolicyAddress(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForSigningPolicyAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVoterForSubmitAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89292b0f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForSubmitAddress(submitAddress common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForSubmitAddress", submitAddress, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForSubmitAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89292b0f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForSubmitAddress(submitAddress common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForSubmitAddress", submitAddress, blockNumber)
}

// UnpackGetVoterForSubmitAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x89292b0f.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForSubmitAddress(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForSubmitAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVoterForSubmitSignaturesAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61bc320f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) PackGetVoterForSubmitSignaturesAddress(submitSignaturesAddress common.Address, blockNumber *big.Int) []byte {
	enc, err := entityManager.abi.Pack("getVoterForSubmitSignaturesAddress", submitSignaturesAddress, blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterForSubmitSignaturesAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61bc320f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) TryPackGetVoterForSubmitSignaturesAddress(submitSignaturesAddress common.Address, blockNumber *big.Int) ([]byte, error) {
	return entityManager.abi.Pack("getVoterForSubmitSignaturesAddress", submitSignaturesAddress, blockNumber)
}

// UnpackGetVoterForSubmitSignaturesAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x61bc320f.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (entityManager *EntityManager) UnpackGetVoterForSubmitSignaturesAddress(data []byte) (common.Address, error) {
	out, err := entityManager.abi.Unpack("getVoterForSubmitSignaturesAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProposeDelegationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5712b9f3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeDelegationAddress(address _delegationAddress) returns()
func (entityManager *EntityManager) PackProposeDelegationAddress(delegationAddress common.Address) []byte {
	enc, err := entityManager.abi.Pack("proposeDelegationAddress", delegationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeDelegationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5712b9f3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeDelegationAddress(address _delegationAddress) returns()
func (entityManager *EntityManager) TryPackProposeDelegationAddress(delegationAddress common.Address) ([]byte, error) {
	return entityManager.abi.Pack("proposeDelegationAddress", delegationAddress)
}

// PackProposeSigningPolicyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69b5fed9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeSigningPolicyAddress(address _signingPolicyAddress) returns()
func (entityManager *EntityManager) PackProposeSigningPolicyAddress(signingPolicyAddress common.Address) []byte {
	enc, err := entityManager.abi.Pack("proposeSigningPolicyAddress", signingPolicyAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeSigningPolicyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69b5fed9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeSigningPolicyAddress(address _signingPolicyAddress) returns()
func (entityManager *EntityManager) TryPackProposeSigningPolicyAddress(signingPolicyAddress common.Address) ([]byte, error) {
	return entityManager.abi.Pack("proposeSigningPolicyAddress", signingPolicyAddress)
}

// PackProposeSubmitAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87607e41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeSubmitAddress(address _submitAddress) returns()
func (entityManager *EntityManager) PackProposeSubmitAddress(submitAddress common.Address) []byte {
	enc, err := entityManager.abi.Pack("proposeSubmitAddress", submitAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeSubmitAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87607e41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeSubmitAddress(address _submitAddress) returns()
func (entityManager *EntityManager) TryPackProposeSubmitAddress(submitAddress common.Address) ([]byte, error) {
	return entityManager.abi.Pack("proposeSubmitAddress", submitAddress)
}

// PackProposeSubmitSignaturesAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36edd357.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeSubmitSignaturesAddress(address _submitSignaturesAddress) returns()
func (entityManager *EntityManager) PackProposeSubmitSignaturesAddress(submitSignaturesAddress common.Address) []byte {
	enc, err := entityManager.abi.Pack("proposeSubmitSignaturesAddress", submitSignaturesAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeSubmitSignaturesAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36edd357.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeSubmitSignaturesAddress(address _submitSignaturesAddress) returns()
func (entityManager *EntityManager) TryPackProposeSubmitSignaturesAddress(submitSignaturesAddress common.Address) ([]byte, error) {
	return entityManager.abi.Pack("proposeSubmitSignaturesAddress", submitSignaturesAddress)
}

// PackRegisterNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x019476a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerNodeId(bytes20 _nodeId, bytes _certificateRaw, bytes _signature) returns()
func (entityManager *EntityManager) PackRegisterNodeId(nodeId [20]byte, certificateRaw []byte, signature []byte) []byte {
	enc, err := entityManager.abi.Pack("registerNodeId", nodeId, certificateRaw, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x019476a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerNodeId(bytes20 _nodeId, bytes _certificateRaw, bytes _signature) returns()
func (entityManager *EntityManager) TryPackRegisterNodeId(nodeId [20]byte, certificateRaw []byte, signature []byte) ([]byte, error) {
	return entityManager.abi.Pack("registerNodeId", nodeId, certificateRaw, signature)
}

// PackRegisterPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1863cff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerPublicKey(bytes32 _part1, bytes32 _part2, bytes _verificationData) returns()
func (entityManager *EntityManager) PackRegisterPublicKey(part1 [32]byte, part2 [32]byte, verificationData []byte) []byte {
	enc, err := entityManager.abi.Pack("registerPublicKey", part1, part2, verificationData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1863cff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerPublicKey(bytes32 _part1, bytes32 _part2, bytes _verificationData) returns()
func (entityManager *EntityManager) TryPackRegisterPublicKey(part1 [32]byte, part2 [32]byte, verificationData []byte) ([]byte, error) {
	return entityManager.abi.Pack("registerPublicKey", part1, part2, verificationData)
}

// PackUnregisterNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18c03f94.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unregisterNodeId(bytes20 _nodeId) returns()
func (entityManager *EntityManager) PackUnregisterNodeId(nodeId [20]byte) []byte {
	enc, err := entityManager.abi.Pack("unregisterNodeId", nodeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnregisterNodeId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18c03f94.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unregisterNodeId(bytes20 _nodeId) returns()
func (entityManager *EntityManager) TryPackUnregisterNodeId(nodeId [20]byte) ([]byte, error) {
	return entityManager.abi.Pack("unregisterNodeId", nodeId)
}

// PackUnregisterPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1e3cdac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unregisterPublicKey() returns()
func (entityManager *EntityManager) PackUnregisterPublicKey() []byte {
	enc, err := entityManager.abi.Pack("unregisterPublicKey")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnregisterPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1e3cdac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unregisterPublicKey() returns()
func (entityManager *EntityManager) TryPackUnregisterPublicKey() ([]byte, error) {
	return entityManager.abi.Pack("unregisterPublicKey")
}

// EntityManagerDelegationAddressProposed represents a DelegationAddressProposed event raised by the EntityManager contract.
type EntityManagerDelegationAddressProposed struct {
	Voter             common.Address
	DelegationAddress common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const EntityManagerDelegationAddressProposedEventName = "DelegationAddressProposed"

// ContractEventName returns the user-defined event name.
func (EntityManagerDelegationAddressProposed) ContractEventName() string {
	return EntityManagerDelegationAddressProposedEventName
}

// UnpackDelegationAddressProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DelegationAddressProposed(address indexed voter, address indexed delegationAddress)
func (entityManager *EntityManager) UnpackDelegationAddressProposedEvent(log *types.Log) (*EntityManagerDelegationAddressProposed, error) {
	event := "DelegationAddressProposed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerDelegationAddressProposed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerDelegationAddressRegistrationConfirmed represents a DelegationAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerDelegationAddressRegistrationConfirmed struct {
	Voter             common.Address
	DelegationAddress common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const EntityManagerDelegationAddressRegistrationConfirmedEventName = "DelegationAddressRegistrationConfirmed"

// ContractEventName returns the user-defined event name.
func (EntityManagerDelegationAddressRegistrationConfirmed) ContractEventName() string {
	return EntityManagerDelegationAddressRegistrationConfirmedEventName
}

// UnpackDelegationAddressRegistrationConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DelegationAddressRegistrationConfirmed(address indexed voter, address indexed delegationAddress)
func (entityManager *EntityManager) UnpackDelegationAddressRegistrationConfirmedEvent(log *types.Log) (*EntityManagerDelegationAddressRegistrationConfirmed, error) {
	event := "DelegationAddressRegistrationConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerDelegationAddressRegistrationConfirmed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerMaxNodeIdsPerEntitySet represents a MaxNodeIdsPerEntitySet event raised by the EntityManager contract.
type EntityManagerMaxNodeIdsPerEntitySet struct {
	MaxNodeIdsPerEntity *big.Int
	Raw                 *types.Log // Blockchain specific contextual infos
}

const EntityManagerMaxNodeIdsPerEntitySetEventName = "MaxNodeIdsPerEntitySet"

// ContractEventName returns the user-defined event name.
func (EntityManagerMaxNodeIdsPerEntitySet) ContractEventName() string {
	return EntityManagerMaxNodeIdsPerEntitySetEventName
}

// UnpackMaxNodeIdsPerEntitySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MaxNodeIdsPerEntitySet(uint256 maxNodeIdsPerEntity)
func (entityManager *EntityManager) UnpackMaxNodeIdsPerEntitySetEvent(log *types.Log) (*EntityManagerMaxNodeIdsPerEntitySet, error) {
	event := "MaxNodeIdsPerEntitySet"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerMaxNodeIdsPerEntitySet)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerNodeIdRegistered represents a NodeIdRegistered event raised by the EntityManager contract.
type EntityManagerNodeIdRegistered struct {
	Voter  common.Address
	NodeId [20]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const EntityManagerNodeIdRegisteredEventName = "NodeIdRegistered"

// ContractEventName returns the user-defined event name.
func (EntityManagerNodeIdRegistered) ContractEventName() string {
	return EntityManagerNodeIdRegisteredEventName
}

// UnpackNodeIdRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeIdRegistered(address indexed voter, bytes20 indexed nodeId)
func (entityManager *EntityManager) UnpackNodeIdRegisteredEvent(log *types.Log) (*EntityManagerNodeIdRegistered, error) {
	event := "NodeIdRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerNodeIdRegistered)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerNodeIdUnregistered represents a NodeIdUnregistered event raised by the EntityManager contract.
type EntityManagerNodeIdUnregistered struct {
	Voter  common.Address
	NodeId [20]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const EntityManagerNodeIdUnregisteredEventName = "NodeIdUnregistered"

// ContractEventName returns the user-defined event name.
func (EntityManagerNodeIdUnregistered) ContractEventName() string {
	return EntityManagerNodeIdUnregisteredEventName
}

// UnpackNodeIdUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeIdUnregistered(address indexed voter, bytes20 indexed nodeId)
func (entityManager *EntityManager) UnpackNodeIdUnregisteredEvent(log *types.Log) (*EntityManagerNodeIdUnregistered, error) {
	event := "NodeIdUnregistered"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerNodeIdUnregistered)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerPublicKeyRegistered represents a PublicKeyRegistered event raised by the EntityManager contract.
type EntityManagerPublicKeyRegistered struct {
	Voter common.Address
	Part1 [32]byte
	Part2 [32]byte
	Raw   *types.Log // Blockchain specific contextual infos
}

const EntityManagerPublicKeyRegisteredEventName = "PublicKeyRegistered"

// ContractEventName returns the user-defined event name.
func (EntityManagerPublicKeyRegistered) ContractEventName() string {
	return EntityManagerPublicKeyRegisteredEventName
}

// UnpackPublicKeyRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PublicKeyRegistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (entityManager *EntityManager) UnpackPublicKeyRegisteredEvent(log *types.Log) (*EntityManagerPublicKeyRegistered, error) {
	event := "PublicKeyRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerPublicKeyRegistered)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerPublicKeyUnregistered represents a PublicKeyUnregistered event raised by the EntityManager contract.
type EntityManagerPublicKeyUnregistered struct {
	Voter common.Address
	Part1 [32]byte
	Part2 [32]byte
	Raw   *types.Log // Blockchain specific contextual infos
}

const EntityManagerPublicKeyUnregisteredEventName = "PublicKeyUnregistered"

// ContractEventName returns the user-defined event name.
func (EntityManagerPublicKeyUnregistered) ContractEventName() string {
	return EntityManagerPublicKeyUnregisteredEventName
}

// UnpackPublicKeyUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PublicKeyUnregistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (entityManager *EntityManager) UnpackPublicKeyUnregisteredEvent(log *types.Log) (*EntityManagerPublicKeyUnregistered, error) {
	event := "PublicKeyUnregistered"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerPublicKeyUnregistered)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSigningPolicyAddressProposed represents a SigningPolicyAddressProposed event raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressProposed struct {
	Voter                common.Address
	SigningPolicyAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const EntityManagerSigningPolicyAddressProposedEventName = "SigningPolicyAddressProposed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSigningPolicyAddressProposed) ContractEventName() string {
	return EntityManagerSigningPolicyAddressProposedEventName
}

// UnpackSigningPolicyAddressProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningPolicyAddressProposed(address indexed voter, address indexed signingPolicyAddress)
func (entityManager *EntityManager) UnpackSigningPolicyAddressProposedEvent(log *types.Log) (*EntityManagerSigningPolicyAddressProposed, error) {
	event := "SigningPolicyAddressProposed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSigningPolicyAddressProposed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSigningPolicyAddressRegistrationConfirmed represents a SigningPolicyAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressRegistrationConfirmed struct {
	Voter                common.Address
	SigningPolicyAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const EntityManagerSigningPolicyAddressRegistrationConfirmedEventName = "SigningPolicyAddressRegistrationConfirmed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSigningPolicyAddressRegistrationConfirmed) ContractEventName() string {
	return EntityManagerSigningPolicyAddressRegistrationConfirmedEventName
}

// UnpackSigningPolicyAddressRegistrationConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningPolicyAddressRegistrationConfirmed(address indexed voter, address indexed signingPolicyAddress)
func (entityManager *EntityManager) UnpackSigningPolicyAddressRegistrationConfirmedEvent(log *types.Log) (*EntityManagerSigningPolicyAddressRegistrationConfirmed, error) {
	event := "SigningPolicyAddressRegistrationConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSigningPolicyAddressRegistrationConfirmed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSubmitAddressProposed represents a SubmitAddressProposed event raised by the EntityManager contract.
type EntityManagerSubmitAddressProposed struct {
	Voter         common.Address
	SubmitAddress common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const EntityManagerSubmitAddressProposedEventName = "SubmitAddressProposed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSubmitAddressProposed) ContractEventName() string {
	return EntityManagerSubmitAddressProposedEventName
}

// UnpackSubmitAddressProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubmitAddressProposed(address indexed voter, address indexed submitAddress)
func (entityManager *EntityManager) UnpackSubmitAddressProposedEvent(log *types.Log) (*EntityManagerSubmitAddressProposed, error) {
	event := "SubmitAddressProposed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSubmitAddressProposed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSubmitAddressRegistrationConfirmed represents a SubmitAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSubmitAddressRegistrationConfirmed struct {
	Voter         common.Address
	SubmitAddress common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const EntityManagerSubmitAddressRegistrationConfirmedEventName = "SubmitAddressRegistrationConfirmed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSubmitAddressRegistrationConfirmed) ContractEventName() string {
	return EntityManagerSubmitAddressRegistrationConfirmedEventName
}

// UnpackSubmitAddressRegistrationConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubmitAddressRegistrationConfirmed(address indexed voter, address indexed submitAddress)
func (entityManager *EntityManager) UnpackSubmitAddressRegistrationConfirmedEvent(log *types.Log) (*EntityManagerSubmitAddressRegistrationConfirmed, error) {
	event := "SubmitAddressRegistrationConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSubmitAddressRegistrationConfirmed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSubmitSignaturesAddressProposed represents a SubmitSignaturesAddressProposed event raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressProposed struct {
	Voter                   common.Address
	SubmitSignaturesAddress common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const EntityManagerSubmitSignaturesAddressProposedEventName = "SubmitSignaturesAddressProposed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSubmitSignaturesAddressProposed) ContractEventName() string {
	return EntityManagerSubmitSignaturesAddressProposedEventName
}

// UnpackSubmitSignaturesAddressProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubmitSignaturesAddressProposed(address indexed voter, address indexed submitSignaturesAddress)
func (entityManager *EntityManager) UnpackSubmitSignaturesAddressProposedEvent(log *types.Log) (*EntityManagerSubmitSignaturesAddressProposed, error) {
	event := "SubmitSignaturesAddressProposed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSubmitSignaturesAddressProposed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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

// EntityManagerSubmitSignaturesAddressRegistrationConfirmed represents a SubmitSignaturesAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressRegistrationConfirmed struct {
	Voter                   common.Address
	SubmitSignaturesAddress common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const EntityManagerSubmitSignaturesAddressRegistrationConfirmedEventName = "SubmitSignaturesAddressRegistrationConfirmed"

// ContractEventName returns the user-defined event name.
func (EntityManagerSubmitSignaturesAddressRegistrationConfirmed) ContractEventName() string {
	return EntityManagerSubmitSignaturesAddressRegistrationConfirmedEventName
}

// UnpackSubmitSignaturesAddressRegistrationConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SubmitSignaturesAddressRegistrationConfirmed(address indexed voter, address indexed submitSignaturesAddress)
func (entityManager *EntityManager) UnpackSubmitSignaturesAddressRegistrationConfirmedEvent(log *types.Log) (*EntityManagerSubmitSignaturesAddressRegistrationConfirmed, error) {
	event := "SubmitSignaturesAddressRegistrationConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != entityManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntityManagerSubmitSignaturesAddressRegistrationConfirmed)
	if len(log.Data) > 0 {
		if err := entityManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entityManager.abi.Events[event].Inputs {
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
