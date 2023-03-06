package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"fractal/x/fractal/types"
)

func (k msgServer) CancelExchange(goCtx context.Context, msg *types.MsgCancelExchange) (*types.MsgCancelExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	exchange, found := k.GetExchange(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if exchange.Entity != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the entity")
	}
	if exchange.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongExchangeState, "%v", exchange.State)
	}
	borrower, _ := sdk.AccAddressFromBech32(exchange.Entity)
	amount, _ := sdk.ParseCoinsNormalized(exchange.Amount)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, amount)
	if err != nil {
		return nil, err
	}
	exchange.State = "cancelled"
	k.SetExchange(ctx, exchange)
	return &types.MsgCancelExchangeResponse{}, nil
}
