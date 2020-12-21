package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/x/oracle/exported"
)

// var _ sdk.Msg = &MsgClaim{}

// func NewMsgClaim(creator sdk.AccAddress, blockNumber int32, hash string) *MsgClaim {
// 	return &MsgClaim{
// 		Id:          uuid.New().String(),
// 		Creator:     creator,
// 		BlockNumber: blockNumber,
// 		Hash:        hash,
// 	}
// }

var (
	_ sdk.Msg                       = &MsgCreateClaim{}
	_ types.UnpackInterfacesMessage = MsgCreateClaim{}
	_ exported.MsgCreateClaim       = &MsgCreateClaim{}
)

// NewMsgCreateClaim returns a new MsgCreateClaim with a signer/submitter.
//nolint:interfacer
func NewMsgCreateClaim(s sdk.AccAddress, claim exported.Claim) (*MsgCreateClaim, error) {
	msg, ok := claim.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("cannot proto marshal %T", claim)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	return &MsgCreateClaim{Submitter: s.String(), Claim: any}, nil
}

func (msg *MsgCreateClaim) Route() string {
	return RouterKey
}

func (msg *MsgCreateClaim) Type() string {
	return "CreateClaim"
}

func (msg *MsgCreateClaim) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{accAddr}
}

func (msg *MsgCreateClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClaim) ValidateBasic() error {
	if msg.Submitter == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	claim := msg.GetClaim()
	if claim == nil {
		return sdkerrors.Wrap(ErrInvalidClaim, "missing claim")
	}
	if err := claim.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

func (m MsgCreateClaim) GetClaim() exported.Claim {
	evi, ok := m.Claim.GetCachedValue().(exported.Claim)
	if !ok {
		return nil
	}
	return evi
}

func (m MsgCreateClaim) GetSubmitter() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Submitter)
	if err != nil {
		return nil
	}
	return accAddr
}

func (m MsgCreateClaim) UnpackInterfaces(ctx types.AnyUnpacker) error {
	var claim exported.Claim
	return ctx.UnpackAny(m.Claim, &claim)
}
