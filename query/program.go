package query

import (
	"errors"
	"fmt"
	"quanty/antlr/parser"
	"sync"

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

func (p *Program) AppendModule(m *Module) {
	if p.Modules[m.Name] == nil {
		p.Modules[m.Name] = m
	}
}

func (p *Program) GetModule(name string) *Module {
	return p.Modules[name]
}

func (p *Program) look() {
	files := findFiles("./app", ".qy")
	var wg sync.WaitGroup

	for _, filePath := range files {
		wg.Add(1)

		go func(filePath string) {
			defer wg.Done()
			p.currentFilePath = filePath
			file := newParser(filePath)
			p.VisitFile(file)
		}(filePath)
	}

	wg.Wait()
}

type Import struct {
	Name     string
	Ctx      *parser.ImportDefContext
	Elements []string
}

func NewImport(ctx *parser.ImportDefContext) *Import {

	importer := &Import{
		Name:     ctx.GetModule().GetText(),
		Ctx:      ctx,
		Elements: make([]string, 0),
	}

	for _, i := range ctx.GetImports() {
		importer.Elements = append(importer.Elements, i.GetText())
	}

	return importer
}

type ModuleImport = map[string]*Import

type Module struct {
	Name       string
	Imports    ModuleImport
	Ctx        *parser.ModuleDefContext
	Components ComponentMap
}

func NewModule(ctx *parser.ModuleDefContext) *Module {
	return &Module{
		Name:       ctx.GetName().GetText(),
		Ctx:        ctx,
		Imports:    make(map[string]*Import),
		Components: make(map[string]*Component),
	}
}

func (m *Module) AppendComponent(c *Component) {
	if m.Components[c.Name] != nil {
		err := errors.New(
			fmt.Sprintf("Component %s already exist in module %s", c.Name, m.Name),
		)
		panic(err)
	}

	m.Components[c.Name] = c
}

func (m *Module) AppendImport(i *Import) {
	if m.Imports[i.Name] == nil {
		m.Imports[i.Name] = i
	}
}

func (m *Module) GetComponent(name string) *Component {
	return m.Components[name]
}

type Component struct {
	Name string
	Ctx  *parser.ComponentDefContext
}

func NewComponent(ctx *parser.ComponentDefContext) *Component {
	return &Component{
		Name: ctx.GetName().GetText(),
		Ctx:  ctx,
	}
}
