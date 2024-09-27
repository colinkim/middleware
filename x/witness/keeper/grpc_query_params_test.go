package keeper_test

import (
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "middleware/testutil/keeper"
	"middleware/x/witness/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.WitnessKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
