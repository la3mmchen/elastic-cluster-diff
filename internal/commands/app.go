package commands

import (
	"github.com/urfave/cli"
)

var (
	// PrintConfig <tbd>
	PrintConfig bool
)

// GetApp <tbd>
func GetApp(ConfigurationName string, VersionString string) *cli.App {
	app := cli.NewApp()
	app.Name = "ods-diff-elastic-cluster"
	app.Usage = "Compare indices in different elastic clusters"
	app.Version = ConfigurationName + " " + VersionString

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "print only the version",
	}

	app.Commands = []cli.Command{
		compareCluster(),
	}

	return app
}
