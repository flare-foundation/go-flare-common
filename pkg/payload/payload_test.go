package payload_test

import (
	"encoding/hex"
	"strings"

	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/payload"
	"github.com/stretchr/testify/require"
)

var tx = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac66",
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

var txMultiple = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac6665000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac66",
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

func TestExtractPayloads(t *testing.T) {
	tests := []struct {
		tx           *database.Transaction
		protocol     uint8
		nuOfPayloads int
		votingRound  uint32
		length       uint16
	}{
		{
			tx:           tx,
			protocol:     100,
			nuOfPayloads: 1,
			votingRound:  655419,
			length:       32,
		},
		{
			tx:           txMultiple,
			protocol:     101,
			nuOfPayloads: 2,
			votingRound:  655419,
			length:       32,
		},
	}

	for i, test := range tests {
		payloads, err := payload.ExtractPayloads(test.tx)
		require.NoError(t, err, fmt.Sprintf("error in test %d", i))
		require.Equal(t, test.nuOfPayloads, len(payloads), fmt.Sprintf("wrong number of payloads in test %d", i))

		payloadFTSO, ok := payloads[test.protocol]
		require.True(t, ok, fmt.Sprintf("missing payload in test %d", i))
		require.Equal(t, test.protocol, payloadFTSO.ProtocolID, fmt.Sprintf("wrong protocol ID in test %d", i))
		require.Equal(t, test.votingRound, payloadFTSO.VotingRound, fmt.Sprintf("wrong voting round in test %d", i))
	}
}

var txError1 = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac660", //too long
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

var txError2 = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac6", //too short
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

var txError3 = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac664000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac6", //too short
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

var txError4 = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b002043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac6z", //illegal character
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

var txError5 = &database.Transaction{
	Hash:             "8dd67e88aa6f863aeb5cd62874530efd7dafef2d4a8cdf7fbf71844dab1c7327",
	FunctionSig:      "6c532fae",
	Input:            "6c532fae64000a003b001043a94d3c612d7f5cfd65e53a06d55bac77abbd2a6eb4dff766f51092db36ac66", //wrong length
	BlockNumber:      16143116,
	BlockHash:        "40888ee23c4d7da30c42f826ea187386eac4564b02ce801f0b0b91ef1e71da05",
	TransactionIndex: 0,
	FromAddress:      "6bba3b6fb0dc902845666fdad70b2a814a57b6f3",
	ToAddress:        "2ca6571daa15ce734bbd0bf27d5c9d16787fc33f",
	Status:           1,
	Value:            "0",
	GasPrice:         "37500000000",
	Gas:              2500000,
	Timestamp:        1717417740,
}

func TestExtractPayloadsError(t *testing.T) {
	txs := []*database.Transaction{
		txError1,
		txError2,
		txError3,
		txError4,
		txError5,
	}

	for i, tx := range txs {
		_, err := payload.ExtractPayloads(tx)
		require.Error(t, err, fmt.Sprintf("error in test %d", i))
	}
}

func TestBuildMessage(t *testing.T) {
	tests := []struct {
		protocolID  uint8
		votingRound uint32
		payload     string
		result      string
	}{
		{
			protocolID:  1,
			votingRound: 1,
			payload:     "00",
			result:      "0x0100000001000100",
		},
		{
			protocolID:  255,
			votingRound: 256,
			payload:     "110011",
			result:      "0xff000001000003110011",
		},
	}

	for _, test := range tests {
		payloadBytes, err := hex.DecodeString(test.payload)
		require.NoError(t, err)

		payloadMsg := payload.BuildMessage(test.protocolID, test.votingRound, payloadBytes)
		require.Equal(t, test.result, payloadMsg)
	}
}

func TestBuildMessageForSigning(t *testing.T) {
	tests := []struct {
		protocolID     uint8
		roundID        uint32
		isSecureRandom bool
		hash           common.Hash
		expected       string
	}{
		{
			protocolID:     0,
			roundID:        0,
			isSecureRandom: false,
			hash:           [32]byte{},
			expected:       "0x" + strings.Repeat("0", 38*2),
		},
		{
			protocolID:     200,
			roundID:        12345,
			isSecureRandom: true,
			hash:           [32]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			expected:       "0x" + "c8" + "00003039" + "01" + "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f",
		},
	}

	for i, test := range tests {
		result := payload.BuildMessageForSigning(test.protocolID, test.roundID, test.isSecureRandom, test.hash)
		require.Equal(t, test.expected, result, fmt.Sprintf("Error in test %v", i))
	}
}
