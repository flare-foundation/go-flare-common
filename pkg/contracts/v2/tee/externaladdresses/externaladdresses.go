//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/externaladdresses/externaladdresses.abi --pkg=externaladdresses --type=ExternalAddresses --out=autogen.go

package externaladdresses
