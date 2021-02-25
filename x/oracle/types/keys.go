package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
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
)

// Keys for oracle store, with <prefix><key> -> <value>
var (
	// RoundKey defines the oracle vote store key
	RoundKey = []byte{0x00}

	// VoteKey defines the oracle claim store key
	VoteKey = []byte{0x01}

	// PendingRoundKey defines the oracle claim store key
	PendingRoundKey = []byte{0x02}

	// - Delegation<val_address> -> <delegate_address>
	FeedDelegateKey = []byte{0x03} // key for validator feed delegation

	// - PrevoteKey <prevote_hash> -> <prevote_hash>
	PrevoteKey = []byte{0x04}
)

// KeyPrefix helper
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetFeedDelegateKey returns the validator for a given delegate key
func GetFeedDelegateKey(del sdk.AccAddress) []byte {
	return append(FeedDelegateKey, del.Bytes()...)
}

// GetClaimPrevoteKey returns the key for a validators prevote
func GetClaimPrevoteKey(hash tmbytes.HexBytes) []byte {
	return append(PrevoteKey, hash...)
}

// GetRoundKey helper
func GetRoundKey(claimType string, roundID uint64) string {
	return claimType + "_" + strconv.FormatUint(roundID, 10)
}

// RoundPrefix helper
func RoundPrefix(claimType string, roundID uint64) []byte {
	return KeyPrefix(GetRoundKey(claimType, roundID))
}
