package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/types"
)

// SetValidatorDelegateAddress sets a new address that will have the power to send data on behalf of the validator
func (k Keeper) SetValidatorDelegateAddress(ctx sdk.Context, val, del sdk.AccAddress) {
	ctx.KVStore(k.storeKey).Set(types.GetFeedDelegateKey(del), val.Bytes())
}

// GetValidatorAddressFromDelegate returns the delegate address for a given validator
func (k Keeper) GetValidatorAddressFromDelegate(ctx sdk.Context, del sdk.AccAddress) sdk.AccAddress {
	return sdk.AccAddress(ctx.KVStore(k.storeKey).Get(types.GetFeedDelegateKey(del)))
}

// GetDelegateAddressFromValidator returns the validator address for a given delegate
func (k Keeper) GetDelegateAddressFromValidator(ctx sdk.Context, val sdk.AccAddress) sdk.AccAddress {
	var address sdk.AccAddress
	// TODO: create secondary index
	k.IterateDelegateAddresses(ctx, func(delegatorAddr, validatorAddr sdk.AccAddress) bool {
		if !val.Equals(validatorAddr) {
			return false
		}

		address = delegatorAddr
		return true
	})
	return address
}

// IsDelegateAddress returns true if the validator has delegated their feed to an address
func (k Keeper) IsDelegateAddress(ctx sdk.Context, del sdk.AccAddress) bool {
	return ctx.KVStore(k.storeKey).Has(types.GetFeedDelegateKey(del))
}

// IterateDelegateAddresses iterates over all delegate address pairs in the store
func (k Keeper) IterateDelegateAddresses(ctx sdk.Context, handler func(del, val sdk.AccAddress) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.FeedDelegateKeyPrefix))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		del := sdk.AccAddress(bytes.TrimPrefix(iter.Key(), types.KeyPrefix(types.FeedDelegateKeyPrefix)))
		val := sdk.AccAddress(iter.Value())
		if handler(del, val) {
			break
		}
	}
}
