package ast

type File struct {
	Module    string
	Imports   *ImportFileList
	Component *ComponentDefinition
	// Operations OperationList
	// Fragments  FragmentDefinitionList
	Position *Position `dump:"-" json:"-"`
}
