package hash

import (
	"encoding/binary"
)

const signedPrefix uint32 = 0x54584E00

// HashSigned computes a hash of a signed transaction.
func HashSigned(txBlob []byte) []byte {
	prefixed := make([]byte, 0, len(txBlob)+4)
	prefixed = binary.BigEndian.AppendUint32(prefixed, signedPrefix)
	prefixed = append(prefixed, txBlob...)

	return Sha512Half(prefixed)
}
