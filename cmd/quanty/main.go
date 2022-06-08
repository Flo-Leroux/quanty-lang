package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"quanty/cmd/quanty/build"
	"quanty/cmd/quanty/clear"
	initializer "quanty/cmd/quanty/init"
)

func main() {
	app := &cli.App{
		Name:  "quanty",
		Usage: "Quanty App!",

		EnableBashCompletion: true,
		Commands: []*cli.Command{
			initializer.Register(),
			build.Register(),
			clear.Register(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
