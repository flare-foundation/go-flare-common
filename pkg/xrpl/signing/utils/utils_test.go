package utils

import (
	"bytes"
	"encoding/hex"
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

// rippled include/xrpl/protocol/HashPrefix.h: txSign = 'S','T','X',0 ; txMultiSign = 'S','M','T',0.
func TestHashPrefixConstants(t *testing.T) {
	require.Equal(t, uint32(0x53545800), singlePrefix)
	require.Equal(t, uint32(0x534D5400), multiPrefix)
}

// rippled src/libxrpl/protocol/Sign.cpp buildMultiSigningData and Sign.h finishMultiSigningData:
// single-sign layout is [txSign prefix][tx blob]; multi-sign layout is [txMultiSign prefix][tx blob][accountID].
func TestPrepareExactBytes(t *testing.T) {
	blob := []byte{0x01, 0x02, 0x03, 0x04}
	accountID := bytes.Repeat([]byte{0xAA}, accountIDLength)

	singleWant, err := hex.DecodeString("5354580001020304")
	require.NoError(t, err)

	gotSingle, err := Prepare(blob, false, nil)
	require.NoError(t, err)
	require.Equal(t, singleWant, gotSingle)

	multiWant, err := hex.DecodeString("534D5400" + "01020304" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	require.NoError(t, err)

	gotMulti, err := Prepare(blob, true, accountID)
	require.NoError(t, err)
	require.Equal(t, multiWant, gotMulti)
}
