package cmd

import (
	log "github.com/ckeyer/logrus"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	rootCmd = &cobra.Command{
		Use:   "sponcli",
		Short: "[spongebob] A server pressure measurement tool",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				log.SetFormatter(&log.JSONFormatter{})
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "log debug message.")

	rootCmd.Execute()
}
