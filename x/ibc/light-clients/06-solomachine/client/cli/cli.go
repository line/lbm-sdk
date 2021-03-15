package cli

import (
	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/v2/client"
	"github.com/line/lbm-sdk/v2/x/ibc/light-clients/06-solomachine/types"
)

// NewTxCmd returns a root CLI command handler for all solo machine transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "Solo Machine transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewCreateClientCmd(),
		NewUpdateClientCmd(),
		NewSubmitMisbehaviourCmd(),
	)

	return txCmd
}
