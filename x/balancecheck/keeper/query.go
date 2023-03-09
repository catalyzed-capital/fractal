package keeper

import (
	"fractal/x/balancecheck/types"
)

var _ types.QueryServer = Keeper{}
