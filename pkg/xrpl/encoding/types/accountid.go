package types

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/hash"
)

const accountPrefix = 0x00

// AccountID is used for serialization of AccountID fields. https://xrpl.org/docs/references/protocol/binary-format#accountid-fields
type accountID struct{}

var AccountID = &accountID{}

// ToBytes serializes value of an AccountID field.
func (a *accountID) ToBytes(value any, _ bool) ([]byte, error) {
	address, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	out, err := ID(address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	return out, nil
}

// ID returns accountID of an address. If the address has an invalid checksum an error is returned.
func ID(address string) ([]byte, error) {
	addressBytes, err := base58.XRPLCoder.Decode(address)
	if err != nil {
		return nil, fmt.Errorf("decoding address: %v", err)
	}

	// length
	if l := len(addressBytes); l != 25 {
		return nil, fmt.Errorf("wrong length. Expected %d got %d", 25, l)
	}

	// leading byte
	if addressBytes[0] != 0 {
		return nil, fmt.Errorf("wrong leading byte %v", address)
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

// Address calculates XRPL Address from accountID.
// It does not check that id is 20 bytes long.
func Address(id []byte) (string, error) {
	if len(id) != 20 {
		return "", fmt.Errorf("invalid id length %d, should be 20", len(id))
	}

	augmented := make([]byte, 0, len(id)+1+4)

	augmented = append(augmented, accountPrefix) // account prefix
	augmented = append(augmented, id...)

	cs := hash.Checksum(augmented) // checksum

	augmented = append(augmented, cs...)

	return base58.XRPLCoder.Encode(augmented), nil
}

func (a *accountID) ToJson(b *bytes.Buffer, _ int) (any, error) {
	value := make([]byte, 20)

	_, err := b.Read(value)
	if err != nil {
		return nil, fmt.Errorf("cannot read account id from buffer: %v", err)
	}

	addr, err := Address(value)
	if err != nil {
		return nil, fmt.Errorf("deserializing accountID %v: %v", hex.EncodeToString(value), err)
	}

	return addr, nil
}
