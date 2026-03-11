package op

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
)

type Type string

const (
	Reg    Type = "F_REG"
	Wallet Type = "F_WALLET"
	Get    Type = "F_GET"
	Policy Type = "F_POLICY"
	XRP    Type = "F_XRP"
	BTC    Type = "F_BTC"
	FTDC   Type = "F_FTDC"
)

type Command string

const (
	// ReplicateFrom    OPCommand = "REPLICATE_FROM"

	TEEAttestation Command = "TEE_ATTESTATION"

	// ToPauseForUpdate OPCommand = "TO_PAUSE_FOR_UPGRADE".

	KeyDataProviderRestore     Command = "KEY_DATA_PROVIDER_RESTORE"
	KeyDataProviderRestoreTest Command = "KEY_DATA_PROVIDER_RESTORE_TEST"
	KeyDelete                  Command = "KEY_DELETE"
	KeyGenerate                Command = "KEY_GENERATE"
	VRF                        Command = "VRF"

	KeyInfo   Command = "KEY_INFO"
	TEEBackup Command = "TEE_BACKUP"
	TEEInfo   Command = "TEE_INFO"

	InitializePolicy Command = "INITIALIZE_POLICY"
	UpdatePolicy     Command = "UPDATE_POLICY"

	Pay     Command = "PAY"
	Reissue Command = "REISSUE"

	Prove Command = "PROVE"
)

const (
	flarePrefix = "F_"
)

var validSystemPairs = map[Type]map[Command]bool{
	Reg: {
		TEEAttestation: true,
	},
	Wallet: {
		KeyDataProviderRestore:     true,
		KeyDataProviderRestoreTest: true,
		KeyDelete:                  true,
		KeyGenerate:                true,
		VRF:                        true,
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

// HashToOPType converts hash to op.Type.
func HashToOPType(h common.Hash) Type {
	s := convert.CommonHashToString(h)
	return Type(s)
}

// Hash returns utf8 encoding of t padded to 32 bytes.
func (t Type) Hash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(t), 32))
}

func (t Type) IsSystem() bool {
	_, exists := validSystemPairs[t]
	return exists
}

func (t Type) isF() bool {
	if len(t) < 2 {
		return false
	}
	return t[0:2] == flarePrefix
}

// HashToOPCommand converts hash to OpCommand and indicates whether it is a valid OpCommand.
func HashToOPCommand(h common.Hash) Command {
	s := convert.CommonHashToString(h)
	return Command(s)
}

// Hash returns utf8 encoding of c padded to 32 bytes.
func (c Command) Hash() common.Hash {
	return common.BytesToHash(common.RightPadBytes([]byte(c), 32))
}

func IsValid(t Type, c Command) bool {
	if t.IsSystem() {
		return validSystemPairs[t][c]
	}

	return !t.isF()
}

func IsValidPair(opType common.Hash, opCommand common.Hash) bool {
	return IsValid(HashToOPType(opType), HashToOPCommand(opCommand))
}
