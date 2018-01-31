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
	opt := agent.AgentOption{}

	cmd := &cobra.Command{
		Use:   "join",
		Short: "j",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("join ", managerEndpoint)
			log.Fatalln(agent.Start(opt))
		},
	}

	cmd.Flags().StringVarP(&opt.Name, "agent-name", "n", "", "agent name.")
	cmd.Flags().StringVarP(&opt.HTTPAddr, "http-listen", "l", ":8092", "http listen address.")
	cmd.Flags().StringVarP(&opt.GRPCAddr, "grpc-listen", "g", ":8093", "grpc listen address.")
	cmd.Flags().StringVarP(&opt.ServerEndpoint, "manager-endpoint", "s", "localhost:8091", "gRPC server endpoint.")
	return cmd
}
