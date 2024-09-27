package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	keepertest "middleware/testutil/keeper"
	"middleware/x/witness/keeper"
	"middleware/x/witness/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WitnessKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
