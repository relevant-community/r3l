package cli

// func CmdCreateClaim() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "create-claim [blockNumber] [hash]",
// 	Short: "Creates a new claim",
// 	Args:  cobra.ExactArgs(2),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		argsBlockNumber, _ := strconv.ParseInt(args[0], 10, 64)
// 		argsHash := string(args[1])

// 		clientCtx := client.GetClientContextFromCmd(cmd)
// 		clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
// 		if err != nil {
// 			return err
// 		}

// 		msg := types.NewMsgCreateClaim(clientCtx.GetFromAddress(), int64(argsBlockNumber), string(argsHash))
// 		if err := msg.ValidateBasic(); err != nil {
// 			return err
// 		}
// 		return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 	},
// }

// flags.AddTxFlagsToCmd(cmd)

// return cmd
// }
