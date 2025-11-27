package main

import (
	"github.com/urfave/cli/v2"
)

func NewConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Read configuration file (default: feishu2mkdocs.yaml)",
		Subcommands: []*cli.Command{
			// config init <path>
			{
				Name:      "init",
				Usage:     "Initialize a new config file",
				ArgsUsage: "[path]",
				Action: func(c *cli.Context) error {
					// path := c.Args().First()
					// return ConfigInit(path)
					return nil
				},
			},
			{
				Name:  "path",
				Usage: "Show the path of the active configuration file",
				Action: func(c *cli.Context) error {
					// p, err := ConfigPath()
					// if err != nil {
					// 	return err
					// }
					// fmt.Println(p)
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "Print the current configuration",
				Action: func(c *cli.Context) error {
					// s, err := ConfigShow()
					// if err != nil {
					// 	return err
					// }
					// fmt.Println(s)
					return nil
				},
			},
			{
				Name:      "set",
				Usage:     "Set a configuration field",
				ArgsUsage: "<key> <value>",
				Action: func(c *cli.Context) error {
					// if c.Args().Len() < 2 {
					// 	return fmt.Errorf("missing key or value")
					// }
					// key := c.Args().Get(0)
					// value := c.Args().Get(1)
					// return ConfigSet(key, value)
					return nil
				},
			},
			{
				Name:      "get",
				Usage:     "Get a configuration field",
				ArgsUsage: "<key>",
				Action: func(c *cli.Context) error {
					// key := c.Args().First()
					// v, err := ConfigGet(key)
					// if err != nil {
					// 	return err
					// }
					// fmt.Println(v)
					return nil
				},
			},
		},
	}
}
