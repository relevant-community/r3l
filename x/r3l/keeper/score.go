package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func (k Keeper) CreateScore(ctx sdk.Context, score types.MsgScore) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoreKey))
	b := k.cdc.MustMarshalBinaryBare(&score)
	store.Set(types.KeyPrefix(types.ScoreKey+score.Id), b)
}

func (k Keeper) SetScores(ctx sdk.Context, scoresMsg types.MsgScores) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoreKey))
	// for _, score := range scoresMsg.Scores {
	// 	score.BlockHeight = scoresMsg.BlockHeight
	// 	b := k.cdc.MustMarshalBinaryBare(&score)
	// 	store.Set(types.KeyPrefix(types.ScoreKey+score.Id), b)
	// }
	// k.oracleKeeper.CreateClaim(ctx, &scoresMsg)
}

func (k Keeper) GetAllScore(ctx sdk.Context) (msgs []types.MsgScore) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoreKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ScoreKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgScore
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
