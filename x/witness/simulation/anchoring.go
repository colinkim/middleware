package simulation

import (
	"math/rand"

	"github.com/reapchain/cosmos-sdk/baseapp"
	sdk "github.com/reapchain/cosmos-sdk/types"
	simtypes "github.com/reapchain/cosmos-sdk/types/simulation"
	"middleware/x/witness/keeper"
	"middleware/x/witness/types"
)

func SimulateMsgAnchoring(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAnchoring{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Anchoring simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Anchoring simulation not implemented"), nil, nil
	}
}
