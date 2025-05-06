package address

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountID(t *testing.T) {
	addr := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S"

	id, err := ID(addr)
	require.NoError(t, err)

	addrBack, err := Address(id)
	require.NoError(t, err)

	require.Equal(t, addr, addrBack)
}
