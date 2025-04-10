package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayFail(t *testing.T) {
	tests := []string{
		`[{"Memo": {
                    "MemoType": "584D4D2076616C7565",
                    "MemoData": "322E3230393635"
                }},
                {"TakerPays": {
                    "currency": "BTC",
                    "value": "0.01262042643559221",
                    "issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
                }}
                    ]`,
		`[{"Memo": {
                    "MemoType": "584D4D2076616C7565",
                    "MemoData": "322E3230393635"
                },
        "TakerPays": {
                    "currency": "BTC",
                    "value": "0.01262042643559221",
                    "issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
                }} ]`,
		`[1,2,3]`,
		`123`,
	}

	for j := range tests {

		var value any

		err := json.Unmarshal([]byte(tests[j]), &value)
		require.NoError(t, err)

		_, err = STArray.ToBytes(value, false)
		require.Error(t, err)
	}

}
