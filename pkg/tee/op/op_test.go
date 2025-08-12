package op

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOPTypes(t *testing.T) {
	opTypesSystem := []Type{
		"F_REG",
		"F_GET",
	}

	for _, ot := range opTypesSystem {
		otfh := HashToOPType(ot.Hash())

		require.Equal(t, ot, otfh)

		require.True(t, ot.isF())
		require.True(t, ot.IsSystem())
	}

	opTypesNonSystem := []Type{
		"reg",
		"GET",
		"f_GET",
	}

	for _, ot := range opTypesNonSystem {
		otfh := HashToOPType(ot.Hash())

		require.Equal(t, ot, otfh)

		require.False(t, ot.isF())
		require.False(t, ot.IsSystem())
	}

	opTypesNonSystemF := []Type{
		"F_reg",
		"F_",
	}

	for _, ot := range opTypesNonSystemF {
		otfh := HashToOPType(ot.Hash())

		require.Equal(t, ot, otfh)

		require.True(t, ot.isF())
		require.False(t, ot.IsSystem())
	}
}

func TestOPCommand(t *testing.T) {
	opCommands := []Command{
		"X",
		"x",
		"12",
	}

	for _, ot := range opCommands {
		ocfh := HashToOPCommand(ot.Hash())

		require.Equal(t, ot, ocfh)
	}
}

func TestValidPair(t *testing.T) {
	validPairs := []struct {
		t Type
		c Command
	}{
		{"x", "x"},
		{"F_FTDC", "PROVE"},
		{"F", "123"},
	}

	for _, p := range validPairs {
		require.True(t, IsValid(p.t, p.c))
		require.True(t, IsValidPair(p.t.Hash(), p.c.Hash()))
	}

	invalidPairs := []struct {
		t Type
		c Command
	}{
		{"F_FTDC", "x"},
		{"F_X", "PROVE"},
		{"F_", "123"},
	}
	for _, p := range invalidPairs {
		require.False(t, IsValid(p.t, p.c))
		require.False(t, IsValidPair(p.t.Hash(), p.c.Hash()))
	}
}
