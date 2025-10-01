package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountIDFailToByte(t *testing.T) {
	inputs := []any{"", 123, "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAee", "QLbzfJH5BT1FS9apRLKV3G8dWEAjwnKaa", "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAeĆ"}

	for _, input := range inputs {
		_, err := AccountID.ToBytes(input, false)
		require.Error(t, err)
	}
}

func TestAccountIDFailToJSON(t *testing.T) {
	inputs := [][]byte{
		{},
		{1},
	}

	for _, input := range inputs {
		b := bytes.NewBuffer(input)
		_, err := AccountID.ToJSON(b, 0)
		require.Error(t, err)
	}
}

func TestAccountID(t *testing.T) {
	inAddresses := []any{"rrrrrrrrrrrrrrrrrrrrrhoLvTp", "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", "rrrrrrrrrrrrrrrrrrrn5RM1rHd", "rfmdBKhtJw2J22rw1JxQcchQTM68qzE4N2"}

	for _, add := range inAddresses {
		enc, err := AccountID.ToBytes(add, false)
		require.NoError(t, err)

		b := bytes.NewBuffer(enc)
		dec, err := AccountID.ToJSON(b, 0)
		require.NoError(t, err)

		require.Equal(t, add, dec)

		require.Equal(t, 0, b.Len())
	}
}
