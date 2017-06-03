package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/spongebob/api"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	var (
		debug bool
	)
	app := &cli.App{
		Name:  "sponcli",
		Usage: "",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"D"},
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &debug,
			},
		},
		Commands: []*cli.Command{
			ServeCommand(),
			JoinCommand(),
		},
		Before: func(ctx *cli.Context) error {
			log.SetFormatter(&log.JSONFormatter{})
			if debug {
				log.SetLevel(log.DebugLevel)
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("main Action.")

			return nil
		},
	}

	app.Run(os.Args)
}

func ServeCommand() *cli.Command {
	var addr string

	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"address"},
				EnvVars:     []string{"ADDR", "ADDRESS"},
				Value:       ":8080",
				Destination: &addr,
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("serve ", addr)

			return api.Serve("addr")
		},
	}
}

func JoinCommand() *cli.Command {
	var endpoint string

	return &cli.Command{
		Name:    "join",
		Aliases: []string{"j"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "e",
				Aliases:     []string{"endpoint"},
				EnvVars:     []string{"ENDPOINT"},
				Value:       ":8080",
				Destination: &endpoint,
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("join ", endpoint)
			return nil
		},
	}
}
