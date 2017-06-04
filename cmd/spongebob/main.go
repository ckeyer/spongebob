package main

import (
	"github.com/ckeyer/spongebob/agent"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/spongebob/server"
	"github.com/ckeyer/spongebob/server/api"
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
	var (
		addr string
		web  string
	)

	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"address"},
				EnvVars:     []string{"ADDR", "ADDRESS"},
				Value:       ":8091",
				Destination: &addr,
			},
			&cli.StringFlag{
				Name:        "web",
				Aliases:     []string{"web_address"},
				EnvVars:     []string{"WEB_ADDR"},
				Value:       ":8090",
				Destination: &web,
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Infof("server listenning at: %s", addr)
			log.Infof("web ui listenning at: %s", web)

			if err := server.Start(addr); err != nil {
				return err
			}

			return api.Serve(web)
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
				Value:       "localhost:8091",
				Destination: &endpoint,
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("join ", endpoint)
			return agent.Start(endpoint)
		},
	}
}
