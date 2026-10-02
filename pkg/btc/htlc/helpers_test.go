package htlc

import (
	"bytes"
	"slices"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
)

// makeFakePubKey derives a valid compressed secp256k1 key from a seed byte, so
// vectors can name distinct keys without shipping fixtures.
//
// Valid, not arbitrary bytes: BuildScript rejects anything that is not a point
// on the curve, so a test using 33 bytes of filler would exercise the rejection
// path instead of the one it means to.
func makeFakePubKey(seed byte) []byte {
	var scalar [32]byte
	for i := range scalar {
		scalar[i] = seed + byte(i)
	}
	// The all-zero scalar is invalid for secp256k1, and a seed that nulls the
	// low bytes would produce it; pin the LSB.
	scalar[31] |= 1
	priv, _ := btcec.PrivKeyFromBytes(scalar[:])
	return priv.PubKey().SerializeCompressed()
}

// walletKeys is a BIP-67-sorted set of compressed keys, standing in for the
// wallet's derived multisig set.
//
// Sorted, not merely distinct: BuildScript takes the keys in the order they
// will appear in the script, and OP_CHECKMULTISIG requires the signatures to
// match that order — an unsorted fixture would build a script that no correctly
// ordered witness can satisfy.
func walletKeys(t *testing.T, n int) [][]byte {
	t.Helper()
	keys := make([][]byte, 0, n)
	for i := range n {
		keys = append(keys, makeFakePubKey(byte(0x10+i)))
	}
	slices.SortFunc(keys, bytes.Compare)
	return keys
}
