//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../entitymanager/entitymanager.abi --pkg=entitymanager --type=EntityManager --out=autogen.go

package entitymanager
