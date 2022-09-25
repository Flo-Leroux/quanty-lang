package ast

// Node -
type Node interface {
	node()
}

// NodeList T -
type NodeList = List[Node]

// New Node list
func NewNodeList() *NodeList {
	return NewList[Node]()
}

// Expression -
type Expression interface {
	Node
	expressionNode()
}

// ExpressionList T -
type ExpressionList = List[Expression]

// New Expression list
func NewExpressionList() *ExpressionList {
	return NewList[Expression]()
}

// Statement -
type Statement interface {
	Node
	statementNode()
}

// StatementList T -
type StatementList = List[Statement]

// New Statement list
func NewStatementList() *StatementList {
	return NewList[Statement]()
}
