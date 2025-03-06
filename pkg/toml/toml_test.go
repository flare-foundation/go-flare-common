package toml

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type testStruct1 struct {
	A     int            `toml:"a"`
	C     *int           `toml:"c"`
	Addr1 common.Address `toml:"addr1"`
}

type testStruct2 struct {
	A     int            `toml:"a"`
	C     *int           `toml:"c"`
	Addr1 common.Address `toml:"addr1"`
	Addr2 common.Address `toml:"addr2"`
}

func TestReadToml(t *testing.T) {
	addr := common.HexToAddress("0xE3D78e09f8464f82d1F248d74E2FdEb21010e3Ec")

	path := "./test.toml"
	a, err := ReadToml[testStruct1](path, true)
	require.NoError(t, err)

	expected := testStruct1{10, nil, addr}
	require.Equal(t, expected, a)

	_, err = ReadToml[testStruct2](path, true)
	require.Error(t, err)

	_, err = ReadToml[testStruct1](path, false)
	require.Error(t, err)
}
