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
	if msg.Creator != exchange.Provider {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot settle: not the provider")
	}
	amount, _ := sdk.ParseCoinsNormalized(exchange.Amount)
	unitRatio, _ := sdk.NewIntFromString(exchange.Unitratio)

	// Calculate the amount of coins to be sent from the entity to the provider
	startAmount := amount.AmountOf(exchange.Finalunit).Quo(unitRatio)
	startAmountCoins := sdk.NewCoins(sdk.NewCoin(exchange.Startunit, startAmount))

	// Send coins from entity to provider
	err := k.bankKeeper.SendCoins(ctx, lender, borrower, startAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to send coins from entity account to provider account")
	}

	// Send coins from provider to entity
	finalAmountCoins := sdk.NewCoins(sdk.NewCoin(exchange.Finalunit, amount.AmountOf(exchange.Finalunit)))
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, finalAmountCoins)
	if err != nil {
		// Roll back startAmount transaction if sending finalAmount fails
		err = k.bankKeeper.SendCoins(ctx, lender, borrower, startAmountCoins)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "failed to roll back startAmount transaction")
		}
		return nil, sdkerrors.Wrap(err, "failed to send coins from provider account to entity account")
	}

	exchange.State = "settled"
	k.SetExchange(ctx, exchange)

	return &types.MsgSettleExchangeResponse{}, nil
}
