package main

import (
	log "github.com/ckeyer/logrus"
	"github.com/ckeyer/spongebob/agent"
	"github.com/ckeyer/spongebob/server"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	rootCmd = &cobra.Command{
		Use:   "sponcli",
		Short: "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				log.SetFormatter(&log.JSONFormatter{})
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("...")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "log debug message.")

	// rootCmd.Flags().StringVarP(p, name, shorthand, value, usage)
}

// Main
func main() {
	rootCmd.Execute()
}

// children commands.
func init() {
	rootCmd.AddCommand(ServeCommand())
	rootCmd.AddCommand(JoinCommand())
}

func ServeCommand() *cobra.Command {
	var (
		addr string
		web  string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "s",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("server listenning at %s", addr)
			log.Infof("web ui listenning at %s", web)

			if err := server.Run(addr, web); err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().StringVarP(&addr, "rpc-addr", "s", ":8091", "gRPC listenning address.")
	cmd.Flags().StringVarP(&web, "web-addr", "w", ":8090", "web UI address.")
	return cmd
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
