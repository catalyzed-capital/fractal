package keeper

import (
	"context"

	"fractal/x/fractal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExchangeAll(goCtx context.Context, req *types.QueryAllExchangeRequest) (*types.QueryAllExchangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exchanges []types.Exchange
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	exchangeStore := prefix.NewStore(store, types.KeyPrefix(types.ExchangeKey))

	pageRes, err := query.Paginate(exchangeStore, req.Pagination, func(key []byte, value []byte) error {
		var exchange types.Exchange
		if err := k.cdc.Unmarshal(value, &exchange); err != nil {
			return err
		}

		exchanges = append(exchanges, exchange)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExchangeResponse{Exchange: exchanges, Pagination: pageRes}, nil
}

func (k Keeper) Exchange(goCtx context.Context, req *types.QueryGetExchangeRequest) (*types.QueryGetExchangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	exchange, found := k.GetExchange(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetExchangeResponse{Exchange: exchange}, nil
}
