package html

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

func (e *Engine) tab() {
	e.sb.WriteString(
		strings.Repeat("\t", e.indentLvl),
	)
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

func (e *Engine) write(str string) *Engine {
	e.sb.WriteString(str)
	return e
}

func (e *Engine) openTag(str string) {
	e.sb.WriteString(
		fmt.Sprintf("<%s>", str),
	)
}

func (e *Engine) closeTag(str string) {
	e.sb.WriteString(
		fmt.Sprintf("</%s>", str),
	)
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

		if i+1 < stmtsTotal {
			e.cReturn()
		}
	}
}

// VisitComponentStatement -
func (e *Engine) VisitComponentStatement(ctx *ast.ComponentStatement) {
	e.openTag(ctx.Name())

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

	e.closeTag(ctx.Name())
}

// VisitField -
func (e *Engine) VisitField(ctx *ast.Field) {
	e.openTag(ctx.Name())

	if !ctx.Selections().IsEmpty() {
		e.
			cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			e.tab()
			sel.Accept(e)
			e.cReturn()
		}

		e.
			dedent().
			tab()
	}

	e.closeTag(ctx.Name())
}

// VisitStringValue -
func (e *Engine) VisitStringValue(ctx *ast.StringValue) {
	e.write(ctx.Value())
}
