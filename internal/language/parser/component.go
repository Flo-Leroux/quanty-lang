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
			if token.IDENT == p.currentToken.Type {
				stmt.Selections = append(stmt.Selections, p.parseField())
			}
		},
	)

	return stmt
}
