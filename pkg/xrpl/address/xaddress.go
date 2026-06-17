package address

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
)

const (
	xPayloadLen = 31 // 2 prefix + 20 accountID + 1 flag + 8 tag
	xDecodedLen = 35 // payload + 4-byte checksum
	xEncodedCap = 48 // 35 bytes encode to at most 48 base58 characters

	flagNoTag byte = 0x00
	flagTag   byte = 0x01 // 32-bit tag
)

// X-address network prefixes, see https://github.com/XRPLF/XRPL-Standards/discussions/6.
var (
	xPrefixMain = []byte{0x05, 0x44}
	xPrefixTest = []byte{0x04, 0x93}
)

// XAddress encodes an accountID and an optional tag into an X-address. The embedded tag is a
// source tag when the address identifies the sender and a destination tag when it identifies
// the recipient. A nil tag produces a tagless X-address. test selects the testnet prefix.
func XAddress(id []byte, tag *uint32, test bool) (string, error) {
	if len(id) != 20 {
		return "", fmt.Errorf("invalid id length %d, should be 20", len(id))
	}

	payload := make([]byte, 0, xPayloadLen)
	if test {
		payload = append(payload, xPrefixTest...)
	} else {
		payload = append(payload, xPrefixMain...)
	}
	payload = append(payload, id...)

	tagBytes := make([]byte, 8)
	if tag != nil {
		payload = append(payload, flagTag)
		binary.LittleEndian.PutUint32(tagBytes, *tag)
	} else {
		payload = append(payload, flagNoTag)
	}
	payload = append(payload, tagBytes...)

	return base58.XRPLCoder.Encode(append(payload, hash.Checksum(payload)...)), nil
}

// XID decodes an X-address into its accountID, optional tag (source or destination depending on
// use, see XAddress), and network. A nil tag indicates a tagless X-address. An invalid checksum
// or structure returns an error.
func XID(xAddress string) (id []byte, tag *uint32, test bool, err error) {
	if len(xAddress) > xEncodedCap {
		return nil, nil, false, fmt.Errorf("x-address too long, got %d chars", len(xAddress))
	}

	decoded, err := base58.XRPLCoder.Decode(xAddress)
	if err != nil {
		return nil, nil, false, fmt.Errorf("decoding: %w", err)
	}

	if l := len(decoded); l != xDecodedLen {
		return nil, nil, false, fmt.Errorf("wrong length, expected %d got %d", xDecodedLen, l)
	}

	provided := decoded[xPayloadLen:]
	if computed := hash.Checksum(decoded[:xPayloadLen]); !bytes.Equal(provided, computed) {
		return nil, nil, false, errors.New("invalid checksum")
	}

	switch prefix := decoded[:2]; {
	case bytes.Equal(prefix, xPrefixMain):
		test = false
	case bytes.Equal(prefix, xPrefixTest):
		test = true
	default:
		return nil, nil, false, fmt.Errorf("invalid prefix %x", prefix)
	}

	switch flag := decoded[22]; flag {
	case flagNoTag:
		if !allZero(decoded[23:xPayloadLen]) {
			return nil, nil, false, errors.New("tag bytes set without tag flag")
		}
	case flagTag:
		if !allZero(decoded[27:xPayloadLen]) {
			return nil, nil, false, errors.New("tag exceeds 32 bits")
		}
		t := binary.LittleEndian.Uint32(decoded[23:27])
		tag = &t
	default:
		return nil, nil, false, fmt.Errorf("unsupported tag flag %d", flag)
	}

	return decoded[2:22], tag, test, nil
}

// ClassicToX converts a classic r-address and an optional tag (see XAddress) into an X-address.
func ClassicToX(rAddress string, tag *uint32, test bool) (string, error) {
	id, err := ID(rAddress)
	if err != nil {
		return "", fmt.Errorf("decoding classic address: %w", err)
	}

	return XAddress(id, tag, test)
}

// XToClassic converts an X-address into its classic r-address and optional tag (see XAddress).
func XToClassic(xAddress string) (rAddress string, tag *uint32, test bool, err error) {
	var id []byte
	id, tag, test, err = XID(xAddress)
	if err != nil {
		return "", nil, false, fmt.Errorf("decoding x-address: %w", err)
	}

	return IDToAddress(id), tag, test, nil
}

// Normalize accepts either a classic r-address or an X-address and returns the canonical
// classic r-address together with the embedded tag (nil if none; see XAddress for its meaning).
// The X-address network flag is not surfaced, as classic addresses are network-agnostic.
func Normalize(addr string) (rAddress string, tag *uint32, err error) {
	if addr == "" {
		return "", nil, errors.New("empty address")
	}

	switch addr[0] {
	case 'X', 'T':
		rAddress, tag, _, err = XToClassic(addr)
		return rAddress, tag, err
	case 'r':
		var id []byte
		id, err = ID(addr)
		if err != nil {
			return "", nil, fmt.Errorf("decoding classic address: %w", err)
		}

		return IDToAddress(id), nil, nil
	default:
		return "", nil, fmt.Errorf("unrecognized address prefix %q", addr[0])
	}
}

// allZero reports whether every byte in b is zero.
func allZero(b []byte) bool {
	for _, x := range b {
		if x != 0 {
			return false
		}
	}

	return true
}
