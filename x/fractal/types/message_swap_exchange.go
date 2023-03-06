package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwapExchange = "swap_exchange"

var _ sdk.Msg = &MsgSwapExchange{}

func NewMsgSwapExchange(creator string, id uint64) *MsgSwapExchange {
	return &MsgSwapExchange{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgSwapExchange) Route() string {
	return RouterKey
}

func (msg *MsgSwapExchange) Type() string {
	return TypeMsgSwapExchange
}

func (msg *MsgSwapExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwapExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
