package address

import (
	"bytes"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
)

const (
	accountPrefix = 0x00
)

// ID returns accountID of an address. If the address has an invalid checksum, an error is returned.
func ID(address string) ([]byte, error) {
	addressBytes, err := base58.XRPLCoder.Decode(address)
	if err != nil {
		return nil, fmt.Errorf("decoding: %v", err)
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

	computed := Checksum(addressBytes[:21])
	if !bytes.Equal(provided, computed) {
		return nil, fmt.Errorf("invalid checksum")
	}

	// trim leading byte and checksum
	return addressBytes[1:21], nil
}

// Address calculates XRPL Address from accountID.
// It checks that id is 20 bytes long. Use IDToAddress to bypass the check.
func Address(id []byte) (string, error) {
	if len(id) != 20 {
		return "", fmt.Errorf("invalid id length %d, should be 20", len(id))
	}

	return IDToAddress(id), nil
}

// IDToAddress calculates XRPL address from accountID.
// It does not check that id is 20 bytes long.
func IDToAddress(id []byte) string {
	augmented := make([]byte, 0, len(id)+1+4)

	augmented = append(augmented, accountPrefix) // account prefix
	augmented = append(augmented, id...)

	cs := Checksum(augmented) // checksum

	augmented = append(augmented, cs...)

	return base58.XRPLCoder.Encode(augmented)
}

// Checksum returns a checksum for the byte string (first 4 bytes of double sha256 hash).
func Checksum(b []byte) []byte {
	return hash.DoubleSha256(b)[:4]
}
