package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRankSource{}

func NewMsgRankSource(creator sdk.AccAddress, account string) *MsgRankSource {
	return &MsgRankSource{
		Creator: creator,
		Account: account,
	}
}

func (msg *MsgRankSource) Route() string {
	return RouterKey
}

func (msg *MsgRankSource) Type() string {
	return "CreateRankSource"
}

func (msg *MsgRankSource) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgRankSource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRankSource) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
