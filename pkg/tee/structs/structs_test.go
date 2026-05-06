package structs

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/fdc2"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/machine"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payments"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/tee"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/wallet"
	"github.com/stretchr/testify/require"
)

// This test assures that all inits in structs directory are called and do not panic.
func TestInit(t *testing.T) {
	require.NotNil(t, payments.MessageArguments)
	require.NotNil(t, wallet.MessageArguments)
	require.NotNil(t, fdc2.MessageArguments)
	require.NotNil(t, fdc2.AttestationTypeArguments)
	require.NotNil(t, tee.StructArg)
	require.NotNil(t, machine.TeeMachineDataStructArg)
}
