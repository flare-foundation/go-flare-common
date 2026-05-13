package payload_test

import (
	"bytes"
	"encoding/binary"
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

// TestExtractPayloadsDeclaredLength varies the on-wire length field
// independently of the bytes that follow.
func TestExtractPayloadsDeclaredLength(t *testing.T) {
	const selector = "6c532fae"

	frame := func(proto uint8, round uint32, declaredLen uint16, payload []byte) []byte {
		var buf bytes.Buffer
		buf.WriteByte(proto)
		_ = binary.Write(&buf, binary.BigEndian, round)
		_ = binary.Write(&buf, binary.BigEndian, declaredLen)
		buf.Write(payload)
		return buf.Bytes()
	}

	mkInput := func(frames ...[]byte) string {
		out := selector
		for _, f := range frames {
			out += hex.EncodeToString(f)
		}
		return out
	}

	p32 := bytes.Repeat([]byte{0xaa}, 32)
	pMax := bytes.Repeat([]byte{0xbb}, 0xFFFF)

	tests := []struct {
		name    string
		input   string
		wantErr bool
		wantLen int
		wantPID map[uint8]int
	}{
		{
			name:    "declared matches payload",
			input:   mkInput(frame(100, 1, 32, p32)),
			wantLen: 1,
			wantPID: map[uint8]int{100: 32},
		},
		{
			name:    "declared zero with no trailing bytes",
			input:   mkInput(frame(50, 7, 0, nil)),
			wantLen: 1,
			wantPID: map[uint8]int{50: 0},
		},
		{
			name:    "two frames, both declared correctly",
			input:   mkInput(frame(100, 1, 1, []byte{0x11}), frame(101, 2, 1, []byte{0x22})),
			wantLen: 2,
			wantPID: map[uint8]int{100: 1, 101: 1},
		},
		{
			name:    "declared exceeds remaining",
			input:   mkInput(frame(100, 1, 64, p32)),
			wantErr: true,
		},
		{
			// cb61c39 regression: 7 + 0xFFFF overflowed uint16 to 6.
			name:    "declared max uint16 with insufficient data",
			input:   mkInput(frame(100, 1, 0xFFFF, p32)),
			wantErr: true,
		},
		{
			name:    "declared 0xFFFE with insufficient data",
			input:   mkInput(frame(100, 1, 0xFFFE, p32)),
			wantErr: true,
		},
		{
			// Positive side of cb61c39: 65535 bytes must round-trip.
			name:    "declared max uint16 with exact data",
			input:   mkInput(frame(100, 1, 0xFFFF, pMax)),
			wantLen: 1,
			wantPID: map[uint8]int{100: 0xFFFF},
		},
		{
			// Under-claim: tail re-parsed as a new frame, decodes garbage -> error.
			name:    "declared under-claims, tail mis-frames into error",
			input:   mkInput(frame(100, 1, 7, p32)),
			wantErr: true,
		},
		{
			// Declared 0 with trailing bytes: parser re-enters and mis-frames the tail.
			name:    "declared zero with trailing bytes mis-frames",
			input:   mkInput(frame(100, 1, 0, p32)),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &database.Transaction{Hash: "h", FunctionSig: "f", Input: tt.input}
			msgs, err := payload.ExtractPayloads(tx)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantLen, len(msgs))
			for pid, n := range tt.wantPID {
				m, ok := msgs[pid]
				require.True(t, ok, "missing protocolID %d", pid)
				require.Equal(t, n, len(m.Payload), "payload length for protocolID %d", pid)
			}
		})
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
