// jq renames _offers: abigen v2 turns it into "offers", clashing with the method receiver.
// Remove (back to the plain directive) once abigen v2 checks parameters against the receiver name;
// still broken in go-ethereum v1.17.7.
//go:generate sh -c "jq '(.[] | select(.type==\"function\") | .inputs[]? | select(.name==\"_offers\") | .name) = \"_rewardOffers\"' ../../offers/offers.abi | go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.2 --v2 --abi=- --pkg=offers --type=Offers --out=autogen.go"

package offers
