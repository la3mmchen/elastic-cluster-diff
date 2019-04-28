package main

import (
	"os"

	"github.com/la3mmchen/elastic-cluster-diff/internal/commands"
)

var (
	ConfigurationName string
	Env               string
	VersionString     = "0.0.1"
)

func main() {
	app := commands.GetApp(ConfigurationName, VersionString)

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
