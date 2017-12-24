package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(StatusCommand())
}

func StatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "status of daemon",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
