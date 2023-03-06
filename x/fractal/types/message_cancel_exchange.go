package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelExchange = "cancel_exchange"

var _ sdk.Msg = &MsgCancelExchange{}

func NewMsgCancelExchange(creator string, id uint64) *MsgCancelExchange {
	return &MsgCancelExchange{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelExchange) Route() string {
	return RouterKey
}

func (msg *MsgCancelExchange) Type() string {
	return TypeMsgCancelExchange
}

func (msg *MsgCancelExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
