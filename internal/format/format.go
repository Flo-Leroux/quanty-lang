package format

import (
	"fmt"
	"io"
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
)

// Fmt - implements ast.Visitor
type Fmt struct {
	strings.Builder
	indentLvl int
}

func (f *Fmt) tab() *Fmt {
	f.WriteString(
		strings.Repeat("\t", f.indentLvl),
	)
	return f
}

func (f *Fmt) indent() {
	f.indentLvl += 1
}

func (f *Fmt) dedent() *Fmt {
	f.indentLvl -= 1
	return f
}

func (f *Fmt) cReturn() *Fmt {
	f.WriteString("\n")
	return f
}

func (f *Fmt) space() *Fmt {
	f.WriteString(" ")
	return f
}

func (f *Fmt) write(str string) *Fmt {
	f.WriteString(str)
	return f
}

func (f *Fmt) fmt(format string, a ...any) *Fmt {
	f.WriteString(
		fmt.Sprintf(format, a...),
	)
	return f
}

// New -
func New() *Fmt {
	return &Fmt{
		indentLvl: 0,
	}
}

// String -
func (f *Fmt) String(v ast.Visitable) string {
	v.Accept(f)
	return f.Builder.String()
}

// Write -
func (f *Fmt) Write(w io.Writer, v ast.Visitable) (n int, err error) {
	v.Accept(f)
	return w.Write([]byte(f.Builder.String()))
}

// VisitSchema -
func (f *Fmt) VisitSchema(ctx *ast.Schema) {
	if ctx.Statements().IsEmpty() {
		return
	}

	stmts := ctx.Statements()
	stmtsTotal := stmts.Len()
	for i, stmt := range stmts.Items() {
		stmt.Accept(f)

		f.cReturn()
		if i+1 < stmtsTotal {
			f.cReturn()
		}
	}
}

// VisitComponentStatement -
func (f *Fmt) VisitComponentStatement(ctx *ast.ComponentStatement) {
	f.write("component").
		space().
		write(ctx.Name()).
		space().
		write("{")

	if !ctx.Selections().IsEmpty() {
		f.cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			f.tab()
			sel.Accept(f)
			f.cReturn()
		}

		f.dedent()
	}

	f.write("}")
}

// VisitField -
func (f *Fmt) VisitField(ctx *ast.Field) {
	f.write(ctx.Name())

	if !ctx.Selections().IsEmpty() {
		f.
			space().
			write("{").
			cReturn().
			indent()

		for _, sel := range ctx.Selections().Items() {
			f.tab()
			sel.Accept(f)
			f.cReturn()
		}

		f.
			dedent().
			tab().
			write("}")
	}
}

// VisitStringValue -
func (f *Fmt) VisitStringValue(ctx *ast.StringValue) {
	f.fmt(`"%s"`, ctx.Value())
}
