package keeper

import (
	"context"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/collection"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServer returns an implementation of the collection MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper Keeper) collection.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ collection.MsgServer = (*msgServer)(nil)

func (s msgServer) Send(c context.Context, req *collection.MsgSend) (*collection.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		To:         req.To,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgSendResponse{}, nil
}

func (s msgServer) OperatorSend(c context.Context, req *collection.MsgOperatorSend) (*collection.MsgOperatorSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Operator)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.Operator,
		From:       req.From,
		To:         req.To,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgOperatorSendResponse{}, nil
}

func (s msgServer) TransferFT(c context.Context, req *collection.MsgTransferFT) (*collection.MsgTransferFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		To:         req.To,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgTransferFTResponse{}, nil
}

func (s msgServer) TransferFTFrom(c context.Context, req *collection.MsgTransferFTFrom) (*collection.MsgTransferFTFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.Proxy,
		From:       req.From,
		To:         req.To,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgTransferFTFromResponse{}, nil
}

func (s msgServer) TransferNFT(c context.Context, req *collection.MsgTransferNFT) (*collection.MsgTransferNFTResponse, error) {
	amount := make([]collection.Coin, len(req.TokenIds))
	for i, id := range req.TokenIds {
		amount[i] = collection.Coin{TokenId: id, Amount: sdk.OneInt()}
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		To:         req.To,
		Amount:     amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgTransferNFTResponse{}, nil
}

func (s msgServer) TransferNFTFrom(c context.Context, req *collection.MsgTransferNFTFrom) (*collection.MsgTransferNFTFromResponse, error) {
	amount := make([]collection.Coin, len(req.TokenIds))
	for i, id := range req.TokenIds {
		amount[i] = collection.Coin{TokenId: id, Amount: sdk.OneInt()}
	}

	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.SendCoins(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.To), amount); err != nil {
		return nil, err
	}

	event := collection.EventSent{
		ContractId: req.ContractId,
		Operator:   req.Proxy,
		From:       req.From,
		To:         req.To,
		Amount:     amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}
	return &collection.MsgTransferNFTFromResponse{}, nil
}

func (s msgServer) AuthorizeOperator(c context.Context, req *collection.MsgAuthorizeOperator) (*collection.MsgAuthorizeOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.AuthorizeOperator(ctx, req.ContractId, sdk.AccAddress(req.Holder), sdk.AccAddress(req.Operator)); err != nil {
		return nil, err
	}

	event := collection.EventAuthorizedOperator{
		ContractId: req.ContractId,
		Holder:     req.Holder,
		Operator:   req.Operator,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgAuthorizeOperatorResponse{}, nil
}

func (s msgServer) RevokeOperator(c context.Context, req *collection.MsgRevokeOperator) (*collection.MsgRevokeOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.RevokeOperator(ctx, req.ContractId, sdk.AccAddress(req.Holder), sdk.AccAddress(req.Operator)); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&collection.EventRevokedOperator{
		ContractId: req.ContractId,
		Holder:     req.Holder,
		Operator:   req.Operator,
	}); err != nil {
		return nil, err
	}

	return &collection.MsgRevokeOperatorResponse{}, nil
}

func (s msgServer) Approve(c context.Context, req *collection.MsgApprove) (*collection.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.AuthorizeOperator(ctx, req.ContractId, sdk.AccAddress(req.Approver), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, err
	}

	event := collection.EventAuthorizedOperator{
		ContractId: req.ContractId,
		Holder:     req.Approver,
		Operator:   req.Proxy,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgApproveResponse{}, nil
}

func (s msgServer) Disapprove(c context.Context, req *collection.MsgDisapprove) (*collection.MsgDisapproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := s.keeper.RevokeOperator(ctx, req.ContractId, sdk.AccAddress(req.Approver), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, err
	}

	event := collection.EventRevokedOperator{
		ContractId: req.ContractId,
		Holder:     req.Approver,
		Operator:   req.Proxy,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgDisapproveResponse{}, nil
}

func (s msgServer) CreateContract(c context.Context, req *collection.MsgCreateContract) (*collection.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	contract := collection.Contract{
		Name:       req.Name,
		BaseImgUri: req.BaseImgUri,
		Meta:       req.Meta,
	}
	id := s.keeper.CreateContract(ctx, sdk.AccAddress(req.Owner), contract)

	return &collection.MsgCreateContractResponse{Id: id}, nil
}

func (s msgServer) IssueFT(c context.Context, req *collection.MsgIssueFT) (*collection.MsgIssueFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.Owner), collection.Permission_Issue); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	class := &collection.FTClass{
		Name:     req.Name,
		Meta:     req.Meta,
		Decimals: req.Decimals,
		Mintable: req.Mintable,
	}
	id, err := s.keeper.CreateTokenClass(ctx, req.ContractId, class)
	if err != nil {
		return nil, err
	}

	return &collection.MsgIssueFTResponse{Id: *id}, nil
}

