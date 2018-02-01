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
		opt = daemon.DaemonOption{}
	)

	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "daemon process for mamager all client's job.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("daemon process %+v", opt)
			log.Fatalln(daemon.Serve(opt))
		},
	}
	cmd.Flags().StringVarP(&opt.GRPCAddr, "grpc-addr", "l", "unix:///var/run/spongebob.sock", "host daemon listenning address for manager all agents.")
	cmd.Flags().StringVarP(&opt.HTTPAddr, "http-addr", "w", ":8090", "web UI address.")
	cmd.Flags().StringVarP(&opt.PromBin, "prom-bin", "p", "", "Prometheus binary path.")
	return cmd
}
