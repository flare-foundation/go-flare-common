package aggregator

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/stretchr/testify/require"
)

func TestAddSignatures(t *testing.T) {
	tx := make(map[string]any)

	tx["TransactionType"] = "TrustSet"
	tx["Account"] = "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC"
	tx["Fee"] = "30000"
	tx["Flags"] = 262144

	tx["LimitAmount"] = map[string]any{
		"currency": "USD",
		"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		"value":    "100",
	}
	tx["Sequence"] = 2
	tx["SigningPubKey"] = ""

	s1 := &signer.Signer{
		Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
		TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
	}

	s2 := &signer.Signer{
		Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		TxnSignature:  "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
		SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
	}

	signers := []*signer.Signer{s1, s2}

	blob, err := signing.JoinMultisig(tx, signers)
	require.NoError(t, err)

	acc := Account{
		Address:    "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true},
		Quorum:     1,
		txs:        make(map[common.Hash]*TX),
	}

	transaction, dispatch, err := acc.AddSignatures(blob)
	require.NoError(t, err)
	require.True(t, dispatch)

	require.Len(t, transaction.signers, 1, "signatures")
	require.Len(t, acc.txs, 1, "transactions")
}
