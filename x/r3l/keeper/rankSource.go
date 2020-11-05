package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func (k Keeper) CreateRankSource(ctx sdk.Context, rankSource types.MsgRankSource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RankSourceKey))
	b := k.cdc.MustMarshalBinaryBare(&rankSource)
	store.Set(types.KeyPrefix(types.RankSourceKey), b)
}

func (k Keeper) GetAllRankSource(ctx sdk.Context) (msgs []types.MsgRankSource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RankSourceKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.RankSourceKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgRankSource
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
