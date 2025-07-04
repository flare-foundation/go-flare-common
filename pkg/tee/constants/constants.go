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
	Policy OPType = "POLICY"
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

	InitializePolicy OPCommand = "INITIALIZE_POLICY"
	UpdatePolicy     OPCommand = "UPDATE_POLICY"

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
	Policy: {
		InitializePolicy: true,
		UpdatePolicy:     true,
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

// IsValid checks whether t is a valid OPType.
func (t OPType) IsValid() bool {
	_, ok := validPairs[t]
	return ok
}

// Hash returns utf8 encoding of t padded to 32 bytes.
func (t OPType) Hash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(t), 32))
}

// IsValid checks whether c is a valid OPCommand.
func (c OPCommand) IsValid() bool {
	for _, m := range validPairs {
		_, ok := m[c]
		if ok {
			return true
		}
	}

	return false
}

// Hash returns utf8 encoding of c padded to 32 bytes.
func (c OPCommand) Hash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(c), 32))
}

// StringToOPType converts string to OPType and indicates whether it is a valid OPType.
func StringToOPType(s string) (OPType, bool) {
	t := OPType(s)

	return t, t.IsValid()
}

// HashToOPType converts hash to OPType and indicates whether it is a valid OPType.
func HashToOPType(h common.Hash) (OPType, bool) {
	s := strings.TrimRight(string(h.Bytes()), "\x00")

	return StringToOPType(s)
}

// StringToOPCommand converts string to OpCommand and indicates whether it is a valid OpCommand.
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

// HashToOPCommand converts hash to OpCommand and indicates whether it is a valid OpCommand.
func HashToOPCommand(h common.Hash) (OPCommand, bool) {
	s := strings.TrimRight(string(h.Bytes()), "\x00")

	return StringToOPCommand(s)
}

// IsValidPair checks whether (t,c) is a valid pair of OPType and OPcommand.
func IsValidPair(t OPType, c OPCommand) bool {
	cs, ok := validPairs[t]
	if !ok {
		return false
	}
	_, ok = cs[c]
	return ok
}
