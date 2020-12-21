package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// func (k Keeper) CreateClaim(ctx sdk.Context, claim types.MsgClaim) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
// 	b := k.cdc.MustMarshalBinaryBare(&claim)
// 	store.Set(types.KeyPrefix(types.ClaimKey), b)
// }

// CreateClaim sets Claim by hash in the module's KVStore.
func (k Keeper) CreateClaim(ctx sdk.Context, claim exported.Claim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	store.Set(claim.Hash(), k.MustMarshalClaim(claim))
}

// GetClaim retrieves Claim by hash if it exists. If no Claim exists for
// the given hash, (nil, false) is returned.
func (k Keeper) GetClaim(ctx sdk.Context, hash tmbytes.HexBytes) (exported.Claim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))

	bz := store.Get(hash)
	if len(bz) == 0 {
		return nil, false
	}

	return k.MustUnmarshalClaim(bz), true
}

// GetAllClaim returns all claim
func (k Keeper) GetAllClaim(ctx sdk.Context) (msgs []exported.Claim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ClaimKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		// k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		claim := k.MustUnmarshalClaim(iterator.Value())
		fmt.Println(claim)
		msgs = append(msgs, claim)
	}

	return
}

// MustUnmarshalClaim attempts to decode and return an Claim object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalClaim(bz []byte) exported.Claim {
	Claim, err := k.UnmarshalClaim(bz)
	if err != nil {
		panic(fmt.Errorf("failed to decode Claim: %w", err))
	}

	return Claim
}

// MustMarshalClaim attempts to encode an Claim object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalClaim(Claim exported.Claim) []byte {
	bz, err := k.MarshalClaim(Claim)
	if err != nil {
		panic(fmt.Errorf("failed to encode Claim: %w", err))
	}

	return bz
}

// MarshalClaim marshals a Claim interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalClaim(claimI exported.Claim) ([]byte, error) {
	return codec.MarshalAny(k.cdc, claimI)
}

// UnmarshalClaim returns a Claim interface from raw encoded Claim
// bytes of a Proto-based Claim type. An error is returned upon decoding
// failure.
func (k Keeper) UnmarshalClaim(bz []byte) (exported.Claim, error) {
	var claim exported.Claim
	if err := codec.UnmarshalAny(k.cdc, &claim, bz); err != nil {
		return nil, err
	}

	return claim, nil
}
