package utils

import (
	"encoding/binary"
)

const (
	singlePrefix uint32 = 0x53545800
	multiPrefix  uint32 = 0x534D5400
)

// Prepare creates a tx message for signing.
// If multiSig is true, txBlob is prefixed with multi-signing prefix and postfixed with accountID.
// For multi-signing, accountID of the signer should be provided.
// If multiSig is false, txBlob is prefixed with single-signing prefix.
func Prepare(txBlob []byte, multiSig bool, accountID []byte) []byte {
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
