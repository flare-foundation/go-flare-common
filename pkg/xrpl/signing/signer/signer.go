package signer

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/base58"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

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

// Sort sorts signers according to the numeric value of the account (accountID).
// The signers with invalid addresses are discarded and returned in a separate slice.
func Sort(signers []*Signer) ([]*Signer, []*Signer) {
	in := make([]*Signer, 0, len(signers))
	out := make([]*Signer, 0)

	for _, s := range signers {
		_, err := s.Value()
		if err != nil {
			out = append(out, s)
		} else {
			in = append(in, s)
		}
	}

	slices.SortFunc(in, Compare)

	return in, out
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

func extractString(values map[string]any, key string) (string, error) {
	value, ok := values[key]
	if !ok {
		return "", fmt.Errorf("missing '%s'", key)
	}
	valueStr, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("value for key `%s` is not string: %v", key, value)
	}

	return valueStr, nil
}

// Parse builds Signer.
func Parse(arrayObject types.ArrayObject) (*Signer, error) {
	s, ok := arrayObject["Signer"]
	if !ok {
		return nil, fmt.Errorf("not a Signer %v", arrayObject)
	}

	sMap, ok := s.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("not a Signer item %v", s)
	}

	if len(sMap) != 3 {
		return nil, fmt.Errorf("invalid signer")
	}

	acc, err := extractString(sMap, "Account")
	if err != nil {
		return nil, err
	}
	txS, err := extractString(sMap, "TxnSignature")
	if err != nil {
		return nil, err
	}
	pub, err := extractString(sMap, "SigningPubKey")
	if err != nil {
		return nil, err
	}

	signer := &Signer{
		Account:       acc,
		TxnSignature:  txS,
		SigningPubKey: pub,
	}

	_, err = signer.Value()
	if err != nil {
		return nil, err
	}

	return signer, nil
}

// Compare compares values signers
//   - -1 if s1 is lesser than s2
//   - 0 if s1 is equal to s2
//   - 1 if s1 is greater than s2
//
// The function assumes:
//
//   - s1 and s2 are non nil.
//   - values of s1 and s2 are defined.
func Compare(s1, s2 *Signer) int {
	return s1.value.Cmp(s2.value)
}
