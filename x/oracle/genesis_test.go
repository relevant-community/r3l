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

	ctx    sdk.Context
	keeper keeper.Keeper
}

func (suite *GenesisTestSuite) SetupTest() {
	checkTx := false

	app, _ := testoracle.CreateTestInput()

	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.keeper = app.OracleKeeper
}

func (suite *GenesisTestSuite) TestInitGenesis() {
	var (
		genesisState *types.GenesisState
		testClaim    []exported.Claim
	)
	roundVotes := []types.RoundVotes{}
	params := types.DefaultParams()

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func()
	}{
		{
			"valid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i + 1),
						Content:     "test",
						ClaimType:   "test",
					}
				}
				genesisState = types.NewGenesisState(params, roundVotes, testClaim)
			},
			true,
			func() {
				for _, c := range testClaim {
					res := suite.keeper.GetClaim(suite.ctx, c.Hash())
					suite.NotNil(res)
				}
			},
		},
		{
			"invalid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						// Content:   "test",
						ClaimType: "test",
					}
				}
				genesisState = types.NewGenesisState(params, roundVotes, testClaim)
			},
			false,
			func() {
				suite.Empty(suite.keeper.GetAllClaim(suite.ctx))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()

			if tc.expPass {
				suite.NotPanics(func() {
					oracle.InitGenesis(suite.ctx, suite.keeper, *genesisState)
				})
			} else {
				suite.Panics(func() {
					oracle.InitGenesis(suite.ctx, suite.keeper, *genesisState)
				})
			}

			tc.posttests()
		})
	}
}

func (suite *GenesisTestSuite) TestExportGenesis() {

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func()
	}{
		{
			"success",
			func() {
				suite.keeper.CreateClaim(suite.ctx, &types.TestClaim{
					BlockHeight: 1,
					Content:     "test",
					ClaimType:   "test",
				})
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
					oracle.ExportGenesis(suite.ctx, suite.keeper)
				})
			} else {
				suite.Panics(func() {
					oracle.ExportGenesis(suite.ctx, suite.keeper)
				})
			}

			tc.posttests()
		})
	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
