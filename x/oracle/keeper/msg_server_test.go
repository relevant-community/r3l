package keeper_test

import (
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func testMsgSubmitClaim(r *require.Assertions, claim exported.Claim, sender sdk.AccAddress) exported.MsgCreateClaimI {
	msg, err := types.NewMsgCreateClaim(sender, claim)
	r.NoError(err)
	return msg
}

func (suite *KeeperTestSuite) TestMsgSubmitClaim() {
	nonValidator := suite.addrs[3]
	validator := suite.validators[0]
	val1 := suite.validators[1]
	delegator := suite.addrs[4]

	suite.k.SetValidatorDelegateAddress(suite.ctx, sdk.AccAddress(val1), delegator)

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
				types.NewTestClaim(11, "test", "test"),
				sdk.AccAddress(delegator),
			),
			false,
		},
		{
			testMsgSubmitClaim(
				suite.Require(),
				types.NewTestClaim(12, "test", "test"),
				nonValidator,
			),
			true,
		},
	}

	for i, tc := range testCases {
		ctx := suite.ctx
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
