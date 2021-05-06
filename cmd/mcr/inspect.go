package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/shizhMSFT/mcr-go"
	"github.com/urfave/cli/v2"
)

var inspectCommand = &cli.Command{
	Name:      "inspect",
	Usage:     "Inspect manifest",
	ArgsUsage: "<repository_name>:<reference>",
	Action:    runInspect,
}

func runInspect(ctx *cli.Context) error {
	name := ctx.Args().First()
	if name == "" {
		return errors.New("missing repository name")
	}

	var repo, ref string
	if parts := strings.SplitN(name, "@", 2); len(parts) == 2 {
		repo = parts[0]
		ref = parts[1]
	} else {
		parts = strings.SplitN(name, ":", 2)
		repo = parts[0]
		ref = "latest"
		if len(parts) == 2 && parts[1] != "" {
			ref = parts[1]
		}
	}

	cli := mcr.NewClient(nil)
	_, manifest, err := cli.Manifest(ctx.Context, repo, ref)
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, bytes.NewReader(manifest))
	return err
}
