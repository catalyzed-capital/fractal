package keeper

import (
	"context"

	"fractal/x/balancecheck/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) ShowUsdcBalance(goCtx context.Context, req *types.QueryShowUsdcBalanceRequest) (*types.QueryShowUsdcBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var usdcbalances []types.Usdcbalance
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	balanceStore := prefix.NewStore(store, types.KeyPrefix(types.UsdcKey))


	pageRes, err := query.Paginate(balanceStore, req.Pagination, func(key []byte, value []byte) error{
	
		var usdcbalance types.Usdcbalance
		if err := k.cdc.Unmarshal(value, &usdcbalance); err != nil{
			return err
		}

		usdcbalances = append(usdcbalances, usdcbalance)
		return nil

	})

	if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }
	
	return &types.QueryShowUsdcBalanceResponse{Usdcbalance: usdcbalances, Pagination: pageRes}, nil
}
