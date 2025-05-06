package types

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

// AccountID is used for serialization of AccountID fields. https://xrpl.org/docs/references/protocol/binary-format#accountid-fields
type accountID struct{}

var AccountID = &accountID{}

// ToBytes serializes value of an AccountID field.
func (a *accountID) ToBytes(value any, _ bool) ([]byte, error) {
	addr, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	out, err := address.ID(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	return out, nil
}

// ToJson reads next 20 bytes and converts them to a string xrpl address.
func (a *accountID) ToJson(b *bytes.Buffer, _ int) (any, error) {
	const l = 20
	value := make([]byte, l)

	n, err := b.Read(value)
	if err != nil {
		return nil, fmt.Errorf("cannot read account id from buffer: %v", err)
	}
	if n != l {
		return nil, outOfBytes("account id", l, n)
	}

	addr, err := address.Address(value)
	if err != nil {
		return nil, fmt.Errorf("deserializing accountID %v: %v", hex.EncodeToString(value), err)
	}

	return addr, nil
}
