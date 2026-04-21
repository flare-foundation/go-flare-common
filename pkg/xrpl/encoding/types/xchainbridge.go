package types

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

type XChainBridge struct{}

func (*XChainBridge) ToBytes(value any, _ bool) ([]byte, error) {
	valueMap, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid XChainBridge %v, %v", value, reflect.TypeOf(value))
	}

	lockingChainDoor, err := extractString(valueMap, "LockingChainDoor")
	if err != nil {
		return nil, fmt.Errorf("invalid XChainBridge %v: %w", value, err)
	}

	lockingChainIssue, ok := valueMap["LockingChainIssue"]
	if !ok {
		return nil, fmt.Errorf("XChainBridge %v, missing %s", value, "LockingChainIssue")
	}

	issuingChainDoor, err := extractString(valueMap, "IssuingChainDoor")
	if err != nil {
		return nil, fmt.Errorf("invalid XChainBridge %v: %w", value, err)
	}

	issuingChainIssue, ok := valueMap["IssuingChainIssue"]
	if !ok {
		return nil, fmt.Errorf("XChainBridge %v, missing %s", value, "IssuingChainIssue")
	}

	out := bytes.NewBuffer(nil)

	lockingChainDoorID, err := address.ID(lockingChainDoor)
	if err != nil {
		return nil, fmt.Errorf("invalid LockingChainDoor, %v: %w", lockingChainDoor, err)
	}
	if len(lockingChainDoorID) >= 256 {
		return nil, fmt.Errorf("lockingChainDoorID length overflow, %v: %w", issuingChainDoor, err)
	}

	err = out.WriteByte(uint8(len(lockingChainDoorID))) //nolint:gosec // checked in the if above
	if err != nil {
		return nil, fmt.Errorf("writing length to buffer, %v: %w", lockingChainDoorID, err)
	}
	_, err = out.Write(lockingChainDoorID)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer lockingChainDoorID, %v: %w", lockingChainDoorID, err)
	}

	lockingChain, err := Issue.ToBytes(lockingChainIssue, false)
	if err != nil {
		return nil, fmt.Errorf("encoding, %v: %w", value, err)
	}
	_, err = out.Write(lockingChain)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer, %v: %w", lockingChain, err)
	}

	issuingChainDoorID, err := address.ID(issuingChainDoor)
	if err != nil {
		return nil, fmt.Errorf("invalid IssuingChainDoor, %v: %w", issuingChainDoor, err)
	}
	if len(issuingChainDoorID) >= 256 {
		return nil, fmt.Errorf("issuingChainDoorID length overflow, %v: %w", issuingChainDoor, err)
	}

	err = out.WriteByte(uint8(len(issuingChainDoorID))) //nolint:gosec // checked in the if above
	if err != nil {
		return nil, fmt.Errorf("writing length to buffer, %v: %w", issuingChainDoorID, err)
	}
	_, err = out.Write(issuingChainDoorID)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer issuingChainDoorID, %v: %w", issuingChainDoorID, err)
	}

	issuingChain, err := Issue.ToBytes(issuingChainIssue, false)
	if err != nil {
		return nil, fmt.Errorf("encoding, %v: %w", value, err)
	}
	_, err = out.Write(issuingChain)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer issuingChain, %v: %w", issuingChain, err)
	}

	return out.Bytes(), nil
}

func (*XChainBridge) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	out := make(map[string]any)

	lockingChainDoor, err := readLengthPrefixedAccount(b)
	if err != nil {
		return nil, fmt.Errorf("reading LockingChainDoor: %w", err)
	}
	out["LockingChainDoor"] = lockingChainDoor

	lockingChainIssue, err := Issue.ToJSON(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading LockingChainIssue: %w", err)
	}
	out["LockingChainIssue"] = lockingChainIssue

	issuingChainDoor, err := readLengthPrefixedAccount(b)
	if err != nil {
		return nil, fmt.Errorf("reading IssuingChainDoor: %w", err)
	}
	out["IssuingChainDoor"] = issuingChainDoor

	issuingChainIssue, err := Issue.ToJSON(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading IssuingChainIssue: %w", err)
	}
	out["IssuingChainIssue"] = issuingChainIssue

	return out, nil
}

func readLengthPrefixedAccount(b *bytes.Buffer) (any, error) {
	l, err := b.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("reading length byte: %w", err)
	}
	if l != 20 {
		return nil, fmt.Errorf("invalid length byte expected %x is %x", 20, l)
	}

	return AccountID.ToJSON(b, 0)
}
