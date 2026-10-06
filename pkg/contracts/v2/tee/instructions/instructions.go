//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/instructions/instructions.abi --pkg=instructions --type=Instructions --out=autogen.go

package instructions
