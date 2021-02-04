package keeper

import (
	// this line is used by starport scaffolding # 1

	"fmt"

	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a querier for oracle REST endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryParameters:
			res, err = queryParams(ctx, k, legacyQuerierCdc)

		case types.QueryClaim:
			res, err = queryClaim(ctx, req, k, legacyQuerierCdc)

		case types.QueryListClaim:
			res, err = queryAllClaims(ctx, req, k, legacyQuerierCdc)

		// this line is used by starport scaffolding # 1
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

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

	var c exported.Claim
	undo := legacyQuerierCdc.UnmarshalJSON(res, &c)
	fmt.Println("RES", undo)

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
