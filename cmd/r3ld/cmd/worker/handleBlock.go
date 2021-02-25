package worker

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	oracle "github.com/relevant-community/r3l/x/oracle/types"

	"github.com/spf13/cobra"
)

// store claim and salt of previous round to submit after prevote
// TODO: a non-ephemeral storage option?
var (
	CurrentClaim exported.Claim
	CurrentSalt  string
)

// HandleBlock is our custom worker code
func HandleBlock(cmd *cobra.Command, blockHeight int64) error {
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
			// use VotePeriod to determine frequency of submitting claims
			blockHeight := uint64(clientCtx.Height)
			if blockHeight%param.VotePeriod == 0 {
				// run process for the given claimType
				scores, err := ComputeReputation(cmd, clientCtx)
				if err != nil {
					fmt.Println("Error computing rep", err)
					continue
				}
				SubmitTx(cmd, scores, param.Prevote)
			}
		}
	}

	return nil
}

// SubmitTx submits the correct transactions based on oracle params
func SubmitTx(cmd *cobra.Command, claim exported.Claim, usePrevote bool) {

	///////////////
	/// NO PREVOTE
	///////////////

	// use the most recent data and no salt = ""
	if usePrevote == false {
		err := SubmitVote(cmd, claim, "")
		if err != nil {
			fmt.Println("Error submitting claim vote", err)
		}
		return
	}

	///////////////
	/// YES PREVOTE
	///////////////

	// 1. submit data for prev round
	// 2. submit prevote for current round
	if CurrentClaim != nil && CurrentSalt != "" {
		err := SubmitVote(cmd, CurrentClaim, CurrentSalt)
		if err != nil {
			fmt.Println("Error submitting claim vote", err)
		}
	}

	err := SubmitPrevote(cmd, claim)
	if err != nil {
		fmt.Println("Error submitting prevote", err)
	}
}

// SubmitPrevote submits prevote of current claim
func SubmitPrevote(cmd *cobra.Command, claim exported.Claim) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	newSalt := genrandstr(9)
	prevoteHash := oracle.VoteHash(newSalt, claim.Hash().String(), clientCtx.FromAddress)
	prevoteMsg := oracle.NewMsgPrevote(clientCtx.GetFromAddress(), prevoteHash)
	if err := prevoteMsg.ValidateBasic(); err != nil {
		return err
	}
	err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), prevoteMsg)
	if err != nil {
		return err
	}
	CurrentClaim = claim
	CurrentSalt = newSalt
	return nil
}

// SubmitVote submits a vote of the given claim + salt
func SubmitVote(cmd *cobra.Command, claim exported.Claim, salt string) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	voteMsg, err := oracle.NewMsgVote(clientCtx.GetFromAddress(), claim, salt)
	if err != nil {
		return err
	}
	if err := voteMsg.ValidateBasic(); err != nil {
		return err
	}

	// Submit Claim from previous round
	err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), voteMsg)

	// TODO retry if the TX fails
	if err != nil {
		return err
	}
	return nil
}
