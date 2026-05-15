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
		// Index is encoded as uint16 below. Without this guard, an Index >
		// math.MaxUint16 passes the monotonicity check (compared as int)
		// and is then truncated to its low 16 bits — silently producing a
		// non-monotonic, possibly-duplicate on-the-wire index sequence.
		if signature.Index > math.MaxUint16 {
			return nil, fmt.Errorf("payload index %d exceeds uint16 range", signature.Index)
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
// vrs must be 65 bytes and vrs[0] (the V byte) must be 27 or 28 — the only
// Ethereum-style values that map to a valid secp256k1 recovery id (0 or 1).
// An already-normalised V (0 or 1) underflows on subtraction; EIP-155
// values (chainID*2 + 35 + {0,1}) and garbage bytes produce out-of-range
// recids that downstream ecrecover would reject. Both are caught up front.
func TransformSignatureVRStoRSV(vrs []byte) ([]byte, error) {
	if len(vrs) != 65 {
		return nil, fmt.Errorf("signature must be 65 bytes, got %d", len(vrs))
	}
	if vrs[0] < 27 || vrs[0] > 28 {
		return nil, fmt.Errorf("invalid V byte %d, expected 27 or 28", vrs[0])
	}

	rsv := make([]byte, 65)
	copy(rsv, vrs[1:33])
	copy(rsv[32:], vrs[33:65])
	rsv[64] = vrs[0] - 27

	return rsv, nil
}

// TransformSignatureRSVtoVRS transforms [R || S || V - 27] to [V || R || S].
// rsv must be 65 bytes and rsv[64] (the recid) must be 0 or 1 — the canonical
// secp256k1 recovery id values. Any other value would produce an out-of-spec
// V byte (29-255) or, at the far edge, wrap to a low byte (e.g. recid 229 →
// V 0) that masks the caller's input bug behind a "successful" transform.
func TransformSignatureRSVtoVRS(rsv []byte) ([]byte, error) {
	if len(rsv) != 65 {
		return nil, fmt.Errorf("signature must be 65 bytes, got %d", len(rsv))
	}
	if rsv[64] > 1 {
		return nil, fmt.Errorf("invalid recid %d, expected 0 or 1", rsv[64])
	}

	vrs := make([]byte, 65)
	vrs[0] = rsv[64] + 27
	copy(vrs[1:], rsv[0:32])
	copy(vrs[33:], rsv[32:64])

	return vrs, nil
}
