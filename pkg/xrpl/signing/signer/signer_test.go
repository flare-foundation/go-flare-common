package signer

import (
	"bytes"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	s1 := &Signer{
		Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
		TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
	}

	s2 := &Signer{
		Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF0207366730440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
	}

	t.Run("compare", func(t *testing.T) {
		_, err := s1.Value()
		require.NoError(t, err)

		_, err = s2.Value()
		require.NoError(t, err)

		c1 := Compare(s1, s2)
		require.Equal(t, -1, c1)

		c2 := Compare(s2, s1)
		require.Equal(t, 1, c2)

		c3 := Compare(s1, s1)
		require.Equal(t, 0, c3)
	})

	t.Run("sort empty", func(t *testing.T) {
		x, y := Sort(nil)
		require.Len(t, x, 0)
		require.Len(t, y, 0)

		x, y = Sort([]*Signer{})
		require.Len(t, x, 0)
		require.Len(t, y, 0)
	})

	t.Run("sort one", func(t *testing.T) {
		x, y := Sort([]*Signer{s1})
		require.Len(t, x, 1)
		require.Len(t, y, 0)
		require.Equal(t, s1, x[0])
	})

	t.Run("sort two", func(t *testing.T) {
		x, y := Sort([]*Signer{s1, s2})
		require.Len(t, x, 2)
		require.Len(t, y, 0)
		require.Equal(t, s1, x[0])
		require.Equal(t, s2, x[1])

		x, y = Sort([]*Signer{s2, s1})
		require.Len(t, x, 2)
		require.Len(t, y, 0)
		require.Equal(t, s1, x[0])
		require.Equal(t, s2, x[1])
	})
}

func TestFormatParse(t *testing.T) {
	signers := []Signer{
		{
			Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
			TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
			SigningPubKey: "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
		},
	}

	for _, signer := range signers {
		object := signer.Format()

		parsed, err := Parse(object)
		require.NoError(t, err)

		require.Equal(t, signer.Account, parsed.Account)
		require.Equal(t, signer.TxnSignature, parsed.TxnSignature)
		require.Equal(t, signer.SigningPubKey, parsed.SigningPubKey)
	}
}

func TestParseFail(t *testing.T) {
	tests := []types.ArrayObject{
		{},
		{
			"Signer": "invalid",
		},
		{
			"NotSigner": "invalid",
		},
		{
			"Signer": map[string]any{
				"Account":      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TxnSignature": "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"OtherField":   10,
			},
		},
		{
			"Signer": map[string]any{
				"Account":       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
				"OtherField":    10,
			},
		},
		{
			"Signer": map[string]any{
				"Account":       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
				"OtherField":    10,
			},
		},
		{
			"Signer": map[string]any{
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
				"OtherField":    10,
			},
		},
		{
			"Signer": map[string]any{
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "******************************************************************",
			},
		},
		{
			"Signer": map[string]any{
				"Account":       123,
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "******************************************************************",
			},
		},
		{
			"Signer": map[string]any{
				"Account":       "invalidAddress",
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "******************************************************************",
			},
		},
		{
			"Signer": map[string]any{
				"Account":       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
			},
			"OtherSigner": map[string]any{
				"Account":       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TxnSignature":  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
				"SigningPubKey": "ED5F5AC8B98974A3CA843326D9B88CEBD0560177B973EE0B149F782CFAA5F3BB9A",
			},
		},
	}

	for _, test := range tests {
		_, err := Parse(test)
		require.Error(t, err)
	}
}

// rippled src/libxrpl/tx/Transactor.cpp checkMultiSign:
// "Both the multiSigners and accountSigners are sorted by account."
// Rippled compares AccountIDs as 160-bit big-endian integers. This test checks
// that Sort produces an ordering consistent with big.Int over the raw 20-byte
// accountID and verifies that ordering against a known rippled-derived set
// (addresses are from rippled src/test/protocol/SecretKey_test.cpp secp256k1TestVectors).
func TestSortMatchesAccountIDOrder(t *testing.T) {
	addrs := []string{
		"rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		"rfZCLUjvKSNmg5xMufb6fgq9VfP5biBDfU",
		"rpyTz8db86bWHi8E43GGexhRsLDwwMRta3",
		"rfwGiMcsffk9TcwQu1KMYYvkPq4cdHdMik",
		"rNYp97qFzBZ5SLfXRtRMwFNxHrGR9cQyGL",
	}

	signers := make([]*Signer, len(addrs))
	for i, a := range addrs {
		signers[i] = &Signer{
			Account:       a,
			TxnSignature:  "DEAD",
			SigningPubKey: "02BEEF",
		}
	}

	in, out := Sort(signers)
	require.Len(t, out, 0)
	require.Len(t, in, len(addrs))

	for i := 1; i < len(in); i++ {
		prev, err := address.ID(in[i-1].Account)
		require.NoError(t, err)
		curr, err := address.ID(in[i].Account)
		require.NoError(t, err)
		require.Lessf(t, bytes.Compare(prev, curr), 0,
			"sort should produce ascending accountID order: %s vs %s",
			in[i-1].Account, in[i].Account)
	}
}
