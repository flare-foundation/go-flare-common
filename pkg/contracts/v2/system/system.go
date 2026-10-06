//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../system/system.abi --pkg=system --type=FlareSystemsManager --out=autogen.go

package system
