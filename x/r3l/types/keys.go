package types

const (
	// ModuleName defines the module name
	ModuleName = "r3l"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	VoteKey = "Vote"
)

const (
	ScoreKey = "Score"
)

const (
	NamespaceKey = "Namespace"
)

const (
	RankSourceKey = "RankSource"
)

const (
	NextNamespaceIdKey = "Namespace"
)
