package keeper

import (
	"context"

	"fractal/x/fractal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SettleExchange(goCtx context.Context, msg *types.MsgSettleExchange) (*types.MsgSettleExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	exchange, found := k.GetExchange(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if exchange.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongExchangeState, "%v", exchange.State)
	}
	lender, _ := sdk.AccAddressFromBech32(exchange.Provider)
	borrower, _ := sdk.AccAddressFromBech32(exchange.Entity)
	if msg.Creator != exchange.Entity {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot settle: not the entity")
	}
	amount, _ := sdk.ParseCoinsNormalized(exchange.Amount)
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	exchange.State = "settled"
	k.SetExchange(ctx, exchange)
	return &types.MsgSettleExchangeResponse{}, nil
}
