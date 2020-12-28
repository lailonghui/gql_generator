package cmd

import (
	"github.com/urfave/cli"
	"os"
)

// Execute ...
func Execute() {
	app := cli.NewApp()
	app.Name = "gqlgen-generator"
	app.Usage = "This tool is for generating GraphQL API using gqlgen and gorm"

	app.Action = genCmd.Action
	app.Usage = genCmd.Usage
	app.Flags = genCmd.Flags

	app.Commands = []cli.Command{
		genCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
