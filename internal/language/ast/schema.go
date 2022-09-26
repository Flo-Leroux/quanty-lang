package ast

// Schema -
type Schema struct {
	statements *List[Statement]
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
func (s *Schema) Statements() *List[Statement] {
	return s.statements
}

// Accept -
func (s *Schema) Accept(v Visitor) {
	v.VisitSchema(s)
}
