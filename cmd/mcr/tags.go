package main

import (
	"errors"
	"fmt"

	"github.com/shizhMSFT/mcr-go"
	"github.com/urfave/cli/v2"
)

var tagsCommand = &cli.Command{
	Name:      "tags",
	Usage:     "List tags",
	ArgsUsage: "<repository_name>",
	Flags: []cli.Flag{
		jsonFlag,
	},
	Action: runTags,
}

func runTags(ctx *cli.Context) error {
	repo := ctx.Args().First()
	if repo == "" {
		return errors.New("missing repository name")
	}

	cli := mcr.NewClient(nil)
	tags, err := cli.Tags(ctx.Context, repo)
	if err != nil {
		return err
	}

	if ctx.Bool(jsonFlag.Name) {
		return printJSON(tags)
	}
	for _, tag := range tags {
		fmt.Println(tag)
	}
	return nil
}
