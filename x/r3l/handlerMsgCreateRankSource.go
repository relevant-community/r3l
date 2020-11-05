package r3l

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/relevant-community/r3l/x/r3l/keeper"
)

func handleMsgCreateRankSource(ctx sdk.Context, k keeper.Keeper, rankSource *types.MsgRankSource) (*sdk.Result, error) {
	k.CreateRankSource(ctx, *rankSource)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
