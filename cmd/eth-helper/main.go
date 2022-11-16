package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "eth-helper"
	app.Usage = "ethereum helper tool"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "flag",
			Usage: "flag (no spaces or hyphens, please)",
		},
	}
	app.Action = runWizard
	app.Run(os.Args)
}

func runWizard(c *cli.Context) error {
	flag := c.String("flag")
	if strings.Contains(flag, " ") || strings.Contains(flag, "-") || strings.ToLower(flag) != flag {
		fmt.Println("No spaces, hyphens or capital letters allowed in config")
	}
	makeWizard(c.String("flag")).run()
	return nil
}
