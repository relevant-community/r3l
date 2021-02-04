package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func testMsgCreateClaim(t *testing.T, c exported.Claim, s sdk.AccAddress) exported.MsgCreateClaimI {
	msg, err := types.NewMsgCreateClaim(s, c)
	require.NoError(t, err)
	return msg
}

func TestMsgCreateClaim(t *testing.T) {
	submitter := sdk.AccAddress("test________________")

	testCases := []struct {
		msg       sdk.Msg
		submitter sdk.AccAddress
		expectErr bool
	}{
		{
			testMsgCreateClaim(t, &types.TestClaim{
				BlockHeight: 0,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			true,
		},
		{
			testMsgCreateClaim(t, &types.TestClaim{
				BlockHeight: 10,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			false,
		},
	}

	for i, tc := range testCases {
		require.Equal(t, tc.msg.Route(), types.RouterKey, "unexpected result for tc #%d", i)
		require.Equal(t, tc.msg.Type(), types.TypeMsgCreateClaim, "unexpected result for tc #%d", i)
		require.Equal(t, tc.expectErr, tc.msg.ValidateBasic() != nil, "unexpected result for tc #%d", i)

		if !tc.expectErr {
			require.Equal(t, tc.msg.GetSigners(), []sdk.AccAddress{tc.submitter}, "unexpected result for tc #%d", i)
		}
	}
}
