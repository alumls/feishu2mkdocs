package main

import (
	"github.com/urfave/cli/v2"
)

func NewCleanCommand() *cli.Command {
	return &cli.Command{
		Name:  "clean",
		Usage: "Clean generated files",
		Action: func(c *cli.Context) error {
			//return ServiceClean()
			return nil
		},
	}
}
