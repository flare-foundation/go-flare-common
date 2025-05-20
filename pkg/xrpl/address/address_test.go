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

func TestAccountIDFail(t *testing.T) {
	failAddrLengthAlphabet := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1l"
	_, err := ID(failAddrLengthAlphabet)
	require.Error(t, err)

	failAddrLength := "rKTyuZsNuQbL"
	_, err = ID(failAddrLength)
	require.Error(t, err)

	failAddrLeading := "hhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1S"
	_, err = ID(failAddrLeading)
	require.Error(t, err)

	failAddrChecksum := "rhbQ2PoSsmzh5XnmWnutCa6dmAC2qj1z1h"
	_, err = ID(failAddrChecksum)
	require.Error(t, err)

	invalidID0 := []byte{}
	invalidID1 := make([]byte, 21)

	_, err = Address(invalidID0)
	require.Error(t, err)

	_, err = Address(invalidID1)
	require.Error(t, err)
}
