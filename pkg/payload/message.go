// Package payload provides utilities for building and extracting Flare protocol submission messages.
package payload

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

type ResponseStatus string

const (
	Ok    ResponseStatus = "OK"
	Empty ResponseStatus = "EMPTY"
	Retry ResponseStatus = "RETRY"
)

const (
	msgHeaderLength             = 7  // protocolID 1, votingRoundID 4, payloadLength 2
	submitSignatureHeaderLength = 6  // protocolID 1, roundID 4, randomQualityScore 1
	submitSignatureMsgLength    = 38 // protocolID 1, roundID 4, randomQualityScore 1, merkleRoot 32
)

type SubprotocolResponse struct {
	Status         ResponseStatus `json:"status"`
	Data           string         `json:"data"`
	AdditionalData string         `json:"additionalData"`
}

// BuildMessage builds a submit message string from protocolID, votingRoundID and payload.
// The message string is in the format: 0x<protocolID(1 byte)><votingRoundID(4 byte)><payloadLength(2 byte)><payload>.
func BuildMessage(protocolID uint8, votingRoundID uint32, payload []byte) string {
	message := make([]byte, msgHeaderLength, msgHeaderLength+len(payload))

	message[0] = protocolID

	binary.BigEndian.PutUint32(message[1:5], votingRoundID)
	binary.BigEndian.PutUint16(message[5:7], uint16(len(payload)))

	message = append(message, payload...)

	return "0x" + hex.EncodeToString(message)
}

// BuildMessageForSigning builds payload message for submitSignatures.
//
//   - protocolID (1 byte),
//   - roundID (4 bytes),
//   - randomQualityScore (1 byte),
//   - merkleRoot (32 bytes).
func BuildMessageForSigning(protocolID uint8, roundID uint32, isSecureRandom bool, merkleRoot common.Hash) string {
	data := make([]byte, submitSignatureHeaderLength, submitSignatureMsgLength)

	data[0] = protocolID
	binary.BigEndian.PutUint32(data[1:5], roundID)

	if isSecureRandom {
		data[5] = 1
	} else {
		data[5] = 0
	}

	data = append(data, merkleRoot[:]...)

	return "0x" + hex.EncodeToString(data)
}
