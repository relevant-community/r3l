package worker

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/relevant-community/r3l/x/oracle/client/cli"
	"github.com/relevant-community/r3l/x/oracle/types"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"
	tendermint "github.com/tendermint/tendermint/types"

	"github.com/spf13/cobra"
)

// HandleTx is our custom tx handler
func HandleTx(cmd *cobra.Command, txEvent rpctypes.ResultEvent) error {
	return nil
}

// HandleBlock is our custom block handler
func HandleBlock(cmd *cobra.Command, blockEvent rpctypes.ResultEvent) error {
	blockHeight :=
		blockEvent.Data.(tendermint.EventDataNewBlock).Block.Header.Height

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
