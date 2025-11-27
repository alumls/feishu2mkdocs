package main

import (
	"log"
	"strings"
	"os"
	"github.com/urfave/cli/v2"
)

var version = "v0"

func main() {
	app := &cli.App{
		Name:    "feishu2mkdocs",
		Version: strings.TrimSpace(string(version)),
		Usage:   "Convert feishu/larksuite wiki to MkDocs-compatible markdown files",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{
			NewConfigCommand(),
			NewSyncCommand(),
			NewExportCommand(),
			NewListCommand(),
			NewCleanCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
