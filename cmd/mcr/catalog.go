package main

import (
	"fmt"

	"github.com/shizhMSFT/mcr-go"
	"github.com/urfave/cli/v2"
)

var catalogCommand = &cli.Command{
	Name:   "catalog",
	Usage:  "List repositories",
	Action: runCatalog,
	Flags: []cli.Flag{
		jsonFlag,
	},
}

func runCatalog(ctx *cli.Context) error {
	cli := mcr.NewClient(nil)
	repos, err := cli.Repositories(ctx.Context)
	if err != nil {
		return err
	}

	if ctx.Bool(jsonFlag.Name) {
		return printJSON(repos)
	}
	for _, repo := range repos {
		fmt.Println(repo)
	}
	return nil
}
