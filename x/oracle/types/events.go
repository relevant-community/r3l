package types

// claim module events
const (
	AttributeValueCategory = "oracle"
	EventTypeCreateClaim   = "create_claim"
	AttributeKeyClaimHash  = "claim_hash"

	EventTypeDelegateFeed = "delegate_feed"
	AttributeKeyDelegate  = "delegate"
	AttributeKeyValidator = "validator"

	EventTypePrevote        = "prevote"
	AttributeKeyPrevoteHash = "prevote_hash"
)
