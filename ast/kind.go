package ast

type Kind int64

func (k Kind) String() string {
	switch k {
	case name:
		return "Name"
	case document:
		return "Document"
	case componentDefinition:
		return "ComponentDefinition"
	case variableDefinition:
		return "VariableDefinition"

	case variable:
		return "Variable"
	case nonNullType:
		return "NonNullType"
	case namedType:
		return "NamedType"
	default:
		return "unknown"
	}
}

const (
	name Kind = iota
	document
	componentDefinition
	variableDefinition
	variable
	nonNullType
	namedType
)
