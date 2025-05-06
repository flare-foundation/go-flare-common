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
