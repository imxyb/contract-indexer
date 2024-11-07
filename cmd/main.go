package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/imxyb/contract-indexer/tasker"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	app := &cli.App{
		Name: "contract indexer cli",
		Commands: []*cli.Command{
			{
				Name:  "start-task",
				Usage: "start user api server",
				Action: func(c *cli.Context) error {
					return tasker.Start(c.Context, c.String("conf"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "conf",
						Usage:    "config file path",
						Required: true,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
