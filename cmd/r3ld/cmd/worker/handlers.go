package worker

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/oracle/x/oracle/client/cli"
	"github.com/relevant-community/oracle/x/oracle/types"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/spf13/cobra"
)

// HandleTx is our custom tx handler
func HandleTx(cmd *cobra.Command, txEvent rpctypes.ResultEvent) error {
	return nil
}

// HandleBlock is our custom block handler
func HandleBlock(cmd *cobra.Command, blockEvent rpctypes.ResultEvent) error {
	helper, err := cli.NewWorkerHelper(cmd, blockEvent)
	if err != nil {
		return err
	}

	// for each claim type in the array, run the approriate handler
	for _, param := range helper.OracleParams.ClaimParams {
		switch param.ClaimType {
		case "ScoreClaim":
			handleScoreClaim(cmd, helper, param)
		}
	}

	return nil
}

func handleScoreClaim(cmd *cobra.Command, helper *cli.WorkerHelper, param types.ClaimParams) error {
	// use VotePeriod to check if its time to submit a claim
	if helper.IsRoundStart(param.ClaimType) == false {
		return nil
	}

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	// Make sure we querey the data from the right height
	clientCtx.Height = helper.BlockHeight

	// run process for the given claimType
	scores, err := ComputeReputation(clientCtx)
	if err != nil {
		fmt.Println("Error computing reputaion", err)
		return nil
	}
	helper.SubmitWorkerTx(scores)
	return nil
}
