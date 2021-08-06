package ast

import "fmt"

type ComponentDefinition struct {
	Kind                DefinitionKind
	Name                string
	VariableDefinitions VariableDefinitionList
	Directives          DirectiveList
	SelectionSet        SelectionSet
	Position            *Position `dump:"-" json:"-"`
}

func (c *ComponentDefinition) String() string {
	str := fmt.Sprintf("{{ define \"%s\" }}", c.Name)
	for _, set := range c.SelectionSet {
		str += set.String()
		// html += c.SelectionSet[0].Html(html)
		// html += "\n"
	}
	str += "{{ end }}"
	return str
}
