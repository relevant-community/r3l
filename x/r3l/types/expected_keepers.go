package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	oracle "github.com/relevant-community/r3l/x/oracle/exported"
	oracletypes "github.com/relevant-community/r3l/x/oracle/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

type (
	// OracleKeeper interface
	OracleKeeper interface {
		DeleteVotesForRound(ctx sdk.Context, claimType string, roundID uint64)
		DeletePendingRound(ctx sdk.Context, claimType string, roundID uint64)
		GetPendingRounds(ctx sdk.Context, roundType string) (rounds []uint64)
		TallyVotes(ctx sdk.Context, claimType string, roundID uint64) *oracletypes.RoundResult
		GetClaim(ctx sdk.Context, hash tmbytes.HexBytes) oracle.Claim
	}
	// StakingKeeper interface
	StakingKeeper interface {
		// ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingexported.ValidatorI
		Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI // get validator by operator address; nil when validator not found
	}
)
