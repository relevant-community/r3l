package worker

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/oracle/client/cli"
	"github.com/relevant-community/r3l/x/oracle/types"

	"github.com/spf13/cobra"
)

// HandleBlock is our custom worker code
func HandleBlock(cmd *cobra.Command, blockHeight int64) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}
	// Make sure we querey the data from the right block
	clientCtx.Height = blockHeight
	queryClient := types.NewQueryClient(clientCtx)

	params := &types.QueryParamsRequest{}

	// we query the array of claim parameters for different claims
	res, err := queryClient.Params(context.Background(), params)
	if err != nil {
		return err
	}

	// for each claim type in the array, we run the approriate process
	for _, param := range res.Params.ClaimParams {
		switch param.ClaimType {
		case "ScoreClaim":
			// use VotePeriod to check if its time to submit a claim
			if uint64(blockHeight)%param.VotePeriod != 0 {
				continue
			}

			// run process for the given claimType
			scores, err := ComputeReputation(clientCtx)
			if err != nil {
				fmt.Println("Error computing rep", err)
				continue
			}
			cli.SubmitWorkerTx(cmd, scores, param.Prevote)
		}
	}

	return nil
}
