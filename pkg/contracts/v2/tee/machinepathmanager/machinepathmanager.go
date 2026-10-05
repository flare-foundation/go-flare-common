//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/machinepathmanager/machinepathmanager.abi --pkg=machinepathmanager --type=MachinePathManager --out=autogen.go

package machinepathmanager
