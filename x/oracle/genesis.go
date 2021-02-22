package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	exported "github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/keeper"
	"github.com/relevant-community/r3l/x/oracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	k.SetParams(ctx, genState.Params)

	for _, round := range genState.Rounds {
		k.CreateRound(ctx, round)
	}

	for _, c := range genState.Claims {
		claim, ok := c.GetCachedValue().(exported.Claim)
		if !ok {
			panic("expected claim")
		}
		if cliamExists := k.GetClaim(ctx, claim.Hash()); cliamExists != nil {
			panic(fmt.Sprintf("claim with hash %s already exists", claim.Hash()))
		}

		k.CreateClaim(ctx, claim)
	}

	for claimType, pendingRounds := range genState.Pending {
		for _, roundID := range pendingRounds.Pending {
			k.AddPendingRound(ctx, claimType, roundID)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	rounds := k.GetAllRounds(ctx)
	claims := k.GetAllClaims(ctx)
	pending := k.GetAllPendingRounds(ctx)
	return types.NewGenesisState(params, rounds, claims, pending)
}
