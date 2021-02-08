package rest_test

// import (
// 	"fmt"
// 	"testing"

// 	proto "github.com/gogo/protobuf/proto"
// 	"github.com/relevant-community/r3l/app"
// 	"github.com/relevant-community/r3l/x/oracle/types"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/cosmos/cosmos-sdk/testutil"
// 	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
// 	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
// )

// type IntegrationTestSuite struct {
// 	suite.Suite

// 	cfg     testnet.Config
// 	network *testnet.Network
// }

// func (s *IntegrationTestSuite) SetupSuite() {
// 	s.T().Log("setting up integration test suite")

// 	cfg := app.DefaultConfig()
// 	cfg.NumValidators = 1

// 	s.cfg = cfg
// 	s.network = testnet.New(s.T(), cfg)

// 	_, err := s.network.WaitForHeight(1)
// 	s.Require().NoError(err)
// }

// func (s *IntegrationTestSuite) TearDownSuite() {
// 	s.T().Log("tearing down integration test suite")
// 	s.network.Cleanup()
// }

// func TestIntegrationTestSuite(t *testing.T) {
// 	suite.Run(t, new(IntegrationTestSuite))
// }

// func (s *IntegrationTestSuite) TestGRPCQueries() {
// 	val := s.network.Validators[0]
// 	baseURL := val.APIAddress

// 	// consAddr := sdk.ConsAddress(val.PubKey.Address()).String()

// 	testCases := []struct {
// 		name     string
// 		url      string
// 		headers  map[string]string
// 		expErr   bool
// 		respType proto.Message
// 		expected proto.Message
// 	}{
// 		{
// 			"Get all claims",
// 			fmt.Sprintf("%s/custom/oracle/"+types.QueryParameters+"/", baseURL),
// 			map[string]string{
// 				grpctypes.GRPCBlockHeightHeader: "1",
// 			},
// 			false,
// 			&types.QueryParamsResponse{},
// 			&types.QueryParamsResponse{},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc

// 		s.Run(tc.name, func() {

// 			resp, err := testutil.GetRequestWithHeaders(tc.url, tc.headers)
// 			s.Require().NoError(err)
// 			fmt.Println(tc.url)

// 			err = val.ClientCtx.JSONMarshaler.UnmarshalJSON(resp, tc.respType)

// 			if tc.expErr {
// 				s.Require().Error(err)
// 			} else {
// 				fmt.Println(err)
// 				s.Require().NoError(err)
// 				fmt.Println(tc.respType.String())
// 				s.Require().Equal(tc.expected.String(), tc.respType.String())
// 			}
// 		})
// 	}
// }
