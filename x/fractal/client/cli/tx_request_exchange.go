package cli

import (
	"strconv"

	"fractal/x/fractal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestExchange() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-exchange [finalunit] [amount] [startunit] [unitratio] [settledate]",
		Short: "Broadcast message request-exchange",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFinalunit := args[0]
			argAmount := args[1]
			argStartunit := args[2]
			argUnitratio := args[3]
			argSettledate := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestExchange(
				clientCtx.GetFromAddress().String(),
				argFinalunit,
				argAmount,
				argStartunit,
				argUnitratio,
				argSettledate,
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
