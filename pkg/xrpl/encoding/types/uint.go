package types

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/defs"
)

// TODO: ??Allow hexadecimal string??

type UInt8 struct {
}

type InvalidType struct {
	t string
	v any
}

func (i *InvalidType) Error() string {
	return fmt.Sprintf("invalid %s: %v", i.t, i.v)
}

func InvalidUInt8(v any) *InvalidType {
	return &InvalidType{t: "UInt8", v: v}
}

// ToBytes serializes values of UInt8 fields.
func (u *UInt8) ToBytes(value any, _ bool) ([]byte, error) {
	tempInt, err := convertInt64(value, "UInt8")
	if err != nil {
		return nil, err
	}
	if tempInt >= 1<<8 || tempInt < 0 {
		return nil, InvalidUInt8(value)
	}
	return []byte{byte(tempInt)}, nil
}

type UInt16 struct {
}

// ToBytes serializes values of UInt16 fields.
func (u *UInt16) ToBytes(value any, _ bool) ([]byte, error) {
	var valueUint uint16

	valueStr, ok := value.(string)
	if ok {
		txType, ok := defs.TxTypeToValue[valueStr]
		if !ok || txType < 0 || txType >= 1<<16 {
			return nil, &InvalidType{
				t: "TransactionType",
				v: value,
			}
		}
		valueUint = uint16(txType)
	} else {
		tempInt64, err := convertInt64(value, "UInt16")
		if err != nil {
			return nil, err
		}

		if tempInt64 >= 1<<16 || tempInt64 < 0 {
			return nil,
				&InvalidType{
					t: "UInt16",
					v: value,
				}
		}

		valueUint = uint16(tempInt64)
	}

	out := make([]byte, 2)
	binary.BigEndian.PutUint16(out, valueUint)

	return out, nil
}

type UInt32 struct {
}

// ToBytes serializes values of UInt32 fields.
func (u *UInt32) ToBytes(value any, _ bool) ([]byte, error) {
	tempInt, err := convertInt64(value, "UInt32")
	if err != nil {
		return nil, err
	}
	if tempInt >= 1<<32 || tempInt < 0 {
		return nil, &InvalidType{
			t: "UInt32",
			v: value,
		}
	}

	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, uint32(tempInt))

	return out, nil
}

type UInt64 struct {
}

// ToBytes serializes values of UInt64 fields.
func (u *UInt64) ToBytes(value any, _ bool) ([]byte, error) {
	var valueUint uint64
	var err error

	switch value := value.(type) {
	case string:
		valueUint, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid UInt64: %v, error: %v", value, err)
		}
	case uint64:
		valueUint = value
	default:
		valueInt, err := convertInt64(value, "int64")
		if err != nil || valueInt < 0 {
			return nil, fmt.Errorf("invalid UInt64: %v, error: %v", value, err)
		}
		valueUint = uint64(valueInt)
	}

	out := make([]byte, 8)
	binary.BigEndian.PutUint64(out, valueUint)

	return out, nil
}

// convertInt64 converts a value of a number type to int64.
//
// a value of type uint64 has to fit in int64, otherwise an error is returned.
// a value of type float32 or float64 has to be an integer and fit in int64, otherwise an error is returned.
func convertInt64(value any, t string) (int64, error) {
	switch value := value.(type) {
	case uint8:
		return int64(value), nil
	case uint:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return value, nil
	case uint64:
		if value > 1<<63 {
			return 0, &InvalidType{
				t: t,
				v: value,
			}
		}
		return int64(value), nil
	case float32:
		if float64(value) != math.Ceil(float64(value)) || value > 1<<63 || -value > 1<<63 {
			return 0, &InvalidType{
				t: t,
				v: value,
			}
		}
		return int64(value), nil
	case float64:
		if value != math.Ceil(value) || value > 1<<63 || -value > 1<<63 {
			return 0, &InvalidType{
				t: t,
				v: value,
			}
		}
		return int64(value), nil
	default:
		return 0, &InvalidType{
			t: t,
			v: value,
		}
	}
}
