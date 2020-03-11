package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	iam "github.com/line/link/x/iam/exported"
	"github.com/line/link/x/safetybox/internal/types"
	"github.com/tendermint/tendermint/crypto"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	iamKeeper     iam.IamKeeper
	bankKeeper    types.BankKeeper
	hooks         types.SafetyBoxHooks
	accountKeeper auth.AccountKeeper
}

func NewKeeper(cdc *codec.Codec, iamKeeper iam.IamKeeper, bankKeeper types.BankKeeper, accountKeeper auth.AccountKeeper, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		iamKeeper:     iamKeeper,
		bankKeeper:    bankKeeper,
		hooks:         nil,
		accountKeeper: accountKeeper,
	}
}

// Set the hooks
func (k *Keeper) SetHooks(sh types.SafetyBoxHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set the safety box hooks twice")
	}
	k.hooks = sh
	return k
}

func (k Keeper) NewSafetyBox(ctx sdk.Context, msg types.MsgSafetyBoxCreate) (types.SafetyBox, error) {
	// create new safety box account
	newSafetyBoxAccount, err := k.newSafetyBoxAccount(ctx, msg.SafetyBoxID)
	if err != nil {
		return types.SafetyBox{}, err
	}

	if len(msg.SafetyBoxDenoms) > 1 {
		return types.SafetyBox{}, sdkerrors.Wrapf(types.ErrSafetyBoxTooManyCoinDenoms, "Requested: %v", msg.SafetyBoxDenoms)
	}

	// create new safety box
	sb := types.NewSafetyBox(msg.SafetyBoxOwner, msg.SafetyBoxID, newSafetyBoxAccount, msg.SafetyBoxDenoms)

	// reject if the safety box id exists
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.SafetyBoxKey(sb.ID)) {
		return types.SafetyBox{}, sdkerrors.Wrapf(types.ErrSafetyBoxIDExist, "ID: %s", sb.ID)
	}
	store.Set(types.SafetyBoxKey(sb.ID), k.cdc.MustMarshalBinaryBare(sb))

	// grant the owner a permission to whitelist operators
	k.iamKeeper.GrantPermission(ctx, sb.Owner, types.NewWhitelistOperatorsPermission(sb.ID))

	// call the after-creation hooks if any
	if k.hooks != nil {
		k.hooks.AfterSafetyBoxCreated(ctx, sb.Address)
	}

	return sb, nil
}

func (k Keeper) newSafetyBoxAccount(ctx sdk.Context, safetyBoxID string) (sdk.AccAddress, error) {
	// hash safety box id
	newAddress := sdk.AccAddress(crypto.AddressHash(types.SafetyBoxKey(safetyBoxID)))

	// check if exist
	acc := k.accountKeeper.GetAccount(ctx, newAddress)
	if acc != nil {
		return nil, sdkerrors.Wrapf(types.ErrSafetyBoxAccountExist, "ID: %s", safetyBoxID)
	}

	// create new account and return its address
	newAccount := k.accountKeeper.NewAccountWithAddress(ctx, newAddress)
	k.accountKeeper.SetAccount(ctx, newAccount)

	return newAccount.GetAddress(), nil
}

func (k Keeper) GetSafetyBox(ctx sdk.Context, safetyBoxID string) (types.SafetyBox, error) {
	sb, err := k.get(ctx, safetyBoxID)
	if err != nil {
		return types.SafetyBox{}, err
	}
	return sb, nil
}

func (k Keeper) validDenom(coins sdk.Coins, denoms []string) error {
	// safety box accepts only one type of coins
	if len(coins) != 1 || len(denoms) != 1 {
		return sdkerrors.Wrapf(types.ErrSafetyBoxTooManyCoinDenoms, "Requested: %v", denoms)
	}
	if coins[0].Denom != denoms[0] {
		return sdkerrors.Wrapf(types.ErrSafetyBoxIncorrectDenom, "Expected: %s, Requested: %s", denoms[0], coins[0].Denom)
	}
	return nil
}

