package r3l

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/keeper"
	"github.com/relevant-community/r3l/x/r3l/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	//TODO
	// genState.Votes = types.SanitizeGenesisBalances(genState.Votes)
	for _, vote := range genState.Votes {
		k.CreateVote(ctx, vote)
	}
	for _, rankSource := range genState.RankSource {
		k.CreateRankSource(ctx, rankSource)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
