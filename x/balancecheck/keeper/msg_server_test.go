package keeper_test

import (
	"context"
	"testing"

	keepertest "fractal/testutil/keeper"
	"fractal/x/balancecheck/keeper"
	"fractal/x/balancecheck/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BalancecheckKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
