package keeper

import (
	"github.com/relevant-community/r3l/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
