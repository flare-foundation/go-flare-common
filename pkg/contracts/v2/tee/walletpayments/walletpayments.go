//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/walletpayments/walletpayments.abi --pkg=walletpayments --type=WalletPayments --out=autogen.go

package walletpayments
