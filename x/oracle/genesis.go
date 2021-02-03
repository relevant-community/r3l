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
	k.SetParams(ctx, genState.Params)

	for _, roundVote := range genState.RoundVotes {
		k.SetRoundVote(ctx, roundVote)
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

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	roundVotes := k.GetRoundVotes(ctx)
	claims := k.GetAllClaim(ctx)
	return types.NewGenesisState(params, roundVotes, claims)
}
