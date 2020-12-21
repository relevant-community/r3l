// Package exported package is based on evidence module
package exported

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Claim defines the methods that all oracle claims must implement
type Claim interface {
	Route() string
	Type() string
	String() string
	Hash() tmbytes.HexBytes
	ValidateBasic() error

	// Height of inputs
	GetHeight() int64

	// GetConsensusAddress() sdk.ConsAddress
	// GetValidatorPower() int64
	// GetTotalPower() int64
}

// MsgCreateClaim defines the specific interface a concrete message must
// implement in order to process an oracle claim. The concrete MsgSubmitClaim
// must be defined at the application-level.
type MsgCreateClaim interface {
	sdk.Msg

	GetClaim() Claim
	GetSubmitter() sdk.AccAddress
}
