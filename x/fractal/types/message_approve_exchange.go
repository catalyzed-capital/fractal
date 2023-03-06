package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveExchange = "approve_exchange"

var _ sdk.Msg = &MsgApproveExchange{}

func NewMsgApproveExchange(creator string, id uint64) *MsgApproveExchange {
	return &MsgApproveExchange{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveExchange) Route() string {
	return RouterKey
}

func (msg *MsgApproveExchange) Type() string {
	return TypeMsgApproveExchange
}

func (msg *MsgApproveExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
