package encoding

import (
	"encoding/binary"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/hash"
)

const (
	singlePrefix uint32 = 0x53545800
	multiPrefix  uint32 = 0x534D5400
	signedPrefix uint32 = 0x54584E00
)

// MessageToSign creates a tx message for signing.
// If multiSig is true, txBlob is prefixed with multi-signing prefix and postfixed with accountID. For multi-signing, accountID of the signer should be provided.
// If multiSig is false, txBlob is prefixed with single-signing prefix
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

	return hash.Sha512Half(prefixed)
}
