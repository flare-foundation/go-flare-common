package structs

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// DecodeTo decodes abi encoded data and writes it to destination.
//
// dest has to be a pointer to a struct that is structured as arg describes.
func DecodeTo[T any](arg abi.Argument, data []byte, dest *T) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("dest is not a non nil pointer")
	}

	var wDest any

	if rv.Elem().Kind() == reflect.Struct {
		wDest = &struct{ X any }{dest}
	} else {
		wDest = dest
	}

	args := abi.Arguments{arg}

	decodedSlice, err := args.Unpack(data)
	if err != nil {
		return err
	}

	err = args.Copy(wDest, decodedSlice)
	if err != nil {
		return err
	}

	return nil
}

// DecodeTo2 decodes abi encoded data and writes it to destination.
//
// dest has to be a pointer to a struct that is structured as arg describes.
func DecodeTo2[T any](arg abi.Argument, data []byte, dest *T) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("dest is not a non nil pointer")
	}

	args := abi.Arguments{arg}

	decodedSlice, err := args.Unpack(data)
	if err != nil {
		return err
	}

	*dest = *abi.ConvertType(decodedSlice[0], new(T)).(*T)
	return nil
}

// DecodeTo decodes abi encoded data and returns it as
func Decode[T any](arg abi.Argument, data []byte) (T, error) {
	var t T
	args := abi.Arguments{arg}

	decodedSlice, err := args.Unpack(data)
	if err != nil {
		return t, err
	}

	t = *abi.ConvertType(decodedSlice[0], new(T)).(*T)

	return t, nil
}
