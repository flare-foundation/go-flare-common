//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../contractregistry/contractregistry.abi --pkg=contractregistry --type=ContractRegistry --out=autogen.go

package contractregistry
