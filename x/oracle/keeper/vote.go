package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
)

// CastVote casts a vote for a given claim.
func (k Keeper) CastVote(ctx sdk.Context, claim exported.Claim, validator sdk.ValAddress) {
	k.CreateClaim(ctx, claim)
	roundID := claim.GetRoundID()
	claimType := claim.Type()
	vote := types.NewVote(roundID, claim, validator, claimType)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))

	var votes types.Round
	bz := store.Get(types.RoundPrefix(claimType, roundID))
	if len(bz) == 0 {
		votes = types.Round{
			Votes:   []types.Vote{*vote},
			RoundId: roundID,
			Type:    claimType,
		}
	} else {
		k.cdc.MustUnmarshalBinaryBare(bz, &votes)
		votes.Votes = append(votes.Votes, *vote)
	}

	k.AddPendingRound(ctx, vote.Type, vote.RoundId)
	store.Set(types.RoundPrefix(claimType, roundID), k.cdc.MustMarshalBinaryBare(&votes))
}

// DeleteVotesForRound deletes all votes and claims for a given round and claimType
func (k Keeper) DeleteVotesForRound(ctx sdk.Context, claimType string, roundID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	roundKey := types.GetRoundKey(claimType, roundID)

	round := k.GetRound(ctx, claimType, roundID)

	for _, vote := range round.Votes {
		k.DeleteClaim(ctx, []byte(vote.ClaimHash))
	}
	store.Delete(types.KeyPrefix(roundKey))
}

// TallyVotes tallies up the votes for a given Claim and returns the result with the maximum claim
// vote.ClaimHash
func (k Keeper) TallyVotes(ctx sdk.Context, claimType string, roundID uint64) *types.RoundResult {
	votes := k.GetRound(ctx, claimType, roundID)

	voteMap := make(map[string]*types.RoundResult)
	var maxVotePower int64
	var maxVoteKey string
	for _, vote := range votes.Votes {
		validator := k.StakingKeeper.Validator(ctx, vote.Validator)
		if validator == nil || !validator.IsBonded() || validator.IsJailed() {
			continue
		}
		weight := validator.GetConsensusPower()
		key := vote.GetConsensusId()

		resultObject := voteMap[key]
		if resultObject == nil {
			resultObject = &types.RoundResult{}
		}

		resultObject.UpsertClaim(vote.ClaimHash, weight)
		resultObject.VotePower += weight

		voteMap[key] = resultObject

		if resultObject.VotePower > maxVotePower {
			maxVoteKey = key
		}
	}

	result := voteMap[maxVoteKey]
	if result == nil {
		return nil
	}

	passesThreshold, totalBondedPower := k.VotePassedThreshold(ctx, result)

	if passesThreshold {
		result.TotalPower = totalBondedPower
		return result
	}
	return nil
}

// VotePassedThreshold checks if a given claim has cleared the vote threshold
func (k Keeper) VotePassedThreshold(ctx sdk.Context, roundResult *types.RoundResult) (bool, int64) {
	totalBondedPower := sdk.TokensToConsensusPower(k.StakingKeeper.TotalBondedTokens(ctx))
	voteThreshold := k.VoteThreshold(ctx)
	thresholdVotes := voteThreshold.MulInt64(totalBondedPower).RoundInt()
	votePower := sdk.NewInt(roundResult.VotePower)
	return !votePower.IsZero() && votePower.GT(thresholdVotes), totalBondedPower
}
