package simulation

import (
	"bytes"
	"fmt"

	"github.com/Finschia/finschia-rdk/codec"
	"github.com/Finschia/finschia-rdk/types/kv"
	"github.com/Finschia/finschia-rdk/x/feegrant"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding feegrant type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], feegrant.FeeAllowanceKeyPrefix):
			var grantA, grantB feegrant.Grant
			cdc.MustUnmarshal(kvA.Value, &grantA)
			cdc.MustUnmarshal(kvB.Value, &grantB)
			return fmt.Sprintf("%v\n%v", grantA, grantB)
		default:
			panic(fmt.Sprintf("invalid feegrant key %X", kvA.Key))
		}
	}
}
