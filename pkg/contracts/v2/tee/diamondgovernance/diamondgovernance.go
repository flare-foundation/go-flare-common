//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/diamondgovernance/diamondgovernance.abi --pkg=diamondgovernance --type=DiamondGovernance --out=autogen.go

package diamondgovernance
