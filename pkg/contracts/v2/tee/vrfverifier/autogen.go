// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfverifier

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
var VRFVerifierMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"COutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GammaNotOnCurve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GammaXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashToCurveExceededIterationLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCGammaWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVWitness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZInv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PkNotOnCurve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PkXUnverifiable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gammaX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gammaY\",\"type\":\"uint256\"}],\"name\":\"randomnessFromProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"gamma\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"u\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"cGamma\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Point\",\"name\":\"v\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structIVrfVerifier.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_pkX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pkY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"verifyRandomness\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "VRFVerifier",
	Bin: "0x60808060405234601557610ab8908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806313a0c07c1461003457631fc57fbd1461002f575f80fd5b610074565b346100705760403660031901126100705760043560a05260243560c05260406080526100606060610111565b602060805160a020604051908152f35b5f80fd5b346100705736600319016101c08112610070576101601361007057610184356101a4356101643567ffffffffffffffff821161007057366023830112156100705781600401359067ffffffffffffffff8211610070573660248385010111610070576100f99360246100e79401916102bc565b60405190151581529081906020820190565b0390f35b634e487b7160e01b5f52604160045260245ffd5b601f80199101166080016080811067ffffffffffffffff82111761013457604052565b6100fd565b6040810190811067ffffffffffffffff82111761013457604052565b90601f8019910116810190811067ffffffffffffffff82111761013457604052565b60405190610186604083610155565b565b1561018f57565b630b097aa560e41b5f5260045ffd5b9190826040910312610070576040516101b681610139565b6020808294803584520135910152565b156101cd57565b63094d477160e41b5f5260045ffd5b156101e357565b63040107d160e11b5f5260045ffd5b156101f957565b63ec6e729760e01b5f5260045ffd5b1561020f57565b63263091cb60e11b5f5260045ffd5b1561022557565b6304605f3560e21b5f5260045ffd5b92919267ffffffffffffffff8211610134576040519161025e601f8201601f191660200184610155565b829481845281830111610070578281602093845f960137010152565b1561028157565b63ff8c457560e01b5f5260045ffd5b1561029757565b63de86b44560e01b5f5260045ffd5b156102ad57565b63125ecc0760e31b5f5260045ffd5b61040861041a946102e46102df6102d1610177565b858152866020820152610483565b610188565b6102ff6102fa6102f536600461019e565b610483565b6101c6565b60443580151580610454575b610314906101dc565b6103d861038a6103836064359861033e70014551231950b75fc4402da1732fc9bebe198b106101f2565b61035b70014551231950b75fc4402da1732fc9bebe198910610208565b6004359561037c70014551231950b75fc4402da1732fc9bebe19881061021e565b3691610234565b87876104ec565b966103a970014551231950b75fc4402da1732fc9bebe1989511061027a565b856001600160a01b036103c060a4356084356105bb565b168489821515938461043a575b505050509050610290565b6103f56103e960e43560c4356105bb565b6001600160a01b031690565b90811515928361041d575b5050506102a6565b6104138360046107cc565b6004610904565b90565b610431929350906103e99160243590610626565b145f8080610400565b61044994506103e993956106ef565b14805f8489896103cd565b5070014551231950b75fc4402da1732fc9bebe19811061030b565b634e487b7160e01b5f52601260045260245ffd5b602081015190516401000003d019906007908290818180090908906401000003d0199080091490565b604051906104b982610139565b5f6020838281520152565b805191908290602001825e015f815290565b61041a93926040928252602082015201906104c4565b919061051d906104fa6104ac565b5061050f6040519384926020840196876104d6565b03601f198101835282610155565b5190205f916401000003d019820691835b61010085106105465763424c4e9d60e01b5f5260045ffd5b6105b6575f926105656401000003d01960078184818180090908610a1e565b8061059d5750506040516105858161050f60208201948560209181520190565b5190206001909301926401000003d01981069261052e565b9294509250506105ab610177565b918252602082015290565b61046f565b60408051602081019283528082019390935282526105da606083610155565b905190206001600160a01b031690565b634e487b7160e01b5f52601160045260245ffd5b6401000003d01903906401000003d019821161061657565b6105ea565b6040513d5f823e3d90fd5b90919070014551231950b75fc4402da1732fc9bebe195f820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161061657600116601b019182601b11610616576106d75f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b838052039060015afa156106ea575f5190565b61061b565b91929170014551231950b75fc4402da1732fc9bebe1990820970014551231950b75fc4402da1732fc9bebe19039270014551231950b75fc4402da1732fc9bebe19841161061657600116601b019182601b11610616576106d75f9360ff9360209660405195869570014551231950b75fc4402da1732fc9bebe1990840970014551231950b75fc4402da1732fc9bebe19909206865290921660ff166020850152604084015260608301526080820190565b156107a757565b63083cb12560e41b5f5260045ffd5b156107bd57565b6308330cb360e11b5f5260045ffd5b60c0810135916101008201356107e1816105fe565b5f94906401000003d01990820880151590816108e4575b50610802906107a0565b61080f60e08501356105fe565b9061012085013592610820846105fe565b966105b657610186966108aa936401000003d01991610140890135918391900809916401000003d019610852836105fe565b6401000003d0198580090861086d6401000003d019926105fe565b9008916401000003d019918290610883856105fe565b900890096001600160a01b03936401000003d01991906108a2906105fe565b9008906105bb565b169081151592836108be575b5050506107b6565b6108db929350906103e99160606020835193015191013591610626565b145f80806108b6565b95505061080260016101408601355f976401000003d019910914906107f8565b926109f8610a17926109926109f29361050f60409760208351930151948951958694602086019094939260a09260c08301967f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179884527f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b860208501526040840152606083015260808201520152565b845186356020828101919091528701356040820152608080880135606083015260a080890135918301919091526101008801359082015261012087013560c082015261050f906109e38160e0810184565b865194859360208501906104c4565b906104c4565b70014551231950b75fc4402da1732fc9bebe1990602081519101200690565b9101351490565b908115610a7d57604051602081526020808201526020604082015282606082015263400000f4600160fe1b0360808201526401000003d01960a082015260208160c08160055afa156100705751916401000003d01983800903610a7d57565b5f915056fea2646970667358221220929c79697fd9ce9ddacf8f88bebb3bbaa511080bbd75d0b7aa961f0e434cd1db64736f6c63430008230033",
}

