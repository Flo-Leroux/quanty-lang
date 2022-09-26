package ast

// ComponentStatement -
type ComponentStatement struct {
	name       string
	selections *List[Statement]
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

// Accept -
func (cs *ComponentStatement) Accept(v Visitor) {
	v.VisitComponentStatement(cs)
}

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
func (cs *ComponentStatement) Selections() *List[Statement] {
	return cs.selections
}
