package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestExchange{}, "fractal/RequestExchange", nil)
	cdc.RegisterConcrete(&MsgApproveExchange{}, "fractal/ApproveExchange", nil)
	cdc.RegisterConcrete(&MsgSettleExchange{}, "fractal/SettleExchange", nil)
	cdc.RegisterConcrete(&MsgSwapExchange{}, "fractal/SwapExchange", nil)
	cdc.RegisterConcrete(&MsgCancelExchange{}, "fractal/CancelExchange", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestExchange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveExchange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSettleExchange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwapExchange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelExchange{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
