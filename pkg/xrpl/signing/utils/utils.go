// Package utils provides utilities for preparing XRPL transaction messages for signing.
package utils

import (
	"encoding/binary"
	"errors"
)

const (
	singlePrefix uint32 = 0x53545800
	multiPrefix  uint32 = 0x534D5400

	prefixLength    = 4
	accountIDLength = 20
)

// Prepare creates a tx message for signing.
// If multiSig is true, txBlob is prefixed with multi-signing prefix and postfixed with accountID.
// For multi-signing, accountID of the signer should be provided. AccountID should be 20 bytes long
// If multiSig is false, txBlob is prefixed with single-signing prefix.
func Prepare(txBlob []byte, multiSig bool, accountID []byte) ([]byte, error) {
	length := prefixLength + len(txBlob)
	if multiSig {
		if len(accountID) != accountIDLength {
			return nil, errors.New("invalid accountID length")
		}
		length += accountIDLength
	}

	forSigning := make([]byte, 0, length)

	if multiSig {
		forSigning = binary.BigEndian.AppendUint32(forSigning, multiPrefix)
	} else {
		forSigning = binary.BigEndian.AppendUint32(forSigning, singlePrefix)
	}

	forSigning = append(forSigning, txBlob...)

	if multiSig {
		forSigning = append(forSigning, accountID...)
	}

	return forSigning, nil
}
