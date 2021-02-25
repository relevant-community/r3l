package cli

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	"github.com/spf13/cobra"
)

// store claim and salt of previous round to submit after prevote
// TODO: a non-ephemeral storage option?
var (
	CurrentClaim exported.Claim
	CurrentSalt  string
)

// SubmitWorkerTx submits the correct transactions based on oracle params
func SubmitWorkerTx(cmd *cobra.Command, claim exported.Claim, usePrevote bool) {

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
	prevoteHash := types.VoteHash(newSalt, claim.Hash().String(), clientCtx.FromAddress)
	prevoteMsg := types.NewMsgPrevote(clientCtx.GetFromAddress(), prevoteHash)
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

	voteMsg, err := types.NewMsgVote(clientCtx.GetFromAddress(), claim, salt)
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

func genrandstr(s int) string {
	b := make([]byte, s)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
