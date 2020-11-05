package r3l

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/r3l/types"
	"github.com/relevant-community/r3l/x/r3l/keeper"
)

func handleMsgCreateNamespace(ctx sdk.Context, k keeper.Keeper, namespace *types.MsgNamespace) (*sdk.Result, error) {
	k.CreateNamespace(ctx, *namespace)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
