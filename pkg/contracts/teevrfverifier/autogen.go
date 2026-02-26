// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teevrfverifier

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

// VRFVerifierPoint is an auto generated low-level Go binding around an user-defined struct.
type VRFVerifierPoint struct {
	X *big.Int
	Y *big.Int
}

// VRFVerifierProof is an auto generated low-level Go binding around an user-defined struct.
type VRFVerifierProof struct {
	Gamma  VRFVerifierPoint
	C      *big.Int
	S      *big.Int
	U      VRFVerifierPoint
	CGamma VRFVerifierPoint
	V      VRFVerifierPoint
	ZInv   *big.Int
}

// TeeVRFVerifierMetaData contains all meta data concerning the TeeVRFVerifier contract.
var TeeVRFVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gammaX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gammaY\",\"type\":\"uint256\"}],\"name\":\"randomnessFromProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structVRFVerifier.Point\",\"name\":\"gamma\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structVRFVerifier.Point\",\"name\":\"u\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structVRFVerifier.Point\",\"name\":\"cGamma\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structVRFVerifier.Point\",\"name\":\"v\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRFVerifier.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"pkX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pkY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"}],\"name\":\"verifyRandomness\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610b8e908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806313a0c07c1461003457631fc57fbd1461002f575f80fd5b610074565b346100705760403660031901126100705760043560a05260243560c05260406080526100606060610111565b602060805160a020604051908152f35b5f80fd5b346100705736600319016101c08112610070576101601361007057610184356101a4356101643567ffffffffffffffff821161007057366023830112156100705781600401359067ffffffffffffffff8211610070573660248385010111610070576100f99360246100e79401916104af565b60405190151581529081906020820190565b0390f35b634e487b7160e01b5f52604160045260245ffd5b601f80199101166080016080811067ffffffffffffffff82111761013457604052565b6100fd565b6040810190811067ffffffffffffffff82111761013457604052565b90601f8019910116810190811067ffffffffffffffff82111761013457604052565b60405190610186604083610155565b565b1561018f57565b60405162461bcd60e51b81526020600482015260146024820152735652463a20706b206e6f74206f6e20637572766560601b6044820152606490fd5b9190826040910312610070576040516101e381610139565b6020808294803584520135910152565b156101fa57565b60405162461bcd60e51b815260206004820152601760248201527f5652463a2067616d6d61206e6f74206f6e2063757276650000000000000000006044820152606490fd5b1561024657565b60405162461bcd60e51b81526020600482015260136024820152725652463a2063206f7574206f662072616e676560681b6044820152606490fd5b1561028857565b60405162461bcd60e51b81526020600482015260136024820152725652463a2073206f7574206f662072616e676560681b6044820152606490fd5b92919267ffffffffffffffff821161013457604051916102ed601f8201601f191660200184610155565b829481845281830111610070578281602093845f960137010152565b1561031057565b606460405162461bcd60e51b815260206004820152602060248201527f5652463a20682e78203e3d204e2028646567656e657261746520696e707574296044820152fd5b1561035b57565b60405162461bcd60e51b81526020600482015260166024820152755652463a20696e76616c69642075207769746e65737360501b6044820152606490fd5b156103a057565b60405162461bcd60e51b815260206004820152601b60248201527f5652463a20696e76616c6964206347616d6d61207769746e65737300000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b6401000003d01903906401000003d019821161041157565b6103e5565b634e487b7160e01b5f52601260045260245ffd5b1561043157565b60405162461bcd60e51b81526020600482015260116024820152702b29231d1034b73b30b634b2103d24b73b60791b6044820152606490fd5b1561047157565b60405162461bcd60e51b81526020600482015260166024820152755652463a20696e76616c69642076207769746e65737360501b6044820152606490fd5b9291926104d36104ce6104c0610177565b838152846020820152610810565b610188565b6104ee6104e96104e43660046101cb565b610810565b6101f3565b61053761053260443595861515806107f5575b61050a9061023f565b6064359561052b70014551231950b75fc4402da1732fc9bebe198810610281565b36916102c3565b610851565b61055570014551231950b75fc4402da1732fc9bebe19825110610309565b60a435916084356105826001600160a01b036105718684610940565b1680151590816107dd575b50610354565b60c4359060e435946105b66105a661059a8886610940565b6001600160a01b031690565b80151590816107c2575b50610399565b61010435916105c4836103f9565b5f97906401000003d019908608801515908161079e575b50906105e96105ee9261042a565b6103f9565b9461012435946105fd866103f9565b98610799576107769a61068d61059a61069e936107579c6107499b61014435906401000003d019926401000003d019910809906401000003d0196106408b6103f9565b6401000003d0198480090861065b6401000003d019926103f9565b9008906401000003d0199081610670846103f9565b8c0890096401000003d019906106858c6103f9565b900890610940565b801515918261077a575b505061046a565b602087519701519860043590602435926040519b8c9a60208c019693909a9998959261016098959261018089019c7f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f817988a527f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b860208b015260408a01526060890152608088015260a087015260c086015260e08501526101008401526101208301526101408201520152565b03601f198101835282610155565b70014551231950b75fc4402da1732fc9bebe1990602081519101200690565b1490565b61079191925061059a908b5160208d01519061097a565b145f80610697565b610416565b9850506105ee906105e96001610144355f9b6401000003d0199109149192506105db565b90506107d661059a8c60243560043561097a565b145f6105b0565b90506107ee61059a8a89878c610a43565b145f61057c565b5070014551231950b75fc4402da1732fc9bebe198710610501565b602081015190516401000003d019906007908290818180090908906401000003d0199080091490565b6040519061084682610139565b5f6020838281520152565b610859610839565b5080516020909101205f916401000003d019820691835b61010085106108d05760405162461bcd60e51b815260206004820152602960248201527f5652463a2048617368546f4375727665206578636565646564206974657261746044820152681a5bdb881b1a5b5a5d60ba1b6064820152608490fd5b610799575f926108ef6401000003d01960078184818180090908610af4565b8061092757505060405161090f8161074960208201948560209181520190565b5190206001909301926401000003d019810692610870565b929450925050610935610177565b918252602082015290565b604080516020810192835280820193909352825261095f606083610155565b905190206001600160a01b031690565b6040513d5f823e3d90fd5b90919070014551231950b75fc4402da1732fc9bebe195f820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161041157600116601b019182601b1161041157610a2b5f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b838052039060015afa15610a3e575f5190565b61096f565b91929170014551231950b75fc4402da1732fc9bebe1990820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161041157600116601b019182601b1161041157610a2b5f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b908115610b5357604051602081526020808201526020604082015282606082015263400000f4600160fe1b0360808201526401000003d01960a082015260208160c08160055afa156100705751916401000003d01983800903610b5357565b5f915056fea2646970667358221220e289f4a4776df552b6573650fd1bf26fce1bc5cc2de465819caa0af7e442fec564736f6c634300081b0033",
}

// TeeVRFVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVRFVerifierMetaData.ABI instead.
var TeeVRFVerifierABI = TeeVRFVerifierMetaData.ABI

// TeeVRFVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeeVRFVerifierMetaData.Bin instead.
var TeeVRFVerifierBin = TeeVRFVerifierMetaData.Bin

// DeployTeeVRFVerifier deploys a new Ethereum contract, binding an instance of TeeVRFVerifier to it.
func DeployTeeVRFVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeeVRFVerifier, error) {
	parsed, err := TeeVRFVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeeVRFVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeeVRFVerifier{TeeVRFVerifierCaller: TeeVRFVerifierCaller{contract: contract}, TeeVRFVerifierTransactor: TeeVRFVerifierTransactor{contract: contract}, TeeVRFVerifierFilterer: TeeVRFVerifierFilterer{contract: contract}}, nil
}

// TeeVRFVerifier is an auto generated Go binding around an Ethereum contract.
type TeeVRFVerifier struct {
	TeeVRFVerifierCaller     // Read-only binding to the contract
	TeeVRFVerifierTransactor // Write-only binding to the contract
	TeeVRFVerifierFilterer   // Log filterer for contract events
}

// TeeVRFVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVRFVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVRFVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVRFVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVRFVerifierSession struct {
	Contract     *TeeVRFVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVRFVerifierCallerSession struct {
	Contract *TeeVRFVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TeeVRFVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVRFVerifierTransactorSession struct {
	Contract     *TeeVRFVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TeeVRFVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVRFVerifierRaw struct {
	Contract *TeeVRFVerifier // Generic contract binding to access the raw methods on
}

// TeeVRFVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVRFVerifierCallerRaw struct {
	Contract *TeeVRFVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVRFVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVRFVerifierTransactorRaw struct {
	Contract *TeeVRFVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVRFVerifier creates a new instance of TeeVRFVerifier, bound to a specific deployed contract.
func NewTeeVRFVerifier(address common.Address, backend bind.ContractBackend) (*TeeVRFVerifier, error) {
	contract, err := bindTeeVRFVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVerifier{TeeVRFVerifierCaller: TeeVRFVerifierCaller{contract: contract}, TeeVRFVerifierTransactor: TeeVRFVerifierTransactor{contract: contract}, TeeVRFVerifierFilterer: TeeVRFVerifierFilterer{contract: contract}}, nil
}

// NewTeeVRFVerifierCaller creates a new read-only instance of TeeVRFVerifier, bound to a specific deployed contract.
func NewTeeVRFVerifierCaller(address common.Address, caller bind.ContractCaller) (*TeeVRFVerifierCaller, error) {
	contract, err := bindTeeVRFVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVerifierCaller{contract: contract}, nil
}

// NewTeeVRFVerifierTransactor creates a new write-only instance of TeeVRFVerifier, bound to a specific deployed contract.
func NewTeeVRFVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVRFVerifierTransactor, error) {
	contract, err := bindTeeVRFVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVerifierTransactor{contract: contract}, nil
}

// NewTeeVRFVerifierFilterer creates a new log filterer instance of TeeVRFVerifier, bound to a specific deployed contract.
func NewTeeVRFVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVRFVerifierFilterer, error) {
	contract, err := bindTeeVRFVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVerifierFilterer{contract: contract}, nil
}

// bindTeeVRFVerifier binds a generic wrapper to an already deployed contract.
func bindTeeVRFVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVRFVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRFVerifier *TeeVRFVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRFVerifier.Contract.TeeVRFVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRFVerifier *TeeVRFVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRFVerifier.Contract.TeeVRFVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRFVerifier *TeeVRFVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRFVerifier.Contract.TeeVRFVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRFVerifier *TeeVRFVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRFVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRFVerifier *TeeVRFVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRFVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRFVerifier *TeeVRFVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRFVerifier.Contract.contract.Transact(opts, method, params...)
}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 gammaX, uint256 gammaY) pure returns(bytes32)
func (_TeeVRFVerifier *TeeVRFVerifierCaller) RandomnessFromProof(opts *bind.CallOpts, gammaX *big.Int, gammaY *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeeVRFVerifier.contract.Call(opts, &out, "randomnessFromProof", gammaX, gammaY)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 gammaX, uint256 gammaY) pure returns(bytes32)
func (_TeeVRFVerifier *TeeVRFVerifierSession) RandomnessFromProof(gammaX *big.Int, gammaY *big.Int) ([32]byte, error) {
	return _TeeVRFVerifier.Contract.RandomnessFromProof(&_TeeVRFVerifier.CallOpts, gammaX, gammaY)
}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 gammaX, uint256 gammaY) pure returns(bytes32)
func (_TeeVRFVerifier *TeeVRFVerifierCallerSession) RandomnessFromProof(gammaX *big.Int, gammaY *big.Int) ([32]byte, error) {
	return _TeeVRFVerifier.Contract.RandomnessFromProof(&_TeeVRFVerifier.CallOpts, gammaX, gammaY)
}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) proof, uint256 pkX, uint256 pkY, bytes nonce) view returns(bool valid)
func (_TeeVRFVerifier *TeeVRFVerifierCaller) VerifyRandomness(opts *bind.CallOpts, proof VRFVerifierProof, pkX *big.Int, pkY *big.Int, nonce []byte) (bool, error) {
	var out []interface{}
	err := _TeeVRFVerifier.contract.Call(opts, &out, "verifyRandomness", proof, pkX, pkY, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) proof, uint256 pkX, uint256 pkY, bytes nonce) view returns(bool valid)
func (_TeeVRFVerifier *TeeVRFVerifierSession) VerifyRandomness(proof VRFVerifierProof, pkX *big.Int, pkY *big.Int, nonce []byte) (bool, error) {
	return _TeeVRFVerifier.Contract.VerifyRandomness(&_TeeVRFVerifier.CallOpts, proof, pkX, pkY, nonce)
}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) proof, uint256 pkX, uint256 pkY, bytes nonce) view returns(bool valid)
func (_TeeVRFVerifier *TeeVRFVerifierCallerSession) VerifyRandomness(proof VRFVerifierProof, pkX *big.Int, pkY *big.Int, nonce []byte) (bool, error) {
	return _TeeVRFVerifier.Contract.VerifyRandomness(&_TeeVRFVerifier.CallOpts, proof, pkX, pkY, nonce)
}
