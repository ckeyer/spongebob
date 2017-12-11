package cmd

import (
	log "github.com/ckeyer/logrus"
	"github.com/ckeyer/spongebob/agent"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(JoinCommand())
}

func JoinCommand() *cobra.Command {
	var endpoint string

	cmd := &cobra.Command{
		Use:   "join",
		Short: "j",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("join ", endpoint)
			log.Fatalln(agent.Start(endpoint))
		},
	}

	cmd.Flags().StringVarP(&endpoint, "endpoint", "s", "localhost:8091", "gRPC server endpoint.")
	return cmd
}