func (k Keeper) Allocate(ctx sdk.Context, msg types.MsgSafetyBoxAllocateCoins) error {
	sb, err := k.get(ctx, msg.SafetyBoxID)
	if err != nil {
		return err
	}

	// safety box accepts only one type of coins
	if err = k.validDenom(msg.Coins, sb.Denoms); err != nil {
		return err
	}

	// from the allocator, to the safety box
	fromAddress := msg.AllocatorAddress
	toAddress := sb.Address

	// only allocator could allocate
	allocatePermission := types.NewAllocatePermission(msg.SafetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, fromAddress, allocatePermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionAllocate, "Account: %s", fromAddress.String())
	}

	// allocation
	err = k.sendCoins(ctx, fromAddress, toAddress, msg.Coins)
	if err != nil {
		return err
	}

	// increase the total allocation and cumulative allocation
	sb.TotalAllocation = sb.TotalAllocation.Add(msg.Coins...)
	sb.CumulativeAllocation = sb.CumulativeAllocation.Add(msg.Coins...)

	return k.set(ctx, msg.SafetyBoxID, sb)
}

//nolint:dupl
func (k Keeper) Recall(ctx sdk.Context, msg types.MsgSafetyBoxRecallCoins) error {
	sb, err := k.get(ctx, msg.SafetyBoxID)
	if err != nil {
		return err
	}

	// safety box accepts only one type of coins
	if err = k.validDenom(msg.Coins, sb.Denoms); err != nil {
		return err
	}

	// from the safety box, to the allocator
	fromAddress := sb.Address
	toAddress := msg.AllocatorAddress

	// only allocator could recall
	recallPermission := types.NewRecallPermission(msg.SafetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, toAddress, recallPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionRecall, "Account: %s", toAddress.String())
	}

	// check not to recall more than allocated
	if msg.Coins.IsAnyGT(sb.TotalAllocation) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxRecallMoreThanAllocated, "Has: %v, Requested: %v", sb.TotalAllocation, msg.Coins)
	}

	// recall
	err = k.sendCoins(ctx, fromAddress, toAddress, msg.Coins)
	if err != nil {
		return err
	}

	// decrease the total allocation
	sb.TotalAllocation = sb.TotalAllocation.Sub(msg.Coins)

	return k.set(ctx, msg.SafetyBoxID, sb)
}

func (k Keeper) Issue(ctx sdk.Context, msg types.MsgSafetyBoxIssueCoins) error {
	sb, err := k.get(ctx, msg.SafetyBoxID)
	if err != nil {
		return err
	}

	// safety box accepts only one type of coins
	if err = k.validDenom(msg.Coins, sb.Denoms); err != nil {
		return err
	}

	// both issuer and issuee must be issuers
	issuerAddress := msg.FromAddress
	toAddress := msg.ToAddress

	issuePermission := types.NewIssuePermission(msg.SafetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, issuerAddress, issuePermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionIssue, "Account: %s", issuerAddress.String())
	}
	if !k.iamKeeper.HasPermission(ctx, toAddress, issuePermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionIssue, "Account: %s", toAddress.String())
	}

	// issue from the safety box to an issuer
	fromAddress := sb.Address
	err = k.sendCoins(ctx, fromAddress, toAddress, msg.Coins)
	if err != nil {
		return err
	}

	// increase the total issuance
	sb.TotalIssuance = sb.TotalIssuance.Add(msg.Coins...)

	return k.set(ctx, msg.SafetyBoxID, sb)
}

//nolint:dupl
func (k Keeper) Return(ctx sdk.Context, msg types.MsgSafetyBoxReturnCoins) error {
	sb, err := k.get(ctx, msg.SafetyBoxID)
	if err != nil {
		return err
	}

	// safety box accepts only one type of coins
	if err = k.validDenom(msg.Coins, sb.Denoms); err != nil {
		return err
	}

	// from the returner, to the safety box
	fromAddress := msg.ReturnerAddress
	toAddress := sb.Address

	// only returner could return
	returnPermission := types.NewReturnPermission(msg.SafetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, fromAddress, returnPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionReturn, "Account: %s", fromAddress.String())
	}

	// check not to return more than issued
	if msg.Coins.IsAnyGT(sb.TotalIssuance) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxReturnMoreThanIssued, "Has: %v, Requested: %v", sb.TotalIssuance, msg.Coins)
	}

	// return
	err = k.sendCoins(ctx, fromAddress, toAddress, msg.Coins)
	if err != nil {
		return err
	}

	// decrease the total issuance
	sb.TotalIssuance = sb.TotalIssuance.Sub(msg.Coins)

	return k.set(ctx, msg.SafetyBoxID, sb)
}

func (k Keeper) sendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error {
	_, err := k.bankKeeper.SubtractCoins(ctx, fromAddr, coins)
	if err != nil {
		return err
	}

	_, err = k.bankKeeper.AddCoins(ctx, toAddr, coins)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeySender, fromAddr.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, toAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
		),
	})

	return nil
}

