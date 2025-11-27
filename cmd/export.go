package main

import (
	"github.com/urfave/cli/v2"
)

func NewExportCommand() *cli.Command {
	return &cli.Command{
		Name:      "export",
		Usage:     "Export a single wiki page by node token",
		ArgsUsage: "<node_token>",
		Action: func(c *cli.Context) error {
			// token := c.Args().First()
			// return ServiceExportNode(token)
			return nil
		},
	}
}
