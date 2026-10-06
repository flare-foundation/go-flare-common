//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../fumanager/fumanager.abi --pkg=fumanager --type=FUManager --out=autogen.go

package fumanager
