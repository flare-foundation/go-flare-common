//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../preregistry/preregistry.abi --pkg=preregistry --type=Preregistry --out=autogen.go

package preregistry
