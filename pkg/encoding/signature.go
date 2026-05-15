// Package encoding provides signature encoding and transformation utilities for Flare protocol finalization.
package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
)

// IndexedSignature pairs a signature with the index of its provider.
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

		sig, err := TransformSignatureRSVtoVRS(signature.Signature)
		if err != nil {
			return nil, fmt.Errorf("transforming signature at index %d: %w", signature.Index, err)
		}
		buffer.Write(sig)

		indexBytes := convert.Uint16ToBytes(uint16(signature.Index))
		buffer.Write(indexBytes)

		prevIndex = signature.Index
	}

	return buffer.Bytes(), nil
}

// TransformSignatureVRStoRSV transforms [V || R || S] to [R || S || V - 27].
// vrs must be 65 bytes and vrs[0] (the V byte) must be at least 27; an
// already-normalised V of 0 or 1 would underflow on subtraction and is
// rejected with an error.
func TransformSignatureVRStoRSV(vrs []byte) ([]byte, error) {
	if len(vrs) != 65 {
		return nil, fmt.Errorf("signature must be 65 bytes, got %d", len(vrs))
	}
	if vrs[0] < 27 {
		return nil, fmt.Errorf("invalid V byte %d, expected >= 27 (input may already be normalised)", vrs[0])
	}

	rsv := make([]byte, 65)
	copy(rsv, vrs[1:33])
	copy(rsv[32:], vrs[33:65])
	rsv[64] = vrs[0] - 27

	return rsv, nil
}

// TransformSignatureRSVtoVRS transforms [R || S || V - 27] to [V || R || S].
// rsv must be 65 bytes.
func TransformSignatureRSVtoVRS(rsv []byte) ([]byte, error) {
	if len(rsv) != 65 {
		return nil, fmt.Errorf("signature must be 65 bytes, got %d", len(rsv))
	}

	vrs := make([]byte, 65)
	vrs[0] = rsv[64] + 27
	copy(vrs[1:], rsv[0:32])
	copy(vrs[33:], rsv[32:64])

	return vrs, nil
}
