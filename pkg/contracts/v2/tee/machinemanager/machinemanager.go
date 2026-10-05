//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/machinemanager/machinemanager.abi --pkg=machinemanager --type=MachineManager --out=autogen.go

package machinemanager
