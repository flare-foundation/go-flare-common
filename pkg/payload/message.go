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

type SubprotocolResponse struct {
	Status         ResponseStatus `json:"status"`
	Data           string         `json:"data"`
	AdditionalData string         `json:"additionalData"`
}

// BuildMessage builds a submit message string from protocolID, votingRoundID and payload.
// The message string is in the format: 0x<protocolID(1 byte)><votingRoundID(4 byte)><payloadLength(2 byte)><payload>
func BuildMessage(protocolID uint8, votingRoundID uint32, payload []byte) string {
	message := make([]byte, 7)
	message[0] = protocolID

	binary.BigEndian.PutUint32(message[1:5], votingRoundID)
	binary.BigEndian.PutUint16(message[5:7], uint16(len(payload)))

	message = append(message, payload...)

	return "0x" + hex.EncodeToString(message)
}

// buildMessageForSigning builds payload message for submitSignatures.
//
// protocolID (1 byte)
// roundID (4 bytes)
// randomQualityScore (1 byte)
// merkleRoot (32 bytes)
func BuildMessageForSigning(protocolID uint8, roundID uint32, isSecureRandom bool, merkleRoot common.Hash) string {
	data := make([]byte, 38)

	data[0] = uint8(protocolID)
	binary.BigEndian.PutUint32(data[1:5], uint32(roundID))

	if isSecureRandom {
		data[5] = 1
	} else {
		data[5] = 0
	}
	copy(data[6:38], merkleRoot[:])

	return "0x" + hex.EncodeToString(data)
}
