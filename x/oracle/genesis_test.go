package oracle_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle"
	exported "github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/keeper"
	"github.com/relevant-community/r3l/x/oracle/testoracle"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx        sdk.Context
	k          keeper.Keeper
	validators []sdk.ValAddress
	addrs      []sdk.AccAddress
}

func (suite *GenesisTestSuite) SetupTest() {
	checkTx := false

	app, _ := testoracle.CreateTestInput()

	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.k = app.OracleKeeper

	powers := []int64{10, 10, 10}
	addrs, validators, _ := testoracle.CreateValidators(suite.T(), suite.ctx, app, powers)
	suite.validators = validators
	suite.addrs = addrs
}

func (suite *GenesisTestSuite) TestGenesis() {
	var (
		genesisState *types.GenesisState
		testClaim    []exported.Claim
	)

	claimType := "test"
	round := []types.Round{}
	params := types.DefaultParams()
	params.ClaimParams = map[string](types.ClaimParams){
		claimType: {
			ClaimType: claimType,
		},
	}
	delegations := []types.MsgDelegateFeedConsent{
		{
			Validator: suite.addrs[0].String(),
			Delegate:  suite.addrs[4].String(),
		},
	}

	numClaims := 10

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func()
	}{
		{
			"valid",
			func() {

				testClaim = make([]exported.Claim, numClaims)
				for i := 0; i < numClaims; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i + 1),
						Content:     "test",
						ClaimType:   claimType,
					}
				}
				pending := map[string][]uint64{
					"test": {98, 99},
				}
				genesisState = types.NewGenesisState(params, round, testClaim, pending, delegations)
			},
			true,
			func() {
				for _, c := range testClaim {
					res := suite.k.GetClaim(suite.ctx, c.Hash())
					suite.NotNil(res)
				}
				p := suite.k.GetPendingRounds(suite.ctx, claimType)
				suite.Equal(p[0], uint64(98))
				suite.Equal(p[1], uint64(99))
			},
		},
		{
			"invalid",
			func() {
				testClaim = make([]exported.Claim, numClaims)
				for i := 0; i < numClaims; i++ {
					testClaim[i] = &types.TestClaim{
						Content:   "test",
						ClaimType: "test",
					}
				}
				pending := map[string][]uint64{}
				genesisState = types.NewGenesisState(params, round, testClaim, pending, delegations)
			},
			false,
			func() {
				suite.Empty(suite.k.GetAllClaims(suite.ctx))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()

			if tc.expPass {
				suite.NotPanics(func() {
					oracle.InitGenesis(suite.ctx, suite.k, *genesisState)
				})
			} else {
				suite.Panics(func() {
					oracle.InitGenesis(suite.ctx, suite.k, *genesisState)
				})
			}

			tc.posttests()
		})
	}
}

func (suite *GenesisTestSuite) TestImportExportGenesis() {
	claimType := "test"
	roundID := int64(1)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func()
	}{
		{
			"success",
			func() {
				suite.k.CreateVote(suite.ctx, &types.TestClaim{
					BlockHeight: roundID,
					Content:     "test",
					ClaimType:   claimType,
				}, suite.validators[0])
			},
			true,
			func() {},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()

			if tc.expPass {
				suite.NotPanics(func() {
					genesisState := oracle.ExportGenesis(suite.ctx, suite.k)

					// Test Import
					// clear store
					suite.k.DeleteVotesForRound(suite.ctx, claimType, uint64(roundID))
					oracle.InitGenesis(suite.ctx, suite.k, *genesisState)
					newGenesis := oracle.ExportGenesis(suite.ctx, suite.k)

					suite.Require().Equal(genesisState, newGenesis)
				})
			} else {
				suite.Panics(func() {
					oracle.ExportGenesis(suite.ctx, suite.k)
				})
			}

			tc.posttests()
		})
	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
