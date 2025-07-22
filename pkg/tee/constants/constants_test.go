package constants

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOPTypes(t *testing.T) {
	opTypesOK := []string{
		"REG",
	}

	for _, opType := range opTypesOK {
		ot, ok := StringToOPType(opType)
		require.True(t, ok)

		ok = ot.IsValid()
		require.True(t, ok)

		h := ot.Hash()

		ot2, ok := HashToOPType(h)
		require.True(t, ok)

		require.Equal(t, ot, ot2)
		require.Equal(t, opType, string(ot))
	}

	opTypesFail := []string{
		"reg",
	}

	for _, opType := range opTypesFail {
		ot, ok := StringToOPType(opType)
		require.False(t, ok)

		ok = ot.IsValid()
		require.False(t, ok)

		h := ot.Hash()

		ot2, ok := HashToOPType(h)
		require.False(t, ok)

		require.Equal(t, ot, ot2)
		require.Equal(t, opType, string(ot))
	}
}

func TestOPTCommands(t *testing.T) {
	opCommandsOK := []string{
		"PAY",
	}

	for _, opCommands := range opCommandsOK {
		ot, ok := StringToOPCommand(opCommands)
		require.True(t, ok)

		ok = ot.IsValid()
		require.True(t, ok)

		h := ot.Hash()

		ot2, ok := HashToOPCommand(h)
		require.True(t, ok)

		require.Equal(t, ot, ot2)
		require.Equal(t, opCommands, string(ot))
	}

	opCommandsFail := []string{
		"pay",
	}

	for _, opCommands := range opCommandsFail {
		ot, ok := StringToOPCommand(opCommands)
		require.False(t, ok)

		ok = ot.IsValid()
		require.False(t, ok)

		h := ot.Hash()

		ot2, ok := HashToOPCommand(h)
		require.False(t, ok)

		require.Equal(t, ot, ot2)
		require.Equal(t, opCommands, string(ot))
	}
}

func TestValidPair(t *testing.T) {
	ok := IsValidPair(XRP, Reissue)
	require.True(t, ok)

	ok = IsValid(XRP.Hash(), Reissue.Hash())
	require.True(t, ok)

	ok = IsValidPair(Reg, KeyGenerate)
	require.False(t, ok)

	ok = IsValid(Reg.Hash(), KeyGenerate.Hash())
	require.False(t, ok)
}
