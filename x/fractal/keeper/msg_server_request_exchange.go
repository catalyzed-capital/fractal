package keeper

import (
	"context"

	"fractal/x/fractal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestExchange(goCtx context.Context, msg *types.MsgRequestExchange) (*types.MsgRequestExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var exchange = types.Exchange{
		Finalunit:  msg.Finalunit,
		Amount:     msg.Amount,
		Startunit:  msg.Startunit,
		Unitratio:  msg.Unitratio,
		Settledate: msg.Settledate,
		State:      "requested",
		Entity:     msg.Creator,
	}
	k.AppendExchange(ctx, exchange)
	return &types.MsgRequestExchangeResponse{}, nil
}
