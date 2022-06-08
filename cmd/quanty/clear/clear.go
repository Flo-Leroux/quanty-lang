package clear

import "github.com/urfave/cli/v2"

func Register() *cli.Command {
	return &cli.Command{
		Name:    "clear",
		Aliases: []string{"c"},
		Usage:   "clear .quanty folder",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