// VRFVerifier is an auto generated Go binding around an Ethereum contract.
type VRFVerifier struct {
	abi abi.ABI
}

// NewVRFVerifier creates a new instance of VRFVerifier.
func NewVRFVerifier() *VRFVerifier {
	parsed, err := VRFVerifierMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &VRFVerifier{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *VRFVerifier) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackRandomnessFromProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13a0c07c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (vRFVerifier *VRFVerifier) PackRandomnessFromProof(gammaX *big.Int, gammaY *big.Int) []byte {
	enc, err := vRFVerifier.abi.Pack("randomnessFromProof", gammaX, gammaY)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRandomnessFromProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13a0c07c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (vRFVerifier *VRFVerifier) TryPackRandomnessFromProof(gammaX *big.Int, gammaY *big.Int) ([]byte, error) {
	return vRFVerifier.abi.Pack("randomnessFromProof", gammaX, gammaY)
}

// UnpackRandomnessFromProof is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13a0c07c.
//
// Solidity: function randomnessFromProof(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (vRFVerifier *VRFVerifier) UnpackRandomnessFromProof(data []byte) ([32]byte, error) {
	out, err := vRFVerifier.abi.Unpack("randomnessFromProof", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackVerifyRandomness is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1fc57fbd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (vRFVerifier *VRFVerifier) PackVerifyRandomness(proof IVrfVerifierProof, pkX *big.Int, pkY *big.Int, nonce []byte) []byte {
	enc, err := vRFVerifier.abi.Pack("verifyRandomness", proof, pkX, pkY, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyRandomness is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1fc57fbd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (vRFVerifier *VRFVerifier) TryPackVerifyRandomness(proof IVrfVerifierProof, pkX *big.Int, pkY *big.Int, nonce []byte) ([]byte, error) {
	return vRFVerifier.abi.Pack("verifyRandomness", proof, pkX, pkY, nonce)
}

// UnpackVerifyRandomness is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1fc57fbd.
//
// Solidity: function verifyRandomness(((uint256,uint256),uint256,uint256,(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256) _proof, uint256 _pkX, uint256 _pkY, bytes _nonce) view returns(bool _valid)
func (vRFVerifier *VRFVerifier) UnpackVerifyRandomness(data []byte) (bool, error) {
	out, err := vRFVerifier.abi.Unpack("verifyRandomness", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (vRFVerifier *VRFVerifier) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["COutOfRange"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackCOutOfRangeError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["GammaNotOnCurve"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackGammaNotOnCurveError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["GammaXUnverifiable"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackGammaXUnverifiableError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["HXUnverifiable"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackHXUnverifiableError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["HashToCurveExceededIterationLimit"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackHashToCurveExceededIterationLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["InvalidCGammaWitness"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackInvalidCGammaWitnessError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["InvalidUWitness"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackInvalidUWitnessError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["InvalidVWitness"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackInvalidVWitnessError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["InvalidZInv"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackInvalidZInvError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["PkNotOnCurve"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackPkNotOnCurveError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["PkXUnverifiable"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackPkXUnverifiableError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRFVerifier.abi.Errors["SOutOfRange"].ID.Bytes()[:4]) {
		return vRFVerifier.UnpackSOutOfRangeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VRFVerifierCOutOfRange represents a COutOfRange error raised by the VRFVerifier contract.
type VRFVerifierCOutOfRange struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error COutOfRange()
func VRFVerifierCOutOfRangeErrorID() common.Hash {
	return common.HexToHash("0x08020fa2ea6423eac5402d46dbddde59da24c343b8abd1a39c9071bd55ca3e74")
}

// UnpackCOutOfRangeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error COutOfRange()
func (vRFVerifier *VRFVerifier) UnpackCOutOfRangeError(raw []byte) (*VRFVerifierCOutOfRange, error) {
	out := new(VRFVerifierCOutOfRange)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "COutOfRange", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierGammaNotOnCurve represents a GammaNotOnCurve error raised by the VRFVerifier contract.
type VRFVerifierGammaNotOnCurve struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GammaNotOnCurve()
func VRFVerifierGammaNotOnCurveErrorID() common.Hash {
	return common.HexToHash("0x94d477104190264149db1543e02044f8f73e056014aad0dbfac726e1eea6e3ce")
}

// UnpackGammaNotOnCurveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GammaNotOnCurve()
func (vRFVerifier *VRFVerifier) UnpackGammaNotOnCurveError(raw []byte) (*VRFVerifierGammaNotOnCurve, error) {
	out := new(VRFVerifierGammaNotOnCurve)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "GammaNotOnCurve", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierGammaXUnverifiable represents a GammaXUnverifiable error raised by the VRFVerifier contract.
type VRFVerifierGammaXUnverifiable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GammaXUnverifiable()
func VRFVerifierGammaXUnverifiableErrorID() common.Hash {
	return common.HexToHash("0x11817cd44befacbf074aa667e75b7d5cbce0d0df78bb17fc713cc0be7a10545d")
}

// UnpackGammaXUnverifiableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GammaXUnverifiable()
func (vRFVerifier *VRFVerifier) UnpackGammaXUnverifiableError(raw []byte) (*VRFVerifierGammaXUnverifiable, error) {
	out := new(VRFVerifierGammaXUnverifiable)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "GammaXUnverifiable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierHXUnverifiable represents a HXUnverifiable error raised by the VRFVerifier contract.
type VRFVerifierHXUnverifiable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HXUnverifiable()
func VRFVerifierHXUnverifiableErrorID() common.Hash {
	return common.HexToHash("0xff8c4575cefb2cef88dd8b13c094db9d1e2f9476be2670127e482bf7555186c9")
}

// UnpackHXUnverifiableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HXUnverifiable()
func (vRFVerifier *VRFVerifier) UnpackHXUnverifiableError(raw []byte) (*VRFVerifierHXUnverifiable, error) {
	out := new(VRFVerifierHXUnverifiable)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "HXUnverifiable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierHashToCurveExceededIterationLimit represents a HashToCurveExceededIterationLimit error raised by the VRFVerifier contract.
type VRFVerifierHashToCurveExceededIterationLimit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HashToCurveExceededIterationLimit()
func VRFVerifierHashToCurveExceededIterationLimitErrorID() common.Hash {
	return common.HexToHash("0x424c4e9d03be4c3707316edd346c64f5e4bbe9630b77481322ae6164ee6d0ca0")
}

// UnpackHashToCurveExceededIterationLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HashToCurveExceededIterationLimit()
func (vRFVerifier *VRFVerifier) UnpackHashToCurveExceededIterationLimitError(raw []byte) (*VRFVerifierHashToCurveExceededIterationLimit, error) {
	out := new(VRFVerifierHashToCurveExceededIterationLimit)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "HashToCurveExceededIterationLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierInvalidCGammaWitness represents a InvalidCGammaWitness error raised by the VRFVerifier contract.
type VRFVerifierInvalidCGammaWitness struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCGammaWitness()
func VRFVerifierInvalidCGammaWitnessErrorID() common.Hash {
	return common.HexToHash("0x92f660380fcfb78286a9fed48d2738420257c42c384046f38f362ccf5c65924b")
}

// UnpackInvalidCGammaWitnessError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCGammaWitness()
func (vRFVerifier *VRFVerifier) UnpackInvalidCGammaWitnessError(raw []byte) (*VRFVerifierInvalidCGammaWitness, error) {
	out := new(VRFVerifierInvalidCGammaWitness)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "InvalidCGammaWitness", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierInvalidUWitness represents a InvalidUWitness error raised by the VRFVerifier contract.
type VRFVerifierInvalidUWitness struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidUWitness()
func VRFVerifierInvalidUWitnessErrorID() common.Hash {
	return common.HexToHash("0xde86b445764b3d842275b7948c7576b95bf768412afe5243a21e629c89a16b0f")
}

// UnpackInvalidUWitnessError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidUWitness()
func (vRFVerifier *VRFVerifier) UnpackInvalidUWitnessError(raw []byte) (*VRFVerifierInvalidUWitness, error) {
	out := new(VRFVerifierInvalidUWitness)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "InvalidUWitness", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierInvalidVWitness represents a InvalidVWitness error raised by the VRFVerifier contract.
type VRFVerifierInvalidVWitness struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidVWitness()
func VRFVerifierInvalidVWitnessErrorID() common.Hash {
	return common.HexToHash("0x106619662374bc3f6768d0373a20c937eec5d84588c552a01d6cad0fde4f70e8")
}

// UnpackInvalidVWitnessError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidVWitness()
func (vRFVerifier *VRFVerifier) UnpackInvalidVWitnessError(raw []byte) (*VRFVerifierInvalidVWitness, error) {
	out := new(VRFVerifierInvalidVWitness)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "InvalidVWitness", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierInvalidZInv represents a InvalidZInv error raised by the VRFVerifier contract.
type VRFVerifierInvalidZInv struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidZInv()
func VRFVerifierInvalidZInvErrorID() common.Hash {
	return common.HexToHash("0x83cb1250dafb539dd7e2da8cc2bee3f969cb82c3f89dd88171c30e6e39619597")
}

// UnpackInvalidZInvError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidZInv()
func (vRFVerifier *VRFVerifier) UnpackInvalidZInvError(raw []byte) (*VRFVerifierInvalidZInv, error) {
	out := new(VRFVerifierInvalidZInv)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "InvalidZInv", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierPkNotOnCurve represents a PkNotOnCurve error raised by the VRFVerifier contract.
type VRFVerifierPkNotOnCurve struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PkNotOnCurve()
func VRFVerifierPkNotOnCurveErrorID() common.Hash {
	return common.HexToHash("0xb097aa503b78216c526e4885770603f68f0dee72bcd2398f9ac7c87b8986e841")
}

// UnpackPkNotOnCurveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PkNotOnCurve()
func (vRFVerifier *VRFVerifier) UnpackPkNotOnCurveError(raw []byte) (*VRFVerifierPkNotOnCurve, error) {
	out := new(VRFVerifierPkNotOnCurve)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "PkNotOnCurve", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierPkXUnverifiable represents a PkXUnverifiable error raised by the VRFVerifier contract.
type VRFVerifierPkXUnverifiable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PkXUnverifiable()
func VRFVerifierPkXUnverifiableErrorID() common.Hash {
	return common.HexToHash("0x4c61239609343d3d423c09b4d999bdf3269bd00e338911e8732db7e819b82779")
}

// UnpackPkXUnverifiableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PkXUnverifiable()
func (vRFVerifier *VRFVerifier) UnpackPkXUnverifiableError(raw []byte) (*VRFVerifierPkXUnverifiable, error) {
	out := new(VRFVerifierPkXUnverifiable)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "PkXUnverifiable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFVerifierSOutOfRange represents a SOutOfRange error raised by the VRFVerifier contract.
type VRFVerifierSOutOfRange struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SOutOfRange()
func VRFVerifierSOutOfRangeErrorID() common.Hash {
	return common.HexToHash("0xec6e72977645fb8627439418f2e770ea19e938bcd091ed92d0603a74d21c0dcf")
}

// UnpackSOutOfRangeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SOutOfRange()
func (vRFVerifier *VRFVerifier) UnpackSOutOfRangeError(raw []byte) (*VRFVerifierSOutOfRange, error) {
	out := new(VRFVerifierSOutOfRange)
	if err := vRFVerifier.abi.UnpackIntoInterface(out, "SOutOfRange", raw); err != nil {
		return nil, err
	}
	return out, nil
}
