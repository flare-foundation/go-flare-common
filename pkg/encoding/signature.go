package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
)

type IndexedSignature struct {
	Index     int
	Signature hexutil.Bytes
}

// EncodeSignatures encodes indexed signature to be used in the finalization transaction input.
// Signatures should be ordered by the indexes of their providers.
// Signatures should be 65 bytes long, otherwise the function will panic.
func EncodeSignatures(signatures []IndexedSignature) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if len(signatures) > math.MaxUint16 {
		return nil, fmt.Errorf("too many payloads: %d", len(signatures))
	}

	sizeBytes := convert.Uint16ToBytes(uint16(len(signatures)))
	buffer.Write(sizeBytes)
	prevIndex := -1
	for _, signature := range signatures {
		if signature.Index < 0 {
			return nil, errors.New("payload index not set")
		}
		if prevIndex >= signature.Index {
			return nil, errors.New("payloads not sorted by index")
		}

		sig := TransformSignatureRSVtoVRS(signature.Signature)
		buffer.Write(sig)

		indexBytes := convert.Uint16ToBytes(uint16(signature.Index))
		buffer.Write(indexBytes)

		prevIndex = signature.Index
	}

	return buffer.Bytes(), nil
}

// TransformSignatureVRStoRSV transforms [V || R || S] to [R || S || V - 27].
// No checks are performed, we assume that signature array has length 65.
func TransformSignatureVRStoRSV(vrs []byte) (rsv []byte) {
	rsv = make([]byte, 65)

	copy(rsv, vrs[1:33])
	copy(rsv[32:], vrs[33:65])
	rsv[64] = vrs[0] - 27

	return rsv
}

// TransformSignatureRSVtoVRS transforms [R || S || V - 27] to [V || R || S].
// No checks are performed, we assume that signature array has length 65.
func TransformSignatureRSVtoVRS(rsv []byte) (vrs []byte) {
	vrs = make([]byte, 65)

	vrs[0] = rsv[64] + 27
	copy(vrs[1:], rsv[0:32])
	copy(vrs[33:], rsv[32:64])

	return vrs
}
