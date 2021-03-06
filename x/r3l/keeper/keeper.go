package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// appparams "github.com/relevant-community/r3l/app/params"
	"github.com/relevant-community/r3l/x/r3l/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		oracleKeeper  types.OracleKeeper
		stakingKeeper types.StakingKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, oracleKeeper types.OracleKeeper, stakingKeeper types.StakingKeeper) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		oracleKeeper:  oracleKeeper,
		stakingKeeper: stakingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
