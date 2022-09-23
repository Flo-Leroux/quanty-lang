package main

import (
	"fmt"
	"os/user"
)

var version = "0.0.0"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`
Hello %s!
This is the Quanty Query Language!

Version: %s`,
		user.Username, version)
	fmt.Println()
}
