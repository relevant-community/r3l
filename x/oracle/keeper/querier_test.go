package keeper_test

import (
	"strings"

	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	custom = "custom"
)

func (suite *KeeperTestSuite) TestQuerier_QueryClaim_Existing() {
	ctx := suite.ctx.WithIsCheckTx(false)
	num := 3
	cdc := suite.app.LegacyAmino()

	claim := suite.populateClaims(ctx, num)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryClaim}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryClaimRequest(claim[0].Hash())),
	}

	bz, err := suite.querier(ctx, []string{types.QueryClaim}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c exported.Claim
	suite.Nil(cdc.UnmarshalJSON(bz, &c))
	suite.Equal(claim[0], c)
}

func (suite *KeeperTestSuite) TestQuerier_QueryClaim_NonExisting() {
	ctx := suite.ctx.WithIsCheckTx(false)
	cdc := suite.app.LegacyAmino()
	num := 100

	suite.populateClaims(ctx, num)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryClaim}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryClaimRequest([]byte("0000000000000000000000000000000000000000000000000000000000000000"))),
	}

	bz, err := suite.querier(ctx, []string{types.QueryClaim}, query)
	suite.NotNil(err)
	suite.Nil(bz)
}

func (suite *KeeperTestSuite) TestQuerier_QueryAllEvidence() {
	ctx := suite.ctx.WithIsCheckTx(false)
	cdc := suite.app.LegacyAmino()
	num := 100

	suite.populateClaims(ctx, num)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryListClaim}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryAllClaimParams(1, num)),
	}

	bz, err := suite.querier(ctx, []string{types.QueryListClaim}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c []exported.Claim
	suite.Nil(cdc.UnmarshalJSON(bz, &c))
	suite.Len(c, num)
}

func (suite *KeeperTestSuite) TestQuerier_QueryAllClaim_InvalidPagination() {
	ctx := suite.ctx.WithIsCheckTx(false)
	cdc := suite.app.LegacyAmino()
	num := 100

	suite.populateClaims(ctx, num)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryListClaim}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryAllClaimParams(0, num)),
	}

	bz, err := suite.querier(ctx, []string{types.QueryListClaim}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c []exported.Claim
	suite.Nil(cdc.UnmarshalJSON(bz, &c))
	suite.Len(c, 0)
}

func (suite *KeeperTestSuite) TestQuerier_Params() {
	ctx := suite.ctx.WithIsCheckTx(false)
	cdc := suite.app.LegacyAmino()

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryParameters}, "/"),
		Data: cdc.MustMarshalJSON(types.QueryParamsRequest{}),
	}

	bz, err := suite.querier(ctx, []string{types.QueryParameters}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var p types.Params
	suite.Nil(cdc.UnmarshalJSON(bz, &p))
	suite.Equal(p, types.DefaultParams())
}
