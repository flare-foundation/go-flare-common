package structs

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// DecodeTo decodes abi encoded data and writes it to destination.
//
// dest has to be a pointer to a struct that is structured as arg describes.
//
// If type T and abi are not aligned, it will panic.
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

	err = checkEncodeDecode(&args, data, decodedSlice)
	if err != nil {
		return err
	}

	err = args.Copy(wDest, decodedSlice)
	if err != nil {
		return err
	}

	return nil
}

// checkEncodeDecode encoded decodedSlice using args and compares it to data.
func checkEncodeDecode(args *abi.Arguments, data []byte, decodedSlice []any) error {
	encoded, err := args.Pack(decodedSlice...)
	if err != nil {
		return err
	}

	if !bytes.Equal(data, encoded) {
		return fmt.Errorf("initial data not equal to decoded and encoded data %v, %v", data, encoded)
	}

	return nil
}

// DecodeTo2 decodes abi encoded data and writes it to destination.
//
// dest has to be a pointer to a struct that is structured as arg describes.
//
// If abi and struct T are not aligned, it will panic.
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

	err = checkEncodeDecode(&args, data, decodedSlice)
	if err != nil {
		return err
	}

	*dest = *abi.ConvertType(decodedSlice[0], new(T)).(*T)
	return nil
}

// Decode decodes abi encoded data and converts it to a go struct of type T.
//
// If type T and abi are not aligned, it will panic.
func Decode[T any](arg abi.Argument, data []byte) (T, error) {
	var t T
	args := abi.Arguments{arg}

	decodedSlice, err := args.Unpack(data)
	if err != nil {
		return t, err
	}

	err = checkEncodeDecode(&args, data, decodedSlice)
	if err != nil {
		return t, err
	}

	t = *abi.ConvertType(decodedSlice[0], new(T)).(*T)

	return t, nil
}
