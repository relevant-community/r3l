package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func CmdCreateScore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-score [blockHeight] [score]",
		Short: "Creates a new score",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockHeight, _ := strconv.ParseInt(args[0], 10, 64)
			argsPRank, _ := strconv.ParseInt(args[1], 10, 64)
			argsNRank, _ := strconv.ParseInt(args[2], 10, 64)

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgScore(clientCtx.GetFromAddress(), int64(argsBlockHeight), int64(argsPRank), int64(argsNRank))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
