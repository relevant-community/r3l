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

func (k Keeper) AllNamespace(c context.Context, req *types.QueryAllNamespaceRequest) (*types.QueryAllNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var namespaces []*types.MsgNamespace
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	namespaceStore := prefix.NewStore(store, types.KeyPrefix(types.NamespaceKey))

	pageRes, err := query.Paginate(namespaceStore, req.Pagination, func(key []byte, value []byte) error {
		var namespace types.MsgNamespace
		if err := k.cdc.UnmarshalBinaryBare(value, &namespace); err != nil {
			return err
		}

		namespaces = append(namespaces, &namespace)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNamespaceResponse{Namespace: namespaces, Pagination: pageRes}, nil
}
