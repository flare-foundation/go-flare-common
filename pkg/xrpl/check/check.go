package check

import (
	"context"

	"github.com/flare-foundation/go-flare-common/pkg/call"
)

const accountInfoMethod = "account_info"

type AccountInfoResponseWrapped struct {
	Result AccountInfoResponse `json:"result"`
}

type AccountInfoResponse struct {
	AccountData  AccountData  `json:"account_data"`
	AccountFlags AccountFlags `json:"account_flags"`
	Validated    bool         `json:"validated"`
}

// AccountData has data from https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/accountroot#accountroot-fields
type AccountData struct {
	Account         string
	Balance         string `json:",omitempty"`
	Flags           uint64
	LedgerEntryType string
	RegularKey      string `json:",omitempty"`
	Sequence        uint64
	SignersLists    []SignerList `json:"signer_lists"`
	TicketCount     uint32       `json:",omitempty"`
}

type AccountFlags struct {
	DepositAuth           bool `json:"depositAuth"`
	DisableMasterKey      bool `json:"disableMasterKey"`
	DisallowIncomingXRP   bool `json:"disallowIncomingXRP"`
	RequireDestinationTag bool `json:"requireDestinationTag"`
}

type SignerList struct {
	LedgerEntryType string
	Flags           uint64
	SignerEntries   []SignerEntryWrapped
	SignerQuorum    uint64
}

type SignerEntryWrapped struct {
	SignerEntry SignerEntry
}

type SignerEntry struct {
	Account      string `json:"Account"`
	SignerWeight uint16 `json:"SignerWeight"`
}

type AccountInfoRequest struct {
	Method string              `json:"method"`
	Params []AccountInfoParams `json:"params"`
}

type AccountInfoParams struct {
	Account    string `json:"account"`
	SignerList bool   `json:"signer_lists"`
}

type JSONRPC struct {
	URL string
}

func (jr JSONRPC) Info(address string) (AccountInfoResponse, error) {
	request := AccountInfoRequest{
		Method: accountInfoMethod,
		Params: []AccountInfoParams{{
			Account:    address,
			SignerList: true,
		}},
	}

	ctx := context.Background()

	res, err := call.Post[AccountInfoRequest, AccountInfoResponseWrapped](ctx, jr.URL, call.NoAPIKey, request, call.Params{
		Timeout:         0,
		MaxResponseSize: 10000,
	})

	return res.Result, err
}
