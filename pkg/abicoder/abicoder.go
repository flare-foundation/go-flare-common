package abicoder

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/panics"
)

type Coder[T any] struct {
	args abi.Arguments
}

func New[T any](arg abi.Argument) (*Coder[T], error) {
	if err := checkArgType(arg.Type); err != nil {
		return nil, err
	}

	if err := checkShape[T](arg.Type); err != nil {
		return nil, err
	}

	return &Coder[T]{args: abi.Arguments{arg}}, nil
}

func MustNew[T any](arg abi.Argument) *Coder[T] {
	c, err := New[T](arg)
	if err != nil {
		panic(err)
	}

	return c
}

func (c *Coder[T]) Encode(data T) (encoded []byte, err error) {
	defer panics.Recover(&err)

	if err := checkRange(&c.args[0].Type, reflect.ValueOf(data)); err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	encoded, err = c.args.Pack(data)
	if err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	if err := checkLayout(c.args, encoded); err != nil {
		return nil, fmt.Errorf("encoding type %T: %w", data, err)
	}

	return encoded, nil
}

func (c *Coder[T]) Decode(data []byte) (t T, err error) {
	defer panics.Recover(&err)

	if err := checkLayout(c.args, data); err != nil {
		return t, err
	}

	values, err := c.args.Unpack(data)
	if err != nil {
		return t, err
	}

	if err := checkEncodeDecode(&c.args, data, values); err != nil {
		return t, err
	}

	x := abi.ConvertType(values[0], new(T))

	tp, ok := x.(*T)
	if !ok {
		return t, errors.New("invalid type assertion")
	}
	t = *tp

	return t, nil
}

func (c *Coder[T]) DecodeTo(data []byte, dest *T) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errNonPointer
	}

	t, err := c.Decode(data)
	if err != nil {
		return err
	}

	*dest = t

	return nil
}
