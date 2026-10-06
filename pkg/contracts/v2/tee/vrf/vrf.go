//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/vrf/vrf.abi --pkg=vrf --type=VRF --out=autogen.go

package vrf