func (s msgServer) IssueNFT(c context.Context, req *collection.MsgIssueNFT) (*collection.MsgIssueNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.Owner), collection.Permission_Issue); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	class := &collection.NFTClass{
		Name: req.Name,
		Meta: req.Meta,
	}
	id, err := s.keeper.CreateTokenClass(ctx, req.ContractId, class)
	if err != nil {
		return nil, err
	}

	return &collection.MsgIssueNFTResponse{Id: *id}, nil
}

func (s msgServer) MintFT(c context.Context, req *collection.MsgMintFT) (*collection.MsgMintFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.From), collection.Permission_Mint); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.mintFT(ctx, req.ContractId, sdk.AccAddress(req.To), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventMintedFT{
		ContractId: req.ContractId,
		Operator:   req.From,
		To:         req.To,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgMintFTResponse{}, nil
}

func (s msgServer) MintNFT(c context.Context, req *collection.MsgMintNFT) (*collection.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.From), collection.Permission_Mint); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	tokens, err := s.keeper.mintNFT(ctx, req.ContractId, sdk.AccAddress(req.To), req.Params)
	if err != nil {
		return nil, err
	}

	event := collection.EventMintedNFT{
		ContractId: req.ContractId,
		Operator:   req.From,
		To:         req.To,
		Tokens:     tokens,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	tokenIDs := make([]string, 0, len(tokens))
	for _, token := range tokens {
		tokenIDs = append(tokenIDs, token.Id)
	}
	return &collection.MsgMintNFTResponse{Ids: tokenIDs}, nil
}

func (s msgServer) Burn(c context.Context, req *collection.MsgBurn) (*collection.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.From), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgBurnResponse{}, nil
}

func (s msgServer) OperatorBurn(c context.Context, req *collection.MsgOperatorBurn) (*collection.MsgOperatorBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Operator)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.Operator), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.Operator,
		From:       req.From,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgOperatorBurnResponse{}, nil
}

func (s msgServer) BurnFT(c context.Context, req *collection.MsgBurnFT) (*collection.MsgBurnFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.From), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgBurnFTResponse{}, nil
}

func (s msgServer) BurnFTFrom(c context.Context, req *collection.MsgBurnFTFrom) (*collection.MsgBurnFTFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.Proxy), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), req.Amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.Proxy,
		From:       req.From,
		Amount:     req.Amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgBurnFTFromResponse{}, nil
}

func (s msgServer) BurnNFT(c context.Context, req *collection.MsgBurnNFT) (*collection.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.From), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	amount := make([]collection.Coin, 0, len(req.TokenIds))
	for _, id := range req.TokenIds {
		amount = append(amount, collection.NewCoin(id, sdk.OneInt()))
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.From,
		From:       req.From,
		Amount:     amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgBurnNFTResponse{}, nil
}

func (s msgServer) BurnNFTFrom(c context.Context, req *collection.MsgBurnNFTFrom) (*collection.MsgBurnNFTFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if _, err := s.keeper.GetAuthorization(ctx, req.ContractId, sdk.AccAddress(req.From), sdk.AccAddress(req.Proxy)); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	if _, err := s.keeper.GetGrant(ctx, req.ContractId, sdk.AccAddress(req.Proxy), collection.Permission_Burn); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrap(err.Error())
	}

	amount := make([]collection.Coin, 0, len(req.TokenIds))
	for _, id := range req.TokenIds {
		amount = append(amount, collection.NewCoin(id, sdk.OneInt()))
	}

	if err := s.keeper.burnCoins(ctx, req.ContractId, sdk.AccAddress(req.From), amount); err != nil {
		return nil, err
	}

	event := collection.EventBurned{
		ContractId: req.ContractId,
		Operator:   req.Proxy,
		From:       req.From,
		Amount:     amount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return nil, err
	}

	return &collection.MsgBurnNFTFromResponse{}, nil
}

