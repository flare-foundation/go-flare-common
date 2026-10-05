//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../submission/submission.abi --pkg=submission --type=Submission --out=autogen.go

package submission
