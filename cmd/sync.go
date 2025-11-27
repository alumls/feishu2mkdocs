package main

import (
	"github.com/urfave/cli/v2"
)

func NewSyncCommand() *cli.Command {
	return &cli.Command{
		Name:  "sync",
		Usage: "Sync Feishu wiki to local Markdown",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "force",
				Usage: "Force full sync",
			},
		},
		Action: func(c *cli.Context) error {
			// force := c.Bool("force")
			//return ServiceSync(force)
			return nil
		},
	}
}
