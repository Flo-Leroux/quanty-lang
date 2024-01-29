package ast

type (
	Schema struct {
		Components []*Component
	}

	Component struct {
		Name       string
		Selections []Selection
	}

	Field struct {
		Name       string
		Selections []Selection
	}

	StringValue struct {
		Value string
	}

	Selection interface {
		selection()
	}

	Value interface {
		value()
	}
)

// Implement Selection -
func (x *Field) selection()       {}
func (x *StringValue) selection() {}

// Implement Value -
func (x *StringValue) value() {}
