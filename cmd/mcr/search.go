package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/shizhMSFT/mcr-go"
	"github.com/urfave/cli/v2"
)

var searchCommand = &cli.Command{
	Name:  "search",
	Usage: "Search repositories",
	Flags: []cli.Flag{
		jsonFlag,
	},
	Action: runSearch,
}

func runSearch(ctx *cli.Context) error {
	repo := ctx.Args().First()
	if repo == "" {
		return errors.New("missing repository pattern")
	}
	pattern, err := regexp.Compile(repo)
	if err != nil {
		return err
	}

	cli := mcr.NewClient(nil)
	repos, err := cli.Repositories(ctx.Context)
	if err != nil {
		return err
	}

	if ctx.Bool(jsonFlag.Name) {
		var result []string
		for _, repo := range repos {
			if pattern.MatchString(repo) {
				result = append(result, repo)
			}
		}
		return printJSON(result)
	}
	for _, repo := range repos {
		if pattern.MatchString(repo) {
			fmt.Println(repo)
		}
	}
	return nil
}
