//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/walletkeymanager/walletkeymanager.abi --pkg=walletkeymanager --type=WalletKeyManager --out=autogen.go

package walletkeymanager
