//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../calculator/calculator.abi --pkg=calculator --type=Calculator --out=autogen.go

package calculator
