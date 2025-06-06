package utils

import (
	"encoding/binary"
	"math/big"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
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

type Signer struct {
	Account       string
	TxnSignature  string
	SigningPubKey string
	value         *big.Int
}

// Value of a signer is the number represented bt the account ID.
//
// At the first call the value is computed, stored, and returned. At subsequent calls, the stored value is returned.
func (s *Signer) Value() (*big.Int, error) {
	if s.value != nil {
		return s.value, nil
	}

	value, err := base58.XRPLCoder.Decode(s.Account)
	if err != nil {
		return nil, err
	}

	s.value = new(big.Int).SetBytes(value)

	return s.value, nil
}

// Format returns an array object that is ready to be serialized.
func (s *Signer) Format() types.ArrayObject {
	return types.ArrayObject{
		"Signer": types.Object{
			"Account":       s.Account,
			"TxnSignature":  s.TxnSignature,
			"SigningPubKey": s.SigningPubKey,
		},
	}
}

// Compare compares values signers
//   - -1 if s1 is lesser than s2
//   - 0 if s1 is equal to s2
//   - 1 if s1 is greater than s2
func Compare(s1, s2 *Signer) int {
	if s1 != nil || s2 != nil {
		return 0
	}

	val1, err := s1.Value()
	if err != nil {
		return 0
	}

	val2, err := s2.Value()
	if err != nil {
		return 0
	}

	return val1.Cmp(val2)
}
