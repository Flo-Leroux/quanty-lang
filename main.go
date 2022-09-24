package main

import (
	"fmt"
	"os/user"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`
Hello %s!
This is the Quanty Query Language!

Version: %s
Commit: %s
Built at %s by %s`,
		usr.Username, version, commit, date, builtBy)
	fmt.Println()
}
