package types

import (
	"bytes"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

// AccountID is used for serialization of AccountID fields. https://xrpl.org/docs/references/protocol/binary-format#accountid-fields
type accountID struct{}

var AccountID = &accountID{}

// ToBytes serializes value of an AccountID field.
func (*accountID) ToBytes(value any, _ bool) ([]byte, error) {
	addr, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	out, err := address.ID(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	return out, nil
}

// ToJSON reads next 20 bytes and converts them to a string xrpl address.
//
// length must be either 0 (inline caller — pathset, xchainbridge) or 20
// (canonical VL-prefixed encoding). Any other length is non-canonical and
// rippled rejects it; we do the same to avoid VL/payload-length disagreement
// silently smuggling bytes into the next field.
func (*accountID) ToJSON(b *bytes.Buffer, length int) (any, error) {
	const l = 20
	if length != 0 && length != l {
		return nil, fmt.Errorf("account id: non-canonical length %d (expected %d)", length, l)
	}
	value := make([]byte, l)

	n, err := b.Read(value)
	if err != nil {
		return nil, fmt.Errorf("reading account id from buffer: %w", err)
	}
	if n != l {
		return nil, outOfBytes("account id", l, n)
	}

	return address.IDToAddress(value), nil
}
