package component

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/list"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
)

// Component -
type Component struct {
	name       string
	selections *list.List[node.Statement]
}

func New(name string) *Component {
	return &Component{
		name:       name,
		selections: list.New[node.Statement](node.STATEMENT),
	}
}

// Node -
func (cs *Component) Node() {}

// StatementNode -
func (cs *Component) StatementNode() {}

// Name -
func (cs *Component) Name() string {
	return cs.name
}

// WithSelections -
func (cs *Component) WithSelections(stmts ...node.Statement) *Component {

	cs.selections.Append(stmts...)

	return cs
}

// Selections -
func (cs *Component) Selections() *list.List[node.Statement] {
	return cs.selections
}
