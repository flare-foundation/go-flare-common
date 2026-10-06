//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/machineemergencypause/machineemergencypause.abi --pkg=machineemergencypause --type=MachineEmergencyPause --out=autogen.go

package machineemergencypause
