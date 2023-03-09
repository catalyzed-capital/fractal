package keeper

import (
	"context"
	"fmt"
	"fractal/x/balancecheck/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UsdcBalanceCheck(goCtx context.Context, msg *types.MsgUsdcBalanceCheck) (*types.MsgUsdcBalanceCheckResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var usdcbalance types.Usdcbalance
	balance := k.AppendUsdcBalance(ctx, usdcbalance)
	
	fmt.Println(balance)

	return &types.MsgUsdcBalanceCheckResponse{}, nil
}
