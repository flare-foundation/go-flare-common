//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=../../../tee/vrfverifier/vrfverifier.abi --bin=../../../tee/vrfverifier/vrfverifier.bin --pkg=vrfverifier --type=VRFVerifier --out=autogen.go

package vrfverifier
