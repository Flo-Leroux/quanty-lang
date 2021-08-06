package ast

import (
	"fmt"
)

type SelectionSet []Selection

type Selection interface {
	isSelection()
	GetPosition() *Position
	String() string
}

func (*Field) isSelection()          {}
func (*FragmentSpread) isSelection() {}
func (*InlineFragment) isSelection() {}

func (s *Field) GetPosition() *Position          { return s.Position }
func (s *FragmentSpread) GetPosition() *Position { return s.Position }
func (s *InlineFragment) GetPosition() *Position { return s.Position }

func (s *Field) String() string {

	if s.IsTextContent {
		return s.Name
	}

	str := fmt.Sprintf("<%s ", s.Name)
	str += s.Arguments.String()
	str += ">"

	for _, set := range s.SelectionSet {
		str += set.String()
	}

	str += fmt.Sprintf("</%s>", s.Name)
	return str
}
func (s *FragmentSpread) String() string { return s.Name }
func (s *InlineFragment) String() string { return "Inline" }

func addNewLine(html string) string {
	return html + "\n"
}

type Field struct {
	Alias         string
	Name          string
	IsTextContent bool
	Arguments     ArgumentList
	Directives    DirectiveList
	SelectionSet  SelectionSet
	Position      *Position `dump:"-" json:"-"`

	// Require validation
	Definition       *FieldDefinition
	ObjectDefinition *Definition
}

type Argument struct {
	Name     string
	Value    *Value
	Position *Position `dump:"-" json:"-"`
}

func (arg *Argument) String() string {
	return fmt.Sprintf("%s=%s", arg.Name, arg.Value.String())
}

func (f *Field) ArgumentMap(vars map[string]interface{}) map[string]interface{} {
	return arg2map(f.Definition.Arguments, f.Arguments, vars)
}
