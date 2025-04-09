package types

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
)

// https://xrpl.org/docs/references/protocol/binary-format#accountid-fields

type AccountID struct{}

func (a *AccountID) ToBytes(value any, _ bool) ([]byte, error) {
	address, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	out, err := trimAddress(address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	return out, nil
}

func trimAddress(addressHex string) ([]byte, error) {
	address, err := base58.XRPLCoder.Decode(addressHex)
	if err != nil {
		return nil, fmt.Errorf("decoding address: %v", err)
	}

	// length
	if l := len(address); l != 25 {
		return nil, fmt.Errorf("wrong length. Expected %d got %d", 25, l)
	}

	// checksum
	provided := address[21:]

	computed := doubleSha256(address[:21])[:4]
	if !bytes.Equal(provided, computed) {
		return nil, fmt.Errorf("invalid checksum")
	}

	// trim leading byte and checksum
	return address[1:21], nil
}

func doubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sha := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sha)
	return hasher.Sum(nil)
}
