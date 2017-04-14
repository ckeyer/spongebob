package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:     "spongebob",
		Usage:    "",
		Commands: []*cli.Command{},
		Flags:    []cli.Flag{},
		Before: func(ctx *cli.Context) error {
			log.SetLevel(log.DebugLevel)
			log.SetFormatter(&log.JSONFormatter{})

			log.Debug("Before.")
			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("Action.")
			return nil
		},
	}

	app.Run(os.Args)
}
