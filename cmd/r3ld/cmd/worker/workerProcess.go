package r3lworker

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/oracle/types"

	"github.com/spf13/cobra"
)

// RunWorkerProcess is our custom worker code
func RunWorkerProcess(cmd *cobra.Command, clientCtx client.Context) error {
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
			// use VotePeriod to determin frequency of submitting claims
			if clientCtx.Height%param.VotePeriod != 0 {
				// run process for the given claimType
				err = ComputeReputation(cmd, clientCtx)
				if err != nil {
					fmt.Println("Error processing worker command", err)
				}
			}
		}
	}

	return nil
}
