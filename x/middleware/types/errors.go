package types

// DONTCOVER

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// x/middleware module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
