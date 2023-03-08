package keeper

import (
	"context"

	"fractal/x/fractal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SwapExchange(goCtx context.Context, msg *types.MsgSwapExchange) (*types.MsgSwapExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message, function to change the provider on the exchange such that a new provider can settle
	_ = ctx

	return &types.MsgSwapExchangeResponse{}, nil
}
