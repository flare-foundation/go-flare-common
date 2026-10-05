//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/walletmanager/walletmanager.abi --pkg=walletmanager --type=WalletManager --out=autogen.go

package walletmanager
