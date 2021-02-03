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

func (k Keeper) AllScore(c context.Context, req *types.QueryAllScoreRequest) (*types.QueryAllScoreResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scores []*types.Score
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	scoreStore := prefix.NewStore(store, types.KeyPrefix(types.ScoreKey))

	pageRes, err := query.Paginate(scoreStore, req.Pagination, func(key []byte, value []byte) error {
		var score types.Score
		if err := k.cdc.UnmarshalBinaryBare(value, &score); err != nil {
			return err
		}

		scores = append(scores, &score)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScoreResponse{Score: scores, Pagination: pageRes}, nil
}
