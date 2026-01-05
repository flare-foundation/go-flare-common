package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrepare(t *testing.T) {
	txBlob := []byte{0x01, 0x02, 0x03, 0x04}

	t.Run("single-signing", func(t *testing.T) {
		prepared, err := Prepare(txBlob, false, nil)
		require.NoError(t, err)

		require.Len(t, prepared, len(txBlob)+prefixLength)
	})

	t.Run("multi-signing success", func(t *testing.T) {
		accountID := bytes.Repeat([]byte{0xAA}, accountIDLength)
		prepared, err := Prepare(txBlob, true, accountID)
		require.NoError(t, err)

		require.Len(t, prepared, len(txBlob)+prefixLength+accountIDLength)
	})

	t.Run("multi-signing invalid accountID", func(t *testing.T) {
		testCases := []struct {
			name      string
			accountID []byte
		}{
			{
				name:      "nil accountID",
				accountID: nil,
			},
			{
				name:      "short accountID",
				accountID: bytes.Repeat([]byte{0xBB}, accountIDLength-1),
			},
			{
				name:      "long accountID",
				accountID: bytes.Repeat([]byte{0xCC}, accountIDLength+1),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := Prepare(txBlob, true, tc.accountID)
				require.Error(t, err)
			})
		}
	})
}
