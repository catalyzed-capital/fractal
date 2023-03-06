package simulation

import (
	"math/rand"

	"fractal/x/fractal/keeper"
	"fractal/x/fractal/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgApproveExchange(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgApproveExchange{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ApproveExchange simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ApproveExchange simulation not implemented"), nil, nil
	}
}
