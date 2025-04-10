package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountIDFail(t *testing.T) {
	addresses := []any{"", 123, "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAee", "QLbzfJH5BT1FS9apRLKV3G8dWEAjwnKaa", "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAeĆ"}

	for _, address := range addresses {
		_, err := AccountID.ToBytes(address, false)
		require.Error(t, err)
	}
}
