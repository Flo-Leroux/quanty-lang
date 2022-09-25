package ast

// ComponentStatement -
type ComponentStatement struct {
	name       string
	selections *StatementList
}

func NewComponentStatement(name string) *ComponentStatement {
	return &ComponentStatement{
		name:       name,
		selections: NewStatementList(),
	}
}

// node -
func (cs *ComponentStatement) node() {}

// statementNode -
func (cs *ComponentStatement) statementNode() {}

// Name -
func (cs *ComponentStatement) Name() string {
	return cs.name
}

// WithSelections -
func (cs *ComponentStatement) WithSelections(stmts ...Statement) *ComponentStatement {

	cs.selections.Append(stmts...)

	return cs
}

// Selections -
func (cs *ComponentStatement) Selections() *StatementList {
	return cs.selections
}
