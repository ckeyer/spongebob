package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(ReloadCommand())
}

func ReloadCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "reload",
		Short: "reload daemon.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
