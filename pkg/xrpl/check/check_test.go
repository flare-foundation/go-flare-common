package check

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckFlags(t *testing.T) {
	tests := []struct {
		name    string
		flags   uint64
		wantErr bool
	}{
		{
			name:  "disable master only",
			flags: lsfDisableMaster,
		},
		{
			name:  "disable master with unrelated bits",
			flags: lsfDisableMaster | 0x00000001,
		},
		{
			name:    "master key not disabled",
			flags:   0,
			wantErr: true,
		},
		{
			name:    "requireDestTag set",
			flags:   lsfDisableMaster | lsfRequireDestTag,
			wantErr: true,
		},
		{
			name:    "disallowXRP set",
			flags:   lsfDisableMaster | lsfDisallowXRP,
			wantErr: true,
		},
		{
			name:    "depositAuth set",
			flags:   lsfDisableMaster | lsfDepositAuth,
			wantErr: true,
		},
		{
			name:    "all disallowed flags without master disabled",
			flags:   lsfRequireDestTag | lsfDisallowXRP | lsfDepositAuth,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := checkFlags(test.flags)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCheckSigners(t *testing.T) {
	tests := []struct {
		name    string
		sl      SignerList
		quorum  uint64
		signers []string
		wantErr bool
	}{
		{
			name: "valid single signer",
			sl: SignerList{
				SignerQuorum: 1,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
				},
			},
			quorum:  1,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"},
		},
		{
			name: "valid multiple signers",
			sl: SignerList{
				SignerQuorum: 2,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
					{SignerEntry: SignerEntry{Account: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v", SignerWeight: 1}},
				},
			},
			quorum:  2,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v"},
		},
		{
			name: "wrong quorum",
			sl: SignerList{
				SignerQuorum: 1,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
				},
			},
			quorum:  2,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"},
			wantErr: true,
		},
		{
			name: "wrong number of signers",
			sl: SignerList{
				SignerQuorum: 2,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
					{SignerEntry: SignerEntry{Account: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v", SignerWeight: 1}},
				},
			},
			quorum:  2,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"},
			wantErr: true,
		},
		{
			name: "signer weight not 1",
			sl: SignerList{
				SignerQuorum: 2,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 2}},
				},
			},
			quorum:  2,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"},
			wantErr: true,
		},
		{
			name: "signer not in expected list",
			sl: SignerList{
				SignerQuorum: 1,
				SignerEntries: []SignerEntryWrapped{
					{SignerEntry: SignerEntry{Account: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v", SignerWeight: 1}},
				},
			},
			quorum:  1,
			signers: []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := checkSigners(test.sl, test.quorum, test.signers)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	validSignerList := []SignerList{
		{
			SignerQuorum: 1,
			SignerEntries: []SignerEntryWrapped{
				{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
			},
		},
	}

	t.Run("valid", func(t *testing.T) {
		info := AccountInfoResponse{
			AccountData: AccountData{
				Flags:        lsfDisableMaster,
				SignersLists: validSignerList,
			},
		}
		require.NoError(t, info.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
	})

	t.Run("regular key set", func(t *testing.T) {
		info := AccountInfoResponse{
			AccountData: AccountData{
				Flags:        lsfDisableMaster,
				RegularKey:   "rSomeOtherKey",
				SignersLists: validSignerList,
			},
		}
		require.Error(t, info.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
	})

	t.Run("no signer list", func(t *testing.T) {
		info := AccountInfoResponse{
			AccountData: AccountData{
				Flags: lsfDisableMaster,
			},
		}
		require.Error(t, info.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
	})

	t.Run("multiple signer lists", func(t *testing.T) {
		info := AccountInfoResponse{
			AccountData: AccountData{
				Flags:        lsfDisableMaster,
				SignersLists: []SignerList{validSignerList[0], validSignerList[0]},
			},
		}
		require.Error(t, info.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
	})

	t.Run("bad flags", func(t *testing.T) {
		info := AccountInfoResponse{
			AccountData: AccountData{
				Flags:        lsfDisableMaster | lsfRequireDestTag,
				SignersLists: validSignerList,
			},
		}
		require.Error(t, info.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
	})
}

// Source: rippled include/xrpl/protocol/LedgerFormats.h lines 125-132 (AccountRoot LSF_FLAG macros).
func TestAccountRootFlagConstants(t *testing.T) {
	assert.Equal(t, uint64(0x00020000), lsfRequireDestTag, "lsfRequireDestTag")
	assert.Equal(t, uint64(0x00080000), lsfDisallowXRP, "lsfDisallowXRP")
	assert.Equal(t, uint64(0x00100000), lsfDisableMaster, "lsfDisableMaster")
	assert.Equal(t, uint64(0x01000000), lsfDepositAuth, "lsfDepositAuth")
}

// Source: xrpl.js packages/ripple-binary-codec/test/fixtures/signerlistset-tx.json
// meta.AffectedNodes[1].ModifiedNode.FinalFields (the SignerList ledger entry).
// Verifies that the Go SignerList JSON shape matches the rippled SignerList ledger-entry layout.
func TestSignerListFromXrplFixture(t *testing.T) {
	raw := `{
		"Flags": 0,
		"OwnerNode": "0000000000000000",
		"SignerEntries": [
			{"SignerEntry": {"Account": "rH7KDR67MZR7LDV7gesmEMXtaqU3FaK7Lr", "SignerWeight": 1}},
			{"SignerEntry": {"Account": "r4PQv7BCpp4SAJx3isNpQM8T2BuGrMQs5U", "SignerWeight": 1}},
			{"SignerEntry": {"Account": "rPqHsX34XApKSfE4UxKbqVXb3WRmmgMY2u", "SignerWeight": 1}}
		],
		"SignerListID": 0,
		"SignerQuorum": 3
	}`

	var sl SignerList
	require.NoError(t, json.Unmarshal([]byte(raw), &sl))

	require.EqualValues(t, 3, sl.SignerQuorum)
	require.Len(t, sl.SignerEntries, 3)

	signers := []string{
		"rH7KDR67MZR7LDV7gesmEMXtaqU3FaK7Lr",
		"r4PQv7BCpp4SAJx3isNpQM8T2BuGrMQs5U",
		"rPqHsX34XApKSfE4UxKbqVXb3WRmmgMY2u",
	}

	require.NoError(t, checkSigners(sl, 3, signers))

	// Wrong quorum expectation.
	require.Error(t, checkSigners(sl, 2, signers))

	// Missing one signer in the expected set.
	require.Error(t, checkSigners(sl, 3, signers[:2]))
}

// Deviation note: rippled's SignerList ledger entry includes sfOwner, sfOwnerNode,
// sfSignerListID, sfPreviousTxnID, sfPreviousTxnLgrSeq as additional required/optional
// fields (ledger_entries.macro lines 103-111). The Go SignerList struct only tracks the
// fields needed for the narrow "check multisig configuration" use case. This test pins
// the Go struct to the fields it currently exposes.
func TestSignerListStructFields(t *testing.T) {
	sl := SignerList{
		LedgerEntryType: "SignerList",
		Flags:           0,
		SignerEntries: []SignerEntryWrapped{
			{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
		},
		SignerQuorum: 1,
	}
	require.Equal(t, "SignerList", sl.LedgerEntryType)
	require.EqualValues(t, 1, sl.SignerQuorum)
}

// TestInfoHTTPStub exercises Info against an in-process JSON-RPC stub so the code path
// (request marshaling, response unmarshaling) is covered without hitting the network.
func TestInfoHTTPStub(t *testing.T) {
	expected := AccountInfoResponseWrapped{
		Result: AccountInfoResponse{
			AccountData: AccountData{
				Account:         "rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8",
				Balance:         "1000000000",
				Flags:           lsfDisableMaster,
				LedgerEntryType: "AccountRoot",
				Sequence:        42,
				SignersLists: []SignerList{{
					LedgerEntryType: "SignerList",
					Flags:           0,
					SignerEntries: []SignerEntryWrapped{
						{SignerEntry: SignerEntry{Account: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW", SignerWeight: 1}},
					},
					SignerQuorum: 1,
				}},
			},
			Validated: true,
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var req AccountInfoRequest
		require.NoError(t, json.Unmarshal(body, &req))
		require.Equal(t, accountInfoMethod, req.Method)
		require.Len(t, req.Params, 1)
		require.Equal(t, "rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8", req.Params[0].Account)
		require.True(t, req.Params[0].SignerList)

		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(expected))
	}))
	defer srv.Close()

	rpc := JSONRPC{URL: srv.URL}
	resp, err := rpc.Info("rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8")
	require.NoError(t, err)

	require.Equal(t, expected.Result.AccountData.Account, resp.AccountData.Account)
	require.EqualValues(t, 42, resp.Sequence())
	require.Len(t, resp.AccountData.SignersLists, 1)
	require.NoError(t, resp.Check(1, []string{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"}))
}

// TestInfoHTTPStubRejectsBadStatus covers the non-200 error branch.
func TestInfoHTTPStubRejectsBadStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadGateway)
		_, _ = fmt.Fprint(w, "upstream unavailable")
	}))
	defer srv.Close()

	rpc := JSONRPC{URL: srv.URL}
	_, err := rpc.Info("rpo6E7mHvQ4xzeEBy8ViVzbG8q251ztKB8")
	require.Error(t, err)
}

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
