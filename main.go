package main

import (
	"fmt"
	"os/user"

	"github.com/Flo-Leroux/quanty-lang/pkg/complete"
	"github.com/Flo-Leroux/quanty-lang/pkg/complete/caret"
	"github.com/Flo-Leroux/quanty-lang/pkg/sdl/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type vocabulary struct {
	p antlr.Parser
}

func (v *vocabulary) GetDisplayName(index int) string {
	return v.p.GetSymbolicNames()[index]
}

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

	source := `component Main {
h1 {
"Hello World! My name is"
    strong {
      "Quan-team!"
    }
  }
}`

	caretPosition := caret.CaretPosition{
		Line:   1,
		Column: 11,
	}

	input := antlr.NewInputStream(source)
	lexer := parser.NewQuantyLexer(input)
	lexer.RemoveErrorListeners()

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewQuantyParser(stream)
	p.RemoveErrorListeners()
	p.BuildParseTrees = true

	p.Schema()

	// tree := p.Schema()

	// caretTokenIndex := caret.ComputeTokenPosition(
	// 	tree,
	// 	stream,
	// 	caretPosition,
	// 	[]int{
	// 		parser.QuantyParserRULE_componentName,
	// 		parser.QuantyParserRULE_fieldName,
	// 	},
	// )

	idx := caret.FindCursorTokenIndex(stream, caretPosition)
	fmt.Println(idx)
	if idx == -1 {
		// if caretTokenIndex == nil {
		fmt.Println("TokenPosition not valid")
		return
	}

	voc := &vocabulary{
		p: p,
	}

	core := complete.NewCodeCompletionCore(
		p.GetATN(),
		voc,
		p.GetRuleNames(),
		"Quanty",
		p,
	)
	// core.ShowResult()

	core.SetPreferredRules(
		parser.QuantyParserRULE_componentName,
		parser.QuantyParserRULE_fieldName,
	)

	candidates := core.CollectCandidates(p.GetTokenStream(), idx, nil)

	for r := range candidates.Rules().ToMap() {
		fmt.Println("Rule", voc.GetDisplayName(r))
	}

	for t := range candidates.Tokens().ToMap() {
		fmt.Println("Tok", voc.GetDisplayName(t))
	}

	// 	l := html.New("locale string")

	// 	sdl.Run(src, l)

	// 	tmpl := l.Lookup("Nav")

	// 	tmpl.Execute(os.Stdout, "COUCOU")
	fmt.Println()

}
