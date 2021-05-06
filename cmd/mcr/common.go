package main

import "github.com/urfave/cli/v2"

var (
	jsonFlag = &cli.BoolFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "output as JSON",
	}
)
