package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keys for oracle store, with <prefix><key> -> <value>
const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// RoundKey defines the oracle vote store key
	RoundKey = "vote"

	// ClaimKey defines the oracle claim store key
	ClaimKey = "Claim"

	// PendingRoundKey defines the oracle claim store key
	PendingRoundKey = "PendingRound"

	// - Delegation<val_address> -> <delegate_address>
	FeedDelegateKeyPrefix = "Delegation" // key for validator feed delegation
)

// KeyPrefix helper
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetFeedDelegateKey returns the validator for a given delegate key
func GetFeedDelegateKey(del sdk.AccAddress) []byte {
	return append(KeyPrefix(FeedDelegateKeyPrefix), del.Bytes()...)
}

// GetRoundKey helper
func GetRoundKey(claimType string, roundID uint64) string {
	return claimType + "_" + strconv.FormatUint(roundID, 10)
}

// RoundPrefix helper
func RoundPrefix(claimType string, roundID uint64) []byte {
	return KeyPrefix(GetRoundKey(claimType, roundID))
}
