package keeper

import (
	"context"

	"fractal/x/balancecheck/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UsdcBalanceCheck(goCtx context.Context, msg *types.MsgUsdcBalanceCheck) (*types.MsgUsdcBalanceCheckResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUsdcBalanceCheckResponse{}, nil
}
