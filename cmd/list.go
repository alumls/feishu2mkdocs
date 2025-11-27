package main

import (
	"github.com/urfave/cli/v2"
)

func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List Feishu wiki nodes",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "flat", Usage: "Flat output"},
		},
		ArgsUsage: "[spaceId]",
		Action: func(c *cli.Context) error {
			// spaceID := c.Args().First()
			// flat := c.Bool("flat")
			// res, err := ServiceList(spaceID, flat)
			// if err != nil {
			// 	return err
			// }
			// fmt.Println(res)
			return nil
		},
	}
}
