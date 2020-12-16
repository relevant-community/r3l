package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	r3l "github.com/relevant-community/r3l/x/r3l/types"
	"github.com/spf13/cobra"
	rpcClient "github.com/tendermint/tendermint/rpc/client/http"
	tendermint "github.com/tendermint/tendermint/types"
)

// StartCmd overwrites the default server.StartCmd and launches the worker process
func StartCmd(appCreator types.AppCreator, defaultNodeHome string) *cobra.Command {
	cmd := server.StartCmd(appCreator, defaultNodeHome)
	return cmd
}

func StartWorkerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-worker",
		Short: "Starts offchain worker",
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx := client.GetClientContextFromCmd(cmd)

			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			httpClient, _ := rpcClient.New(clientCtx.NodeURI, "/websocket")
			err = httpClient.Start()
			if err != nil {
				return err
			}

			defer httpClient.Stop()
			ctx := context.Background()

			// TODO defer cleanup?
			query := "tm.event = 'NewBlock'"
			e, err := httpClient.Subscribe(ctx, "test-client", query)
			if err != nil {
				return err
			}

			// We don't need a prompt
			clientCtx.SkipConfirm = true
			clientCtx.OutputFormat = "text"

			// Loop on received messages.
			for {
				time.Sleep(100 * time.Millisecond)
				// _, err := httpClient.Status(ctx)
				// if err != nil {
				// 	fmt.Println("Fail to fetch status", err.Error())
				// 	break
				// }
				go func() {
					block := <-e

					blockHeight :=
						block.Data.(tendermint.EventDataNewBlock).Block.Header.Height

					// Important to update the bock height here
					// Data will be be fetch from clientCtx.Height - 1
					clientCtx.Height = blockHeight

					fmt.Println("Height", clientCtx.Height)

					queryData, err := QueryData(cmd, clientCtx)
					if err != nil {
						fmt.Println("Error feting data", err)
						return
					}
					// fmt.Println("SCORES", queryData.scores)
					updatedScores := ComputeRank(queryData.votes, queryData.scores, queryData.rankSources)

					msg := r3l.NewMsgScores(clientCtx.GetFromAddress(), blockHeight, updatedScores)

					err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
					if err != nil {
						fmt.Println("TX ERROR", err)
					}

					fmt.Println("Finished blockHeight ", blockHeight)
				}()
			}
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.PersistentFlags().String(flags.FlagOutputDocument, "", "Output type")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}
