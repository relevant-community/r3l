package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/types"
)

const claimType = "test"

func (suite *KeeperTestSuite) TestCastVote() {
	ctx := suite.ctx.WithIsCheckTx(false)
	claim := types.NewTestClaim(99, "test", claimType)
	val0 := suite.validators[0]
	val1 := suite.validators[1]

	suite.k.CastVote(ctx, claim, val0)

	savedClaim := suite.k.GetClaim(ctx, claim.Hash())
	suite.NotNil(savedClaim)

	pending := suite.k.GetPendingRounds(ctx, claimType)
	suite.NotNil(pending)

	roundID := pending[len(pending)-1]
	suite.NotNil(roundID)
	suite.Equal(roundID, claim.GetRoundID())

	roundVotes := suite.k.GetVotesForRound(ctx, claimType, roundID)
	suite.NotNil(roundVotes)

	vote := roundVotes.Votes[len(roundVotes.Votes)-1]
	suite.NotNil(vote)

	suite.Equal(vote.Validator, val0)
	suite.Equal(vote.ClaimHash, claim.Hash())
	suite.Equal(vote.Type, claim.ClaimType)
	suite.Equal(vote.ConsensusId, claim.GetConcensusKey())
	suite.Equal(vote.RoundId, roundID)

	// Add second vote
	suite.k.CastVote(ctx, claim, val1)
	roundVotes = suite.k.GetVotesForRound(ctx, claimType, roundID)
	suite.NotNil(roundVotes)
	suite.Equal(len(roundVotes.Votes), 2)

	// Test cleanup

	// Remove pending round
	suite.k.DeletePendingRound(ctx, claimType, roundID)

	pending = suite.k.GetPendingRounds(ctx, claimType)
	suite.Equal(len(pending), 0)

	// Remove votes + claims
	suite.k.DeleteVotesForRound(ctx, claimType, roundID)
	roundVotes = suite.k.GetVotesForRound(ctx, claimType, roundID)
	suite.Nil(roundVotes)

	savedClaim = suite.k.GetClaim(ctx, claim.Hash())
	suite.Nil(savedClaim)
}

func (suite *KeeperTestSuite) TestVoteTally() {
	ctx := suite.ctx.WithIsCheckTx(false)
	claim := types.NewTestClaim(99, "test", claimType)
	roundID := claim.GetRoundID()

	val0 := suite.validators[0]
	val1 := suite.validators[1]

	suite.k.CastVote(ctx, claim, val0)

	// Haven't reached threshold
	roundResult := suite.k.TallyVotes(ctx, claimType, roundID)
	suite.Nil(roundResult)

	// Haven't reached threshold (50%)
	suite.k.CastVote(ctx, claim, val1)
	roundResult = suite.k.TallyVotes(ctx, claimType, roundID)
	suite.NotNil(roundResult)

	totalBondedPower := sdk.TokensToConsensusPower(suite.k.StakingKeeper.TotalBondedTokens(ctx))
	suite.Equal(roundResult.TotalPower, totalBondedPower)

	suite.Equal(roundResult.VotePower, suite.pow[0]+suite.pow[1])

	suite.Equal(len(roundResult.Claims), 1)
	suite.Equal(roundResult.Claims[0].ClaimHash, claim.Hash())
}

// TEST Getters & Setters used in Genesis

func (suite *KeeperTestSuite) TestIterateVotes() {
	ctx := suite.ctx.WithIsCheckTx(false)
	num := 20
	suite.populateVotes(ctx, num)

	votes := suite.k.GetRoundVotes(ctx)
	suite.Len(votes, num)
}

func (suite *KeeperTestSuite) populateVotes(ctx sdk.Context, num int) []types.RoundVotes {
	roundVotes := make([]types.RoundVotes, num)

	for i := 0; i < num; i++ {
		roundID := uint64(i)

		claim := types.NewTestClaim(int64(roundID), "test", claimType)
		vote := types.NewVote(roundID, claim, suite.validators[0], claimType)
		roundVote := types.RoundVotes{
			Votes:   []types.Vote{*vote},
			RoundId: roundID,
			Type:    claimType,
		}
		roundVotes[i] = roundVote
		suite.k.SetRoundVote(ctx, roundVote)
	}
	return roundVotes
}
