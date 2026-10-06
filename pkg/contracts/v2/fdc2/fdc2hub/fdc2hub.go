//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../fdc2/fdc2hub/fdc2hub.abi --pkg=fdc2hub --type=Fdc2Hub --out=autogen.go

package fdc2hub
