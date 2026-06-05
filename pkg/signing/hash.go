package signing

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
)

// Payload is the Go representation of the Solidity Payload struct.
//
//	struct Payload {
//	    bytes32 prefix;
//	    uint256 chainId;
//	    bytes32 dataHash;
//	}
type Payload struct {
	Prefix   [32]byte `abi:"prefix"`
	ChainID  *big.Int `abi:"chainId"`
	DataHash [32]byte `abi:"dataHash"`
}

// PayloadArgument is the abi.Argument matching the Solidity Payload struct:
var PayloadArgument abi.Argument

func init() {
	payloadType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "prefix", Type: "bytes32"},
		{Name: "chainId", Type: "uint256"},
		{Name: "dataHash", Type: "bytes32"},
	})
	if err != nil {
		panic(err)
	}

	PayloadArgument = abi.Argument{Type: payloadType}
}

// Hash ABI-encodes the payload and returns its keccak256 hash.
func (p Payload) Hash() ([32]byte, error) {
	encoded, err := abicoder.Encode(PayloadArgument, p)
	if err != nil {
		return [32]byte{}, fmt.Errorf("packing payload: %w", err)
	}

	return crypto.Keccak256Hash(encoded), nil
}
