package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestExchange = "request_exchange"

var _ sdk.Msg = &MsgRequestExchange{}

func NewMsgRequestExchange(creator string, finalunit string, amount string, startunit string, unitratio string, settledate string) *MsgRequestExchange {
	return &MsgRequestExchange{
		Creator:    creator,
		Finalunit:  finalunit,
		Amount:     amount,
		Startunit:  startunit,
		Unitratio:  unitratio,
		Settledate: settledate,
	}
}

func (msg *MsgRequestExchange) Route() string {
	return RouterKey
}

func (msg *MsgRequestExchange) Type() string {
	return TypeMsgRequestExchange
}

func (msg *MsgRequestExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	settledate, err := strconv.ParseInt(msg.Settledate, 10, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deadline is not an integer")
	}
	if settledate <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be a positive integer")
	}
	return nil
}
