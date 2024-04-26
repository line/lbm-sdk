package types

import sdk "github.com/Finschia/finschia-sdk/types"

var (
	_ sdk.Msg = &MsgTransfer{}
	_ sdk.Msg = &MsgProvision{}
	_ sdk.Msg = &MsgHoldTransfer{}
	_ sdk.Msg = &MsgReleaseTransfer{}
	_ sdk.Msg = &MsgRemoveProvision{}
	_ sdk.Msg = &MsgClaimBatch{}
	_ sdk.Msg = &MsgClaim{}
	_ sdk.Msg = &MsgUpdateRole{}
	_ sdk.Msg = &MsgHalt{}
	_ sdk.Msg = &MsgResume{}
)

func (m MsgTransfer) ValidateBasic() error { return nil }

func (m MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgProvision) ValidateBasic() error { return nil }

func (m MsgProvision) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgProvision) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgHoldTransfer) ValidateBasic() error { return nil }

func (m MsgHoldTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgHoldTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgReleaseTransfer) ValidateBasic() error { return nil }

func (m MsgReleaseTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgReleaseTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgRemoveProvision) ValidateBasic() error { return nil }

func (m MsgRemoveProvision) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgRemoveProvision) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgClaimBatch) ValidateBasic() error { return nil }

func (m MsgClaimBatch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgClaimBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgClaim) ValidateBasic() error { return nil }

func (m MsgClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateRole) ValidateBasic() error { return nil }

func (m MsgUpdateRole) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgUpdateRole) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgHalt) ValidateBasic() error { return nil }

func (m MsgHalt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Guardian)}
}

func (m MsgHalt) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgResume) ValidateBasic() error { return nil }

func (m MsgResume) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.From)}
}

func (m MsgResume) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
