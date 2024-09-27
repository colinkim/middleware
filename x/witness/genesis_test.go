package witness_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "middleware/testutil/keeper"
	"middleware/testutil/nullify"
	"middleware/x/witness"
	"middleware/x/witness/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WitnessKeeper(t)
	witness.InitGenesis(ctx, *k, genesisState)
	got := witness.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
