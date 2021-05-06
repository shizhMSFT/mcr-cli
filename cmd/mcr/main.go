package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "mcr",
		Usage:   "Microsoft Container Registry",
		Version: "0.1.0",
		Authors: []*cli.Author{
			{
				Name:  "Shiwei Zhang",
				Email: "shizh@microsoft.com",
			},
		},
		Commands: []*cli.Command{
			treeCommand,
			tagsCommand,
			inspectCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
