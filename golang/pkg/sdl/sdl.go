package sdl

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/sdl/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func Run[T parser.QuantyParserListener](source string, listener T) T {
	input := antlr.NewInputStream(source)
	lexer := parser.NewQuantyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewQuantyParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Schema()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener
}
