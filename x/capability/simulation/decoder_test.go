package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/line/lbm-sdk/v2/simapp"
	sdk "github.com/line/lbm-sdk/v2/types"
	"github.com/line/lbm-sdk/v2/types/kv"
	"github.com/line/lbm-sdk/v2/x/capability/simulation"
	"github.com/line/lbm-sdk/v2/x/capability/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := simapp.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	capOwners := types.CapabilityOwners{
		Owners: []types.Owner{{Module: "transfer", Name: "ports/transfer"}},
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   types.KeyIndex,
				Value: sdk.Uint64ToBigEndian(10),
			},
			{
				Key:   types.KeyPrefixIndexCapability,
				Value: cdc.MustMarshalBinaryBare(&capOwners),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Index", "Index A: 10\nIndex B: 10\n"},
		{"CapabilityOwners", fmt.Sprintf("CapabilityOwners A: %v\nCapabilityOwners B: %v\n", capOwners, capOwners)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
