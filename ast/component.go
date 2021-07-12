package ast

type ComponentDefinition struct {
	Kind                DefinitionKind
	Name                string
	VariableDefinitions VariableDefinitionList
	Directives          DirectiveList
	SelectionSet        SelectionSet
	Position            *Position `dump:"-"`
}
