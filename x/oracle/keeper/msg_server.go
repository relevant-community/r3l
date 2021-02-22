package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/relevant-community/r3l/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateClaim(goCtx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claim := msg.GetClaim()
	validatorAddress := sdk.ValAddress(msg.GetSubmitter())

	claimParams := k.GetParams(ctx).ClaimParams
	var claimTypeExists bool
	for _, param := range claimParams {
		if param.ClaimType == claim.Type() {
			claimTypeExists = true
		}
	}
	if claimTypeExists != true {
		return nil, sdkerrors.Wrap(types.ErrNoClaimTypeExists, claim.Type())
	}

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
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSubmitter().String()),
			sdk.NewAttribute(types.AttributeKeyClaimHash, claim.Hash().String()),
		),
	)

	return &types.MsgCreateClaimResponse{
		Hash: claim.Hash(),
	}, nil
}
