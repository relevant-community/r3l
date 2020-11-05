package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/relevant-community/r3l/x/r3l/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllRankSource(c context.Context, req *types.QueryAllRankSourceRequest) (*types.QueryAllRankSourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rankSources []*types.MsgRankSource
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rankSourceStore := prefix.NewStore(store, types.KeyPrefix(types.RankSourceKey))

	pageRes, err := query.Paginate(rankSourceStore, req.Pagination, func(key []byte, value []byte) error {
		var rankSource types.MsgRankSource
		if err := k.cdc.UnmarshalBinaryBare(value, &rankSource); err != nil {
			return err
		}

		rankSources = append(rankSources, &rankSource)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRankSourceResponse{RankSource: rankSources, Pagination: pageRes}, nil
}
