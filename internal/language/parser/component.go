package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// parseComponentStatement returns a COMPONENT Statement AST Node
func (p *Parser) parseComponentStatement() *ast.ComponentStatement {
	stmt := &ast.ComponentStatement{Token: p.currentToken, Fields: []*ast.Field{}}
	if !p.expectAndNext(token.IDENT) {
		return nil
	}
	stmt.Name = p.currentToken

	p.wrapWith(
		BraceWrapper,
		func() {
			switch p.currentToken.Type {
			case token.IDENT:
				stmt.Fields = append(stmt.Fields, p.parseField())
			}
		},
	)

	// if p.peekTokenIs(token.OPTION) {
	// 	p.next()
	// 	stmt.Option = p.currentToken
	// }

	return stmt
}
