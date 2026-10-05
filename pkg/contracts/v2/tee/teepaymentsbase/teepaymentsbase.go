//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/teepaymentsbase/teepaymentsbase.abi --pkg=teepaymentsbase --type=TeePaymentsBase --out=autogen.go

package teepaymentsbase
