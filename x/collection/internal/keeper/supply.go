package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/line/link/x/collection/internal/types"
)

type SupplyKeeper interface {
	GetTotalInt(ctx sdk.Context, contractID, tokenID, target string) (supply sdk.Int, err error)
	GetSupply(ctx sdk.Context, contractID string) (supply types.Supply)
	SetSupply(ctx sdk.Context, supply types.Supply)
	MintSupply(ctx sdk.Context, contractID string, to sdk.AccAddress, amt types.Coins) error
	BurnSupply(ctx sdk.Context, contractID string, from sdk.AccAddress, amt types.Coins) error
}

var _ SupplyKeeper = (*Keeper)(nil)

func (k Keeper) GetSupply(ctx sdk.Context, contractID string) (supply types.Supply) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.SupplyKey(contractID))
	if b == nil {
		panic("stored supply should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &supply)
	return
}

func (k Keeper) SetSupply(ctx sdk.Context, supply types.Supply) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(supply)
	store.Set(types.SupplyKey(supply.GetContractID()), b)
}

func (k Keeper) GetTotalInt(ctx sdk.Context, contractID, tokenID, target string) (supply sdk.Int, err error) {
	if _, err = k.GetToken(ctx, contractID, tokenID); err != nil {
		return sdk.NewInt(0), err
	}

	s := k.GetSupply(ctx, contractID)
	switch target {
	case types.QuerySupply:
		return s.GetTotalSupply().AmountOf(tokenID), nil
	case types.QueryBurn:
		return s.GetTotalBurn().AmountOf(tokenID), nil
	case types.QueryMint:
		return s.GetTotalMint().AmountOf(tokenID), nil
	default:
		return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid request target to query total %s", target)
	}
}

// MintCoins creates new coins from thin air and adds it to the module account.
// Panics if the name maps to a non-minter module account or if the amount is invalid.
func (k Keeper) MintSupply(ctx sdk.Context, contractID string, to sdk.AccAddress, amt types.Coins) error {
	_, err := k.AddCoins(ctx, contractID, to, amt)
	if err != nil {
		return err
	}
	supply := k.GetSupply(ctx, contractID)
	supply = supply.Inflate(amt)
	// supply should never be negative. Big.Int.Add will be panic if it becomes overflow

	k.SetSupply(ctx, supply)
	return nil
}

// BurnCoins burns coins deletes coins from the balance of the module account.
// Panics if the name maps to a non-burner module account or if the amount is invalid.
func (k Keeper) BurnSupply(ctx sdk.Context, contractID string, from sdk.AccAddress, amt types.Coins) error {
	_, err := k.SubtractCoins(ctx, contractID, from, amt)
	if err != nil {
		return err
	}
	supply := k.GetSupply(ctx, contractID)
	supply = supply.Deflate(amt)
	if supply.GetTotalSupply().IsAnyNegative() {
		return sdkerrors.Wrapf(types.ErrInsufficientSupply, "insufficient supply for token [%s]", contractID)
	}
	k.SetSupply(ctx, supply)

	return nil
}
