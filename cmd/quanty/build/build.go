package build

import (
	"log"
	"os"
	"quanty/language"
	"time"

	"github.com/urfave/cli/v2"
)

var rootDir string

func init() {
	var err error
	rootDir, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
}

func Register() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build application",
		Action: func(ctx *cli.Context) error {

			startAt := time.Now()

			language.NewTree(rootDir)

			endAt := time.Since(startAt)

			log.Printf("Build End in %s\n", endAt)

			return nil
		},
	}
}
