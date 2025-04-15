package types

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObjectOrder(t *testing.T) {
	json1 := `{
    "Account":       "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd",
    "TxnSignature":  "304402204427efa8e33c9907391548bd1748356477e93824458956bec9564aa9dcf9a29f0220686717a7af7c3fb37ba7ecd43c299c98fec9f412db76de3d795add38a059954c",
    "SigningPubKey": "038940a036ee380369b9fcc5929a0d431abe789c8a44e5f48f564e2f6eb608b543"
    }`

	json2 := `{
    "TxnSignature":  "304402204427efa8e33c9907391548bd1748356477e93824458956bec9564aa9dcf9a29f0220686717a7af7c3fb37ba7ecd43c299c98fec9f412db76de3d795add38a059954c",
    "Account":       "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd",
    "SigningPubKey": "038940a036ee380369b9fcc5929a0d431abe789c8a44e5f48f564e2f6eb608b543"
    }`

	var obj1, obj2 any

	err := json.Unmarshal([]byte(json1), &obj1)
	require.NoError(t, err)

	err = json.Unmarshal([]byte(json2), &obj2)
	require.NoError(t, err)

	enc1, err := (&STObject{}).ToBytes(obj1, false)
	require.NoError(t, err)

	enc2, err := (&STObject{}).ToBytes(obj2, false)
	require.NoError(t, err)

	require.Equal(t, enc1, enc2)
}

func TestObjectEncodeDecode(t *testing.T) {
	maps := []map[string]any{
		{
			"Account":       "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd",
			"TxnSignature":  "304402204427efa8e33c9907391548bd1748356477e93824458956bec9564aa9dcf9a29f0220686717a7af7c3fb37ba7ecd43c299c98fec9f412db76de3d795add38a059954c",
			"SigningPubKey": "038940a036ee380369b9fcc5929a0d431abe789c8a44e5f48f564e2f6eb608b543",
		},
		{
			"Fee":           "10",
			"Flags":         uint32(524288),
			"OfferSequence": uint32(1752791),
			"TakerGets":     "150000000000",
		}}

	for j, mapIn := range maps {
		serialized, err := (&STObject{}).ToBytes(mapIn, false)
		require.NoError(t, err, j)

		b := bytes.NewBuffer(serialized)

		deserialized, err := (&STObject{}).ToJson(b, 0)
		require.NoError(t, err, j)

		mapOut, ok := deserialized.(map[string]any)
		require.True(t, ok, j)

		require.Equal(t, mapIn, mapOut, j)

		require.Equal(t, 0, b.Len())
	}
}
