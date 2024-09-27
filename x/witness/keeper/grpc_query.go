package keeper

import (
	"middleware/x/witness/types"
)

var _ types.QueryServer = Keeper{}
