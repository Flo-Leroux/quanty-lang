package format

import (
	"fmt"
	"io"
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
)

// Engine - implements ast.Visitor
type Engine struct {
	sb        strings.Builder
	indentLvl int
}

func (e *Engine) tab() *Engine {
	e.sb.WriteString(
		strings.Repeat("\t", e.indentLvl),
	)
	return e
}

func (e *Engine) indent() {
	e.indentLvl += 1
}

func (e *Engine) dedent() *Engine {
	e.indentLvl -= 1
	return e
}

func (e *Engine) cReturn() *Engine {
	e.sb.WriteString("\n")
	return e
}

func (e *Engine) space() *Engine {
	e.sb.WriteString(" ")
	return e
}

func (e *Engine) write(str string) *Engine {
	e.sb.WriteString(str)
	return e
}

func (e *Engine) fmt(format string, a ...any) *Engine {
	e.sb.WriteString(
		fmt.Sprintf(format, a...),
	)
	return e
}

// New -
func New() *Engine {
	return &Engine{
		indentLvl: 0,
	}
}

// String -
func (e *Engine) String(v ast.Visitable) string {
	v.Accept(e)
	return e.sb.String()
}

// Write -
func (e *Engine) Write(w io.Writer, v ast.Visitable) (n int, err error) {
	v.Accept(e)
	return w.Write([]byte(e.sb.String()))
}

// VisitSchema -
func (e *Engine) VisitSchema(ctx *ast.Schema) {
	if ctx.Statements().IsEmpty() {
		return
	}

	stmts := ctx.Statements()
	stmtsTotal := stmts.Len()
	for i, stmt := range stmts.Items() {
		stmt.Accept(e)

		e.cReturn()
		if i+1 < stmtsTotal {
			e.cReturn()
		}
	}
}

// VisitComponentStatement -
func (e *Engine) VisitComponentStatement(ctx *ast.ComponentStatement) {
	e.write("component").
		space().
		write(ctx.Name()).
		space().
		write("{")

	if !ctx.Selections().IsEmpty() {
		e.cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			e.tab()
			sel.Accept(e)
			e.cReturn()
		}

		e.dedent()
	}

	e.write("}")
}

// VisitField -
func (e *Engine) VisitField(ctx *ast.Field) {
	e.write(ctx.Name())

	if !ctx.Selections().IsEmpty() {
		e.
			space().
			write("{").
			cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			e.tab()
			sel.Accept(e)
			e.cReturn()
		}

		e.
			dedent().
			tab().
			write("}")
	}
}

// VisitStringValue -
func (e *Engine) VisitStringValue(ctx *ast.StringValue) {
	e.fmt(`"%s"`, ctx.Value())
}
