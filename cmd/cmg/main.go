package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/modecode/cmg/internal/commit"
)

func main() {
	app := &cli.App{
		Name:  "conventional-commit",
		Usage: "Generate conventional commits with ease",
		Action: func(c *cli.Context) error {
			return commit.RunCLI()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red("Error: %v", err)
	}
}