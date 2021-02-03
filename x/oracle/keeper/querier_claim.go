package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func queryClaim(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryClaimRequest

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	claim := k.GetClaim(ctx, params.ClaimHash)
	if claim == nil {
		return nil, sdkerrors.Wrap(types.ErrNoClaimExists, params.ClaimHash.String())
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, claim)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllClaims(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllClaimParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	claims := k.GetAllClaim(ctx)

	start, end := client.Paginate(len(claims), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		claims = []exported.Claim{}
	} else {
		claims = claims[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, claims)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
