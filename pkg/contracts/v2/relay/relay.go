//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../relay/relay.abi --pkg=relay --type=Relay --out=autogen.go

package relay
