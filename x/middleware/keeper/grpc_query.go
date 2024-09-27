package keeper

import (
	"middleware/x/middleware/types"
)

var _ types.QueryServer = Keeper{}
