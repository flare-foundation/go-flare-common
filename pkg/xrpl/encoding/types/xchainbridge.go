package types

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

type XChainBridge struct{}

func (x *XChainBridge) ToBytes(value any, _ bool) ([]byte, error) {
	valueMap, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid XChainBridge %v, %v", value, reflect.TypeOf(value))
	}

	lockingChainDoor, err := extractString(valueMap, "LockingChainDoor")
	if err != nil {
		return nil, fmt.Errorf("invalid XChainBridge %v: %v", value, err)
	}

	lockingChainIssue, ok := valueMap["LockingChainIssue"]
	if !ok {
		return nil, fmt.Errorf("XChainBridge %v, missing %s", value, "LockingChainIssue")
	}

	issuingChainDoor, err := extractString(valueMap, "IssuingChainDoor")
	if err != nil {
		return nil, fmt.Errorf("invalid XChainBridge %v: %v", value, err)
	}

	issuingChainIssue, ok := valueMap["IssuingChainIssue"]
	if !ok {
		return nil, fmt.Errorf("XChainBridge %v, missing %s", value, "IssuingChainIssue")
	}

	out := bytes.NewBuffer(nil)

	lockingChainDoorID, err := address.ID(lockingChainDoor)
	if err != nil {
		return nil, fmt.Errorf("invalid LockingChainDoor, %v: %v", lockingChainDoor, err)
	}
	err = out.WriteByte(uint8(len(lockingChainDoorID)))
	if err != nil {
		return nil, fmt.Errorf("writing length to buffer, %v: %v", lockingChainDoorID, err)
	}
	_, err = out.Write(lockingChainDoorID)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer lockingChainDoorID, %v: %v", lockingChainDoorID, err)
	}

	lockingChain, err := Issue.ToBytes(lockingChainIssue, false)
	if err != nil {
		return nil, fmt.Errorf("encoding, %v: %v", value, err)
	}
	_, err = out.Write(lockingChain)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer, %v: %v", lockingChain, err)
	}

	issuingChainDoorID, err := address.ID(issuingChainDoor)
	if err != nil {
		return nil, fmt.Errorf("invalid IssuingChainDoor, %v: %v", issuingChainDoor, err)
	}
	err = out.WriteByte(uint8(len(issuingChainDoorID)))
	if err != nil {
		return nil, fmt.Errorf("writing length to buffer, %v: %v", issuingChainDoorID, err)
	}
	_, err = out.Write(issuingChainDoorID)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer issuingChainDoorID, %v: %v", issuingChainDoorID, err)
	}

	issuingChain, err := Issue.ToBytes(issuingChainIssue, false)
	if err != nil {
		return nil, fmt.Errorf("encoding, %v: %v", value, err)
	}
	_, err = out.Write(issuingChain)
	if err != nil {
		return nil, fmt.Errorf("writing to buffer issuingChain, %v: %v", issuingChain, err)
	}

	return out.Bytes(), nil
}

func (x *XChainBridge) ToJson(b *bytes.Buffer, _ int) (any, error) {
	out := make(map[string]any)

	l, err := b.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("reading length byte: %v", err)
	}
	if l != 20 {
		return nil, fmt.Errorf("invalid length byte expected %x is %x", 20, l)
	}

	lockingChainDoor, err := AccountID.ToJson(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading LockingChainDoor: %v", err)
	}
	out["LockingChainDoor"] = lockingChainDoor

	lockingChainIssue, err := Issue.ToJson(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading LockingChainIssue: %v", err)
	}
	out["LockingChainIssue"] = lockingChainIssue

	issuingChainDoor, err := AccountID.ToJson(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading IssuingChainDoor: %v", err)
	}
	out["IssuingChainDoor"] = issuingChainDoor

	issuingChainIssue, err := Issue.ToJson(b, 0)
	if err != nil {
		return nil, fmt.Errorf("reading IssuingChainIssue: %v", err)
	}
	out["IssuingChainIssue"] = issuingChainIssue

	return out, nil
}
