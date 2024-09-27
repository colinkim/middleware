package witness

import (
	"math/rand"

	"github.com/reapchain/cosmos-sdk/baseapp"
	simappparams "github.com/reapchain/cosmos-sdk/simapp/params"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/module"
	simtypes "github.com/reapchain/cosmos-sdk/types/simulation"
	"github.com/reapchain/cosmos-sdk/x/simulation"
	"middleware/testutil/sample"
	witnesssimulation "middleware/x/witness/simulation"
	"middleware/x/witness/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = witnesssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAnchoring = "op_weight_msg_anchoring"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAnchoring int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	witnessGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&witnessGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAnchoring int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAnchoring, &weightMsgAnchoring, nil,
		func(_ *rand.Rand) {
			weightMsgAnchoring = defaultWeightMsgAnchoring
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAnchoring,
		witnesssimulation.SimulateMsgAnchoring(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
