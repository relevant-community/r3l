package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func (k Keeper) CreateNamespace(ctx sdk.Context, namespace types.MsgNamespace) {
	namespace.Id = k.GetNextNamespaceId(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamespaceKey))
	b := k.cdc.MustMarshalBinaryBare(&namespace)
	store.Set(types.KeyPrefix(types.NamespaceKey+string(namespace.Id)), b)
}

func (k Keeper) GetAllNamespace(ctx sdk.Context) (msgs []types.MsgNamespace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamespaceKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.NamespaceKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgNamespace
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func (k Keeper) GetNextNamespaceId(ctx sdk.Context) uint64 {
	var nextId uint64
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextNamespaceIdKey))

	bz := store.Get(types.KeyPrefix(types.NextNamespaceIdKey))

	if bz == nil {
		// initialize the account numbers
		nextId = 0
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.UnmarshalBinaryBare(bz, &val)
		if err != nil {
			panic(err)
		}

		nextId = val.GetValue()
	}

	bz = k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: nextId + 1})
	store.Set(types.KeyPrefix(types.NextNamespaceIdKey), bz)

	return nextId
}
