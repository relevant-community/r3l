package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgScore{}

func NewMsgScores(creator sdk.AccAddress, blockHeight int64, scores []Score) *MsgScores {
	return &MsgScores{
		Creator:     creator,
		Scores:      scores,
		BlockHeight: blockHeight,
	}
}

func (msg *MsgScores) Route() string {
	return RouterKey
}

func (msg *MsgScores) Type() string {
	return "SetScores"
}

func (msg *MsgScores) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgScores) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgScores) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
