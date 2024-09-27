package keeper

import (
	"context"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"middleware/x/witness/types"
)

func (k msgServer) Anchoring(goCtx context.Context, msg *types.MsgAnchoring) (*types.MsgAnchoringResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAnchoringResponse{}, nil
}
