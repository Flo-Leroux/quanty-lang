package query

import (
	"quanty/antlr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func newParser(filePath string) *parser.FileContext {
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return p.File().(*parser.FileContext)
}

type ModuleMap = map[string]*Module
type ComponentMap = map[string]*Component

type Program struct {
	root    string
	Modules ModuleMap
	parser.BaseParserVisitor
	currentFilePath string
}

func NewProgram(root string) *Program {
	return &Program{
		root:    root,
		Modules: make(map[string]*Module),
	}
}

func (p *Program) look() {
	files := findFiles("./app", ".qy")

	for _, filePath := range files {
		p.currentFilePath = filePath
		file := newParser(filePath)
		p.VisitFile(file)
	}
}

type Module struct {
	Name       string
	Ctx        *parser.ModuleDefContext
	Components ComponentMap
}

func (m *Module) GetComponent(name string) *Component {
	return m.Components[name]
}

type Component struct {
	FilePath string
	Name     string
	Ctx      *parser.ComponentDefContext
}
