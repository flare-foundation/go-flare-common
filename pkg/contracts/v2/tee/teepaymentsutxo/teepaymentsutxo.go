//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/teepaymentsutxo/teepaymentsutxo.abi --pkg=teepaymentsutxo --type=TeePaymentsUtxo --out=autogen.go

package teepaymentsutxo
