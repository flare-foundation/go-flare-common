package check

import (
	"testing"

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
