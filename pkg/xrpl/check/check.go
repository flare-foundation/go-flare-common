// Package check provides XRPL account verification for multi-signature configuration.
package check

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/call"
)

const accountInfoMethod = "account_info"

// lsf = ledgerStateFlag https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/accountroot#accountroot-flags
const (
	lsfRequireDestTag uint64 = 0x00020000
	lsfDisallowXRP    uint64 = 0x00080000
	lsfDisableMaster  uint64 = 0x00100000
	lsfDepositAuth    uint64 = 0x01000000
)

type AccountInfoResponseWrapped struct {
	Result AccountInfoResponse `json:"result"`
}

type AccountInfoResponse struct {
	AccountData AccountData `json:"account_data"`
	// AccountFlags AccountFlags `json:"account_flags"`
	Validated bool `json:"validated"`
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

// type AccountFlags struct {
// 	DepositAuth           bool `json:"depositAuth"`
// 	DisableMasterKey      bool `json:"disableMasterKey"`
// 	DisallowIncomingXRP   bool `json:"disallowIncomingXRP"`
// 	RequireDestinationTag bool `json:"requireDestinationTag"`
// }

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
	// Transport sets the HTTP transport used by Info. Nil falls back to
	// http.DefaultTransport, which permits any reachable address (including
	// private VPC ranges where a co-located rippled typically lives).
	// Callers expecting a public-only URL should pass safeurl.NewTransport()
	// to enable SSRF protection at dial time; callers happy with the default
	// take responsibility for trusting the URL.
	Transport http.RoundTripper
}

// Info gets account info for the XRPL address.
func (jr JSONRPC) Info(ctx context.Context, address string) (AccountInfoResponse, error) {
	request := AccountInfoRequest{
		Method: accountInfoMethod,
		Params: []AccountInfoParams{{
			Account:    address,
			SignerList: true,
		}},
	}

	res, err := call.Post[AccountInfoRequest, AccountInfoResponseWrapped](ctx, jr.URL, call.NoAPIKey, request, call.Params{
		Timeout:         30 * time.Second,
		MaxResponseSize: 10000,
		Transport:       jr.Transport,
	})
	if err != nil {
		return AccountInfoResponse{}, fmt.Errorf("calling %s: %w", jr.URL, err)
	}

	return res.Message.Result, nil
}

// Check checks account info:
//   - response is from a validated ledger
//   - illegal Flags
//   - required Flags
//   - signers list setting
//   - no regular key
//
// Audit M11: a SignerList read from a non-validated ledger can be rolled
// back; trusting it for quorum decisions risks issuing transactions against
// stale or speculative state. Reject responses where rippled has not yet
// committed the read.
func (ai AccountInfoResponse) Check(quorum uint64, signers []string) error {
	if !ai.Validated {
		return errors.New("account_info response is from a non-validated ledger")
	}

	if err := checkFlags(ai.AccountData.Flags); err != nil {
		return fmt.Errorf("flags: %w", err)
	}

	if ai.AccountData.RegularKey != "" {
		return fmt.Errorf("regular key set: %s", ai.AccountData.RegularKey)
	}

	if len(ai.AccountData.SignersLists) != 1 {
		return errors.New("signers not set")
	}

	if err := checkSigners(ai.AccountData.SignersLists[0], quorum, signers); err != nil {
		return fmt.Errorf("signers: %w", err)
	}

	return nil
}

// Sequence return the sequence of the account. TODO: check if this can be undefined.
func (ai AccountInfoResponse) Sequence() uint64 {
	return ai.AccountData.Sequence
}

// checkFlags checks that the following flags are set:
//   - lsfDisableMaster
//
// and the following are not set:
//   - lsfRequireDestTag
//   - lsfDisallowXRP
//   - lsfDepositAuth
func checkFlags(f uint64) error {
	if f&lsfRequireDestTag != 0 {
		return errors.New("requireDestTag set")
	}

	if f&lsfDisallowXRP != 0 {
		return errors.New("disallowXRP set")
	}

	if f&lsfDepositAuth != 0 {
		return errors.New("depositAuth set")
	}

	if f&lsfDisableMaster != lsfDisableMaster {
		return errors.New("master key not disabled")
	}

	return nil
}

// checkSigners checks that SignerList has the right quorum and exactly the specified signers all with weight 1.
// Duplicate accounts in SignerEntries are rejected — without dedup, a malformed list could pass the length
// equality check while missing an expected signer.
func checkSigners(sl SignerList, quorum uint64, signers []string) error {
	if sl.SignerQuorum != quorum {
		return fmt.Errorf("wrong signer quorum. xrpl: %d, expected: %d", sl.SignerQuorum, quorum)
	}

	if len(sl.SignerEntries) != len(signers) {
		return fmt.Errorf("wrong number of signers. xrpl: %d, expected: %d", len(sl.SignerEntries), len(signers))
	}

	seen := make(map[string]struct{}, len(sl.SignerEntries))
	for j := range sl.SignerEntries {
		account := sl.SignerEntries[j].SignerEntry.Account
		if sl.SignerEntries[j].SignerEntry.SignerWeight != 1 {
			return fmt.Errorf("signer %s has weight %d", account, sl.SignerEntries[j].SignerEntry.SignerWeight)
		}
		if _, dup := seen[account]; dup {
			return fmt.Errorf("duplicate signer %s on xrpl", account)
		}
		seen[account] = struct{}{}
		if !slices.Contains(signers, account) {
			return fmt.Errorf("signer on xrpl %s not among expected signers", account)
		}
	}

	return nil
}
