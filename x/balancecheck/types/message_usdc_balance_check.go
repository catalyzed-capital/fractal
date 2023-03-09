package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUsdcBalanceCheck = "usdc_balance_check"

var _ sdk.Msg = &MsgUsdcBalanceCheck{}

func NewMsgUsdcBalanceCheck(creator string, chain string, chainaddress string, amount string) *MsgUsdcBalanceCheck {
	return &MsgUsdcBalanceCheck{
		Creator:      creator,
		Chain:        chain,
		Chainaddress: chainaddress,
		Amount:       amount,
	}
}

func (msg *MsgUsdcBalanceCheck) Route() string {
	return RouterKey
}

func (msg *MsgUsdcBalanceCheck) Type() string {
	return TypeMsgUsdcBalanceCheck
}

func (msg *MsgUsdcBalanceCheck) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUsdcBalanceCheck) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUsdcBalanceCheck) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
