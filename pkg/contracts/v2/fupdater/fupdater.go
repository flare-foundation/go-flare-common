//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../fupdater/fupdater.abi --pkg=fupdater --type=FUpdater --out=autogen.go

package fupdater
