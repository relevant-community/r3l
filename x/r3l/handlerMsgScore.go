package r3l

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/keeper"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func handleMsgCreateScore(ctx sdk.Context, k keeper.Keeper, score *types.MsgScore) (*sdk.Result, error) {
	k.CreateScore(ctx, *score)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgSetScores(ctx sdk.Context, k keeper.Keeper, msgScores *types.MsgScores) (*sdk.Result, error) {
	k.SetScores(ctx, *msgScores)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
