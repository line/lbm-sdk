package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/line/link/x/contract"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func TestMsgBasics(t *testing.T) {
	cdc := ModuleCdc

	{
		msg := NewMsgIssueFT(addr1, addr1, defaultContractID, defaultName, defaultMeta, sdk.NewInt(1), sdk.NewInt(8), true)
		require.Equal(t, "issue_ft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgIssueFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Name, msg2.Name)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.Owner, msg2.Owner)
		require.Equal(t, msg.Amount, msg.Amount)
		require.Equal(t, msg.Decimals, msg2.Decimals)
		require.Equal(t, msg.Mintable, msg2.Mintable)
	}
	{
		msg := NewMsgIssueNFT(addr1, defaultContractID, defaultName, defaultMeta)
		require.Equal(t, "issue_nft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgIssueNFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.Owner, msg2.Owner)
		require.Equal(t, msg.Name, msg2.Name)
	}
	{
		msg := NewMsgMintNFT(addr1, defaultContractID, addr1, defaultName, defaultMeta, defaultTokenType)
		require.Equal(t, "mint_nft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgMintNFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Name, msg2.Name)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.TokenType, msg2.TokenType)
	}
	{
		msg := NewMsgBurnNFT(addr1, defaultContractID, defaultTokenID1)
		require.Equal(t, "burn_nft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgBurnNFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.TokenIDs, msg2.TokenIDs)
	}
	{
		addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
		msg := NewMsgGrantPermission(addr1, addr2, Permission{Action: "issue", Resource: "resource"})
		require.Equal(t, "grant_perm", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgGrantPermission{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.To, msg2.To)
		require.Equal(t, msg.Permission, msg2.Permission)
	}

	{
		msg := NewMsgRevokePermission(addr1, Permission{Action: "issue", Resource: "resource"})
		require.Equal(t, "revoke_perm", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgRevokePermission{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.Permission, msg2.Permission)
	}
	{
		msg := NewMsgTransferFT(addr1, defaultContractID, addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.Equal(t, "transfer_ft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgTransferFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.To, msg2.To)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.Amount, msg2.Amount)
	}

	{
		msg := NewMsgTransferFT(nil, defaultContractID, addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgTransferFT(addr1, defaultContractID, nil, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "To cannot be empty").Error())

		msg = NewMsgTransferFT(addr1, "", addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: ").Error())

		require.Panics(t, func() {
			NewMsgTransferFT(addr1, defaultContractID, addr2, NewCoin("1", sdk.NewInt(defaultAmount)))
		}, "")

		require.Panics(t, func() {
			NewMsgTransferFT(addr1, defaultContractID, addr2, NewCoin("1", sdk.NewInt(-1*defaultAmount)))
		}, "")
	}

	{
		msg := NewMsgTransferNFT(addr1, defaultContractID, addr2, defaultTokenID1)
		require.Equal(t, "transfer_nft", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgTransferNFT{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.To, msg2.To)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenIDs, msg2.TokenIDs)
	}

	{
		msg := NewMsgTransferNFT(nil, defaultContractID, addr2, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgTransferNFT(addr1, defaultContractID, nil, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "To cannot be empty").Error())

		msg = NewMsgTransferNFT(addr1, "", addr2, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: ").Error())

		msg = NewMsgTransferNFT(addr1, defaultContractID, addr2, "1")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(ErrInvalidTokenID, "symbol [1] mismatched to [^[a-f0-9]{16}$]").Error())
	}

	{
		msg := NewMsgTransferFTFrom(addr1, defaultContractID, addr2, addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.Equal(t, "transfer_ft_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgTransferFTFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.To, msg2.To)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.Amount, msg2.Amount)
	}

	{
		msg := NewMsgTransferFTFrom(nil, defaultContractID, addr2, addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgTransferFTFrom(addr1, defaultContractID, nil, addr2, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgTransferFTFrom(addr1, defaultContractID, addr2, nil, NewCoin(defaultTokenIDFT, sdk.NewInt(defaultAmount)))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "To cannot be empty").Error())

		require.Panics(t, func() {
			NewMsgTransferFT(addr1, defaultContractID, addr2, NewCoin("1", sdk.NewInt(defaultAmount)))
		}, "")

		require.Panics(t, func() {
			NewMsgTransferFT(addr1, defaultContractID, addr2, NewCoin("1", sdk.NewInt(-1*defaultAmount)))
		}, "")
	}
	//nolint:dupl
	{
		msg := NewMsgTransferNFTFrom(addr1, defaultContractID, addr2, addr2, defaultTokenID1)
		require.Equal(t, "transfer_nft_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgTransferNFTFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.To, msg2.To)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenIDs, msg2.TokenIDs)
	}

	{
		msg := NewMsgTransferNFTFrom(nil, defaultContractID, addr2, addr2, defaultTokenIDFT)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgTransferNFTFrom(addr1, defaultContractID, nil, addr2, defaultTokenIDFT)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgTransferNFTFrom(addr1, defaultContractID, addr2, nil, defaultTokenIDFT)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "To cannot be empty").Error())

		msg = NewMsgTransferNFTFrom(addr1, defaultContractID, addr2, addr2, "1")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(ErrInvalidTokenID, "symbol [1] mismatched to [^[a-f0-9]{16}$]").Error())
	}

	{
		msg := NewMsgAttach(addr1, defaultContractID, defaultTokenID1, defaultTokenID2)
		require.Equal(t, "attach", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgAttach{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.ToTokenID, msg2.ToTokenID)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenID, msg2.TokenID)
	}

	{
		msg := NewMsgAttach(nil, defaultContractID, defaultTokenID1, defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgAttach(addr1, "s", defaultTokenID1, defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: s").Error())

		msg = NewMsgAttach(addr1, defaultContractID, "1", defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "1").Error())

		msg = NewMsgAttach(addr1, defaultContractID, defaultTokenID1, "2")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "2").Error())

		msg = NewMsgAttach(addr1, defaultContractID, defaultTokenID1, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrapf(ErrCannotAttachToItself, "TokenID: %s", defaultTokenID1).Error())
	}

	{
		msg := NewMsgDetach(addr1, defaultContractID, defaultTokenID1)
		require.Equal(t, "detach", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgDetach{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenID, msg2.TokenID)
	}

	{
		msg := NewMsgDetach(nil, defaultContractID, "item0001")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgDetach(addr1, "s", "item0001")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: s").Error())

		msg = NewMsgDetach(addr1, defaultContractID, "1")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "1").Error())
	}
	//nolint:dupl
	{
		msg := NewMsgAttachFrom(addr1, defaultContractID, addr2, defaultTokenID1, defaultTokenID2)
		require.Equal(t, "attach_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgAttachFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.ToTokenID, msg2.ToTokenID)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenID, msg2.TokenID)
	}

	{
		msg := NewMsgAttachFrom(nil, defaultContractID, addr2, defaultTokenID1, defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgAttachFrom(addr1, defaultContractID, nil, defaultTokenID1, defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgAttachFrom(addr1, "s", addr2, defaultTokenID1, defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: s").Error())

		msg = NewMsgAttachFrom(addr1, defaultContractID, addr2, "1", defaultTokenID2)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "1").Error())

		msg = NewMsgAttachFrom(addr1, defaultContractID, addr2, defaultTokenID1, "2")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "2").Error())

		msg = NewMsgAttachFrom(addr1, defaultContractID, addr2, defaultTokenID1, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrapf(ErrCannotAttachToItself, "TokenID: %s", defaultTokenID1).Error())
	}
	//nolint:dupl
	{
		msg := NewMsgDetachFrom(addr1, defaultContractID, addr2, defaultTokenID1)
		require.Equal(t, "detach_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgDetachFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenID, msg2.TokenID)
	}

	{
		msg := NewMsgDetachFrom(nil, defaultContractID, addr2, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgDetachFrom(addr1, defaultContractID, nil, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())

		msg = NewMsgDetachFrom(addr1, "s", addr2, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(contract.ErrInvalidContractID, "ContractID: s").Error())

		msg = NewMsgDetachFrom(addr1, defaultContractID, addr2, "1")
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "1").Error())
	}

	{
		msg := NewMsgApprove(addr1, defaultContractID, addr2)
		require.Equal(t, "approve_collection", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgApprove{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.Approver, msg2.Approver)
		require.Equal(t, msg.ContractID, msg2.ContractID)
	}

	{
		msg := NewMsgDisapprove(addr1, defaultContractID, addr2)
		require.Equal(t, "disapprove_collection", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgDisapprove{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.Approver, msg2.Approver)
		require.Equal(t, msg.ContractID, msg2.ContractID)
	}

	{
		msg := NewMsgBurnFTFrom(addr1, defaultContractID, addr2, OneCoin(defaultTokenIDFT))
		require.Equal(t, "burn_ft_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgBurnFTFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.Amount, msg2.Amount)
	}

	{
		msg := NewMsgBurnFTFrom(addr1, defaultContractID, addr1, OneCoin(defaultTokenIDFT))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrapf(ErrApproverProxySame, "Approver: %s", addr1.String()).Error())

		msg = NewMsgBurnFTFrom(nil, defaultContractID, addr1, OneCoin(defaultTokenIDFT))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgBurnFTFrom(addr1, defaultContractID, nil, OneCoin(defaultTokenIDFT))
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())
	}

	{
		msg := NewMsgBurnNFTFrom(addr1, defaultContractID, addr2, defaultTokenID1)
		require.Equal(t, "burn_nft_from", msg.Type())
		require.Equal(t, "collection", msg.Route())
		require.Equal(t, sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg)), msg.GetSignBytes())
		require.Equal(t, addr1, msg.GetSigners()[0])
		require.NoError(t, msg.ValidateBasic())

		b := msg.GetSignBytes()

		msg2 := MsgBurnNFTFrom{}

		err := cdc.UnmarshalJSON(b, &msg2)
		require.NoError(t, err)

		require.Equal(t, msg.Proxy, msg2.Proxy)
		require.Equal(t, msg.From, msg2.From)
		require.Equal(t, msg.ContractID, msg2.ContractID)
		require.Equal(t, msg.TokenIDs, msg2.TokenIDs)
	}

	{
		msg := NewMsgBurnNFTFrom(addr1, defaultContractID, addr1, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrapf(ErrApproverProxySame, "Approver: %s", addr1.String()).Error())

		msg = NewMsgBurnNFTFrom(nil, defaultContractID, addr1, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Proxy cannot be empty").Error())

		msg = NewMsgBurnNFTFrom(addr1, defaultContractID, nil, defaultTokenID1)
		require.EqualError(t, msg.ValidateBasic(), sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "From cannot be empty").Error())
	}
}
