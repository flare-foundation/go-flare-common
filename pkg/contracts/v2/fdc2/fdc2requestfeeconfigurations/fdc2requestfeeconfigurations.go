// jq renames _type: abigen v2 turns it into the Go keyword "type".
// Remove (back to the plain directive) once abigen v2 checks for keywords after renaming;
// still broken in go-ethereum v1.17.7.
//go:generate sh -c "jq '(.[] | select(.type==\"function\") | .inputs[]? | select(.name==\"_type\") | .name) = \"_typ\"' ../../../fdc2/fdc2requestfeeconfigurations/fdc2requestfeeconfigurations.abi | go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=- --pkg=fdc2requestfeeconfigurations --type=Fdc2RequestFeeConfigurations --out=autogen.go"

package fdc2requestfeeconfigurations
