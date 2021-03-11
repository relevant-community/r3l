package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	exportedOracle "github.com/relevant-community/oracle/x/oracle/exported"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(MsgRankSource{}, "r3l/CreateRankSource", nil)
	cdc.RegisterConcrete(MsgNamespace{}, "r3l/CreateNamespace", nil)
	cdc.RegisterConcrete(MsgVote{}, "r3l/CreateVote", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRankSource{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNamespace{},
	)
	// registry.RegisterImplementations((*sdk.Msg)(nil),
	// 	&MsgScores{},
	// )
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVote{},
	)
	registry.RegisterInterface(
		"r3l.oracle.v1beta1.Claim",
		(*exportedOracle.Claim)(nil),
		&Scores{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
