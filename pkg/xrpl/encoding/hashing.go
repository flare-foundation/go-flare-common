package encoding

import (
	"crypto/sha512"
	"encoding/binary"
)

const (
	singlePrefix uint32 = 0x53545800
	multiPrefix  uint32 = 0x534D5400
	signedPrefix uint32 = 0x54584E00
)

// MessageToSign computes a hash of a transaction for signing.
// If multiSig is true, hash is for multi-signing. Otherwise for single-signing.
//
// For multi-signing, accountID of the signer should be provided.
func MessageToSign(txBlob []byte, multiSig bool, accountID []byte) []byte {
	length := len(txBlob) + 4
	if multiSig {
		length += 20
	}

	prefixed := make([]byte, 0, length)

	if multiSig {
		prefixed = binary.BigEndian.AppendUint32(prefixed, multiPrefix)
	} else {
		prefixed = binary.BigEndian.AppendUint32(prefixed, singlePrefix)
	}

	prefixed = append(prefixed, txBlob...)

	if multiSig {
		prefixed = append(prefixed, accountID...)
	}

	return prefixed
}

// HashSigned computes a hash of a signed transaction
func HashSigned(txBlob []byte) []byte {
	prefixed := make([]byte, 0, len(txBlob)+4)
	prefixed = binary.BigEndian.AppendUint32(prefixed, signedPrefix)
	prefixed = append(prefixed, txBlob...)

	return Sha512Half(prefixed)
}

// Sha512Half returns the first half of sha512 hash
func Sha512Half(blob []byte) []byte {
	hash := sha512.Sum512(blob)
	return hash[:32]
}
