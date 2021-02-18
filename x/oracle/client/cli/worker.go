package cli

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client/http"
	tendermint "github.com/tendermint/tendermint/types"
)

var once sync.Once

// WorkerProcess is the custom code that the worker must run
type WorkerProcess func(*cobra.Command, client.Context) error

// Worker is a singleton type
type Worker struct {
	runProcess WorkerProcess
}

var instance *Worker

// InitializeWorker intializes the singleton instance and sets the worker process
// This method should be called from the app's cmd/root.go file
func InitializeWorker(_runProcess WorkerProcess) *Worker {
	// We don't really need to do this, but this is a standard singleton pattern
	// to protect from creation of multiple instances
	once.Do(func() {
		instance = &Worker{
			runProcess: _runProcess,
		}
	})
	return instance
}

// StartWorkerCmd starts the off-chain worker process
func StartWorkerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-worker",
		Short: "Starts offchain worker",
		RunE: func(cmd *cobra.Command, args []string) error {

			if instance == nil {
				return fmt.Errorf("Worker process has not been intialized, nothing to run")
			}

			var test bool
			if len(args) > 0 {
				test = true
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if clientCtx.NodeURI == "" {
				return fmt.Errorf("Missing Tendermint Node URI")
			}
			rpcClient, _ := http.New(clientCtx.NodeURI, "/websocket")
			err = rpcClient.Start()

			// Retry if we don't have a connection
			for err != nil {
				fmt.Println("Could not connect to Tendermint node, retrying in 1 second.")
				time.Sleep(1 * time.Second)
				err = rpcClient.Start()
			}

			defer rpcClient.Stop()
			ctx := context.Background()

			query := "tm.event = 'NewBlock'"

			blockEvent, err := rpcClient.Subscribe(ctx, "test-client", query)
			// Retry if we don't have a connection
			for err != nil {
				fmt.Println("Could not subscribe to Tendermint block event, retrying in 1 second.")
				time.Sleep(1 * time.Second)
				blockEvent, err = rpcClient.Subscribe(ctx, "test-client", query)
			}

			// We don't need a prompt
			clientCtx.SkipConfirm = true
			// TODO is json better for automated logging?
			clientCtx.OutputFormat = "text"

			// Loop on received messages.
			var stop bool

			for stop == false {
				go func() {
					block := <-blockEvent
					blockHeight :=
						block.Data.(tendermint.EventDataNewBlock).Block.Header.Height

					// Important to update the bock height here
					// Data will be be fetch from clientCtx.Height - 1
					clientCtx.Height = blockHeight

					// run main worker proces here
					err := instance.runProcess(cmd, clientCtx)

					if err != nil {
						fmt.Println("There was an error running the worker process", err)
					}
					stop = test
				}()
				time.Sleep(100 * time.Millisecond)
			}
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.PersistentFlags().String(flags.FlagOutputDocument, "", "Output type")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	return cmd
}
