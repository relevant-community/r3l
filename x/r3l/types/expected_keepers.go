package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
)

type (
	// OracleKeeper interface
	OracleKeeper interface {
		CreateClaim(ctx sdk.Context, claim exported.Claim)
	}
)
