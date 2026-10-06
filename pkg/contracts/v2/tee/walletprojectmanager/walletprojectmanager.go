//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/walletprojectmanager/walletprojectmanager.abi --pkg=walletprojectmanager --type=WalletProjectManager --out=autogen.go

package walletprojectmanager
