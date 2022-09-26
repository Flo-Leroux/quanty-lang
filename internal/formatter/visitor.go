package formatter

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
)

// Formatter - implements ast.Visitor[string]
type Formatter struct {
	ast *ast.Schema
}

// NewFormatter -
func NewFormatter(schema *ast.Schema) *Formatter {
	return &Formatter{
		ast: schema,
	}
}

// Visit -
func (f *Formatter) Visit() string {
	return f.VisitSchema(f.ast)
}

// VisitSchema T = string -
func (f *Formatter) VisitSchema(ctx *ast.Schema) string {
	return ""
}

// VisitComponentStatement T = string -
func (f *Formatter) VisitComponentStatement(ctx *ast.ComponentStatement) string {
	return ""
}

// VisitStringValue T = string -
func (f *Formatter) VisitStringValue(ctx *ast.StringValue) string {
	return ""
}

// VisitField T = string -
func (f *Formatter) VisitField(ctx *ast.Field) string {
	return ""
}
