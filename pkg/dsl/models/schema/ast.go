package schema

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/list"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
)

// Schema -
type Schema struct {
	statements *list.List[node.Statement]
}

// NewSchema -
func New() *Schema {
	return &Schema{
		statements: list.New[node.Statement](node.STATEMENT),
	}
}

// Node -
func (s *Schema) Node() {}

// WithStatements -
func (s *Schema) WithStatements(stms ...node.Statement) *Schema {

	s.statements.Append(stms...)

	return s
}

// Statements -
func (s *Schema) Statements() *list.List[node.Statement] {
	return s.statements
}
