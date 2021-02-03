package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/x/oracle/exported"
)

var (
	_ sdk.Msg                       = &MsgCreateClaim{}
	_ types.UnpackInterfacesMessage = MsgCreateClaim{}
	_ exported.MsgCreateClaimI      = &MsgCreateClaim{}
)

// NewMsgCreateClaim returns a new MsgCreateClaim with a signer/submitter.
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

// Route get msg route
func (msg *MsgCreateClaim) Route() string {
	return RouterKey
}

// Type get msg type
func (msg *MsgCreateClaim) Type() string {
	return "CreateClaim"
}

// GetSigners get msg get signers
func (msg *MsgCreateClaim) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{accAddr}
}

// GetSignBytes get msg get getsingbytes
func (msg *MsgCreateClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validation
func (msg *MsgCreateClaim) ValidateBasic() error {
	if msg.Submitter == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	claim := msg.GetClaim()
	if claim == nil {
		return sdkerrors.Wrap(ErrInvalidClaim, "missing claim")
	}

	return nil
}

// GetClaim get the claim
func (msg MsgCreateClaim) GetClaim() exported.Claim {
	claim, ok := msg.Claim.GetCachedValue().(exported.Claim)
	if !ok {
		return nil
	}
	return claim
}

// GetSubmitter get the submitter
func (msg MsgCreateClaim) GetSubmitter() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return nil
	}
	return accAddr
}

// UnpackInterfaces unpack
func (msg MsgCreateClaim) UnpackInterfaces(ctx types.AnyUnpacker) error {
	var claim exported.Claim
	return ctx.UnpackAny(msg.Claim, &claim)
}
