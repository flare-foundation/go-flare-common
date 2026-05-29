// Package op defines operation types and commands used in TEE instruction processing.
package op

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
)

// Type represents an operation type in the TEE system.
type Type string

const (
	Reg        Type = "F_REG"
	Wallet     Type = "F_WALLET"
	Get        Type = "F_GET"
	Policy     Type = "F_POLICY"
	Governance Type = "F_GOVERNANCE"
	XRP        Type = "F_XRP"
	BTC        Type = "F_BTC"
	FDC2       Type = "F_FDC2"
)

// Command represents an operation command in the TEE system.
type Command string

const (
	// ReplicateFrom    OPCommand = "REPLICATE_FROM"

	TEEAttestation Command = "TEE_ATTESTATION"

	// ToPauseForUpdate OPCommand = "TO_PAUSE_FOR_UPGRADE".

	KeyDataProviderRestore     Command = "KEY_DATA_PROVIDER_RESTORE"
	KeyDataProviderRestoreTest Command = "KEY_DATA_PROVIDER_RESTORE_TEST"
	KeyDelete                  Command = "KEY_DELETE"
	KeyDirectBackup            Command = "KEY_DIRECT_BACKUP"
	KeyDirectRestore           Command = "KEY_DIRECT_RESTORE"
	KeyGenerate                Command = "KEY_GENERATE"
	VRF                        Command = "VRF"

	KeyInfo Command = "KEY_INFO"
	// KeyProof has no payload binding yet; without one tying the proof to instructionId,
	// attestation hash, and freshness nonce, a signed KeyProof can be replayed across contexts.
	// Do not use in production until the binding is defined in pkg/tee/structs.
	KeyProof  Command = "KEY_PROOF"
	TEEBackup Command = "TEE_BACKUP"
	TEEInfo   Command = "TEE_INFO"

	InitializePolicy   Command = "INITIALIZE_POLICY"
	UpdatePolicy       Command = "UPDATE_POLICY"
	SetMachinePathList Command = "SET_MACHINE_PATH_LIST"

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
		KeyDirectBackup:            true,
		KeyDirectRestore:           true,
		KeyGenerate:                true,
		VRF:                        true,
	},
	Get: {
		KeyInfo:   true,
		KeyProof:  true,
		TEEBackup: true,
		TEEInfo:   true,
	},
	Policy: {
		InitializePolicy: true,
		UpdatePolicy:     true,
	},
	Governance: {
		SetMachinePathList: true,
	},
	XRP: {
		Pay:     true,
		Reissue: true,
	},
	BTC: {
		Pay:     true,
		Reissue: true,
	},
	FDC2: {
		Prove: true,
	},
}

// HashToOPType converts hash to op.Type.
func HashToOPType(h common.Hash) Type {
	s := convert.CommonHashToString(h)
	return Type(s)
}

// Hash returns utf8 encoding of t padded to 32 bytes. Panics if t exceeds 32 bytes.
func (t Type) Hash() common.Hash {
	return toHash(string(t))
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

// Hash returns utf8 encoding of c padded to 32 bytes. Panics if c exceeds 32 bytes.
func (c Command) Hash() common.Hash {
	return toHash(string(c))
}

func toHash(s string) common.Hash {
	b := []byte(s)
	if len(b) > 32 {
		panic("op: identifier too long for 32-byte hash: " + s)
	}
	var h common.Hash
	copy(h[:], b)
	return h
}

// IsValid reports whether the command is valid for the given operation type.
func IsValid(t Type, c Command) bool {
	if t.IsSystem() {
		return validSystemPairs[t][c]
	}

	return !t.isF()
}

// IsValidPair checks whether the hashed operation type and command form a valid pair.
func IsValidPair(opType common.Hash, opCommand common.Hash) bool {
	return IsValid(HashToOPType(opType), HashToOPCommand(opCommand))
}
