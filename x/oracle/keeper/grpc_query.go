package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	proto "github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Keeper{}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Claim implements the Query/Claim gRPC method
func (k Keeper) Claim(c context.Context, req *types.QueryClaimRequest) (*types.QueryClaimResponse, error) {
	empty := &types.QueryClaimRequest{}
	if req == nil || req == empty {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.ClaimHash == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid hash")
	}

	ctx := sdk.UnwrapSDKContext(c)

	claim := k.GetClaim(ctx, req.ClaimHash)
	if claim == nil {
		return nil, status.Errorf(codes.NotFound, "claim %s not found", req.ClaimHash)
	}

	msg, ok := claim.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
	}

	claimAny, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryClaimResponse{Claim: claimAny}, nil
}

// AllClaim implements the Query/AllClaim gRPC method
func (k Keeper) AllClaim(c context.Context, req *types.QueryAllClaimRequest) (*types.QueryAllClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claims []*codectypes.Any
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claimStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimKey))

	pageRes, err := query.Paginate(claimStore, req.Pagination, func(key []byte, value []byte) error {
		result, err := k.UnmarshalClaim(value)
		if err != nil {
			return err
		}

		msg, ok := result.(proto.Message)
		if !ok {
			return status.Errorf(codes.Internal, "can't protomarshal %T", msg)
		}

		claimAny, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			return err
		}
		claims = append(claims, claimAny)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimResponse{Claim: claims, Pagination: pageRes}, nil
}
