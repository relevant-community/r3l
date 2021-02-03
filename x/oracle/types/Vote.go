package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
)

// NewVote creates a MsgVote instance
func NewVote(roundID uint64, claim exported.Claim, validator sdk.ValAddress, claimType string) *Vote {
	return &Vote{
		RoundId:     roundID,
		ClaimHash:   claim.Hash(),
		ConsensusId: claim.GetConcensusKey(),
		Validator:   validator,
		Type:        claimType,
	}
}
