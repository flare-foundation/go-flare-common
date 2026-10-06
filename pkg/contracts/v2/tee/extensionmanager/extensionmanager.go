//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/extensionmanager/extensionmanager.abi --pkg=extensionmanager --type=ExtensionManager --out=autogen.go

package extensionmanager
