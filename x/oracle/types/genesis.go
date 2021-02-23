package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/relevant-community/r3l/x/oracle/exported"
)

var _ types.UnpackInterfacesMessage = GenesisState{}

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params,
	rounds []Round,
	_claims []exported.Claim,
	pending map[string]([]uint64),
	delegations []MsgDelegateFeedConsent,
) *GenesisState {

	claims := make([]*types.Any, len(_claims))
	for i, claim := range _claims {
		msg, ok := claim.(proto.Message)
		if !ok {
			panic(fmt.Errorf("cannot proto marshal %T", claim))
		}
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}
		claims[i] = any
	}

	genPending := map[string]GenesisState_ListOfUint{}
	for i, p := range pending {
		genPending[i] = GenesisState_ListOfUint{p}
	}

	return &GenesisState{
		Params:            params,
		Rounds:            rounds,
		Claims:            claims,
		Pending:           genPending,
		FeederDelegations: delegations,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		Claims:  []*types.Any{},
		Rounds:  []Round{},
		Pending: map[string]GenesisState_ListOfUint{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	for i, feederDelegation := range gs.FeederDelegations {
		if _, err := sdk.AccAddressFromBech32(feederDelegation.Delegate); err != nil {
			return fmt.Errorf("invalid feeder at index %d: %w", i, err)
		}
		if _, err := sdk.AccAddressFromBech32(feederDelegation.Validator); err != nil {
			return fmt.Errorf("invalid feeder at index %d: %w", i, err)
		}
	}

	for _, c := range gs.Claims {
		claim, ok := c.GetCachedValue().(exported.Claim)
		if !ok {
			return fmt.Errorf("expected claim")
		}
		if err := claim.ValidateBasic(); err != nil {
			return err
		}
	}
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, any := range gs.Claims {
		var claim exported.Claim
		err := unpacker.UnpackAny(any, &claim)
		if err != nil {
			return err
		}
	}
	return nil
}
