package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"fractal/x/fractal/types"
)

func (k msgServer) ApproveExchange(goCtx context.Context, msg *types.MsgApproveExchange) (*types.MsgApproveExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	exchange, found := k.GetExchange(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if exchange.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongExchangeState, "%v", exchange.State)
	}

	exchange.Provider = msg.Creator
	exchange.State = "approved"
	k.SetExchange(ctx, exchange)

	return &types.MsgApproveExchangeResponse{}, nil
}
