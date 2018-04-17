// this is the CLI interface to fetiched

package main

import (
	"fmt"
	"github.com/keltia/fetiche"
	"github.com/urfave/cli"
	"os"
	"sort"
)

const (
	cliVersion = "0.0.0"
)

var (
	// MyName is both package & CLI command name
	MyName = "fetiche"

	// All possible commands
	cliCommands []cli.Command

	// Global flags
	fDebug   bool
	fVerbose bool
)

func main() {
	cli.VersionFlag = cli.BoolFlag{Name: "version, V"}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("CLI: %s Fetiche: %s\n", c.App.Version, fetiche.GetVersion())
	}

	app := cli.NewApp()
	app.Name = "fetiche"
	app.Usage = "Fetiche CLI interface"
	app.Author = "Ollivier Robert <roberto@keltia.net>"
	app.Version = cliVersion
	//app.HideVersion = true

	// General flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,D",
			Usage:       "debug mode",
			Destination: &fDebug,
		},
		cli.BoolFlag{
			Name:        "verbose,v",
			Usage:       "verbose mode",
			Destination: &fVerbose,
		},
	}

	//app.Before = finalcheck

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
