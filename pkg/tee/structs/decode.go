package structs

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Decode decodes abi encoded data and writes it to destination.
//
// dest has to be a pointer to a struct that is structured as arg describes.
func Decode(arg abi.Argument, data []byte, dest any) error {
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
