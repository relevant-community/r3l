package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/keeper"
)

func handleMsgCreateClaim(ctx sdk.Context, k keeper.Keeper, msg exported.MsgCreateClaim) (*sdk.Result, error) {
	claim := msg.GetClaim()
	k.CreateClaim(ctx, claim)
	// if err := k.CreateClaim(ctx, claim);
	// err != nil {
	// 	return nil, err
	// }

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
