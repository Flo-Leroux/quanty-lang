package main

import (
	"fmt"
	"os"
	"quanty/ast"

	"github.com/alecthomas/kong"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
	"github.com/alecthomas/repr"
)

var (
	// graphQLLexer = lexer.Must(stateful.NewSimple([]stateful.Rule{
	// 	{"Comment", `(?:#|//)[^\n]*\n?`, nil},
	// 	{"Ident", `[A-Z][a-zA-Z]\w*`, nil},
	// 	{"Number", `(?:\d*\.)?\d+`, nil},
	// 	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	// 	{"Whitespace", `[ \t\n\r]+`, nil},
	// }))
	lexer = stateful.MustSimple([]stateful.Rule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`, nil},
		{"Comment", `[#;][^\n]*`, nil},
		{"Whitespace", `[ \r\t\n]+`, nil},
	})
	parser = participle.MustBuild(&ast.File{},
		participle.Lexer(lexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)

var cli struct {
	EBNF  bool     `help"Dump EBNF."`
	Files []string `arg:"" optional:"" type:"existingfile" help:"GraphQL schema files to parse."`
}

func main() {
	ctx := kong.Parse(&cli)
	if cli.EBNF {
		fmt.Println(parser.String())
		ctx.Exit(0)
	}
	for _, file := range cli.Files {
		ast := &ast.File{}
		r, err := os.Open(file)
		ctx.FatalIfErrorf(err)
		err = parser.Parse("", r, ast)
		r.Close()
		repr.Println(ast)
		ctx.FatalIfErrorf(err)
	}
}
