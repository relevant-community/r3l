package types_test

import (
	fmt "fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/relevant-community/r3l/x/oracle/exported"
	"github.com/relevant-community/r3l/x/oracle/types"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	gs := types.DefaultGenesis()
	require.NotNil(t, gs.Claims)
	require.Len(t, gs.Claims, 0)

	require.NotNil(t, gs.RoundVotes)
	require.Len(t, gs.RoundVotes, 0)

	require.NotNil(t, gs.Params)
	require.Equal(t, gs.Params, types.DefaultParams())
}

func TestNewGenesisState(t *testing.T) {
	var (
		claims     []exported.Claim
		roundVotes []types.RoundVotes
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"can proto marshal",
			func() {
				claims = []exported.Claim{&types.TestClaim{}}
				roundVotes = []types.RoundVotes{}
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {
			tc.malleate()

			if tc.expPass {
				require.NotPanics(t, func() {
					types.NewGenesisState(types.DefaultParams(), roundVotes, claims)
				})
			} else {
				require.Panics(t, func() {
					types.NewGenesisState(types.DefaultParams(), roundVotes, claims)
				})
			}
		})
	}
}

func TestGenesisStateValidate(t *testing.T) {
	var (
		genesisState *types.GenesisState
		testClaim    []exported.Claim
	)
	roundVotes := []types.RoundVotes{}
	params := types.DefaultParams()

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"valid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i + 1),
						Content:     "test",
						ClaimType:   "test",
					}
				}
				genesisState = types.NewGenesisState(params, roundVotes, testClaim)
			},
			true,
		},
		{
			"invalid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i),
						Content:     "test",
						ClaimType:   "test",
					}
				}
				genesisState = types.NewGenesisState(params, roundVotes, testClaim)
			},
			false,
		},
		{
			"expected evidence",
			func() {
				genesisState = &types.GenesisState{
					Claims: []*codectypes.Any{{}},
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {
			tc.malleate()

			if tc.expPass {
				require.NoError(t, genesisState.Validate())
			} else {
				require.Error(t, genesisState.Validate())
			}
		})
	}
}

func TestUnpackInterfaces(t *testing.T) {
	var gs = types.GenesisState{
		Claims: []*codectypes.Any{{}},
	}

	testCases := []struct {
		msg      string
		unpacker codectypes.AnyUnpacker
		expPass  bool
	}{
		{
			"success",
			codectypes.NewInterfaceRegistry(),
			true,
		},
		{
			"error",
			codec.NewLegacyAmino(),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {

			if tc.expPass {
				require.NoError(t, gs.UnpackInterfaces(tc.unpacker))
			} else {
				require.Error(t, gs.UnpackInterfaces(tc.unpacker))
			}
		})
	}
}
