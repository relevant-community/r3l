package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingexported "github.com/cosmos/cosmos-sdk/x/staking/exported"
)

type (
	// StakingKeeper defines the staking module interface contract needed by the
	// evidence module.
	StakingKeeper interface {
		// ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingexported.ValidatorI
		Validator(ctx sdk.Context, address sdk.ValAddress) stakingexported.ValidatorI // get validator by operator address; nil when validator not found
	}
)
