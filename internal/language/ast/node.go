package ast

import "github.com/Flo-Leroux/quanty-lang/internal/language/token"

// Node -
type Node interface {
	TokenLiteral() string
	String() string
}

// Expression -
type Expression interface {
	Node
	expressionNode()
	IsInfix() bool
	Type() string
}

// Statement -
type Statement interface {
	Node
	statementNode()
}

// Identifier -
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// statementNode -
func (ls *Identifier) expressionNode() {}

// TokenLiteral -
func (ls *Identifier) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *Identifier) String() string {
	return ls.Value
}

// IsInfix -
func (ls *Identifier) IsInfix() bool {
	return false
}

// Type -
func (ls *Identifier) Type() string {
	return "identifier"
}
