package keeper

import (
	"context"
	"fractal/x/balancecheck/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UsdcBalanceCheck(goCtx context.Context, msg *types.MsgUsdcBalanceCheck) (*types.MsgUsdcBalanceCheckResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var usdcbalance = types.Usdcbalance{
		Creator: msg.Creator,
		Chain: msg.Chain,
		Requiredamount: msg.Amount,
	}

	id := k.AppendUsdcBalance(
		ctx,
		usdcbalance,
	)
	

	return &types.MsgUsdcBalanceCheckResponse{
		Id: id,
	}, nil
}
