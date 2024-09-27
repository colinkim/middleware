package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	keepertest "middleware/testutil/keeper"
	"middleware/x/middleware/keeper"
	"middleware/x/middleware/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MiddlewareKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
