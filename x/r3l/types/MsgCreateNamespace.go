package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgNamespace{}

func NewMsgNamespace(creator sdk.AccAddress, name string) *MsgNamespace {
	return &MsgNamespace{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgNamespace) Route() string {
	return RouterKey
}

func (msg *MsgNamespace) Type() string {
	return "CreateNamespace"
}

func (msg *MsgNamespace) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgNamespace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNamespace) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
