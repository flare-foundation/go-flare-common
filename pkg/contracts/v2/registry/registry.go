//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../registry/registry.abi --pkg=registry --type=Registry --out=autogen.go

package registry
