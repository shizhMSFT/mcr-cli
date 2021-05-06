package main

import (
	"fmt"
	"strings"

	"github.com/disiqueira/gotree"
	"github.com/shizhMSFT/mcr-go"
	"github.com/urfave/cli/v2"
)

var treeCommand = &cli.Command{
	Name:   "tree",
	Usage:  "List repositories as a tree",
	Action: runTree,
}

func runTree(ctx *cli.Context) error {
	cli := mcr.NewClient(nil)
	repos, err := cli.Repositories(ctx.Context)
	if err != nil {
		return err
	}

	tree := gotree.New(mcr.RegistryName)
	for _, repo := range repos {
		current := tree
		parts := strings.Split(repo, "/")
		for _, part := range parts[:len(parts)-1] {
			if item, ok := findItem(current, part); ok {
				current = item
			} else {
				current = current.Add(part)
			}
		}
		current.Add(parts[len(parts)-1])
	}

	fmt.Println(tree.Print())
	return nil
}

func findItem(tree gotree.Tree, target string) (gotree.Tree, bool) {
	for _, item := range tree.Items() {
		if item.Text() == target {
			return item, true
		}
	}
	return nil, false
}
