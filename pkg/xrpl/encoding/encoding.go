// Package encoding re-exports XRPL binary serialization functions from the types subpackage.
package encoding

import "github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"

var Encode = types.Encode
var Decode = types.Decode
