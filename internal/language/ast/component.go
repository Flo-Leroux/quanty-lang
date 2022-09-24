package ast

import (
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// ComponentStatement -
type ComponentStatement struct {
	Token  token.Token // token.ENTITY
	Name   token.Token // token.IDENT
	Fields []*Field
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

	// // for _, rs := range ls.RelationStatements {
	// // 	sb.WriteString(rs.String())
	// // 	sb.WriteString("\n")
	// // }

	// // sb.WriteString("\n")

	// // for _, rs := range ls.ActionStatements {
	// // 	sb.WriteString(rs.String())
	// // 	sb.WriteString("\n")
	// // }

	sb.WriteString("}")
	// // sb.WriteString(" ")

	// // if ls.Option.Literal != "" {
	// // 	sb.WriteString("`")
	// // 	sb.WriteString(ls.Option.Literal)
	// // 	sb.WriteString("`")
	// // }

	// // sb.WriteString("\n")
	return sb.String()
}
