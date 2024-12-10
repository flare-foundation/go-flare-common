package payload

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/database"
)

type Round struct {
	Messages []Message
	ID       uint32
}

type Message struct {
	From             common.Address
	Selector         string // function selector
	ProtocolID       uint8
	VotingRound      uint32
	Timestamp        uint64
	BlockNumber      uint64
	TransactionIndex uint64
	Payload          []byte
}

// ExtractPayloads extracts Payloads from transactions to submission contract to functions submit1, submit2, submitSignatures and builds a map protocolID -> payload message.
func ExtractPayloads(tx *database.Transaction) (map[uint8]Message, error) {
	messages := make(map[uint8]Message)

	dataStr := strings.TrimPrefix(tx.Input, "0x") //trim 0x if needed

	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return nil, fmt.Errorf("error decoding input of tx: %s, %v", tx.Hash, err)
	}

	data = data[4:] // trim function selector
	for len(data) > 0 {
		if len(data) < 7 { // 7 = 1 + 4 + 2
			return nil, errors.New("wrongly formatted tx input, too short")
		}

		protocol := data[0]                               // 1 byte protocol ID
		votingRound := binary.BigEndian.Uint32(data[1:5]) // 4 bytes votingRoundID
		length := binary.BigEndian.Uint16(data[5:7])      // 2 bytes length of payload in bytes

		end := 7 + length
		if len(data) < int(end) {
			return nil, errors.New("wrongly formatted tx input")
		}

		payload := data[7:end]

		message := Message{
			From:             common.HexToAddress(tx.FromAddress),
			Selector:         tx.FunctionSig,
			ProtocolID:       protocol,
			VotingRound:      votingRound,
			Timestamp:        tx.Timestamp,
			BlockNumber:      tx.BlockNumber,
			TransactionIndex: tx.TransactionIndex,
			Payload:          payload,
		}

		messages[protocol] = message

		data = data[end:] // trim the extracted payload
	}
	return messages, nil
}
