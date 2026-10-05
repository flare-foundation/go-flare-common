//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/teepayments/teepayments.abi --pkg=teepayments --type=TeePayments --out=autogen.go

package teepayments
