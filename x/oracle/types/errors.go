package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	ErrInvalidClaim = sdkerrors.Register(ModuleName, 2, "invalid claim")
)
