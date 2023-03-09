package cli

import (
	"strconv"

	"fractal/x/balancecheck/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUsdcBalanceCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usdc-balance-check [chain] [chainaddress] [amount]",
		Short: "Broadcast message usdc-balance-check",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argChainaddress := args[1]
			argAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUsdcBalanceCheck(
				clientCtx.GetFromAddress().String(),
				argChain,
				argChainaddress,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
