// Package contracts looks up items in the ABIs of the bindings in its subdirectories.
//
// The Must functions are for package-level variables and init functions, where a
// missing item is a programming error:
//
//	relayABI := contracts.MustABI(relay.RelayMetaData)
//	topic0 := contracts.MustEvent(relayABI, "SigningPolicyInitialized").ID
//
// Names are as go-ethereum parses them: an overloaded function's or event's later
// variants get a numeric suffix (foo0, foo1, ...).
package contracts

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bindv2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

var (
	errNilABI      = errors.New("nil ABI")
	errNilMetaData = errors.New("nil metadata")
)

// MetaData is a binding's metadata: an abigen v1 binding's, such as relay.RelayMetaData,
// or a pointer to an abigen v2 binding's, such as &v2/relay.RelayMetaData.
type MetaData interface {
	*bind.MetaData | *bindv2.MetaData
}

// MustABI returns the parsed ABI of md, panicking if it does not parse.
// The ABI is the one md caches, shared with the binding, so it must not be modified.
func MustABI[M MetaData](md M) *abi.ABI {
	var (
		a   *abi.ABI
		err error
	)
	switch md := any(md).(type) {
	case *bind.MetaData:
		if md == nil {
			panic(errNilMetaData)
		}
		a, err = md.GetAbi()
	case *bindv2.MetaData:
		if md == nil {
			panic(errNilMetaData)
		}
		a, err = md.ParseABI()
	}
	if err != nil {
		panic(fmt.Errorf("loading ABI: %w", err))
	}
	return a
}

// MustEvent returns the event name from a, panicking if a has none.
func MustEvent(a *abi.ABI, name string) abi.Event {
	if a == nil {
		panic(errNilABI)
	}
	return must(a.Events, "event", name)
}

// MustMethod returns the method name from a, panicking if a has none.
func MustMethod(a *abi.ABI, name string) abi.Method {
	if a == nil {
		panic(errNilABI)
	}
	return must(a.Methods, "method", name)
}

// MustError returns the custom error name from a, panicking if a has none.
func MustError(a *abi.ABI, name string) abi.Error {
	if a == nil {
		panic(errNilABI)
	}
	return must(a.Errors, "error", name)
}

func must[T any](items map[string]T, kind, name string) T {
	item, ok := items[name]
	if !ok {
		panic(fmt.Errorf("%s %s not found in ABI", kind, name))
	}
	return item
}
