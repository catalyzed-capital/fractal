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
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(exchange.Entity)
	amount, err := sdk.ParseCoinsNormalized(exchange.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongExchangeState, "Cannot parse coins in exchange amount")
	}
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	exchange.Provider = msg.Creator
	exchange.State = "approved"
	k.SetExchange(ctx, exchange)
	return &types.MsgApproveExchangeResponse{}, nil
}
