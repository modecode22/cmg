package main

import (
	"os"
	"github.com/fatih/color"
	"github.com/modecode22/cmg/internal/commit"
	"github.com/urfave/cli/v2"
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