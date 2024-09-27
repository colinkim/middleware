package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "middleware/testutil/keeper"
	"middleware/x/witness/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WitnessKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
