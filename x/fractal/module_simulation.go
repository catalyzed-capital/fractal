package fractal

import (
	"math/rand"

	"fractal/testutil/sample"
	fractalsimulation "fractal/x/fractal/simulation"
	"fractal/x/fractal/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = fractalsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestExchange = "op_weight_msg_request_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestExchange int = 100

	opWeightMsgApproveExchange = "op_weight_msg_approve_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveExchange int = 100

	opWeightMsgSettleExchange = "op_weight_msg_settle_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSettleExchange int = 100

	opWeightMsgSwapExchange = "op_weight_msg_swap_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwapExchange int = 100

	opWeightMsgCancelExchange = "op_weight_msg_cancel_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelExchange int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	fractalGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&fractalGenesis)
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

	var weightMsgRequestExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestExchange, &weightMsgRequestExchange, nil,
		func(_ *rand.Rand) {
			weightMsgRequestExchange = defaultWeightMsgRequestExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestExchange,
		fractalsimulation.SimulateMsgRequestExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveExchange, &weightMsgApproveExchange, nil,
		func(_ *rand.Rand) {
			weightMsgApproveExchange = defaultWeightMsgApproveExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveExchange,
		fractalsimulation.SimulateMsgApproveExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSettleExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSettleExchange, &weightMsgSettleExchange, nil,
		func(_ *rand.Rand) {
			weightMsgSettleExchange = defaultWeightMsgSettleExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSettleExchange,
		fractalsimulation.SimulateMsgSettleExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwapExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwapExchange, &weightMsgSwapExchange, nil,
		func(_ *rand.Rand) {
			weightMsgSwapExchange = defaultWeightMsgSwapExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwapExchange,
		fractalsimulation.SimulateMsgSwapExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelExchange, &weightMsgCancelExchange, nil,
		func(_ *rand.Rand) {
			weightMsgCancelExchange = defaultWeightMsgCancelExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelExchange,
		fractalsimulation.SimulateMsgCancelExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
