package balancecheck_test

import (
	"testing"

	keepertest "fractal/testutil/keeper"
	"fractal/testutil/nullify"
	"fractal/x/balancecheck"
	"fractal/x/balancecheck/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BalancecheckKeeper(t)
	balancecheck.InitGenesis(ctx, *k, genesisState)
	got := balancecheck.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
