package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// parseComponentStatement returns a COMPONENT Statement AST Node
func (p *Parser) parseComponentStatement() *ast.ComponentStatement {
	stmt := &ast.ComponentStatement{Token: p.currentToken, Selections: ast.SelectionList{}}
	if !p.expectAndNext(token.IDENT) {
		return nil
	}
	stmt.Name = p.currentToken

	p.wrapWith(
		BraceWrapper,
		func() {
			switch p.currentToken.Type {
			case token.STRING:
				stmt.Selections = append(stmt.Selections, p.parseStringValue())
			case token.IDENT:
				stmt.Selections = append(stmt.Selections, p.parseField())
			}
		},
	)

	return stmt
}
