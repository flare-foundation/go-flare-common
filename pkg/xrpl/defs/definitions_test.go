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

		fCode, tCode, err := IDDecode(b)
		require.NoError(t, err, field)

		require.Equal(t, field.Nth, fCode, field)
		require.Equal(t, field.Type, tCode, field)

		require.Equal(t, 0, b.Len())
	}
}
