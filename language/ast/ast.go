package ast

type Kind string

const (
	NAME Kind = "Name"

	DOCUMENT             Kind = "Document"
	OPERATION_DEFINITION Kind = "OperationDefinition"
	VARIABLE_DEFINITION  Kind = "VariableDefinition"
	SELECTION_SET        Kind = "SelectionSet"
	FIELD                Kind = "Field"
	ARGUMENT             Kind = "Argument"

	VARIABLE Kind = "Variable"
	STRING   Kind = "StringValue"

	NAMED_TYPE    Kind = "NamedType"
	LIST_TYPE     Kind = "ListType"
	NON_NULL_TYPE Kind = "NonNullType"
)

type DocumentNode struct {
	Kind        Kind
	Module      string
	Imports     []string
	Definitions []DefinitionNode
}

type DefinitionNode interface {
	definitionNode()
}

type OperationTypeNode string

const (
	COMPONENT OperationTypeNode = "component"
)

type OperationDefinitionNode struct {
	Kind                Kind
	Operation           OperationTypeNode
	Name                *NameNode
	VariableDefinitions []*VariableDefinitionNode
	SelectionSet        *SelectionSetNode
}

func (od *OperationDefinitionNode) definitionNode() {}

type NameNode struct {
	Kind  Kind
	Value string
}

type VariableDefinitionNode struct {
	Kind         Kind
	Variable     *VariableNode
	Type         TypeNode
	DefaultValue ConstValueNode
}

type SelectionSetNode struct {
	Kind       Kind
	Selections []SelectionNode
}

type SelectionNode interface {
	selectionNode()
}

type FieldNode struct {
	Kind         Kind
	Alias        *NameNode
	Name         *NameNode
	Arguments    []*ArgumentNode
	SelectionSet *SelectionSetNode
}

func (n *FieldNode) selectionNode() {}

type ArgumentNode struct {
	Kind  Kind
	Name  *NameNode
	Value ValueNode
}

type ValueNode interface {
	valueNode()
}

type StringValueNode struct {
	Kind  Kind
	Value string
}

func (sv *StringValueNode) selectionNode() {}
func (sv *StringValueNode) valueNode()     {}

type VariableNode struct {
	Kind Kind
	Name *NameNode
}

func (v *VariableNode) valueNode() {}

type TypeNode interface {
	typeNode()
	IsNamedType() bool
	IsListType() bool
	IsNonNullType() bool
}

type NamedTypeNode struct {
	Kind Kind
	Name *NameNode
}

func (t *NamedTypeNode) typeNode() {}

func (t *NamedTypeNode) IsNamedType() bool {
	return true
}

func (t *NamedTypeNode) IsListType() bool {
	return false
}

func (t *NamedTypeNode) IsNonNullType() bool {
	return false
}

type ListTypeNode struct {
	Kind Kind
	Type *TypeNode
}

func (t *ListTypeNode) typeNode() {}

func (t *ListTypeNode) IsNamedType() bool {
	return false
}

func (t *ListTypeNode) IsListType() bool {
	return true
}

func (t *ListTypeNode) IsNonNullType() bool {
	return false
}

type NonNullTypeNode struct {
	Kind Kind
	Type *TypeNode
}

func (t *NonNullTypeNode) typeNode() {}

func (t *NonNullTypeNode) IsNamedType() bool {
	return false
}

func (t *NonNullTypeNode) IsListType() bool {
	return false
}

func (t *NonNullTypeNode) IsNonNullType() bool {
	return true
}

type ConstValueNode interface {
	constValueNode() string
}
