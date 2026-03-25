package secp256k1

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
)

// Signature holds the r and s components of a secp256k1 signature.
type Signature struct {
	r [32]byte
	s [32]byte
}

var errInvalidLength = errors.New("invalid signature length")

// MarshalDER marshals DER formatted signature.
//
// The signature should be in the following format:
//   - 1 byte - 0x30 for compound structure
//   - 1 byte - length or the remaining
//   - 1 byte - 0x02 for integer
//   - 1 byte - x = len of the r
//   - x bytes - r
//   - 1 byte - 0x02 for integer
//   - 1 byte - y = len of the s
//   - y bytes - s
func MarshalDER(sig []byte) (Signature, error) {
	var sigOut Signature
	if len(sig) < 4 {
		return sigOut, errInvalidLength
	}

	if sig[0] != 0x30 {
		return sigOut, fmt.Errorf("invalid start byte should be 0x30 is %x", sig[0])
	}

	if int(sig[1]) != (len(sig) - 2) {
		return sigOut, fmt.Errorf("invalid length. Indicate %d is %d", sig[1]+2, len(sig))
	}

	if sig[2] != 0x02 {
		return sigOut, fmt.Errorf("invalid type byte at index %d should be 0x02 is %x", 2, sig[2])
	}

	rLen := int(sig[3])
	rStart := 4
	rEnd := rStart + rLen
	if len(sig) < rEnd+2 {
		return sigOut, errInvalidLength
	}

	if sig[rStart] == 0 {
		rStart++
		rLen--
	}

	copy(sigOut.r[32-rLen:], sig[rStart:rEnd])

	if sig[rEnd] != 0x02 {
		return sigOut, fmt.Errorf("invalid type byte at index %d should be 0x02 is %x", rEnd, sig[rEnd])
	}

	sLen := int(sig[rEnd+1])
	sStart := rEnd + 2
	sEnd := sStart + sLen
	if len(sig) < sEnd {
		return sigOut, errInvalidLength
	}

	if sig[sStart] == 0 {
		sStart++
		sLen--
	}

	copy(sigOut.s[32-sLen:], sig[sStart:sEnd])

	return sigOut, nil
}

// MarshalRS marshals 64 bytes long sig into Signature.
func MarshalRS(sig []byte) (*Signature, error) {
	sigOut := new(Signature)
	if len(sig) != 64 {
		return nil, errInvalidLength
	}

	copy(sigOut.r[:], sig[:32])
	copy(sigOut.s[:], sig[32:64])

	return sigOut, nil
}

// DER returns DER formatted signature.
func (sig *Signature) DER() []byte {
	r := new(big.Int).SetBytes(sig.r[:])
	s := new(big.Int).SetBytes(sig.s[:])

	out := bytes.NewBuffer(nil)

	rb := r.Bytes()
	sb := s.Bytes()

	out.WriteByte(0x30)

	out.WriteByte(0) // to be amended at the end

	out.WriteByte(0x02)

	if rb[0] > 0x7f {
		out.WriteByte(1 + uint8(len(rb)))
		out.WriteByte(0)
	} else {
		out.WriteByte(uint8(len(rb)))
	}
	out.Write(rb)

	out.WriteByte(0x02)
	if sb[0] > 0x7f {
		out.WriteByte(1 + uint8(len(sb)))
		out.WriteByte(0)
	} else {
		out.WriteByte(uint8(len(sb)))
	}
	out.Write(sb)

	o := out.Bytes()

	o[1] = byte(out.Len() - 2)

	return o
}

// RS returns [r||s].
func (sig *Signature) RS() []byte {
	out := make([]byte, 64)

	copy(out[:32], sig.r[:])
	copy(out[32:], sig.s[:])

	return out
}

func (sig *Signature) Verify(hash []byte, pub *ecdsa.PublicKey) bool {
	return crypto.VerifySignature(toBytesCompressed(pub), hash, sig.RS())
}

// Validate checks that sig (in DER format) of Sha512Half hash of msg corresponds to (compressed) public key.
func Validate(msg, sig []byte, pub string) (bool, error) {
	sigParse, err := MarshalDER(sig)
	if err != nil {
		return false, fmt.Errorf("marshaling signature %v", err)
	}

	pubBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false, fmt.Errorf("decoding pub key%v", err)
	}

	ok := sigParse.VerifyHex(hash.Sha512Half(msg), pubBytes)

	return ok, nil
}

func (sig *Signature) VerifyHex(hash []byte, pub []byte) bool {
	return crypto.VerifySignature(pub, hash, sig.RS())
}

// SignatureWithRecovery contains r,s,v.
type SignatureWithRecovery struct {
	Signature
	v byte // recovery ID
}

// MarshalRecID marshals 65 bytes long sig into SignatureWithRecovery.
func MarshalRecID(sig []byte) (*SignatureWithRecovery, error) {
	sigOut := new(SignatureWithRecovery)
	if len(sig) != 65 {
		return nil, errInvalidLength
	}

	copy(sigOut.r[:], sig[:32])
	copy(sigOut.s[:], sig[32:64])

	sigOut.v = sig[64]

	return sigOut, nil
}

// RSV returns [r||s||v].
func (sig *SignatureWithRecovery) RSV() []byte {
	out := make([]byte, 65)

	copy(out[:32], sig.r[:])
	copy(out[32:], sig.s[:])

	out[64] = sig.v

	return out
}

// Recover recovers signer's public key form signature and hash (has to be 32 bytes long).
func (sig *SignatureWithRecovery) Recover(hash []byte) (*ecdsa.PublicKey, error) {
	return crypto.SigToPub(hash, sig.RSV())
}

// toBytesCompressed returns compressed public Key for ECDSA public key in byte slice.
func toBytesCompressed(pub *ecdsa.PublicKey) []byte {
	xrpPub := make([]byte, 33)

	if pub.Y.Bit(0) == 0 {
		xrpPub[0] = pubKeyEvenPrefix
	} else {
		xrpPub[0] = pubKeyOddPrefix
	}

	b := pub.X.Bytes()

	copy(xrpPub[33-len(b):33], b)

	return xrpPub
}
