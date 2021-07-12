package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var cliArguments struct {
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
	ctx := kong.Parse(&cliArguments)

	cli := &Cli{
		Context: ctx,
	}

	// cliArguments.Files = append(cliArguments.Files, "./example.qy")
	cli.checkAndAssignFilesArgument()

	return cli
}

func (cli *Cli) checkAndAssignFilesArgument() {
	if len(cliArguments.Files) == 0 {
		fmt.Println("No file set")
		cli.Context.Exit(0)
	}
	cli.Files = cliArguments.Files
}
