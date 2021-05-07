package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/need-being/go-tree"
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

	root := tree.New(mcr.RegistryName)
	root.Virtual = true
	for _, repo := range repos {
		root.AddPathString(repo)
	}

	printer := tree.NewPrinter(os.Stdout, func(n *tree.Node) string {
		if n.Virtual {
			return color.BlueString("%v", n.Value)
		}
		return fmt.Sprint(n.Value)
	})
	return printer.Print(root)
}
