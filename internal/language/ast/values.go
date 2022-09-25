package ast

// StringValue -
type StringValue struct {
	value string
}

// NewStringValue -
func NewStringValue(value string) *StringValue {
	return &StringValue{
		value: value,
	}
}

// node -
func (sv *StringValue) node() {}

// statementNode -
func (sv *StringValue) statementNode() {}

// Value -
func (sv *StringValue) Value() string {
	return sv.value
}

// Field -
type Field struct {
	name       string
	selections *StatementList
}

func NewField(name string) *Field {
	return &Field{
		name:       name,
		selections: NewStatementList(),
	}
}

// node -
func (f *Field) node() {}

// statementNode -
func (f *Field) statementNode() {}

// Name -
func (f *Field) Name() string {
	return f.name
}

// WithSelections -
func (f *Field) WithSelections(stmts ...Statement) *Field {
	f.selections.Append(stmts...)

	return f
}

// Selections -
func (f *Field) Selections() *StatementList {
	return f.selections
}
