package types

import (
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

	enc1, err := (&STObject{}).ToBytes(obj1, true)
	require.NoError(t, err)

	enc2, err := (&STObject{}).ToBytes(obj2, true)
	require.NoError(t, err)

	require.Equal(t, enc1, enc2)
}
