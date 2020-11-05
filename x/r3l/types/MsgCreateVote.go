package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgVote{}

func NewMsgVote(creator sdk.AccAddress, to string, amount int32) *MsgVote {
  return &MsgVote{
    Id: uuid.New().String(),
		Creator: creator,
    To: to,
    Amount: amount,
	}
}

func (msg *MsgVote) Route() string {
  return RouterKey
}

func (msg *MsgVote) Type() string {
  return "CreateVote"
}

func (msg *MsgVote) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgVote) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgVote) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
