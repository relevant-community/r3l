package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	rpcClient "github.com/tendermint/tendermint/rpc/client/http"
	tendermint "github.com/tendermint/tendermint/types"
)

var once sync.Once

type process func(*cobra.Command, client.Context) error

// Worker is a singleton type
type Worker struct {
	runProcess process
}

var instance *Worker

// Init intializes the singleton instance and worker parameters
func Init(_runProcess process) *Worker {
	// We don't really need to do this, but this is a standard singleton pattern
	// to protect from creation of multiple instances
	once.Do(func() {
		instance = &Worker{
			runProcess: _runProcess,
		}
	})
	return instance
}

// StartWorkerCmd starts the sff-chain worker process
func (worker *Worker) StartWorkerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-worker",
		Short: "Starts offchain worker",
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			httpClient, _ := rpcClient.New(clientCtx.NodeURI, "/websocket")
			err = httpClient.Start()
			// Retry if we don't have a connection
			for err != nil {
				fmt.Println("Could not connect to Tendermint node, retrying in 1 second.")
				time.Sleep(1 * time.Second)
				err = httpClient.Start()
			}

			defer httpClient.Stop()
			ctx := context.Background()

			query := "tm.event = 'NewBlock'"

			blockEvent, err := httpClient.Subscribe(ctx, "test-client", query)
			// Retry if we don't have a connection
			for err != nil {
				fmt.Println("Could not subscribe to Tendermint block event, retrying in 1 second.")
				time.Sleep(1 * time.Second)
				blockEvent, err = httpClient.Subscribe(ctx, "test-client", query)
			}

			// We don't need a prompt
			clientCtx.SkipConfirm = true
			// TODO is json better for automated logging?
			clientCtx.OutputFormat = "text"

			// Loop on received messages.
			for {
				time.Sleep(100 * time.Millisecond)
				go func() {
					block := <-blockEvent
					blockHeight :=
						block.Data.(tendermint.EventDataNewBlock).Block.Header.Height

						// Important to update the bock height here
					// Data will be be fetch from clientCtx.Height - 1
					clientCtx.Height = blockHeight

					// run main worker proces here
					err := worker.runProcess(cmd, clientCtx)

					if err != nil {
						fmt.Println("There was an error running the worker process")
					}

				}()
			}
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.PersistentFlags().String(flags.FlagOutputDocument, "", "Output type")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	return cmd
}
