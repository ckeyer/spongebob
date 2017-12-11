package cmd

import (
	log "github.com/ckeyer/logrus"
	"github.com/ckeyer/spongebob/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ServeCommand())
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
