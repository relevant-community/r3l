package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyClaimParams   = []byte("claimParams")
	KeyVoteThreshold = []byte("voteThreshold")
)

// Default params for testing
var (
	TestClaimType               = "test"
	TestPrevoteClaimType        = "prevoteTest"
	TestVotePeriod       uint64 = 3
)

// Default parameter values
var (
	DefaultVoteThreshold = sdk.NewDecWithPrec(50, 2) // 50%
	DefaultClaimParams   = map[string](ClaimParams){
		TestClaimType: {
			ClaimType: TestClaimType,
		},
		TestPrevoteClaimType: {
			ClaimType:  TestPrevoteClaimType,
			Prevote:    true,
			VotePeriod: TestVotePeriod,
		},
	}
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable for oracle module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {
	return Params{
		VoteThreshold: DefaultVoteThreshold,
		ClaimParams:   DefaultClaimParams,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyClaimParams, &p.ClaimParams, validateClaimParams),
		paramtypes.NewParamSetPair(KeyVoteThreshold, &p.VoteThreshold, validateVoteThreshold),
	}
}

// ValidateBasic performs basic validation on oracle parameters.
func (p Params) ValidateBasic() error {
	if p.VoteThreshold.LTE(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("oracle parameter VoteTheshold must be greater than 33 percent")
	}
	return nil
}

func validateClaimParams(i interface{}) error {
	claimParams, ok := i.(map[string](ClaimParams))
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, param := range claimParams {
		if param.Prevote && param.VotePeriod == 0 {
			return fmt.Errorf("vote period must be positive: %d", param.VotePeriod)
		}
	}

	return nil
}

func validateVoteThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("vote threshold must be bigger than 33%%: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("vote threshold too large: %s", v)
	}

	return nil
}
