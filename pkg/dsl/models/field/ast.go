package field

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/list"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
)

// Field -
type Field struct {
	name       string
	selections *list.List[node.Statement]
}

func New(name string) *Field {
	return &Field{
		name:       name,
		selections: list.New[node.Statement](node.STATEMENT),
	}
}

// node -
func (f *Field) Node() {}

// statementNode -
func (f *Field) StatementNode() {}

// // Accept -
// func (f *Field) Accept(v Visitor) {
// 	v.VisitField(f)
// }

// Name -
func (f *Field) Name() string {
	return f.name
}

// WithSelections -
func (f *Field) WithSelections(stmts ...node.Statement) *Field {
	f.selections.Append(stmts...)

	return f
}

// Selections -
func (f *Field) Selections() *list.List[node.Statement] {
	return f.selections
}
