package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

func (suite *KeeperTestSuite) TestQueryParam() {
	var (
		req *types.QueryParamsRequest
	)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryParamsResponse)
	}{
		{
			"success",
			func() {
				req = &types.QueryParamsRequest{}
			},
			true,
			func(res *types.QueryParamsResponse) {
				suite.Require().NotNil(res)
				suite.Require().Equal(res.Params, types.DefaultParams())
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Params(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryClaim() {
	var (
		req    *types.QueryClaimRequest
		claims []exported.Claim
	)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryClaimResponse)
	}{
		{
			"empty request",
			func() {
				req = &types.QueryClaimRequest{}
			},
			false,
			func(res *types.QueryClaimResponse) {},
		},
		{
			"invalid request with empty claim hash",
			func() {
				req = &types.QueryClaimRequest{ClaimHash: tmbytes.HexBytes{}}
			},
			false,
			func(res *types.QueryClaimResponse) {},
		},
		{
			"success",
			func() {
				num := 1
				claims = suite.populateClaims(suite.ctx, num)
				req = types.NewQueryClaimRequest(claims[0].Hash())
			},
			true,
			func(res *types.QueryClaimResponse) {
				var c exported.Claim
				err := suite.app.InterfaceRegistry().UnpackAny(res.Claim, &c)
				suite.Require().NoError(err)
				suite.Require().NotNil(c)
				suite.Require().Equal(c, claims[0])
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Claim(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryAllClaims() {
	var (
		req *types.QueryAllClaimRequest
	)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryAllClaimResponse)
	}{
		{
			"success without claim",
			func() {
				req = &types.QueryAllClaimRequest{}
			},
			true,
			func(res *types.QueryAllClaimResponse) {
				suite.Require().Empty(res.Claim)
			},
		},
		{
			"success",
			func() {
				num := 100
				_ = suite.populateClaims(suite.ctx, num)
				pageReq := &query.PageRequest{
					Key:        nil,
					Limit:      50,
					CountTotal: false,
				}
				req = types.NewQueryAllClaimRequest(pageReq)
			},
			true,
			func(res *types.QueryAllClaimResponse) {
				suite.Equal(len(res.Claim), 50)
				suite.NotNil(res.Pagination.NextKey)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.AllClaim(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}
