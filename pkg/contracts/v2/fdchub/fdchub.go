//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../fdchub/fdchub.abi --pkg=fdchub --type=FdcHub --out=autogen.go

package fdchub
