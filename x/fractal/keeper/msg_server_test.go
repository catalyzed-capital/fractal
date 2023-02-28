package keeper_test

import (
	"context"
	"testing"

	keepertest "fractal/testutil/keeper"
	"fractal/x/fractal/keeper"
	"fractal/x/fractal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FractalKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
