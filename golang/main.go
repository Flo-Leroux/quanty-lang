package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Flo-Leroux/quanty-lang/pkg/sdl"
	"github.com/Flo-Leroux/quanty-lang/pkg/sdl/listeners/html"
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

	// app := server.Run()

	// log.Fatal(app.Listen(":8081"))

	src := `component Main { div { p { span { strong { "{{ . }}!" }}} span } "hello world!" }`

	l := html.New("locale string")

	sdl.Run(src, l)

	tmpl := l.Lookup("Main")

	tmpl.Execute(os.Stdout, "COUCOU")
}
