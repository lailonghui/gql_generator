package cmd

import "github.com/urfave/cli"

var genCmd = cli.Command{
	Name:  "generate",
	Usage: "generate contents",
	Action: func(ctx *cli.Context) error {
		if err := generate(); err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	},
}

func generate() error {

	return nil
}
