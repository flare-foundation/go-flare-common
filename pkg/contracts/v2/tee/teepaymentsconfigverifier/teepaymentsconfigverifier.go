//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/teepaymentsconfigverifier/teepaymentsconfigverifier.abi --pkg=teepaymentsconfigverifier --type=TeePaymentsConfigVerifier --out=autogen.go

package teepaymentsconfigverifier
