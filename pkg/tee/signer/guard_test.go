package signer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestECDSAPubKeyToECIESRejectsNil(t *testing.T) {
	_, err := ECDSAPubKeyToECIES(nil)
	require.Error(t, err)
}

func TestECDSAPrivKeyToECIESRejectsNil(t *testing.T) {
	_, err := ECDSAPrivKeyToECIES(nil)
	require.Error(t, err)
}
