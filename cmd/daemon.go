package cmd

import (
	log "github.com/ckeyer/logrus"
	"github.com/ckeyer/spongebob/daemon"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DaemonCommond())
}

func DaemonCommond() *cobra.Command {
	var (
		addr string
		web  string
	)

	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "daemon process",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("daemon process", addr)
			log.Fatalln(daemon.Serve(addr))
		},
	}

	cmd.Flags().StringVarP(&addr, "listen", "l", "unix:///var/run/spongebob.sock", "host daemon listenning address.")
	cmd.Flags().StringVarP(&web, "web-addr", "w", ":8090", "web UI address.")
	return cmd
}