func (k Keeper) GrantPermission(ctx sdk.Context, safetyBoxID string, by sdk.AccAddress, acc sdk.AccAddress, role string) error {
	// reject self-grant
	if by.Equals(acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxSelfPermission, "Account: %s", acc.String())
	}

	// grant
	switch role {
	case types.RoleOperator:
		return k.grantOperator(ctx, safetyBoxID, by, acc)
	case types.RoleAllocator:
		return k.grantAllocator(ctx, safetyBoxID, by, acc)
	case types.RoleIssuer:
		return k.grantIssuer(ctx, safetyBoxID, by, acc)
	case types.RoleReturner:
		return k.grantReturner(ctx, safetyBoxID, by, acc)
	default:
		return nil
	}
}

func (k Keeper) grantOperator(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOperatorsPermission := types.NewWhitelistOperatorsPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOperatorsPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	// check if the target is eligible
	if k.IsOperator(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasPermissionAlready, "Account: %s", acc.String())
	}
	if k.IsAllocator(ctx, safetyBoxID, acc) || k.IsIssuer(ctx, safetyBoxID, acc) || k.IsReturner(ctx, safetyBoxID, acc) || k.IsOwner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasOtherPermission, "Account: %s", acc.String())
	}

	// grant
	k.iamKeeper.GrantPermission(ctx, acc, types.NewWhitelistOtherRolesPermission(safetyBoxID))
	return nil
}

func (k Keeper) grantAllocator(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	// check if the target is eligible
	if k.IsAllocator(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasPermissionAlready, "Account: %s", acc.String())
	}
	if k.IsOperator(ctx, safetyBoxID, acc) || k.IsIssuer(ctx, safetyBoxID, acc) || k.IsReturner(ctx, safetyBoxID, acc) || k.IsOwner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasOtherPermission, "Account: %s", acc.String())
	}

	// grant - allocator may allocate and recall
	k.iamKeeper.GrantPermission(ctx, acc, types.NewAllocatePermission(safetyBoxID))
	k.iamKeeper.GrantPermission(ctx, acc, types.NewRecallPermission(safetyBoxID))
	return nil
}

func (k Keeper) grantIssuer(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	// check if the target is eligible
	if k.IsIssuer(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasPermissionAlready, "Account: %s", acc.String())
	}
	if k.IsOperator(ctx, safetyBoxID, acc) || k.IsReturner(ctx, safetyBoxID, acc) || k.IsAllocator(ctx, safetyBoxID, acc) || k.IsOwner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasOtherPermission, "Account: %s", acc.String())
	}

	// grant
	k.iamKeeper.GrantPermission(ctx, acc, types.NewIssuePermission(safetyBoxID))
	return nil
}

func (k Keeper) grantReturner(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	// check if the target is eligible
	if k.IsReturner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasPermissionAlready, "Account: %s", acc.String())
	}
	if k.IsOperator(ctx, safetyBoxID, acc) || k.IsIssuer(ctx, safetyBoxID, acc) || k.IsAllocator(ctx, safetyBoxID, acc) || k.IsOwner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxHasOtherPermission, "Account: %s", acc.String())
	}

	// grant
	k.iamKeeper.GrantPermission(ctx, acc, types.NewReturnPermission(safetyBoxID))
	return nil
}

func (k Keeper) RevokePermission(ctx sdk.Context, safetyBoxID string, by sdk.AccAddress, acc sdk.AccAddress, role string) error {
	// reject self-revoke
	if by.Equals(acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxSelfPermission, "Account: %s", acc.String())
	}

	// revoke
	switch role {
	case types.RoleOperator:
		return k.revokeOperator(ctx, safetyBoxID, by, acc)
	case types.RoleAllocator:
		return k.revokeAllocator(ctx, safetyBoxID, by, acc)
	case types.RoleIssuer:
		return k.revokeIssuer(ctx, safetyBoxID, by, acc)
	case types.RoleReturner:
		return k.revokeReturner(ctx, safetyBoxID, by, acc)
	default:
		return nil
	}
}

func (k Keeper) revokeOperator(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOperatorsPermission := types.NewWhitelistOperatorsPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOperatorsPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	if !k.IsOperator(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxDoesNotHavePermission, "Account: %s", acc.String())
	}
	k.iamKeeper.RevokePermission(ctx, acc, types.NewWhitelistOtherRolesPermission(safetyBoxID))
	return nil
}

