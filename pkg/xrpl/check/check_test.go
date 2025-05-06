package check

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	url := "https://s.altnet.rippletest.net:51234/"

	rpc := JSONRPC{
		URL: url,
	}

	resp, err := rpc.Info("rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8")
	require.NoError(t, err)

	require.Len(t, resp.AccountData.SignersLists, 1)
	require.Len(t, resp.AccountData.SignersLists[0].SignerEntries, 1)
}
