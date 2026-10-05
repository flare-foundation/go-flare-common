//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/verification/verification.abi --pkg=verification --type=Verification --out=autogen.go

package verification
