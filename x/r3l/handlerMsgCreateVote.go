package r3l

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/relevant-community/r3l/x/r3l/keeper"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, vote *types.MsgVote) (*sdk.Result, error) {
	k.CreateVote(ctx, *vote)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