func (k Keeper) revokeAllocator(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	if !k.IsAllocator(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxDoesNotHavePermission, "Account: %s", acc.String())
	}
	k.iamKeeper.RevokePermission(ctx, acc, types.NewAllocatePermission(safetyBoxID))
	k.iamKeeper.RevokePermission(ctx, acc, types.NewRecallPermission(safetyBoxID))
	return nil
}

func (k Keeper) revokeIssuer(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	if !k.IsIssuer(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxDoesNotHavePermission, "Account: %s", acc.String())
	}
	k.iamKeeper.RevokePermission(ctx, acc, types.NewIssuePermission(safetyBoxID))
	return nil
}

func (k Keeper) revokeReturner(ctx sdk.Context, safetyBoxID string, by, acc sdk.AccAddress) error {
	// check whitelist permission
	whitelistOtherRolesPermission := types.NewWhitelistOtherRolesPermission(safetyBoxID)
	if !k.iamKeeper.HasPermission(ctx, by, whitelistOtherRolesPermission) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxPermissionWhitelist, "Account: %s", by.String())
	}

	if !k.IsReturner(ctx, safetyBoxID, acc) {
		return sdkerrors.Wrapf(types.ErrSafetyBoxDoesNotHavePermission, "Account: %s", acc.String())
	}
	k.iamKeeper.RevokePermission(ctx, acc, types.NewReturnPermission(safetyBoxID))
	return nil
}

func (k Keeper) get(ctx sdk.Context, safetyBoxID string) (types.SafetyBox, error) {
	store := ctx.KVStore(k.storeKey)

	// retrieve the safety box
	bz := store.Get(types.SafetyBoxKey(safetyBoxID))
	if bz == nil {
		return types.SafetyBox{}, sdkerrors.Wrapf(types.ErrSafetyBoxNotExist, "ID: %s", safetyBoxID)
	}
	r := &types.SafetyBox{}
	k.cdc.MustUnmarshalBinaryBare(bz, r)
	return *r, nil
}

func (k Keeper) set(ctx sdk.Context, safetyBoxID string, sb types.SafetyBox) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SafetyBoxKey(safetyBoxID), k.cdc.MustMarshalBinaryBare(sb))
	return nil
}

func (k Keeper) GetPermissions(ctx sdk.Context, safetyBoxID, role string, acc sdk.AccAddress) (types.MsgSafetyBoxRoleResponse, error) {
	var hasRole bool
	switch role {
	case types.RoleOwner:
		hasRole = k.IsOwner(ctx, safetyBoxID, acc)
	case types.RoleOperator:
		hasRole = k.IsOperator(ctx, safetyBoxID, acc)
	case types.RoleAllocator:
		hasRole = k.IsAllocator(ctx, safetyBoxID, acc)
	case types.RoleIssuer:
		hasRole = k.IsIssuer(ctx, safetyBoxID, acc)
	case types.RoleReturner:
		hasRole = k.IsReturner(ctx, safetyBoxID, acc)
	default:
		return types.MsgSafetyBoxRoleResponse{HasRole: false}, sdkerrors.Wrapf(types.ErrSafetyBoxInvalidRole, "Role: %s", role)
	}
	return types.MsgSafetyBoxRoleResponse{HasRole: hasRole}, nil
}

func (k Keeper) IsOwner(ctx sdk.Context, safetyBoxID string, acc sdk.Address) bool {
	sb, err := k.get(ctx, safetyBoxID)
	if err != nil {
		return false
	}

	return sb.Owner.Equals(acc)
}

func (k Keeper) IsOperator(ctx sdk.Context, safetyBoxID string, acc sdk.AccAddress) bool {
	return k.iamKeeper.HasPermission(ctx, acc, types.NewWhitelistOtherRolesPermission(safetyBoxID))
}

func (k Keeper) IsAllocator(ctx sdk.Context, safetyBoxID string, acc sdk.AccAddress) bool {
	canAllocate := k.iamKeeper.HasPermission(ctx, acc, types.NewAllocatePermission(safetyBoxID))
	canRecall := k.iamKeeper.HasPermission(ctx, acc, types.NewRecallPermission(safetyBoxID))
	return canAllocate && canRecall
}

func (k Keeper) IsIssuer(ctx sdk.Context, safetyBoxID string, acc sdk.AccAddress) bool {
	return k.iamKeeper.HasPermission(ctx, acc, types.NewIssuePermission(safetyBoxID))
}

func (k Keeper) IsReturner(ctx sdk.Context, safetyBoxID string, acc sdk.AccAddress) bool {
	return k.iamKeeper.HasPermission(ctx, acc, types.NewReturnPermission(safetyBoxID))
}
