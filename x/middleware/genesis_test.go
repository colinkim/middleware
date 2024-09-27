package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "middleware/testutil/keeper"
	"middleware/testutil/nullify"
	"middleware/x/middleware"
	"middleware/x/middleware/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MiddlewareKeeper(t)
	middleware.InitGenesis(ctx, *k, genesisState)
	got := middleware.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
