package ast

// Schema -
type Schema struct {
	statements *StatementList
}

// NewSchema -
func NewSchema() *Schema {
	return &Schema{
		statements: NewStatementList(),
	}
}

// WithStatements -
func (s *Schema) WithStatements(stms ...Statement) *Schema {

	s.statements.Append(stms...)

	return s
}

// Statements -
func (s *Schema) Statements() *StatementList {
	return s.statements
}
