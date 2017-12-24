package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(StopCommand())
}

func StopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "stop the daemon.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
