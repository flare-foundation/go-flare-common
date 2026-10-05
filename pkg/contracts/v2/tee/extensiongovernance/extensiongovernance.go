//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/extensiongovernance/extensiongovernance.abi --pkg=extensiongovernance --type=ExtensionGovernance --out=autogen.go

package extensiongovernance
