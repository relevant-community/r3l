package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/relevant-community/r3l/x/oracle/keeper"
	"github.com/relevant-community/r3l/x/oracle/types"
)

func handleMsgCreateClaim(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateClaim) (*sdk.Result, error) {
	claim := msg.GetClaim()
	validatorAddress := sdk.ValAddress(msg.GetSubmitter())

	// TODO add validator delegation
	// if !msg.Feeder.Equals(msg.Validator) {
	// 	delegate := keeper.GetOracleDelegate(ctx, msg.Validator)
	// 	if !delegate.Equals(msg.Feeder) {
	// 		return nil, sdkerrors.Wrap(ErrNoVotingPermission, msg.Feeder.String())
	// 	}
	// }

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, validatorAddress)
	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, validatorAddress.String())
	}

	// store the validator vote
	k.CastVote(ctx, claim, validatorAddress)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateClaim,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSubmitter().String()),
			sdk.NewAttribute(types.AttributeKeyClaimHash, claim.Hash().String()),
		),
	)

	return &sdk.Result{
		Data:   claim.Hash(),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
