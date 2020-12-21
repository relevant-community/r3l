package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
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

// Claim types needed for oracle

// Hash returns the hash of an Equivocation object.
func (msg *MsgScores) Hash() tmbytes.HexBytes {
	bz, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}

// func (msg *MsgScores) String() string {
// 	bz, _ := yaml.Marshal(msg)
// 	return string(bz)
// }

// GetHeight returns the height at time of the Equivocation infraction.
func (msg *MsgScores) GetHeight() int64 {
	return msg.BlockHeight
}
