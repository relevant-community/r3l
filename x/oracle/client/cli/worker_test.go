package cli_test

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/x/oracle/client/cli"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/spf13/cobra"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func txHandler(cmd *cobra.Command, txEvent ctypes.ResultEvent) error {
	return nil
}

func blockHandler(cmd *cobra.Command, blockEvent ctypes.ResultEvent) error {
	testClaim := types.NewTestClaim(1, "test", "test")
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	// then create the claim message and submit it to the oracle
	voteMsg, err := types.NewMsgVote(clientCtx.GetFromAddress(), testClaim, "")
	if err != nil {
		fmt.Println("Error creating claim", err)
		return err
	}

	if err := voteMsg.ValidateBasic(); err != nil {
		return err
	}

	err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), voteMsg)

	if err != nil {
		fmt.Println("TX ERROR", err)
		return err
	}
	return nil
}

func (s *IntegrationTestSuite) TestWorkerCmd() {
	val := s.network.Validators[0]

	cli.InitializeWorker(blockHandler, txHandler)

	testCases := map[string]struct {
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		"run-worker": {
			[]string{
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			false,
			&sdk.TxResponse{},
			0,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			clientCtx := val.ClientCtx.WithNodeURI(val.RPCAddress)
			clientCtx.OutputFormat = "json"

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.StartWorkerCmd(), tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code)
			}

			// wait for the worker tx to execute and confirm state transition
			time.Sleep(2000 * time.Millisecond)
			s.network.WaitForNextBlock()

			testClaim := types.NewTestClaim(1, "test", "test")
			res, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdClaim(), []string{testClaim.Hash().String()})

			resType := &types.QueryClaimResponse{}
			s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resType))
			resClaim := &types.TestClaim{}
			resClaim.Unmarshal(resType.Claim.Value)
			s.Require().Contains(resClaim.String(), testClaim.String())
		})
	}
}
