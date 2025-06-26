package constants

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type OPType string

const (
	Reg    OPType = "REG"
	Wallet OPType = "WALLET"
	Get    OPType = "GET"
	XRP    OPType = "XRP"
	BTC    OPType = "BTC"
	FTDC   OPType = "FTDC"
)

type OPCommand string

const (
	// ReplicateFrom    OPCommand = "REPLICATE_FROM".

	TEEAttestation OPCommand = "TEE_ATTESTATION"

	// ToPauseForUpdate OPCommand = "TO_PAUSE_FOR_UPGRADE".

	KeyDataProviderRestore     OPCommand = "KEY_DATA_PROVIDER_RESTORE"
	KeyDataProviderRestoreTest OPCommand = "KEY_DATA_PROVIDER_RESTORE_TEST"
	KeyDelete                  OPCommand = "KEY_DELETE"
	KeyGenerate                OPCommand = "KEY_GENERATE"

	KeyInfo   OPCommand = "KEY_INFO"
	TEEBackup OPCommand = "TEE_BACKUP"
	TEEInfo   OPCommand = "TEE_INFO"

	Pay     OPCommand = "PAY"
	Reissue OPCommand = "REISSUE"

	Prove OPCommand = "PROVE"
)

var validPairs = map[OPType]map[OPCommand]bool{
	Reg: {
		TEEAttestation: true,
	},
	Wallet: {
		KeyDataProviderRestore:     true,
		KeyDataProviderRestoreTest: true,
		KeyDelete:                  true,
		KeyGenerate:                true,
	},
	Get: {
		KeyInfo:   true,
		TEEBackup: true,
		TEEInfo:   true,
	},
	XRP: {
		Pay:     true,
		Reissue: true,
	},
	BTC: {
		Pay:     true,
		Reissue: true,
	},
	FTDC: {
		Prove: true,
	},
}

func (t OPType) IsValid() bool {
	_, ok := validPairs[t]
	return ok
}

func (t OPType) ToHash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(t), 32))
}

func (c OPCommand) IsValid() bool {
	for _, m := range validPairs {
		_, ok := m[c]
		if ok {
			return true
		}
	}

	return false
}

func (c OPCommand) ToHash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(c), 32))
}

func StringToOPType(s string) (OPType, bool) {
	t := OPType(s)

	return t, t.IsValid()
}

func HashToOPType(h common.Hash) (OPType, bool) {
	s := strings.TrimRight(string(h.Bytes()), "\x00")

	return StringToOPType(s)
}

func StringToOPCommand(s string) (OPCommand, bool) {
	c := OPCommand(s)

	for _, m := range validPairs {
		_, ok := m[c]
		if ok {
			return c, true
		}
	}

	return c, false
}

func HashToOCommand(h common.Hash) (OPCommand, bool) {
	s := strings.TrimRight(string(h.Bytes()), "\x00")

	return StringToOPCommand(s)
}

func IsValidPair(t OPType, c OPCommand) bool {
	cs, ok := validPairs[t]
	if !ok {
		return false
	}
	_, ok = cs[c]
	return ok
}
