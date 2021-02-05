package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func (k Keeper) SetScores(ctx sdk.Context, scores *types.Scores) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoreKey))
	for _, score := range scores.Scores {
		score.BlockHeight = scores.BlockHeight
		b := k.cdc.MustMarshalBinaryBare(&score)
		store.Set(types.KeyPrefix(types.ScoreKey+score.Id), b)
	}
}

func (k Keeper) GetAllScore(ctx sdk.Context) (msgs []types.Score) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoreKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ScoreKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Score
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func (k Keeper) UpdateScores(ctx sdk.Context) {
	claimType := types.ScoreClaim
	pending := k.oracleKeeper.GetPendingRounds(ctx, claimType)

	// TODO sort the array first, process latest and abort once seccesfull round is found
	for _, roundID := range pending {
		result := k.oracleKeeper.TallyVotes(ctx, claimType, roundID)

		if result == nil || result.Claims[0] == nil {
			continue
		}

		// we only care about the first claimHash since we enforce that they are all the same
		claimResult := result.Claims[0]
		claimHash := claimResult.ClaimHash

		scores, ok := k.oracleKeeper.GetClaim(ctx, claimHash).(*types.Scores)
		if ok {
			k.SetScores(ctx, scores)
		}
		// TODO delete the earlier rounds also
		k.oracleKeeper.DeletePendingRound(ctx, claimType, roundID)
		return
	}
}
