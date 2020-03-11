package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/line/link/x/collection/internal/types"
)

type TokenTypeKeeper interface {
	GetNextTokenType(ctx sdk.Context, contractID string) (tokenType string, err error)
	GetTokenTypes(ctx sdk.Context, contractID string) (types.TokenTypes, error)
	GetTokenType(ctx sdk.Context, contractID, tokenType string) (types.TokenType, error)
	HasTokenType(ctx sdk.Context, contractID, tokenType string) bool
	SetTokenType(ctx sdk.Context, token types.TokenType) error
	UpdateTokenType(ctx sdk.Context, token types.TokenType) error
}

var _ TokenTypeKeeper = (*Keeper)(nil)

func (k Keeper) SetTokenType(ctx sdk.Context, tokenType types.TokenType) error {
	_, err := k.GetCollection(ctx, tokenType.GetContractID())
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.TokenTypeKey(tokenType.GetContractID(), tokenType.GetTokenType())) {
		return sdkerrors.Wrapf(types.ErrTokenTypeExist, "ContractID: %s, TokenType: %s", tokenType.GetContractID(), tokenType.GetTokenType())
	}
	store.Set(types.TokenTypeKey(tokenType.GetContractID(), tokenType.GetTokenType()), k.cdc.MustMarshalBinaryBare(tokenType))
	k.setNextTokenTypeNFT(ctx, tokenType.GetContractID(), tokenType.GetTokenType())
	k.setNextTokenIndexNFT(ctx, tokenType.GetContractID(), tokenType.GetTokenType(), types.ReservedEmpty)
	return nil
}

func (k Keeper) UpdateTokenType(ctx sdk.Context, tokenType types.TokenType) error {
	_, err := k.GetCollection(ctx, tokenType.GetContractID())
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.TokenTypeKey(tokenType.GetContractID(), tokenType.GetTokenType())) {
		return sdkerrors.Wrapf(types.ErrTokenTypeNotExist, "ContractID: %s, TokenType: %s", tokenType.GetContractID(), tokenType.GetTokenType())
	}
	store.Set(types.TokenTypeKey(tokenType.GetContractID(), tokenType.GetTokenType()), k.cdc.MustMarshalBinaryBare(tokenType))
	return nil
}

func (k Keeper) GetTokenType(ctx sdk.Context, contractID string, tokenTypeID string) (types.TokenType, error) {
	store := ctx.KVStore(k.storeKey)
	tokenTypeKey := types.TokenTypeKey(contractID, tokenTypeID)
	bz := store.Get(tokenTypeKey)
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrTokenTypeNotExist, "ContractID: %s, TokenType: %s", contractID, tokenTypeID)
	}
	tokenType := k.mustDecodeTokenType(bz)
	return tokenType, nil
}

func (k Keeper) GetTokenTypes(ctx sdk.Context, contractID string) (tokenTypes types.TokenTypes, err error) {
	_, err = k.GetCollection(ctx, contractID)
	if err != nil {
		return nil, err
	}
	k.iterateTokenTypes(ctx, contractID, "", false, func(t types.TokenType) bool {
		tokenTypes = append(tokenTypes, t)
		return false
	})
	return tokenTypes, nil
}

func (k Keeper) HasTokenType(ctx sdk.Context, contractID, tokenType string) bool {
	store := ctx.KVStore(k.storeKey)
	tokenTypeKey := types.TokenTypeKey(contractID, tokenType)
	return store.Has(tokenTypeKey)
}

func (k Keeper) GetNextTokenType(ctx sdk.Context, contractID string) (tokenType string, err error) {
	tokenType, err = k.getNextTokenTypeNFT(ctx, contractID)
	if err != nil {
		return "", err
	}
	if tokenType[0] == types.FungibleFlag[0] {
		return "", sdkerrors.Wrapf(types.ErrTokenTypeFull, "ContractID: %s", contractID)
	}
	return tokenType, nil
}

func (k Keeper) iterateTokenTypes(ctx sdk.Context, contractID, prefix string, reverse bool, process func(types.TokenType) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	var iter sdk.Iterator
	if reverse {
		iter = sdk.KVStoreReversePrefixIterator(store, types.TokenTypeKey(contractID, prefix))
	} else {
		iter = sdk.KVStorePrefixIterator(store, types.TokenTypeKey(contractID, prefix))
	}
	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}
		val := iter.Value()
		tokenType := k.mustDecodeTokenType(val)
		if process(tokenType) {
			return
		}
		iter.Next()
	}
}

func (k Keeper) mustDecodeTokenType(bz []byte) (tokenType types.TokenType) {
	err := k.cdc.UnmarshalBinaryBare(bz, &tokenType)
	if err != nil {
		panic(err)
	}
	return tokenType
}
