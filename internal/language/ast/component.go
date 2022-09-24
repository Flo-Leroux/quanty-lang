package ast

import (
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// ComponentStatement -
type ComponentStatement struct {
	Token      token.Token // token.ENTITY
	Name       token.Token // token.IDENT
	Selections SelectionList
	// RelationStatements []Statement
	// ActionStatements   []Statement
	// Option             token.Token // token.OPTION
}

// statementNode -
func (ls *ComponentStatement) statementNode() {}

// TokenLiteral -
func (ls *ComponentStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *ComponentStatement) String() string {
	var sb strings.Builder
	sb.WriteString("component")
	sb.WriteString(" ")
	sb.WriteString(ls.Name.Literal)
	sb.WriteString(" {")
	sb.WriteString("\n")
	sb.WriteString("}")

	return sb.String()
}
