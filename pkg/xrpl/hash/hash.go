package hash

import (
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/ripemd160" //nolint:gosec,staticcheck
)

// Sha512Half returns the first half of sha512 hash.
func Sha512Half(b []byte) []byte {
	hash := sha512.Sum512(b)
	return hash[:32]
}

// DoubleSha256 hashes b with sha256 twice.
func DoubleSha256(b []byte) []byte {
	sha := sha256.New()
	sha.Write(b)
	h0 := sha.Sum(nil)
	sha.Reset()
	sha.Write(h0)
	return sha.Sum(nil)
}

// Sha256RipeMD160 computes sha256 hash of a RipeMD160 hash.
//
// ONLY FOR LEGACY PURPOSES.
func Sha256RipeMD160(b []byte) []byte {
	ripe := ripemd160.New() //nolint:gosec
	sha := sha256.New()
	sha.Write(b)
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}

// Checksum returns a checksum for the byte string (first 4 bytes of double sha256 hash).
func Checksum(b []byte) []byte {
	return DoubleSha256(b)[:4]
}
