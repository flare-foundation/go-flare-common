package structs

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/connector"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payment"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/tee"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/wallet"
	"github.com/stretchr/testify/require"
)

// This test assures that all inits in structs directory are called and do not panic.
func TestInit(t *testing.T) {
	require.NotNil(t, payment.MessageArguments)
	// require.NotNil(t, registry.MessageArguments)
	require.NotNil(t, wallet.MessageArguments)
	require.NotNil(t, connector.MessageArguments)
	require.NotNil(t, connector.AttestationTypeArguments)
	require.NotNil(t, tee.StructArg)
}
