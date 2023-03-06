package keeper

import (
	"context"

	"fractal/x/fractal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SwapExchange(goCtx context.Context, msg *types.MsgSwapExchange) (*types.MsgSwapExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapExchangeResponse{}, nil
}
