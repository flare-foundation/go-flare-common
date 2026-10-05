//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/walletbackupmanager/walletbackupmanager.abi --pkg=walletbackupmanager --type=WalletBackupManager --out=autogen.go

package walletbackupmanager
