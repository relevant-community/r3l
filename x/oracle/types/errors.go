package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	ErrInvalidClaim      = sdkerrors.Register(ModuleName, 2, "invalid claim")
	ErrNoClaimExists     = sdkerrors.Register(ModuleName, 3, "no claim exits")
	ErrNoClaimTypeExists = sdkerrors.Register(ModuleName, 4, "claim type is not registered as part of the oracle params")
)
