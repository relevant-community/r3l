package cli

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/spf13/cobra"
)

func CmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim",
		Short: "query claim by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			hash := args[0]
			decodedHash, err := hex.DecodeString(hash)
			if err != nil {
				return fmt.Errorf("invalid claim hash: %w", err)
			}

			params := &types.QueryClaimRequest{ClaimHash: decodedHash}

			res, err := queryClient.Claim(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-claim",
		Short: "list all claims",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllClaimRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllClaim(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
