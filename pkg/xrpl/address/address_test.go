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

func TestPubToAddress(t *testing.T) {
	tests := []struct {
		pub string
		add string
	}{
		{
			pub: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
			add: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		},
		{
			pub: "ED457192C53CC901123BB8C6E75C8636979BD88D46C32F6D915F89D66CE8862635",
			add: "rMuZNV2kjCKs8v8rd8QFizAaPdvCDYTPc7",
		},
		{
			pub: "032315B578DC026DD182F2FC89723E4E43AF812D7D3C86635B80D3608FBC76C224",
			add: "rf27pmNLFaXFwQfZPbTKe2Z665e1V76oC5",
		},
		{
			pub: "02707A7AE05A8DACDB89CC93429949CDA26F68200D9CE8753D4DCB04D6F80CFCB7",
			add: "rN5N6fJbc8xyViPDeQFMQMpYfVHuxSGV2G",
		},
		{
			pub: "035DB05B1CEA82785FB8B3F7E68E5C7429A1B00BE47CC6B1A651AA12C7B8D9592C",
			add: "rJQesZZEQzW9J3Eb1X1Snc7E6YGk7kTMoK",
		},
		{
			pub: "030C141E3E131B1D25B7EA85B10283E315F77F27A2FA085A45D65D47393DB9219F",
			add: "r9cvJhquqeExszdWZSw2rrFP98fsVFLdPe",
		},
	}

	for _, test := range tests {
		addFromPub, err := PubToAddress(test.pub)
		require.NoError(t, err)

		require.Equal(t, test.add, addFromPub)
	}
}
