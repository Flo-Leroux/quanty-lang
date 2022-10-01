package node

// Node -
type Node interface {
	Node()
}

// Expression -
type Expression interface {
	Node
	ExpressionNode()
}

// Statement -
type Statement interface {
	Node
	// Visitable
	StatementNode()
}
