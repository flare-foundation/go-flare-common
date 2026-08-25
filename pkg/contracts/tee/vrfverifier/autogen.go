// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfverifier

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

// IVrfVerifierPoint is an auto generated low-level Go binding around an user-defined struct.
type IVrfVerifierPoint struct {
	X *big.Int
	Y *big.Int
}

// IVrfVerifierProof is an auto generated low-level Go binding around an user-defined struct.
type IVrfVerifierProof struct {
	Gamma  IVrfVerifierPoint
	C      *big.Int
	S      *big.Int
	U      IVrfVerifierPoint
	CGamma IVrfVerifierPoint
	V      IVrfVerifierPoint
	ZInv   *big.Int
}

// VRFVerifierMetaData contains all meta data concerning the VRFVerifier contract.
var VRFVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"COutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GammaNotOnCurve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GammaXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashToCurveExceededIterationLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCGammaWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZInv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PkNotOnCurve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PkXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gammaX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gammaY\",\"type\":\"uint256\"}],\"name\":\"randomnessFromProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"gamma\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"u\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"cGamma\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"v\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_pkX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pkY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"verifyRandomness\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610ab8908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806313a0c07c1461003457631fc57fbd1461002f575f80fd5b610074565b346100705760403660031901126100705760043560a05260243560c05260406080526100606060610111565b602060805160a020604051908152f35b5f80fd5b346100705736600319016101c08112610070576101601361007057610184356101a4356101643567ffffffffffffffff821161007057366023830112156100705781600401359067ffffffffffffffff8211610070573660248385010111610070576100f99360246100e79401916102bc565b60405190151581529081906020820190565b0390f35b634e487b7160e01b5f52604160045260245ffd5b601f80199101166080016080811067ffffffffffffffff82111761013457604052565b6100fd565b6040810190811067ffffffffffffffff82111761013457604052565b90601f8019910116810190811067ffffffffffffffff82111761013457604052565b60405190610186604083610155565b565b1561018f57565b630b097aa560e41b5f5260045ffd5b9190826040910312610070576040516101b681610139565b6020808294803584520135910152565b156101cd57565b63094d477160e41b5f5260045ffd5b156101e357565b63040107d160e11b5f5260045ffd5b156101f957565b63ec6e729760e01b5f5260045ffd5b1561020f57565b63263091cb60e11b5f5260045ffd5b1561022557565b6304605f3560e21b5f5260045ffd5b92919267ffffffffffffffff8211610134576040519161025e601f8201601f191660200184610155565b829481845281830111610070578281602093845f960137010152565b1561028157565b63ff8c457560e01b5f5260045ffd5b1561029757565b63de86b44560e01b5f5260045ffd5b156102ad57565b63125ecc0760e31b5f5260045ffd5b61040861041a946102e46102df6102d1610177565b858152866020820152610483565b610188565b6102ff6102fa6102f536600461019e565b610483565b6101c6565b60443580151580610454575b610314906101dc565b6103d861038a6103836064359861033e70014551231950b75fc4402da1732fc9bebe198b106101f2565b61035b70014551231950b75fc4402da1732fc9bebe198910610208565b6004359561037c70014551231950b75fc4402da1732fc9bebe19881061021e565b3691610234565b87876104ec565b966103a970014551231950b75fc4402da1732fc9bebe1989511061027a565b856001600160a01b036103c060a4356084356105bb565b168489821515938461043a575b505050509050610290565b6103f56103e960e43560c4356105bb565b6001600160a01b031690565b90811515928361041d575b5050506102a6565b6104138360046107cc565b6004610904565b90565b610431929350906103e99160243590610626565b145f8080610400565b61044994506103e993956106ef565b14805f8489896103cd565b5070014551231950b75fc4402da1732fc9bebe19811061030b565b634e487b7160e01b5f52601260045260245ffd5b602081015190516401000003d019906007908290818180090908906401000003d0199080091490565b604051906104b982610139565b5f6020838281520152565b805191908290602001825e015f815290565b61041a93926040928252602082015201906104c4565b919061051d906104fa6104ac565b5061050f6040519384926020840196876104d6565b03601f198101835282610155565b5190205f916401000003d019820691835b61010085106105465763424c4e9d60e01b5f5260045ffd5b6105b6575f926105656401000003d01960078184818180090908610a1e565b8061059d5750506040516105858161050f60208201948560209181520190565b5190206001909301926401000003d01981069261052e565b9294509250506105ab610177565b918252602082015290565b61046f565b60408051602081019283528082019390935282526105da606083610155565b905190206001600160a01b031690565b634e487b7160e01b5f52601160045260245ffd5b6401000003d01903906401000003d019821161061657565b6105ea565b6040513d5f823e3d90fd5b90919070014551231950b75fc4402da1732fc9bebe195f820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161061657600116601b019182601b11610616576106d75f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b838052039060015afa156106ea575f5190565b61061b565b91929170014551231950b75fc4402da1732fc9bebe1990820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161061657600116601b019182601b11610616576106d75f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b156107a757565b63083cb12560e41b5f5260045ffd5b156107bd57565b6308330cb360e11b5f5260045ffd5b60c0810135916101008201356107e1816105fe565b5f94906401000003d01990820880151590816108e4575b50610802906107a0565b61080f60e08501356105fe565b9061012085013592610820846105fe565b966105b657610186966108aa936401000003d01991610140890135918391900809916401000003d019610852836105fe565b6401000003d0198580090861086d6401000003d019926105fe565b9008916401000003d019918290610883856105fe565b900890096001600160a01b03936401000003d01991906108a2906105fe565b9008906105bb565b169081151592836108be575b5050506107b6565b6108db929350906103e99160606020835193015191013591610626565b145f80806108b6565b95505061080260016101408601355f976401000003d019910914906107f8565b926109f8610a17926109926109f29361050f60409760208351930151948951958694602086019094939260a09260c08301967f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179884527f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b860208501526040840152606083015260808201520152565b845186356020828101919091528701356040820152608080880135606083015260a080890135918301919091526101008801359082015261012087013560c082015261050f906109e38160e0810184565b865194859360208501906104c4565b906104c4565b70014551231950b75fc4402da1732fc9bebe1990602081519101200690565b9101351490565b908115610a7d57604051602081526020808201526020604082015282606082015263400000f4600160fe1b0360808201526401000003d01960a082015260208160c08160055afa156100705751916401000003d01983800903610a7d57565b5f915056fea2646970667358221220929c79697fd9ce9ddacf8f88bebb3bbaa511080bbd75d0b7aa961f0e434cd1db64736f6c63430008230033",
}

// VRFVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFVerifierMetaData.ABI instead.
var VRFVerifierABI = VRFVerifierMetaData.ABI

// VRFVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFVerifierMetaData.Bin instead.
var VRFVerifierBin = VRFVerifierMetaData.Bin

// DeployVRFVerifier deploys a new Ethereum contract, binding an instance of VRFVerifier to it.
func DeployVRFVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFVerifier, error) {
	parsed, err := VRFVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFVerifier{VRFVerifierCaller: VRFVerifierCaller{contract: contract}, VRFVerifierTransactor: VRFVerifierTransactor{contract: contract}, VRFVerifierFilterer: VRFVerifierFilterer{contract: contract}}, nil
}

// VRFVerifier is an auto generated Go binding around an Ethereum contract.
type VRFVerifier struct {
	VRFVerifierCaller     // Read-only binding to the contract
	VRFVerifierTransactor // Write-only binding to the contract
	VRFVerifierFilterer   // Log filterer for contract events
}

// VRFVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFVerifierSession struct {
	Contract     *VRFVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFVerifierCallerSession struct {
	Contract *VRFVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VRFVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFVerifierTransactorSession struct {
	Contract     *VRFVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VRFVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFVerifierRaw struct {
	Contract *VRFVerifier // Generic contract binding to access the raw methods on
}

// VRFVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFVerifierCallerRaw struct {
	Contract *VRFVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VRFVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFVerifierTransactorRaw struct {
	Contract *VRFVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFVerifier creates a new instance of VRFVerifier, bound to a specific deployed contract.
func NewVRFVerifier(address common.Address, backend bind.ContractBackend) (*VRFVerifier, error) {
	contract, err := bindVRFVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFVerifier{VRFVerifierCaller: VRFVerifierCaller{contract: contract}, VRFVerifierTransactor: VRFVerifierTransactor{contract: contract}, VRFVerifierFilterer: VRFVerifierFilterer{contract: contract}}, nil
}

// NewVRFVerifierCaller creates a new read-only instance of VRFVerifier, bound to a specific deployed contract.
func NewVRFVerifierCaller(address common.Address, caller bind.ContractCaller) (*VRFVerifierCaller, error) {
	contract, err := bindVRFVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFVerifierCaller{contract: contract}, nil
}

// NewVRFVerifierTransactor creates a new write-only instance of VRFVerifier, bound to a specific deployed contract.
func NewVRFVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFVerifierTransactor, error) {
	contract, err := bindVRFVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFVerifierTransactor{contract: contract}, nil
}

// NewVRFVerifierFilterer creates a new log filterer instance of VRFVerifier, bound to a specific deployed contract.
func NewVRFVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFVerifierFilterer, error) {
	contract, err := bindVRFVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFVerifierFilterer{contract: contract}, nil
}

// bindVRFVerifier binds a generic wrapper to an already deployed contract.
func bindVRFVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFVerifier *VRFVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFVerifier.Contract.VRFVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFVerifier *VRFVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFVerifier.Contract.VRFVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFVerifier *VRFVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFVerifier.Contract.VRFVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFVerifier *VRFVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFVerifier *VRFVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFVerifier *VRFVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFVerifier.Contract.contract.Transact(opts, method, params...)
}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_VRFVerifier *VRFVerifierCaller) RandomnessFromProof(opts *bind.CallOpts, _gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFVerifier.contract.Call(opts, &out, "randomnessFromProof", _gammaX, _gammaY)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_VRFVerifier *VRFVerifierSession) RandomnessFromProof(_gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	return _VRFVerifier.Contract.RandomnessFromProof(&_VRFVerifier.CallOpts, _gammaX, _gammaY)
}

