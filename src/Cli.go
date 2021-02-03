package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// CliArguments test
var CliArguments struct {
	EBNF  bool     `help:"Dump EBNF."`
	Watch bool     `watch:"Watch Files."`
	Files []string `arg:"" optional:"" type:"existingfile" help:"Quantum files to parse."`
}

// Cli test
type Cli struct {
	Context *kong.Context
	Files   []string
}

// NewCli test
func NewCli() *Cli {
	ctx := kong.Parse(&CliArguments)

	cli := &Cli{
		Context: ctx,
	}

	cli.checkAndAssignFilesArgument()

	return cli
}

func (cli *Cli) checkAndAssignFilesArgument() {
	if len(CliArguments.Files) == 0 {
		fmt.Println("No file set")
		cli.Context.Exit(0)
	}
	cli.Files = CliArguments.Files
}
