package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/spf13/cobra"
)

func CmdListNamespace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-namespace",
		Short: "list all namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNamespaceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllNamespace(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
