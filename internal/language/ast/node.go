package ast

// Node -
type Node interface {
	node()
}

// New Node list
func NewNodeList() *List[Node] {
	return NewList[Node]("node")
}

// Expression -
type Expression interface {
	Node
	expressionNode()
}

// New Expression list
func NewExpressionList() *List[Expression] {
	return NewList[Expression]("expression")
}

// Statement -
type Statement interface {
	Node
	Visitable[any]
	statementNode()
}

// New Statement list
func NewStatementList() *List[Statement] {
	return NewList[Statement]("statement")
}
