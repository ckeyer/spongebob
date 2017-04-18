package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/spongebob/api"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	var (
		addr  string
		debug bool
	)
	app := &cli.App{
		Name:     "spongebob",
		Usage:    "",
		Commands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"D"},
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &debug,
			},
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"address"},
				EnvVars:     []string{"ADDR"},
				Value:       ":8080",
				Destination: &addr,
			},
		},
		Before: func(ctx *cli.Context) error {
			log.SetFormatter(&log.JSONFormatter{})
			if debug {
				log.SetLevel(log.DebugLevel)
			}
			log.Debug("Before.")
			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("Action.")
			api.Serve(addr)
			return nil
		},
	}

	app.Run(os.Args)
}
