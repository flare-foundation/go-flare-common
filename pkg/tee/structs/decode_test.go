package structs

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/verification"
	"github.com/stretchr/testify/require"
)

func TestDecodeInstructionMessage(t *testing.T) {
	arg := verification.MessageArguments[op.TEEAttestation]

	id := common.HexToAddress("6e656b69")
	h := common.HexToHash("6e656b69")

	x := common.HexToHash("0x1234")

	pre := verification.IVerificationTeeAttestation{
		TeeMachine: verification.IMachineManagerTeeMachineWithAttestationData{
			TeeId:        id,
			InitialTeeId: id,
			Url:          "moj.com",
			CodeHash:     h,
			Platform:     h,
		},
		Challenge: x,
	}

	encoded, err := abi.Arguments{arg}.Pack(pre)
	require.NoError(t, err)

	var unpacked1 verification.IVerificationTeeAttestation

	err = DecodeTo(arg, encoded, &unpacked1)
	require.NoError(t, err)
	require.Equal(t, pre, unpacked1)

	unpacked2, err := Decode[verification.IVerificationTeeAttestation](arg, encoded)
	require.NoError(t, err)

	require.Equal(t, pre, unpacked2)
}
