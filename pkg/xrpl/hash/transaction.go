package hash

import (
	"encoding/binary"
)

// signedPrefix is rippled's HashPrefix::transactionID, ASCII "TXN\0".
const signedPrefix uint32 = 0x54584E00

// Signed computes the canonical transaction ID of a signed transaction.
// txBlob must be the full serialization (Encode with signing=false, including
// TxnSignature/Signers); a for-signing blob yields a wrong hash without error.
func Signed(txBlob []byte) []byte {
	prefixed := make([]byte, 0, len(txBlob)+4)
	prefixed = binary.BigEndian.AppendUint32(prefixed, signedPrefix)
	prefixed = append(prefixed, txBlob...)

	return Sha512Half(prefixed)
}
