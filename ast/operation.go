package ast

type Operation string

// const (
// 	Import    Operation = "import"
// 	Component Operation = "component"
// )

type OperationDefinition struct {
	Operation           Operation
	Name                string
	VariableDefinitions VariableDefinitionList
	Directives          DirectiveList
	SelectionSet        SelectionSet
	Position            *Position `dump:"-" json:"-"`
}

type ImportDefinition struct {
	Operation Operation
	Files     ImportFileList
	Position  *Position `dump:"-" json:"-"`
}

type VariableDefinition struct {
	Variable     string
	Type         *Type
	DefaultValue *Value
	Directives   DirectiveList
	Position     *Position `dump:"-" json:"-"`

	// Requires validation
	Definition *Definition
	Used       bool `dump:"-" json:"-"`
}