func (s msgServer) ModifyContract(c context.Context, req *collection.MsgModifyContract) (*collection.MsgModifyContractResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) ModifyTokenClass(c context.Context, req *collection.MsgModifyTokenClass) (*collection.MsgModifyTokenClassResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) ModifyNFT(c context.Context, req *collection.MsgModifyNFT) (*collection.MsgModifyNFTResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) Modify(c context.Context, req *collection.MsgModify) (*collection.MsgModifyResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) Grant(c context.Context, req *collection.MsgGrant) (*collection.MsgGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	granter := sdk.AccAddress(req.Granter)
	grantee := sdk.AccAddress(req.Grantee)
	permission := collection.Permission(collection.Permission_value[req.Permission])
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, granter, permission); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not authorized for %s", granter, permission.String())
	}
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, grantee, permission); err == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is already granted for %s", grantee, permission.String())
	}

	if err := s.keeper.Grant(ctx, req.ContractId, granter, grantee, permission); err != nil {
		return nil, err
	}

	return &collection.MsgGrantResponse{}, nil
}

func (s msgServer) Abandon(c context.Context, req *collection.MsgAbandon) (*collection.MsgAbandonResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	grantee := sdk.AccAddress(req.Grantee)
	permission := collection.Permission(collection.Permission_value[req.Permission])
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, grantee, permission); err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("%s is not authorized for %s", grantee, permission)
	}

	if err := s.keeper.Abandon(ctx, req.ContractId, grantee, permission); err != nil {
		return nil, err
	}

	return &collection.MsgAbandonResponse{}, nil
}

func (s msgServer) GrantPermission(c context.Context, req *collection.MsgGrantPermission) (*collection.MsgGrantPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	granter := sdk.AccAddress(req.From)
	grantee := sdk.AccAddress(req.To)
	permission := collection.Permission(collection.Permission_value[req.Permission])
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, granter, permission); err != nil {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not authorized for %s", granter, permission.String())
	}
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, grantee, permission); err == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is already granted for %s", grantee, permission.String())
	}

	if err := s.keeper.Grant(ctx, req.ContractId, granter, grantee, permission); err != nil {
		return nil, err
	}

	return &collection.MsgGrantPermissionResponse{}, nil
}

func (s msgServer) RevokePermission(c context.Context, req *collection.MsgRevokePermission) (*collection.MsgRevokePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	grantee := sdk.AccAddress(req.From)
	permission := collection.Permission(collection.Permission_value[req.Permission])
	if _, err := s.keeper.GetGrant(ctx, req.ContractId, grantee, permission); err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("%s is not authorized for %s", grantee, permission)
	}

	if err := s.keeper.Abandon(ctx, req.ContractId, grantee, permission); err != nil {
		return nil, err
	}

	return &collection.MsgRevokePermissionResponse{}, nil
}

func (s msgServer) Attach(c context.Context, req *collection.MsgAttach) (*collection.MsgAttachResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) Detach(c context.Context, req *collection.MsgDetach) (*collection.MsgDetachResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) OperatorAttach(c context.Context, req *collection.MsgOperatorAttach) (*collection.MsgOperatorAttachResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) OperatorDetach(c context.Context, req *collection.MsgOperatorDetach) (*collection.MsgOperatorDetachResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) AttachFrom(c context.Context, req *collection.MsgAttachFrom) (*collection.MsgAttachFromResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}

func (s msgServer) DetachFrom(c context.Context, req *collection.MsgDetachFrom) (*collection.MsgDetachFromResponse, error) {
	return nil, sdkerrors.ErrNotSupported
}
