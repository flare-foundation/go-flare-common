package toml

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type testStructHappy struct {
	A     int            `toml:"a"`
	B     uint           `tom:"b"`
	C     uint8          `toml:"c"`
	Addr1 common.Address `toml:"addr1"`
	BN    *big.Int       `toml:"bn"`
	BNS   *big.Int       `toml:"bns"`
	BNH   *big.Int       `toml:"bnh"`
	Z     *common.Hash   `toml:"z"`
}

func TestReadTomlHappy(t *testing.T) {
	addr := common.HexToAddress("0xE3D78e09f8464f82d1F248d74E2FdEb21010e3Ec")
	bn, ok := new(big.Int).SetString("184467440737095516160", 10)

	require.True(t, ok)

	path := "./testHappy.toml"

	a, err := Read[testStructHappy](path, true)
	require.NoError(t, err)

	dest := new(testStructHappy)

	err = ReadTo(path, dest, true)
	require.NoError(t, err)

	expected := testStructHappy{
		A:     -10,
		B:     10,
		C:     11,
		Addr1: addr,
		BN:    bn,
		BNS:   bn,
		BNH:   bn,
		Z:     nil,
	}

	require.Equal(t, expected, a)
	require.Equal(t, expected, *dest)
}

type unknown struct{ R int }

func TestReadTomlFail(t *testing.T) {
	path := "./testHappy.toml"
	_, err := Read[unknown](path, false)
	require.Error(t, err)

	path = "./nonExistant.toml"
	_, err = Read[unknown](path, false)
	require.Error(t, err)
}
