package html

import (
	"fmt"
	"io"
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
)

// HTML - implements ast.Visitor
type HTML struct {
	sb        strings.Builder
	indentLvl int
}

func (h *HTML) tab() {
	h.sb.WriteString(
		strings.Repeat("\t", h.indentLvl),
	)
}

func (h *HTML) indent() {
	h.indentLvl += 1
}

func (h *HTML) dedent() *HTML {
	h.indentLvl -= 1
	return h
}

func (h *HTML) cReturn() *HTML {
	h.sb.WriteString("\n")
	return h
}

func (h *HTML) write(str string) *HTML {
	h.sb.WriteString(str)
	return h
}

func (h *HTML) openTag(str string) {
	h.sb.WriteString(
		fmt.Sprintf("<%s>", str),
	)
}

func (h *HTML) closeTag(str string) {
	h.sb.WriteString(
		fmt.Sprintf("</%s>", str),
	)
}

// New -
func New() *HTML {
	return &HTML{
		indentLvl: 0,
	}
}

// String -
func (h *HTML) String(v ast.Visitable) string {
	v.Accept(h)
	return h.sb.String()
}

// Write -
func (h *HTML) Write(w io.Writer, v ast.Visitable) (n int, err error) {
	v.Accept(h)
	return w.Write([]byte(h.sb.String()))
}

// VisitSchema -
func (h *HTML) VisitSchema(ctx *ast.Schema) {
	if ctx.Statements().IsEmpty() {
		return
	}

	stmts := ctx.Statements()
	stmtsTotal := stmts.Len()
	for i, stmt := range stmts.Items() {
		stmt.Accept(h)

		if i+1 < stmtsTotal {
			h.cReturn()
		}
	}
}

// VisitComponentStatement -
func (h *HTML) VisitComponentStatement(ctx *ast.ComponentStatement) {
	h.openTag(ctx.Name())

	if !ctx.Selections().IsEmpty() {
		h.cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			h.tab()
			sel.Accept(h)
			h.cReturn()
		}

		h.dedent()
	}

	h.closeTag(ctx.Name())
}

// VisitField -
func (h *HTML) VisitField(ctx *ast.Field) {
	h.openTag(ctx.Name())

	if !ctx.Selections().IsEmpty() {
		h.
			cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			h.tab()
			sel.Accept(h)
			h.cReturn()
		}

		h.
			dedent().
			tab()
	}

	h.closeTag(ctx.Name())
}

// VisitStringValue -
func (h *HTML) VisitStringValue(ctx *ast.StringValue) {
	h.write(ctx.Value())
}
