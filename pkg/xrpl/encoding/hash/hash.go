package hash

import (
	"crypto/sha256"
	"crypto/sha512"

	//nolint
	"golang.org/x/crypto/ripemd160"
)

// Sha512Half returns the first half of sha512 hash
func Sha512Half(blob []byte) []byte {
	hash := sha512.Sum512(blob)
	return hash[:32]
}

// Sha256RipeMD160 computes sha256 hash of a RipeMD160 hash.
//
// ONLY FOR LEGACY PURPOSES.
func Sha256RipeMD160(b []byte) []byte {
	ripe := ripemd160.New()
	sha := sha256.New()
	sha.Write(b)
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}

// Checksum returns a checksum fot the byte string (last 4 bytes of double sha256 hash).
func Checksum(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])
	return h2[:4]
}
