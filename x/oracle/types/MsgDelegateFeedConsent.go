package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelegateFeedConsent{}

// Message types for the oracle module
const (
	TypeMsgDelegateFeedConsent = "delegate_feed_consent"
)

// NewMsgDelegateFeedConsent returns a new MsgDelegateFeedConsent
func NewMsgDelegateFeedConsent(val, del sdk.AccAddress) *MsgDelegateFeedConsent {
	return &MsgDelegateFeedConsent{
		Validator: val.String(),
		Delegate:  del.String(),
	}
}

// Route implements sdk.Msg
func (m *MsgDelegateFeedConsent) Route() string { return ModuleName }

// Type implements sdk.Msg
func (m *MsgDelegateFeedConsent) Type() string { return TypeMsgDelegateFeedConsent }

// ValidateBasic implements sdk.Msg
func (m *MsgDelegateFeedConsent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Validator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Delegate); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	return nil
}

// GetSignBytes implements sdk.Msg
func (m *MsgDelegateFeedConsent) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners implements sdk.Msg
func (m *MsgDelegateFeedConsent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.MustGetValidator()}
}

// MustGetValidator returns the sdk.AccAddress for the validator
func (m *MsgDelegateFeedConsent) MustGetValidator() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(m.Validator)
	if err != nil {
		panic(err)
	}
	return val
}

// MustGetDelegate returns the sdk.AccAddress for the delegate
func (m *MsgDelegateFeedConsent) MustGetDelegate() sdk.AccAddress {
	val, err := sdk.AccAddressFromBech32(m.Delegate)
	if err != nil {
		panic(err)
	}
	return val
}
