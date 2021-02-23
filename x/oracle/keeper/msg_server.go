package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	signer := msg.MustGetSubmitter()

	// get delegator's validator
	valAddr := sdk.ValAddress(k.GetValidatorAddressFromDelegate(ctx, signer))

	// if there is no delegation it must be the validator
	if valAddr == nil {
		valAddr = sdk.ValAddress(signer)
	}

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, valAddr)

	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, valAddr.String())
	}
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

	// store the validator vote
	k.CastVote(ctx, claim, valAddr)

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

// DelegateFeedConsent implements types.MsgServer
func (k msgServer) DelegateFeedConsent(c context.Context, msg *types.MsgDelegateFeedConsent) (*types.MsgDelegateFeedConsentResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	val, del := msg.MustGetValidator(), msg.MustGetDelegate()

	if k.Keeper.StakingKeeper.Validator(ctx, sdk.ValAddress(val)) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}

	k.SetValidatorDelegateAddress(ctx, val, del)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDelegateFeed),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyDeleagte, msg.Delegate),
		),
	)

	return &types.MsgDelegateFeedConsentResponse{}, nil
}
