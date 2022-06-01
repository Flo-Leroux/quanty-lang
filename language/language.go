package language

import (
	"quanty/language/ast"
	"quanty/language/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type File struct {
	Path     string
	Document *ast.DocumentNode
}

func NewFile(path string) *File {
	file := &File{
		Path:     path,
		Document: FileToDocumentNode(path),
	}

	return file
}

func FileToDocumentNode(filePath string) *ast.DocumentNode {
	p := newParser(filePath)
	document := p.Document().(*parser.DocumentContext)
	visitor := &visitor{}

	return visitor.VisitDocument(document)
}

func newParser(filePath string) *parser.Parser {
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return p
}
