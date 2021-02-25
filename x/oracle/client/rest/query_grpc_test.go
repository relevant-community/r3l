package rest_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/cosmos/cosmos-sdk/types/query"
	proto "github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/app"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/stretchr/testify/suite"

	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     testnet.Config
	network *testnet.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := app.DefaultConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = testnet.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestGRPCQueries() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	testCases := []struct {
		name     string
		url      string
		headers  map[string]string
		expErr   bool
		respType proto.Message
		expected proto.Message
	}{
		{
			"Get Params",
			fmt.Sprintf("%s/r3l/oracle/v1beta1/params", baseURL),
			map[string]string{
				grpctypes.GRPCBlockHeightHeader: "1",
			},
			false,
			&types.QueryParamsResponse{},
			&types.QueryParamsResponse{
				Params: types.DefaultParams(),
			},
		},
		{
			"Get all claims",
			fmt.Sprintf("%s/r3l/oracle/v1beta1/allclaims", baseURL),
			map[string]string{
				grpctypes.GRPCBlockHeightHeader: "1",
			},
			false,
			&types.QueryAllClaimsResponse{},
			&types.QueryAllClaimsResponse{
				Pagination: &query.PageResponse{Total: 0},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			resp, err := testutil.GetRequestWithHeaders(tc.url, tc.headers)
			s.Require().NoError(err)

			err = val.ClientCtx.JSONMarshaler.UnmarshalJSON(resp, tc.respType)

			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.expected.String(), tc.respType.String())
			}
		})
	}
}
