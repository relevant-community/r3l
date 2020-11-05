package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgScore{}

func NewMsgScore(creator sdk.AccAddress, blockHeight int64, pRank int64, nRank int64) *MsgScore {
	return &MsgScore{
		Id:          uuid.New().String(),
		BlockHeight: blockHeight,
		PRank:       pRank,
		NRank:       nRank,
		Creator:     creator,
	}
}

func (msg *MsgScore) Route() string {
	return RouterKey
}

func (msg *MsgScore) Type() string {
	return "SetScores"
}

func (msg *MsgScore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgScore) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgScore) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
