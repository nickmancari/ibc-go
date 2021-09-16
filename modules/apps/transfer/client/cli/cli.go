package cli

import (
	// external library imports
	"github.com/spf13/cobra"

	// ibc-go library imports
	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the query commands for IBC connections
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "ibc-transfer",
		Short:                      "IBC fungible token transfer query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdQueryDenomTrace(),
		GetCmdQueryDenomTraces(),
		GetCmdParams(),
		GetCmdQueryEscrowAddress(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for IBC fungible token transfer
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "ibc-transfer",
		Short:                      "IBC fungible token transfer transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewTransferTxCmd(),
	)

	return txCmd
}