// RandomnessFromProof is a free data retrieval call binding the contract method 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_VRFVerifier *VRFVerifierCallerSession) RandomnessFromProof(_gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	return _VRFVerifier.Contract.RandomnessFromProof(&_VRFVerifier.CallOpts, _gammaX, _gammaY)
}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (_VRFVerifier *VRFVerifierCaller) VerifyRandomness(opts *bind.CallOpts, _proof IVrfVerifierProof, _pkX *big.Int, _pkY *big.Int, _nonce []byte) (bool, error) {
	var out []interface{}
	err := _VRFVerifier.contract.Call(opts, &out, "verifyRandomness", _proof, _pkX, _pkY, _nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (_VRFVerifier *VRFVerifierSession) VerifyRandomness(_proof IVrfVerifierProof, _pkX *big.Int, _pkY *big.Int, _nonce []byte) (bool, error) {
	return _VRFVerifier.Contract.VerifyRandomness(&_VRFVerifier.CallOpts, _proof, _pkX, _pkY, _nonce)
}

// VerifyRandomness is a free data retrieval call binding the contract method 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (_VRFVerifier *VRFVerifierCallerSession) VerifyRandomness(_proof IVrfVerifierProof, _pkX *big.Int, _pkY *big.Int, _nonce []byte) (bool, error) {
	return _VRFVerifier.Contract.VerifyRandomness(&_VRFVerifier.CallOpts, _proof, _pkX, _pkY, _nonce)
}
