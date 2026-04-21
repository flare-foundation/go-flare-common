// Package seed decodes base58-encoded rippled-native XRPL family seeds.
package seed

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
)

// Length of the raw family seed payload.
const Length = 16

// familySeedPrefix is TokenType::FamilySeed from rippled include/xrpl/protocol/tokens.h.
const familySeedPrefix = 0x21

// encodedCap is the longest valid encoded family seed; rippled tokens.cpp caps decoded tokens at 64 bytes.
const encodedCap = 40

// DecodeFamilySeed decodes a rippled-native base58 seed string into its 16-byte payload.
// The string must be base58-encoded with TokenType::FamilySeed (0x21) prefix and a 4-byte double-SHA256 checksum.
func DecodeFamilySeed(s string) ([]byte, error) {
	if len(s) > encodedCap {
		return nil, fmt.Errorf("seed too long, got %d chars", len(s))
	}
	raw, err := base58.XRPLCoder.Decode(s)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}
	if len(raw) != 1+Length+4 {
		return nil, fmt.Errorf("unexpected decoded length %d, want %d", len(raw), 1+Length+4)
	}
	if raw[0] != familySeedPrefix {
		return nil, fmt.Errorf("wrong token prefix %#x, want %#x", raw[0], familySeedPrefix)
	}
	want := hash.Checksum(raw[:1+Length])
	if !bytes.Equal(want, raw[1+Length:]) {
		return nil, errors.New("invalid checksum")
	}
	out := make([]byte, Length)
	copy(out, raw[1:1+Length])
	return out, nil
}

// EncodeFamilySeed encodes a 16-byte seed payload as a rippled-native base58 family seed string.
// Primarily intended for testing the decode path round-trip.
func EncodeFamilySeed(payload []byte) (string, error) {
	if len(payload) != Length {
		return "", fmt.Errorf("invalid seed length %d, want %d", len(payload), Length)
	}
	body := make([]byte, 0, 1+Length+4)
	body = append(body, familySeedPrefix)
	body = append(body, payload...)
	body = append(body, hash.Checksum(body)...)
	return base58.XRPLCoder.Encode(body), nil
}
