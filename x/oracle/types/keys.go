package types

import "strconv"

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

	// VoteKey defines the oracle vote store key
	VoteKey = "vote"

	// ResultKey defines the oracle result store key
	ResultKey = "result"

	// ClaimKey defines the oracle claim store key
	ClaimKey = "Claim"

	// PendingRoundKey defines the oracle claim store key
	PendingRoundKey = "PendingRound"

	// SuccessRoundKey defines the oracle claim store key
	SuccessRoundKey = "SuccessRound"
)

// KeyPrefix helper
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetRoundKey helper
func GetRoundKey(claimType string, roundID uint64) string {
	return claimType + "_" + strconv.FormatUint(roundID, 10)
}

// RoundPrefix helper
func RoundPrefix(claimType string, roundID uint64) []byte {
	return KeyPrefix(GetRoundKey(claimType, roundID))
}
