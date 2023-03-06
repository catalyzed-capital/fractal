package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSettleExchange = "settle_exchange"

var _ sdk.Msg = &MsgSettleExchange{}

func NewMsgSettleExchange(creator string, id uint64) *MsgSettleExchange {
	return &MsgSettleExchange{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgSettleExchange) Route() string {
	return RouterKey
}

func (msg *MsgSettleExchange) Type() string {
	return TypeMsgSettleExchange
}

func (msg *MsgSettleExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSettleExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSettleExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
