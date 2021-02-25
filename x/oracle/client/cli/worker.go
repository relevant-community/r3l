package cli

import (
	"context"
	"fmt"
	"sync"
	"time"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client/http"
	tendermint "github.com/tendermint/tendermint/types"
)

var once sync.Once

// BlockHandler is the code that the worker must run every block
type BlockHandler func(*cobra.Command, int64) error

// Worker is a singleton type
type Worker struct {
	handleBlock BlockHandler
}

var instance *Worker

// InitializeWorker intializes the singleton instance and sets the worker process
// This method should be called from the app's cmd/root.go file
// TODO add event handler?
func InitializeWorker(handler BlockHandler) *Worker {
	// We don't really need to do this, but this is a standard singleton pattern
	// to protect from creation of multiple instances
	once.Do(func() {
		instance = &Worker{
			handleBlock: handler,
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

			// Loop on received messages.
			var stop bool

			for stop == false {
				// go func() {
				block := <-blockEvent
				blockHeight :=
					block.Data.(tendermint.EventDataNewBlock).Block.Header.Height

				// run main worker proces here
				err := instance.handleBlock(cmd, blockHeight)

				if err != nil {
					fmt.Println("There was an error running the worker process", err)
				}
				stop = test
				// }()
				// time.Sleep(100 * time.Millisecond)
			}
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().StringP(tmcli.OutputFlag, "o", "text", "Output format (text|json)")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	return cmd
}
