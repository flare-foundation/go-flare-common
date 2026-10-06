//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/teepaymentsfeeschedulemanager/teepaymentsfeeschedulemanager.abi --pkg=teepaymentsfeeschedulemanager --type=TeePaymentsFeeScheduleManager --out=autogen.go

package teepaymentsfeeschedulemanager
