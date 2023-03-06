package keeper_test

import (
	"testing"

	keepertest "fractal/testutil/keeper"
	"fractal/testutil/nullify"
	"fractal/x/fractal/keeper"
	"fractal/x/fractal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNExchange(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Exchange {
	items := make([]types.Exchange, n)
	for i := range items {
		items[i].Id = keeper.AppendExchange(ctx, items[i])
	}
	return items
}

func TestExchangeGet(t *testing.T) {
	keeper, ctx := keepertest.FractalKeeper(t)
	items := createNExchange(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetExchange(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestExchangeRemove(t *testing.T) {
	keeper, ctx := keepertest.FractalKeeper(t)
	items := createNExchange(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExchange(ctx, item.Id)
		_, found := keeper.GetExchange(ctx, item.Id)
		require.False(t, found)
	}
}

func TestExchangeGetAll(t *testing.T) {
	keeper, ctx := keepertest.FractalKeeper(t)
	items := createNExchange(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExchange(ctx)),
	)
}

func TestExchangeCount(t *testing.T) {
	keeper, ctx := keepertest.FractalKeeper(t)
	items := createNExchange(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetExchangeCount(ctx))
}
