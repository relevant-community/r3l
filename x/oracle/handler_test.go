package oracle_test

import (
	"testing"

	"github.com/relevant-community/r3l/app"
	"github.com/relevant-community/r3l/x/oracle"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/testoracle"
	"github.com/relevant-community/r3l/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type HandlerTestSuite struct {
	suite.Suite
	handler    sdk.Handler
	app        *app.App
	validators []sdk.ValAddress
}

func testMsgSubmitClaim(r *require.Assertions, claim exported.Claim, sender sdk.AccAddress) exported.MsgCreateClaimI {
	msg, err := types.NewMsgCreateClaim(sender, claim)
	r.NoError(err)
	return msg
}

func (suite *HandlerTestSuite) SetupTest() {

	app, ctx := testoracle.CreateTestInput()
	powers := []int64{10, 10, 10}
	_, validators, _ := testoracle.CreateValidators(suite.T(), ctx, app, powers)
	suite.validators = validators
	suite.handler = oracle.NewHandler(app.OracleKeeper)
	suite.app = app
}

func (suite *HandlerTestSuite) TestMsgSubmitClaim() {

	// pk := ed25519.GenPrivKey()
	nonValidator := sdk.AccAddress("test________________")
	validator := suite.validators[0]

	testCases := []struct {
		msg       sdk.Msg
		expectErr bool
	}{
		{
			testMsgSubmitClaim(
				suite.Require(),
				types.NewTestClaim(10, "test", "test"),
				sdk.AccAddress(validator),
			),
			false,
		},
		{
			testMsgSubmitClaim(
				suite.Require(),
				types.NewTestClaim(10, "test", "test"),
				nonValidator,
			),
			true,
		},
	}

	for i, tc := range testCases {
		ctx := suite.app.BaseApp.NewContext(false, tmproto.Header{Height: suite.app.LastBlockHeight() + 1})

		res, err := suite.handler(ctx, tc.msg)
		if tc.expectErr {
			suite.Require().Error(err, "expected error; tc #%d", i)
		} else {
			suite.Require().NoError(err, "unexpected error; tc #%d", i)
			suite.Require().NotNil(res, "expected non-nil result; tc #%d", i)

			msg := tc.msg.(exported.MsgCreateClaimI)

			var resultData types.MsgCreateClaimResponse
			suite.app.AppCodec().UnmarshalBinaryBare(res.Data, &resultData)
			suite.Require().Equal(msg.GetClaim().Hash().Bytes(), resultData.Hash, "invalid hash; tc #%d", i)
		}
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
