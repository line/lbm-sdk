package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/line/lbm-sdk/v2/store/prefix"
	sdk "github.com/line/lbm-sdk/v2/types"
	"github.com/line/lbm-sdk/v2/types/query"
	"github.com/line/lbm-sdk/v2/x/slashing/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) SigningInfo(c context.Context, req *types.QuerySigningInfoRequest) (*types.QuerySigningInfoResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.ConsAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	consAddr, err := sdk.ConsAddressFromBech32(req.ConsAddress)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	signingInfo, found := k.GetValidatorSigningInfo(ctx, consAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "SigningInfo not found for validator %s", req.ConsAddress)
	}

	return &types.QuerySigningInfoResponse{ValSigningInfo: signingInfo}, nil
}

func (k Keeper) SigningInfos(c context.Context, req *types.QuerySigningInfosRequest) (*types.QuerySigningInfosResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	var signInfos []types.ValidatorSigningInfo

	sigInfoStore := prefix.NewStore(store, types.ValidatorSigningInfoKeyPrefix)
	pageRes, err := query.Paginate(sigInfoStore, req.Pagination, func(key []byte, value interface{}) error {
				info := *value.(*types.ValidatorSigningInfo)
				signInfos = append(signInfos, info)
				return nil
			},
			GetValidatorSigningInfoUnmarshalFunc(k.cdc))
	if err != nil {
		return nil, err
	}
	return &types.QuerySigningInfosResponse{Info: signInfos, Pagination: pageRes}, nil
}
