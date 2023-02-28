package fractal_test

import (
	"testing"

	keepertest "fractal/testutil/keeper"
	"fractal/testutil/nullify"
	"fractal/x/fractal"
	"fractal/x/fractal/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FractalKeeper(t)
	fractal.InitGenesis(ctx, *k, genesisState)
	got := fractal.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
