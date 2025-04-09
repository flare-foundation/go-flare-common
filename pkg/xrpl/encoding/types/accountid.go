package types

import (
	"bytes"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/hash"
)

// AccountID is used for serialization of AccountID fields. https://xrpl.org/docs/references/protocol/binary-format#accountid-fields

type AccountID struct{}

// ToBytes serializes value of an AccountID field.
func (a *AccountID) ToBytes(value any, _ bool) ([]byte, error) {
	address, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	out, err := id(address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	return out, nil
}

// id returns accountID of an address. If the address has an invalid checksum an error is returned.
func id(address string) ([]byte, error) {
	addressBytes, err := base58.XRPLCoder.Decode(address)
	if err != nil {
		return nil, fmt.Errorf("decoding address: %v", err)
	}

	// length
	if l := len(addressBytes); l != 25 {
		return nil, fmt.Errorf("wrong length. Expected %d got %d", 25, l)
	}

	// checksum
	provided := addressBytes[21:]

	computed := hash.Checksum(addressBytes[:21])
	if !bytes.Equal(provided, computed) {
		return nil, fmt.Errorf("invalid checksum")
	}

	// trim leading byte and checksum
	return addressBytes[1:21], nil
}
