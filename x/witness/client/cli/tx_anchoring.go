package cli

import (
	"strconv"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"middleware/x/witness/types"
)

var _ = strconv.Itoa(0)

func CmdAnchoring() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anchoring [blockhash] [height]",
		Short: "Broadcast message anchoring",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlockhash := args[0]
			argHeight := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAnchoring(
				clientCtx.GetFromAddress().String(),
				argBlockhash,
				argHeight,
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
