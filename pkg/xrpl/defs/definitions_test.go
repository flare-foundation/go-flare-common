package defs

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDDecode(t *testing.T) {
	for name := range NameToField {

		field := NameToField[name]

		if !field.IsSerialized {
			continue
		}

		id, err := field.ID()
		require.NoError(t, err, field)

		b := bytes.NewBuffer(id)

		pair, err := ReadID(b)
		require.NoError(t, err, field)

		require.Equal(t, field.Nth, pair.F, field)
		require.Equal(t, field.Type, pair.T, field)

		require.Equal(t, 0, b.Len())
	}
}

func TestIDtoName(t *testing.T) {
	for key, name := range IDToName {

		field := NameToField[name]
		require.Equal(t, key.F, field.Nth)
		require.Equal(t, key.T, field.Type)
	}

	for name, field := range NameToField {
		id := IDPair{
			F: field.Nth,
			T: field.Type,
		}

		nameFrom := IDToName[id]
		require.Equal(t, name, nameFrom)
	}
}
