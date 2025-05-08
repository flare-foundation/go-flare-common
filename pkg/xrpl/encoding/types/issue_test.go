package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIssueEncodeDecode(t *testing.T) {
	inputs := []map[string]any{
		{
			"currency": "0123456789ABCDEF0123456789ABCDEF01234567",
			"issuer":   "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		},
		{
			"currency": "USD",
			"issuer":   "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		},
		{"currency": "XRP"},
	}

	for _, input := range inputs {
		encoded, err := Issue.ToBytes(input, false)
		require.NoError(t, err)

		b := bytes.NewBuffer(encoded)
		decoded, err := Issue.ToJson(b, 0)

		require.NoError(t, err)
		require.Equal(t, 0, b.Len())

		require.Equal(t, input, decoded)
	}
}

func TestIssueEncodeFail(t *testing.T) {
	inputs := []map[string]any{
		{
			"currency": "XRP",
			"issuer":   "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		},
		{
			"issuer": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		},
		{
			"currency": "0123456789ABCDEF0123456789ABCDEF01234567",
		},
	}

	for _, input := range inputs {
		_, err := Issue.ToBytes(input, false)
		require.Error(t, err)
	}
}
